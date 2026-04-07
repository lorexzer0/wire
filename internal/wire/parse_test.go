// Copyright 2018 The Wire Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package wire

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestIsSimpleTypeExpr(t *testing.T) {
	tests := []struct {
		name string
		expr string
		want bool
	}{
		{"new call", "new(Foo)", true},
		{"nil conversion", "(*Foo)(nil)", true},
		{"string literal", `"foo"`, false},
		{"variable", "x", false},
		{"function call", "getType()", false},
		{"composite literal", "Foo{}", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fset := token.NewFileSet()
			expr, err := parser.ParseExpr(tt.expr)
			if err != nil {
				// Some expressions like Foo{} won't parse as standalone
				// expressions. Wrap in a function call to parse.
				src := "package p\nvar _ = f(" + tt.expr + ")\n"
				f, err2 := parser.ParseFile(fset, "", src, 0)
				if err2 != nil {
					t.Fatalf("failed to parse %q: %v (and %v)", tt.expr, err, err2)
				}
				call := f.Decls[0].(*ast.GenDecl).Specs[0].(*ast.ValueSpec).Values[0].(*ast.CallExpr)
				expr = call.Args[0]
			}
			if got := isSimpleTypeExpr(expr); got != tt.want {
				t.Errorf("isSimpleTypeExpr(%q) = %v, want %v", tt.expr, got, tt.want)
			}
		})
	}
}
