package parser

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/partite-ai/modelg/cmd/modelg/parser/generated"
)

func ParseQueries(input io.Reader) (map[string]QueryNode, error) {
	lexer := generated.NewSQLQueryLexer(antlr.NewIoStream(io.MultiReader(input, strings.NewReader("\n"))))
	lexer.RemoveErrorListeners()
	var el captureErrorListener
	lexer.AddErrorListener(&el)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := generated.NewSQLQueryParser(stream)
	p.AddErrorListener(&el)
	tree := p.QuerySet()
	if el.err != nil {
		return nil, fmt.Errorf("failed to parse queries: %w", el.err)
	}
	queries := make(map[string]QueryNode)

	err := tree.Accept(newParserVisitor(queries))

	if err != nil {
		return nil, fmt.Errorf("failed to parse queries: %w", err.(error))
	}

	return queries, nil
}

type parserVisitor struct {
	generated.SQLQueryParserVisitor
	queries map[string]QueryNode
}

func newParserVisitor(queries map[string]QueryNode) *parserVisitor {
	return newSQLQueryParserVisitor(func(base generated.SQLQueryParserVisitor) *parserVisitor {
		return &parserVisitor{
			SQLQueryParserVisitor: base,
			queries:               queries,
		}
	})
}

func (v *parserVisitor) VisitQuery(ctx *generated.QueryContext) any {
	queryVisitor := newQueryVisitor()
	result := queryVisitor.Visit(ctx)
	if result != nil {
		return result
	}

	v.queries[ctx.GetQueryName().GetText()] = mergeNode(queryVisitor.nodes)
	return nil
}

type queryVisitor struct {
	generated.SQLQueryParserVisitor
	nodes []QueryNode
}

func newQueryVisitor() *queryVisitor {
	return newSQLQueryParserVisitor(func(base generated.SQLQueryParserVisitor) *queryVisitor {
		return &queryVisitor{
			SQLQueryParserVisitor: base,
			nodes:                 make([]QueryNode, 0),
		}
	})
}

func (v *queryVisitor) VisitSqlText(ctx *generated.SqlTextContext) any {
	tokens := ctx.GetTextFragments()
	var sb strings.Builder
	for _, token := range tokens {
		sb.WriteString(token.GetText())
		sb.WriteString(" ")
	}
	v.nodes = append(v.nodes, textNode(sb.String()))
	return nil
}

func (v *queryVisitor) VisitExpr(ctx *generated.ExprContext) any {
	tokenPath := ctx.GetPath()
	path := make([]string, 0, len(tokenPath))
	for _, token := range tokenPath {
		path = append(path, token.GetText())
	}
	isLiteral := false
	literalMode := ""
	if ctx.GetLiteralFlag() != nil {
		isLiteral = true

		if lm := ctx.GetLiteralMode(); lm != nil {
			literalMode = lm.GetText()
		}
	}

	v.nodes = append(v.nodes, &exprNode{
		path:        path,
		isLiteral:   isLiteral,
		literalMode: literalMode,
	})

	return nil
}

func (v *queryVisitor) VisitWhenDirective(ctx *generated.WhenDirectiveContext) any {
	if ctx.GetWhenCondition().GetLiteralFlag() != nil {
		return fmt.Errorf("the expression used in a when directive must not be a literal expression")
	}

	var when whenNode
	pathTokens := ctx.GetWhenCondition().GetPath()
	path := make([]string, 0, len(pathTokens))
	for _, token := range pathTokens {
		path = append(path, token.GetText())
	}
	when.condition = path

	bodyVisitor := newQueryVisitor()
	if err := bodyVisitor.Visit(ctx.GetBody()); err != nil {
		return err
	}

	when.body = mergeNode(bodyVisitor.nodes)

	var joinWhens []QueryNode
	for _, joinWhenCtxt := range ctx.GetJoinWhens() {
		suffixVisitor := newQueryVisitor()
		if err := suffixVisitor.Visit(joinWhenCtxt); err != nil {
			return err
		}
		joinWhens = append(joinWhens, suffixVisitor.nodes...)
	}
	when.joinWhen = joinWhens

	v.nodes = append(v.nodes, &when)
	return nil
}

func (v *queryVisitor) VisitChompWhenDirective(ctx *generated.ChompWhenDirectiveContext) any {
	if ctx.GetWhenCondition().GetLiteralFlag() != nil {
		return fmt.Errorf("the expression used in a chomp when directive must not be a literal expression")
	}

	var when chompWhenNode

	pathTokens := ctx.GetWhenCondition().GetPath()
	path := make([]string, 0, len(pathTokens))
	for _, token := range pathTokens {
		path = append(path, token.GetText())
	}
	when.condition = path

	prefixVisitor := newQueryVisitor()
	if err := prefixVisitor.Visit(ctx.GetPrefix()); err != nil {
		return err
	}
	when.prefix = mergeNode(prefixVisitor.nodes)

	bodyVisitor := newQueryVisitor()
	if err := bodyVisitor.Visit(ctx.GetBody()); err != nil {
		return err
	}

	when.body = mergeNode(bodyVisitor.nodes)

	var joinWhens []QueryNode
	for _, joinWhenCtxt := range ctx.GetJoinWhens() {
		suffixVisitor := newQueryVisitor()
		if err := suffixVisitor.Visit(joinWhenCtxt); err != nil {
			return err
		}
		joinWhens = append(joinWhens, suffixVisitor.nodes...)
	}
	when.joinWhen = joinWhens

	v.nodes = append(v.nodes, &when)
	return nil
}

func (v *queryVisitor) VisitJoinWhen(ctx *generated.JoinWhenContext) any {
	if ctx.GetWhenCondition().GetLiteralFlag() != nil {
		return fmt.Errorf("the expression used in a chomp when directive must not be a literal expression")
	}

	var when joinWhenNode

	prefixVisitor := newQueryVisitor()
	if err := prefixVisitor.Visit(ctx.GetPrefix()); err != nil {
		return err
	}
	when.prefix = mergeNode(prefixVisitor.nodes)

	pathTokens := ctx.GetWhenCondition().GetPath()
	path := make([]string, 0, len(pathTokens))
	for _, token := range pathTokens {
		path = append(path, token.GetText())
	}
	when.condition = path

	bodyVisitor := newQueryVisitor()
	if err := bodyVisitor.Visit(ctx.GetBody()); err != nil {
		return err
	}

	when.body = mergeNode(bodyVisitor.nodes)

	v.nodes = append(v.nodes, &when)
	return nil
}

type captureErrorListener struct {
	antlr.DefaultErrorListener
	err error
}

func (l *captureErrorListener) SyntaxError(_ antlr.Recognizer, _ any, line, column int, msg string, _ antlr.RecognitionException) {
	l.err = errors.Join(l.err, fmt.Errorf("syntax error at line %d, column %d: %s", line, column, msg))
}
