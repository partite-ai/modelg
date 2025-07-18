package parser

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/partite-ai/modelg/cmd/modelg/parser/generated"
)

type sqlQueryParserVisitor struct {
	delegate generated.SQLQueryParserVisitor
}

func newSQLQueryParserVisitor[T generated.SQLQueryParserVisitor](make func(generated.SQLQueryParserVisitor) T) T {
	base := &sqlQueryParserVisitor{}
	instance := make(base)
	base.delegate = instance
	return instance
}

func (v *sqlQueryParserVisitor) VisitQuerySet(ctx *generated.QuerySetContext) any {
	return v.VisitChildren(ctx)
}

func (v *sqlQueryParserVisitor) VisitQuery(ctx *generated.QueryContext) any {
	return v.VisitChildren(ctx)
}

func (v *sqlQueryParserVisitor) VisitSqlChunk(ctx *generated.SqlChunkContext) any {
	return v.VisitChildren(ctx)
}

func (v *sqlQueryParserVisitor) VisitSqlLine(ctx *generated.SqlLineContext) any {
	return v.VisitChildren(ctx)
}

func (v *sqlQueryParserVisitor) VisitSqlFragment(ctx *generated.SqlFragmentContext) any {
	return v.VisitChildren(ctx)
}

func (v *sqlQueryParserVisitor) VisitSqlText(ctx *generated.SqlTextContext) any {
	return v.VisitChildren(ctx)
}

func (v *sqlQueryParserVisitor) VisitWhenDirective(ctx *generated.WhenDirectiveContext) any {
	return v.VisitChildren(ctx)
}

func (v *sqlQueryParserVisitor) VisitChompWhenDirective(ctx *generated.ChompWhenDirectiveContext) any {
	return v.VisitChildren(ctx)
}

func (v *sqlQueryParserVisitor) VisitJoinWhen(ctx *generated.JoinWhenContext) any {
	return v.VisitChildren(ctx)
}

func (v *sqlQueryParserVisitor) VisitExpr(ctx *generated.ExprContext) any {
	return v.VisitChildren(ctx)
}

func (v *sqlQueryParserVisitor) VisitChildren(node antlr.RuleNode) any {
	for _, child := range node.GetChildren() {
		child.(antlr.ParseTree).Accept(v.delegate)
	}
	return nil
}

func (v *sqlQueryParserVisitor) Visit(tree antlr.ParseTree) any {
	return tree.Accept(v.delegate)
}
func (v *sqlQueryParserVisitor) VisitTerminal(node antlr.TerminalNode) any { return nil }
func (v *sqlQueryParserVisitor) VisitErrorNode(node antlr.ErrorNode) any   { return nil }
