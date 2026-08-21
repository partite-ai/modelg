package modelg

import (
	"fmt"
	"strings"
)

type InValues []any

var _ SQLTexter = InValues{}

func (c InValues) SQLText(ctx *SQLTexterContext) (string, error) {
	var buf strings.Builder
	for i, v := range c {
		if i > 0 {
			_, _ = buf.WriteString(", ")
		}

		_, _ = buf.WriteString(ctx.Vars.CreatePlaceholder(fmt.Sprintf("value%d", i), v))
	}
	return buf.String(), nil
}

func (c InValues) HasValues() bool {
	return len(c) > 0
}

type DynamicFilters struct {
	op      string
	clauses []dynamicFiltersClause
}

func DynamicFiltersAnd() *DynamicFilters {
	return &DynamicFilters{
		op: "AND",
	}
}

func DynamicFiltersOr() *DynamicFilters {
	return &DynamicFilters{
		op: "OR",
	}
}

func (f *DynamicFilters) HasValues() bool {
	return f != nil && len(f.clauses) > 0
}

func (f *DynamicFilters) SQLText(ctx *SQLTexterContext) (string, error) {
	var buf strings.Builder
	_, _ = buf.WriteString("(")
	f.appendClause(&buf, ctx.Vars)
	_, _ = buf.WriteString(")")
	return buf.String(), nil
}

func (f *DynamicFilters) AddFilter(clause *DynamicFilters) {
	f.clauses = append(f.clauses, clause)
}

func (f *DynamicFilters) AddNot(clause *DynamicFilters) {
	f.clauses = append(f.clauses, &notClause{clause: clause})
}

func (f *DynamicFilters) AddIn(columnName string, values []any) {
	f.clauses = append(f.clauses, &inClause{
		col:    columnName,
		values: values,
	})
}

func (f *DynamicFilters) AddEq(columnName string, val any) {
	f.clauses = append(f.clauses, &binOpClause{
		col: columnName,
		op:  "=",
		v:   val,
	})
}

func (f *DynamicFilters) AddNeq(columnName string, val any) {
	f.clauses = append(f.clauses, &binOpClause{
		col: columnName,
		op:  "<>",
		v:   val,
	})
}

func (f *DynamicFilters) AddGt(columnName string, val any) {
	f.clauses = append(f.clauses, &binOpClause{
		col: columnName,
		op:  ">",
		v:   val,
	})
}

func (f *DynamicFilters) AddGte(columnName string, val any) {
	f.clauses = append(f.clauses, &binOpClause{
		col: columnName,
		op:  ">=",
		v:   val,
	})
}

func (f *DynamicFilters) AddLt(columnName string, val any) {
	f.clauses = append(f.clauses, &binOpClause{
		col: columnName,
		op:  "<",
		v:   val,
	})
}

func (f *DynamicFilters) AddLte(columnName string, val any) {
	f.clauses = append(f.clauses, &binOpClause{
		col: columnName,
		op:  "<=",
		v:   val,
	})
}

func (f *DynamicFilters) AddIsNull(columnName string) {
	f.clauses = append(f.clauses, &suffixClause{
		col: columnName,
		op:  "IS NULL",
	})
}

func (f *DynamicFilters) AddIsNotNull(columnName string) {
	f.clauses = append(f.clauses, &suffixClause{
		col: columnName,
		op:  "IS NOT NULL",
	})
}

func (f *DynamicFilters) appendClause(w *strings.Builder, vars QueryVariablesScope) {

	for i, c := range f.clauses {
		if i > 0 {
			_, _ = w.WriteString(" ")
			_, _ = w.WriteString(f.op)
			_, _ = w.WriteString(" ")
		}
		_, _ = w.WriteString("(")
		c.appendClause(w, vars)
		_, _ = w.WriteString(")")
	}
}

type dynamicFiltersClause interface {
	appendClause(w *strings.Builder, vars QueryVariablesScope)
}

type inClause struct {
	col    string
	values []any
}

func (c *inClause) appendClause(w *strings.Builder, vars QueryVariablesScope) {
	_, _ = w.WriteString(c.col)
	_, _ = w.WriteString(" IN (")
	for i, v := range c.values {
		if i > 0 {
			_, _ = w.WriteString(", ")
		}

		_, _ = w.WriteString(vars.CreatePlaceholder(fmt.Sprintf("value%d", i), v))
	}
	_, _ = w.WriteString(")")
}

type binOpClause struct {
	col string
	op  string
	v   any
}

func (c *binOpClause) appendClause(w *strings.Builder, vars QueryVariablesScope) {
	_, _ = w.WriteString(c.col)
	_, _ = w.WriteString(" ")
	_, _ = w.WriteString(c.op)
	_, _ = w.WriteString(" ")
	_, _ = w.WriteString(vars.CreatePlaceholder("value", c.v))
}

type notClause struct {
	clause dynamicFiltersClause
}

func (c *notClause) appendClause(w *strings.Builder, vars QueryVariablesScope) {
	_, _ = w.WriteString("NOT (")
	c.clause.appendClause(w, vars)
	_, _ = w.WriteString(")")
}

type suffixClause struct {
	col string
	op  string
}

func (c *suffixClause) appendClause(w *strings.Builder, vars QueryVariablesScope) {
	_, _ = w.WriteString(c.col)
	_, _ = w.WriteString(" ")
	_, _ = w.WriteString(c.op)

}
