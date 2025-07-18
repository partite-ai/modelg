// Code generated from SQLQueryParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package generated // SQLQueryParser
import "github.com/antlr4-go/antlr/v4"

// SQLQueryParserListener is a complete listener for a parse tree produced by SQLQueryParser.
type SQLQueryParserListener interface {
	antlr.ParseTreeListener

	// EnterQuerySet is called when entering the querySet production.
	EnterQuerySet(c *QuerySetContext)

	// EnterQuery is called when entering the query production.
	EnterQuery(c *QueryContext)

	// EnterSqlChunk is called when entering the sqlChunk production.
	EnterSqlChunk(c *SqlChunkContext)

	// EnterSqlLine is called when entering the sqlLine production.
	EnterSqlLine(c *SqlLineContext)

	// EnterSqlFragment is called when entering the sqlFragment production.
	EnterSqlFragment(c *SqlFragmentContext)

	// EnterSqlText is called when entering the sqlText production.
	EnterSqlText(c *SqlTextContext)

	// EnterWhenDirective is called when entering the whenDirective production.
	EnterWhenDirective(c *WhenDirectiveContext)

	// EnterChompWhenDirective is called when entering the chompWhenDirective production.
	EnterChompWhenDirective(c *ChompWhenDirectiveContext)

	// EnterJoinWhen is called when entering the joinWhen production.
	EnterJoinWhen(c *JoinWhenContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// ExitQuerySet is called when exiting the querySet production.
	ExitQuerySet(c *QuerySetContext)

	// ExitQuery is called when exiting the query production.
	ExitQuery(c *QueryContext)

	// ExitSqlChunk is called when exiting the sqlChunk production.
	ExitSqlChunk(c *SqlChunkContext)

	// ExitSqlLine is called when exiting the sqlLine production.
	ExitSqlLine(c *SqlLineContext)

	// ExitSqlFragment is called when exiting the sqlFragment production.
	ExitSqlFragment(c *SqlFragmentContext)

	// ExitSqlText is called when exiting the sqlText production.
	ExitSqlText(c *SqlTextContext)

	// ExitWhenDirective is called when exiting the whenDirective production.
	ExitWhenDirective(c *WhenDirectiveContext)

	// ExitChompWhenDirective is called when exiting the chompWhenDirective production.
	ExitChompWhenDirective(c *ChompWhenDirectiveContext)

	// ExitJoinWhen is called when exiting the joinWhen production.
	ExitJoinWhen(c *JoinWhenContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)
}
