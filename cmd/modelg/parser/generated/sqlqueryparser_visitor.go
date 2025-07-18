// Code generated from SQLQueryParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package generated // SQLQueryParser
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by SQLQueryParser.
type SQLQueryParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by SQLQueryParser#querySet.
	VisitQuerySet(ctx *QuerySetContext) interface{}

	// Visit a parse tree produced by SQLQueryParser#query.
	VisitQuery(ctx *QueryContext) interface{}

	// Visit a parse tree produced by SQLQueryParser#sqlChunk.
	VisitSqlChunk(ctx *SqlChunkContext) interface{}

	// Visit a parse tree produced by SQLQueryParser#sqlLine.
	VisitSqlLine(ctx *SqlLineContext) interface{}

	// Visit a parse tree produced by SQLQueryParser#sqlFragment.
	VisitSqlFragment(ctx *SqlFragmentContext) interface{}

	// Visit a parse tree produced by SQLQueryParser#sqlText.
	VisitSqlText(ctx *SqlTextContext) interface{}

	// Visit a parse tree produced by SQLQueryParser#whenDirective.
	VisitWhenDirective(ctx *WhenDirectiveContext) interface{}

	// Visit a parse tree produced by SQLQueryParser#chompWhenDirective.
	VisitChompWhenDirective(ctx *ChompWhenDirectiveContext) interface{}

	// Visit a parse tree produced by SQLQueryParser#joinWhen.
	VisitJoinWhen(ctx *JoinWhenContext) interface{}

	// Visit a parse tree produced by SQLQueryParser#expr.
	VisitExpr(ctx *ExprContext) interface{}
}
