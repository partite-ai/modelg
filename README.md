# ModelG

ModelG is a code generation tool that generates Go data structures and database access code from SQL queries and model definitions.

## Overview

ModelG generates `*_gen.go` files containing:
- Go structs for database models
- Type-safe query methods
- Database access utilities
- Result mapping functions

## Required Files

For modelg to work, you need these files in your target directory:

### 1. `models.go`
Contains the base model definitions and any custom methods. This file should include a `//go:generate` directive:

```go
//go:generate go run ../../cmd/modelg
```

### 2. `modelg.yaml`
Configuration file specifying generation parameters:

```yaml
models:
  - name: Agent
  - name: AgentVersion
    skip_create_parameters: true
    skip_update_parameters: true
  - name: ApiKey
    
converters:
  - name: json
    database: sqlite
    scanner: github.com/partite-ai/mesh/pkg/datastore/converters.JSONScanner
    valuer: github.com/partite-ai/mesh/pkg/datastore/converters.JSONValuer
  - scanner: github.com/partite-ai/mesh/pkg/datastore/converters.OptionalScanner
    valuer: github.com/partite-ai/mesh/pkg/datastore/converters.OptionalValuer
```

**Configuration Options:**

**Models:**
- `name`: Model name (must match struct name in models.go)
- `skip_create_parameters`: Skip generating CreateParams struct
- `skip_update_parameters`: Skip generating UpdateParams struct

**Converters:**
- `name`: Optional converter name for reference in struct tags
- `database`: Target database type (e.g., "sqlite", "pgx")
- `scanner`: Go import path for database scanner implementation
- `valuer`: Go import path for database valuer implementation

### 3. SQL Query Files
SQL files with specific syntax for defining queries. ModelG uses named parameters with `:` prefix and conditional directives. Example `queries.sql`:

```sql
-- name: CreateUser
INSERT INTO users (name, email, created_at) 
VALUES (:name, :email, NOW()) 
RETURNING *;

-- name: UpdateUser  
UPDATE users SET 
    updated_at = NOW()
    
    --when :name.IsSet
    ,name = :name
    --endwhen
    
    --when :email.IsSet
    ,email = :email  
    --endwhen
WHERE id = :id
RETURNING *;

-- name: GetUser
SELECT * FROM users WHERE id = :id;

-- name: FindAllUsers
SELECT 
    *,
    :order_by! AS __cursor,
    lead(id) OVER (ORDER BY :order_by! :order!) IS NULL AS __last
FROM users
WHERE --<when :cursor.IsSet
    :order_by! :order!compare :cursor 
--endwhen
ORDER BY :order_by! :order!
LIMIT :limit;
```

ModelG SQL Syntax:
- **Named Parameters**: Use `:parameter_name` for query parameters (passed as bind parameters)
- **Literal Parameters**: Use `:parameter_name!` to write parameter as literal text into the SQL query (not bound)
- **SQLText Mode**: Use `:parameter_name!mode` to pass "mode" to the parameter's `SQLText(mode string)` method
- **Conditional Blocks**: 
  - `--when $EXPR` / `--endwhen` - Include content only when expression is true
  - `--<when $EXPR` / `--endwhen` - Same as `--when`, but chomps the entire line when false (useful for clauses like WHERE)
  - `--+when $EXPR` - Within a conditional group, includes text on same line only if a previous block was rendered AND expression is true (useful for delimiters)

**Conditional Block Examples:**
```sql
-- Simple conditional
--when :name.IsSet
,name = :name
--endwhen

-- Chomping conditional (removes WHERE if no conditions)
WHERE --<when :foo.IsSet
foo = :foo
--endwhen

-- Delimiter conditional (adds AND only if previous condition rendered)
WHERE --<when :foo.IsSet
foo = :foo
AND --+when :bar.IsSet
bar = :bar
--endwhen
```

## Design Rationale

ModelG's syntax is designed with a key principle in mind: **templates remain valid SQL**. This means you can copy any ModelG SQL template and paste it directly into a SQL client for testing, debugging, or generating explain plans.

The conditional syntax uses SQL comments (`--`) so it doesn't interfere with SQL parsing. The somewhat unusual placement of delimiters (like `AND --+when`) serves this purpose - the template stays syntactically correct SQL even with all the conditional logic in place.

This design choice provides several benefits:
- **Easy Testing**: Copy templates to your SQL client to test queries
- **Debugging**: Use EXPLAIN to analyze query plans with sample parameters
- **Validation**: SQL syntax errors are caught immediately
- **Readability**: Templates look like normal SQL with embedded logic

## Query and Parameter Binding

ModelG works by processing both `models.go` and SQL files together. **Go method signatures are defined first** in `models.go`, and ModelG validates that SQL named parameters correctly map to the defined Go structures.

### Method Definition Process
1. **Define Go interfaces** in `models.go` with method signatures
2. **Define parameter structs** in `models.go` (e.g., `*CreateParams`, `*UpdateParams`)
3. **Write SQL queries** with named parameters that match the Go definitions
4. **ModelG validates** that all named parameters in SQL map to fields in the defined structs

### Parameter Mapping Rules
Named parameters in SQL queries must map to Go struct fields or method parameters:

**Parameter struct fields:**
- `:name` in SQL → `Name` field in Go struct
- `:agent_id` in SQL → `AgentID` field in Go struct (snake_case → PascalCase)
- `:new_input_name` in SQL → `NewInputName` field in Go struct

**Method parameters:**
- Single parameters like `:id` → direct method parameters  
- Complex parameters → fields in dedicated parameter structs

**Example from models.go:**
```go
type InputUpdateParams struct {
    NewInputName      optional.Optional[string]
    NewAgentVersionID optional.Optional[common.EntityID]
}

type inputQueries interface {
    UpdateInput(ctx context.Context, meshID common.EntityID, name string, params *InputUpdateParams) (*Input, error)
}
```

**Corresponding SQL:**
```sql
-- name: UpdateInput
UPDATE input SET updated_at = NOW()
    --when :new_input_name.IsSet
    ,name = :new_input_name
    --endwhen
    --when :new_agent_version_id.IsSet
    ,agent_version_id = :new_agent_version_id
    --endwhen
WHERE mesh_id = :mesh_id AND name = :name
```

ModelG ensures `:new_input_name` maps to `NewInputName`, `:new_agent_version_id` maps to `NewAgentVersionID`, etc.
