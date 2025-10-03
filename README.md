# ModelG

ModelG is a code generation tool that generates Go data structures and database access code from SQL queries and model definitions.

## Overview

ModelG generates `*_gen.go` files containing:

- Go structs for database models
- Type-safe query methods
- Database access utilities
- Result mapping functions

## Quick Start

For ModelG to work, you need these files in your target directory:

1. **`models.go`** - Model definitions and query interfaces
2. **`modelg.yaml`** - Configuration file
3. **`<model>.sql`** - SQL query definitions (named after each model)

Add this generate directive to your `models.go`:

```go
//go:generate go run github.com/partite-ai/modelg/cmd/modelg
```

## Models

Models are Go structs that represent database tables. ModelG uses struct tags to control code generation behavior.

### Struct Tags

Use the `modelg` struct tag to control field behavior:

```go
type User struct {
    ID          UserID        `modelg:"pk,computed"`           // Primary key, auto-generated
    Name        string                                         // Regular field
    Email       optional.Optional[string]                     // Optional field
    Preferences *UserPrefs    `modelg:"converter:json"`       // JSON converter
    CreatedAt   time.Time     `modelg:"computed"`             // Database timestamp
    UpdatedAt   time.Time     `modelg:"computed"`             // Database timestamp
    Contact     ContactInfo   `modelg:"flatten"`              // Embed fields
}
```

**Available Tags:**

- `pk` - Primary key field
- `computed` - Database-generated field (auto-increment, timestamps)
- `immutable` - Cannot be updated after creation
- `converter:name` - Use named converter for type conversion
- `flatten` - Embed struct fields into parent

**Tag Effects on Parameter Structs:**

- `computed` fields excluded from Create/Update parameters
- `immutable` fields excluded from Update parameters
- `pk,computed` fields are auto-generated primary keys

### Parameter Structs

ModelG automatically generates parameter structs for create and update operations:

```go
// Generated automatically (unless skip_create_parameters: true)
type UserCreateParams struct {
    Name        string
    Email       optional.Optional[string]
    Preferences optional.Optional[*UserPrefs]
    // computed fields excluded
}

// Generated automatically (unless skip_update_parameters: true)
type UserUpdateParams struct {
    Name        optional.Optional[string]
    Email       optional.Optional[string]
    Preferences optional.Optional[*UserPrefs]
    // computed and immutable fields excluded
}
```

### Flattening

Use `flatten` to embed struct fields directly:

```go
type UserProfile struct {
    UserID   UserID       `modelg:"immutable"`
    Contact  ContactInfo  `modelg:"flatten"`
}

type ContactInfo struct {
    Phone   string  // Becomes part of UserProfile for database mapping
    Address string  // Becomes part of UserProfile for database mapping
    City    string  // Becomes part of UserProfile for database mapping
}
```

## Query Interfaces

Define query interfaces to specify database operations. **Query interfaces must be private** (lowercase) and follow the naming pattern `modelQueries` for model `Model`.

```go
type userQueries interface {
    CreateUser(ctx context.Context, params *UserCreateParams) (*User, error)
    GetUser(ctx context.Context, id UserID) (*User, error)
    UpdateUser(ctx context.Context, id UserID, params *UserUpdateParams) (*User, error)
    DeleteUser(ctx context.Context, id UserID) error
    FindAllUsers(ctx context.Context, page *PageParameters) (PagedRowSet[User], error)
}
```

### Generated Queries Struct

For each model, ModelG generates a `<Model>Queries` struct that implements the query interface using the corresponding SQL file. For example, the `User` model generates:

```go
// Generated in *_gen.go
type UserQueries struct {
    db Queryer
}

func NewUserQueries(db Queryer) *UserQueries {
    return &UserQueries{db: db}
}

// Implements all methods from userQueries interface
func (q *UserQueries) CreateUser(ctx context.Context, params *UserCreateParams) (*User, error) {
    // Generated implementation using user.sql queries
}
```

### Automatic CRUD Method Generation

ModelG can automatically add common CRUD methods to your query interface if the corresponding SQL queries exist in your `.sql` file, even if you don't explicitly define them in the interface. If you explicitly declare a metod int the interface, the automatic method will not be generated. This saves you from having to declare standard CRUD operations:

**Automatically Generated Methods:**

- `Create<Model>(ctx context.Context, params *<Model>CreateParams) (*<Model>, error)` - if `Create<Model>` query exists
- `Get<Model>(ctx context.Context, id <PKType>) (*<Model>, error)` - if `Get<Model>` query exists and model has a primary key
- `Update<Model>(ctx context.Context, id <PKType>, params *<Model>UpdateParams) (*<Model>, error)` - if `Update<Model>` query exists and model has update fields
- `Delete<Model>(ctx context.Context, id <PKType>) error` - if `Delete<Model>` query exists and model has a primary key

**Example:**

```go
// You only need to define custom methods in the interface
type userQueries interface {
    // Custom methods only
    FindUsersByStatus(ctx context.Context, status string) ([]*User, error)
    GetUserCount(ctx context.Context) (int, error)

    // CRUD methods are automatically added if these queries exist in user.sql:
    // - CreateUser (if CreateUser query exists)
    // - GetUser (if GetUser query exists and User has primary key)
    // - UpdateUser (if UpdateUser query exists and User has update fields)
    // - DeleteUser (if DeleteUser query exists and User has a primary key)
}
```

```sql
-- user.sql
-- name: CreateUser
INSERT INTO users (name, email) VALUES (:name, :email) RETURNING *;

-- name: GetUser
SELECT * FROM users WHERE id = :id;

-- name: UpdateUser
UPDATE users SET name = :name WHERE id = :id RETURNING *;

-- name: DeleteUser
DELETE FROM users WHERE id = :id;

-- name: FindUsersByStatus
SELECT * FROM users WHERE status = :status;

-- name: GetUserCount
SELECT COUNT(*) FROM users;
```

With this setup, the generated `UserQueries` struct will have all six methods available, even though only two were explicitly declared in the interface.

### Extending with Custom Methods

To add wrapper methods or multi-step database operations, define private methods in the query interface and implement public methods in a separate file:

**1. Add private methods to the interface:**

```go
type userQueries interface {
    // Generated methods
    CreateUser(ctx context.Context, params *UserCreateParams) (*User, error)
    GetUser(ctx context.Context, id UserID) (*User, error)

    // Private methods for internal use
    getUserByEmail(ctx context.Context, email string) (*User, error)
    updateUserLastLogin(ctx context.Context, id UserID) error
}
```

**2. Create a separate file with public extensions:**

```go
// user_queries.go (separate from generated files)
func (q *UserQueries) LoginUser(ctx context.Context, email, password string) (*User, error) {
    // Multi-step operation using private methods
    user, err := q.getUserByEmail(ctx, email)
    if err != nil {
        return nil, err
    }

    // Verify password logic here
    if !verifyPassword(user, password) {
        return nil, ErrInvalidCredentials
    }

    // Update last login timestamp
    if err := q.updateUserLastLogin(ctx, user.ID); err != nil {
        return nil, err
    }

    return user, nil
}
```

This pattern allows you to:

- Keep complex business logic separate from generated code
- Compose multiple database operations into single public methods
- Add validation, error handling, and business rules around generated queries
- Maintain clean separation between generated and custom code

### Method Signature Patterns

ModelG supports specific method signature patterns that determine query behavior:

#### 1. Execute Queries (No Results)

```go
MethodName(ctx context.Context, ...params) error
```

- **Usage:** INSERT, UPDATE, DELETE operations without returned data
- **Query Type:** `exec`

#### 2. Single Row Queries

```go
MethodName(ctx context.Context, ...params) (T, error)
MethodName(ctx context.Context, ...params) (*T, error)
```

- **Usage:** SELECT operations returning exactly one row
- **Query Type:** `get`

#### 3. Multiple Row Queries

```go
MethodName(ctx context.Context, ...params) ([]T, error)
MethodName(ctx context.Context, ...params) ([]*T, error)
```

- **Usage:** SELECT operations returning multiple rows
- **Query Type:** `list`

#### 4. Database Result Queries

```go
MethodName(ctx context.Context, ...params) (sql.Result, error)
```

- **Usage:** Access to rows affected, last insert ID, etc.
- **Query Type:** `execresult`

**Parameter Patterns:**

```go
// Direct parameters map to SQL: :id, :status
GetUser(ctx context.Context, id UserID, status string) (*User, error)

// Struct parameters map to fields: :new_name, :new_email
UpdateUser(ctx context.Context, id UserID, params *UserUpdateParams) (*User, error)
```

## SQL Queries

SQL files define queries using named parameters and conditional directives. **SQL files should be named according to their corresponding model** (e.g., `user.sql` for `User` model queries).

### Basic Syntax

```sql
-- name: CreateUser
INSERT INTO users (name, email, created_at)
VALUES (:name, :email, NOW())
RETURNING *;

-- name: GetUser
SELECT * FROM users WHERE id = :id;
```

### Parameter Types

- **Named Parameters:** `:parameter_name` - Bound as query parameters
- **Literal Parameters:** `:parameter_name!` - Written directly into SQL
- **Mode Parameters:** `:parameter_name!mode` - Calls `parameter.SQLText(mode)`

### Conditional Syntax

```sql
-- name: UpdateUser
UPDATE users SET updated_at = NOW()
    --when :name.IsSet
    ,name = :name
    --endwhen
    --when :email.IsSet
    ,email = :email
    --endwhen
WHERE id = :id
RETURNING *;
```

**Conditional Types:**

- `--when $EXPR` / `--endwhen` - Include content when expression is true
- `--<when $EXPR` / `--endwhen` - Same as `--when`, but removes entire line when false
- `--+when $EXPR` - Include content only if previous block rendered AND expression is true

### Parameter Mapping

SQL named parameters must map to Go method parameters or struct fields:

- `:name` → `Name` field
- `:user_id` → `UserID` field (snake_case → PascalCase)
- `:new_email` → `NewEmail` field in parameter struct

## Configuration

### modelg.yaml

```yaml
models:
  - name: User
  - name: Product
    skip_create_parameters: true
    skip_update_parameters: true
  - name: Order
    display_name: Purchase Order

converters:
  - name: json
    scanner: myapp/pkg/converters.JSONScanner
    valuer: myapp/pkg/converters.JSONValuer
  - scanner: myapp/pkg/converters.OptionalScanner
    valuer: myapp/pkg/converters.OptionalValuer
```

**Model Options:**

- `name` - Model name (must match struct name)
- `skip_create_parameters` - Skip generating CreateParams struct
- `skip_update_parameters` - Skip generating UpdateParams struct
- `display_name` - Override human-readable name for error messages

## Converters

Converters handle type conversion between Go types and database types.

### Configuration

```yaml
converters:
  - name: json
    scanner: myapp/pkg/converters.JSONScanner
    valuer: myapp/pkg/converters.JSONValuer
```

**Options:**

- `name` - Optional name for struct tag reference (`converter:json`)
- `scanner` - Import path for constructor function returning `sql.Scanner`
- `valuer` - Import path for constructor function returning `driver.Valuer`

### Constructor Functions

```go
// Constructor function signatures
func JSONScanner[T any](dest *T) sql.Scanner { ... }
func JSONValuer[T any](src T) driver.Valuer { ... }
```

### Dependency Injection

Converters can depend on other converters for nested types:

```go
func OptionalScanner[T any](dest *optional.Optional[T], innerScanner func(*T) sql.Scanner) sql.Scanner {
    return &optionalScannerImpl[T]{dest: dest, innerScanner: innerScanner}
}
```

For fields with layered types like `optional.Optional[*UserSettings]`, ModelG automatically injects the appropriate converter functions.

## Error Translation

ModelG provides standardized error types and metadata for user-friendly error messages.

### Generated Metadata

ModelG automatically generates a `GetModelMetadata` function for metadata that can be used in error translation:

```go
// Generated automatically
func GetModelMetadata(tableName string) *modelg.ModelMetadata {
    // Returns metadata about the given table
}
```

### Error Types

- `UniqueConstraintError` - Unique constraint violations
- `NotNullConstraintError` - NOT NULL constraint violations
- `ForeignKeyConstraintError` - Foreign key constraint violations
- `CheckConstraintError` - Check constraint violations

### Usage

```go
func CreateErrorTranslator() modelg.ErrorTranslator {
    return DBErrorTranslator(GetModelMetadata)
}

// Wrap database connection
translatedDB := modelg.NewErrorTranslatingDB(db, CreateErrorTranslator())
```

## Design Philosophy

**Templates Remain Valid SQL** - ModelG's key design principle ensures that SQL templates can be copied directly into SQL clients for testing and debugging. Conditional syntax uses SQL comments so templates stay syntactically correct.

Benefits:

- Easy testing in SQL clients
- EXPLAIN plan analysis with sample parameters
- Immediate SQL syntax error detection
- High readability

## File Structure

```
your-project/
├── models.go           # Model definitions and query interfaces
├── modelg.yaml         # Configuration
├── user.sql           # SQL queries for User model
├── product.sql        # SQL queries for Product model
├── order.sql          # SQL queries for Order model
└── *_gen.go           # Generated files (do not edit)
```
