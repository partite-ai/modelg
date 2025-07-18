// Code generated from SQLQueryParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package generated // SQLQueryParser
import "github.com/antlr4-go/antlr/v4"

// BaseSQLQueryParserListener is a complete listener for a parse tree produced by SQLQueryParser.
type BaseSQLQueryParserListener struct{}

var _ SQLQueryParserListener = &BaseSQLQueryParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSQLQueryParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSQLQueryParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSQLQueryParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSQLQueryParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterQuerySet is called when production querySet is entered.
func (s *BaseSQLQueryParserListener) EnterQuerySet(ctx *QuerySetContext) {}

// ExitQuerySet is called when production querySet is exited.
func (s *BaseSQLQueryParserListener) ExitQuerySet(ctx *QuerySetContext) {}

// EnterQuery is called when production query is entered.
func (s *BaseSQLQueryParserListener) EnterQuery(ctx *QueryContext) {}

// ExitQuery is called when production query is exited.
func (s *BaseSQLQueryParserListener) ExitQuery(ctx *QueryContext) {}

// EnterSqlChunk is called when production sqlChunk is entered.
func (s *BaseSQLQueryParserListener) EnterSqlChunk(ctx *SqlChunkContext) {}

// ExitSqlChunk is called when production sqlChunk is exited.
func (s *BaseSQLQueryParserListener) ExitSqlChunk(ctx *SqlChunkContext) {}

// EnterSqlLine is called when production sqlLine is entered.
func (s *BaseSQLQueryParserListener) EnterSqlLine(ctx *SqlLineContext) {}

// ExitSqlLine is called when production sqlLine is exited.
func (s *BaseSQLQueryParserListener) ExitSqlLine(ctx *SqlLineContext) {}

// EnterSqlFragment is called when production sqlFragment is entered.
func (s *BaseSQLQueryParserListener) EnterSqlFragment(ctx *SqlFragmentContext) {}

// ExitSqlFragment is called when production sqlFragment is exited.
func (s *BaseSQLQueryParserListener) ExitSqlFragment(ctx *SqlFragmentContext) {}

// EnterSqlText is called when production sqlText is entered.
func (s *BaseSQLQueryParserListener) EnterSqlText(ctx *SqlTextContext) {}

// ExitSqlText is called when production sqlText is exited.
func (s *BaseSQLQueryParserListener) ExitSqlText(ctx *SqlTextContext) {}

// EnterWhenDirective is called when production whenDirective is entered.
func (s *BaseSQLQueryParserListener) EnterWhenDirective(ctx *WhenDirectiveContext) {}

// ExitWhenDirective is called when production whenDirective is exited.
func (s *BaseSQLQueryParserListener) ExitWhenDirective(ctx *WhenDirectiveContext) {}

// EnterChompWhenDirective is called when production chompWhenDirective is entered.
func (s *BaseSQLQueryParserListener) EnterChompWhenDirective(ctx *ChompWhenDirectiveContext) {}

// ExitChompWhenDirective is called when production chompWhenDirective is exited.
func (s *BaseSQLQueryParserListener) ExitChompWhenDirective(ctx *ChompWhenDirectiveContext) {}

// EnterJoinWhen is called when production joinWhen is entered.
func (s *BaseSQLQueryParserListener) EnterJoinWhen(ctx *JoinWhenContext) {}

// ExitJoinWhen is called when production joinWhen is exited.
func (s *BaseSQLQueryParserListener) ExitJoinWhen(ctx *JoinWhenContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseSQLQueryParserListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseSQLQueryParserListener) ExitExpr(ctx *ExprContext) {}
