// Code generated from SQLQueryParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package generated // SQLQueryParser
import "github.com/antlr4-go/antlr/v4"

type BaseSQLQueryParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSQLQueryParserVisitor) VisitQuerySet(ctx *QuerySetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLQueryParserVisitor) VisitQuery(ctx *QueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLQueryParserVisitor) VisitSqlChunk(ctx *SqlChunkContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLQueryParserVisitor) VisitSqlLine(ctx *SqlLineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLQueryParserVisitor) VisitSqlFragment(ctx *SqlFragmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLQueryParserVisitor) VisitSqlText(ctx *SqlTextContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLQueryParserVisitor) VisitWhenDirective(ctx *WhenDirectiveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLQueryParserVisitor) VisitChompWhenDirective(ctx *ChompWhenDirectiveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLQueryParserVisitor) VisitJoinWhen(ctx *JoinWhenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLQueryParserVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}
