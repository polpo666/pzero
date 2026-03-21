package genapi

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"
	"strings"
	"testing"
)

func TestRewriteImportedModule(t *testing.T) {
	content := `package helloworld

import (
	"github.com/polpo666/pzero/examples/simpleapi/plugins/helloworld/internal/logic/helloworld"
	"github.com/polpo666/pzero/examples/simpleapi/plugins/helloworld/internal/svc"
	"github.com/polpo666/pzero/examples/simpleapi/plugins/helloworld/internal/types"
)
`

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "handler.go", content, parser.ParseComments)
	if err != nil {
		t.Fatal(err)
	}

	rewriteImportedModule(f, fset,
		"github.com/polpo666/pzero/examples/simpleapi/plugins/helloworld",
		"github.com/jzero-io/examples/simpleapi/plugins/helloworld",
	)

	var imports []string
	for _, spec := range f.Imports {
		imports = append(imports, spec.Path.Value)
	}

	joined := strings.Join(imports, "\n")
	if strings.Contains(joined, "github.com/polpo666/pzero/examples/simpleapi/plugins/helloworld") {
		t.Fatalf("imports were not rewritten: %s", joined)
	}
	if !strings.Contains(joined, "github.com/jzero-io/examples/simpleapi/plugins/helloworld/internal/logic/helloworld") {
		t.Fatalf("logic import not rewritten: %s", joined)
	}
	if !strings.Contains(joined, "github.com/jzero-io/examples/simpleapi/plugins/helloworld/internal/svc") {
		t.Fatalf("svc import not rewritten: %s", joined)
	}
	if !astutilUsesImport(f, "github.com/jzero-io/examples/simpleapi/plugins/helloworld/internal/types") {
		t.Fatalf("types import not rewritten: %s", joined)
	}
}

func astutilUsesImport(f *ast.File, importPath string) bool {
	for _, spec := range f.Imports {
		if spec.Path != nil && spec.Path.Value == strconv.Quote(importPath) {
			return true
		}
	}
	return false
}
