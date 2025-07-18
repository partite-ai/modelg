package modelg

import "fmt"

// QueryVariablesScope is an interface for managing variables bound to a query.
type QueryVariablesScope interface {
	// QueryArguments returns all arguments bound to the query.
	QueryArguments() []any

	// Creates a placeholder in a sql query for a variable with the given name.
	CreatePlaceholder(name string, value any) string
}

type GenericQueryVariablesScope struct {
	vars []any
}

func NewGenericQueryVariablesScope() *GenericQueryVariablesScope {
	return &GenericQueryVariablesScope{
		vars: make([]any, 0),
	}
}

func (s *GenericQueryVariablesScope) QueryArguments() []any {
	return s.vars
}

func (s *GenericQueryVariablesScope) CreatePlaceholder(name string, value any) string {
	s.vars = append(s.vars, value)
	return "?"
}

type PostgresQueryVariablesScope struct {
	vars      []any
	nameToVar map[string]int
}

func NewPostgresQueryVariablesScope() *PostgresQueryVariablesScope {
	return &PostgresQueryVariablesScope{
		vars:      make([]any, 0),
		nameToVar: make(map[string]int),
	}
}

func (s *PostgresQueryVariablesScope) QueryArguments() []any {
	return s.vars
}

func (s *PostgresQueryVariablesScope) CreatePlaceholder(name string, value any) string {
	if idx, ok := s.nameToVar[name]; ok {
		return fmt.Sprintf("$%d", idx+1)
	}
	idx := len(s.vars)
	s.nameToVar[name] = idx
	s.vars = append(s.vars, value)
	return fmt.Sprintf("$%d", idx+1)
}
