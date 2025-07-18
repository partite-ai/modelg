package parser

import (
	"fmt"
	"strings"
)

type QueryEnvironment interface {
	ParamCodeExpr(path []string) (string, error)
}

type QueryNode interface {
	Apply(env QueryEnvironment) ([]string, error)
}

type mergeNode []QueryNode

func (j mergeNode) Apply(env QueryEnvironment) ([]string, error) {
	var allLines []string

	for _, node := range j {
		lines, err := node.Apply(env)
		if err != nil {
			return nil, err
		}
		allLines = append(allLines, lines...)
	}

	return allLines, nil
}

type textNode string

func (t textNode) Apply(env QueryEnvironment) ([]string, error) {
	return []string{fmt.Sprintf("sqlText__ += %q", t)}, nil
}

type exprNode struct {
	path        []string
	isLiteral   bool
	literalMode string
}

func (n *exprNode) Apply(env QueryEnvironment) ([]string, error) {
	if n.isLiteral {
		paramCodeExpr, err := env.ParamCodeExpr(n.path)
		if err != nil {
			return nil, err
		}

		return []string{
			fmt.Sprintf("sqlText__ += %s.SQLText(%q) + \" \"", paramCodeExpr, n.literalMode),
		}, nil
	}

	paramCodeExpr, err := env.ParamCodeExpr(n.path)
	if err != nil {
		return nil, err
	}

	return []string{
		fmt.Sprintf(`sqlText__ += vars__.CreatePlaceholder(%q, %s)`, strings.Join(n.path, "."), paramCodeExpr),
		`sqlText__ += " "`,
	}, nil
}

type whenNode struct {
	condition []string
	body      QueryNode
	joinWhen  []QueryNode
}

func (w *whenNode) Apply(env QueryEnvironment) ([]string, error) {
	conditionExpr, err := env.ParamCodeExpr(w.condition)
	if err != nil {
		return nil, fmt.Errorf("failed to get condition expression: %w", err)
	}
	body, err := w.body.Apply(env)
	if err != nil {
		return nil, fmt.Errorf("failed to apply body: %w", err)
	}
	var joinWhen []string
	for _, j := range w.joinWhen {
		jw, err := j.Apply(env)
		if err != nil {
			return nil, fmt.Errorf("failed to apply joinWhen: %w", err)
		}
		joinWhen = append(joinWhen, jw...)
	}
	ret := make([]string, 0, len(body)+len(joinWhen)+7)
	if len(joinWhen) > 0 {
		ret = append(ret, "func() {")
		ret = append(ret, "needPrefix__ := false")
		ret = append(ret, "beforeWrite__ := func() { }")
		ret = append(ret, fmt.Sprintf("if %s {", conditionExpr))
		ret = append(ret, "needPrefix__ = true")
		ret = append(ret, body...)
		ret = append(ret, "}")
		ret = append(ret, joinWhen...)
		ret = append(ret, "}()")
	} else {
		ret = append(ret, fmt.Sprintf("if %s {", conditionExpr))
		ret = append(ret, body...)
		ret = append(ret, "}")
	}
	return ret, nil
}

type joinWhenNode struct {
	prefix    QueryNode
	condition []string
	body      QueryNode
}

func (w *joinWhenNode) Apply(env QueryEnvironment) ([]string, error) {
	prefix, err := w.prefix.Apply(env)
	if err != nil {
		return nil, fmt.Errorf("failed to apply prefix: %w", err)
	}
	conditionExpr, err := env.ParamCodeExpr(w.condition)
	if err != nil {
		return nil, fmt.Errorf("failed to get condition expression: %w", err)
	}
	body, err := w.body.Apply(env)
	if err != nil {
		return nil, fmt.Errorf("failed to apply body: %w", err)
	}
	ret := make([]string, 0, len(prefix)+len(body)+6)
	ret = append(ret, fmt.Sprintf("if %s {", conditionExpr))
	ret = append(ret, "beforeWrite__()")
	ret = append(ret, "if needPrefix__ {")
	ret = append(ret, prefix...)
	ret = append(ret, "}")
	ret = append(ret, "needPrefix__ = true")
	ret = append(ret, body...)
	ret = append(ret, "}")
	return ret, nil
}

type chompWhenNode struct {
	prefix    QueryNode
	condition []string
	body      QueryNode
	joinWhen  []QueryNode
}

func (w *chompWhenNode) Apply(env QueryEnvironment) ([]string, error) {
	prefix, err := w.prefix.Apply(env)
	if err != nil {
		return nil, fmt.Errorf("failed to apply prefix: %w", err)
	}
	conditionExpr, err := env.ParamCodeExpr(w.condition)
	if err != nil {
		return nil, fmt.Errorf("failed to get condition expression: %w", err)
	}
	body, err := w.body.Apply(env)
	if err != nil {
		return nil, fmt.Errorf("failed to apply body: %w", err)
	}

	var ret []string
	if len(w.joinWhen) > 0 {
		ret = append(ret, "func() {")
		ret = append(ret, "  needPrefix__ := false")
		ret = append(ret, "  wroteHead__ := false")
		ret = append(ret, "  beforeWrite__ := func() {")
		ret = append(ret, "    if !wroteHead__ {")
		ret = append(ret, "      wroteHead__ = true")
		ret = append(ret, prefix...)
		ret = append(ret, "    }")
		ret = append(ret, "  }")
		ret = append(ret, fmt.Sprintf("  if %s {", conditionExpr))
		ret = append(ret, "    beforeWrite__()")
		ret = append(ret, "    needPrefix__ = true")
		ret = append(ret, body...)
		ret = append(ret, "  }")
		for _, j := range w.joinWhen {
			jw, err := j.Apply(env)
			if err != nil {
				return nil, fmt.Errorf("failed to apply joinWhen: %w", err)
			}
			ret = append(ret, jw...)
		}
		ret = append(ret, "}()")
	} else {
		ret = append(ret, fmt.Sprintf("if %s {", conditionExpr))
		ret = append(ret, prefix...)
		ret = append(ret, body...)
		ret = append(ret, "}")
	}
	return ret, nil
}
