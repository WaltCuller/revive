package rule

import (
	"go/ast"

	"github.com/mgechev/revive/lint"
)

// MapCheckRule ...
type MapCheckRule struct{}

// Apply ...
func (m *MapCheckRule) Apply(file *lint.File, arguments lint.Arguments) []lint.Failure {
	var failures []lint.Failure

	onFailure := func(failure lint.Failure) {
		failures = append(failures, failure)
	}

	w := &lintMapCheck{onFailure: onFailure}
	ast.Walk(w, file.AST)
	return failures
}

// Name ...
func (m *MapCheckRule) Name() string {
	return "map-check"
}

type lintMapCheck struct {
	file      *lint.File
	onFailure func(lint.Failure)
}

// Visit ...
func (w *lintMapCheck) Visit(node ast.Node) ast.Visitor {

	return w
}
