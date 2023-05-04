package checkers

import (
	"go/ast"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"
)

func init() {
	var info = linter.CheckerInfo{
		Name:    "errPosition",
		Tags:    []string{"diagnostic", "experimental"},
		Summary: "Detect errors not returned as the last argument",
		Before: `
func fileExists(path string) (error, bool) {
	var ok
	var err

	// ...

	return err, ok
}
`,
		After: `
func fileExists(path string) (bool, error) {
	var ok
	var err

	// ...

	return err, ok
}
`,
	}

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		return astwalk.WalkerForFuncDecl(&errPosition{ctx: ctx}), nil
	})
}

type errPosition struct {
	ctx *linter.CheckerContext
	astwalk.WalkHandler
}

func (m *errPosition) VisitFuncDecl(fnc *ast.FuncDecl) {
	var errCount int
	var lastErr bool

	if fnc.Type == nil {
		return
	}

	if fnc.Type.Results == nil {
		return
	}

	for idx, field := range fnc.Type.Results.List {
		id, ok := field.Type.(*ast.Ident)
		if ok {
			typeName := m.ctx.TypesInfo.ObjectOf(id).Type().String()
			if typeName == `error` {
				errCount++
				if idx == len(fnc.Type.Results.List)-1 {
					lastErr = true
				}
			}
		}
	}

	if errCount > 0 && !lastErr {
		obj := m.ctx.TypesInfo.ObjectOf(fnc.Name)
		m.ctx.Warn(fnc, "The last return parameter is not an error (%v/%v)", obj.Pkg().Name(), obj.Name())
	}
}
