// Copyright 2018 The go-python Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ast

import (
	"testing"

	"github.com/go-python/gpython/py"
)

func TestDump(t *testing.T) {
	for _, test := range []struct {
		in  Ast
		out string
	}{
		{nil, `<nil>`},
		{&Pass{}, `Pass()`},
		{&Constant{Value: py.String("potato")}, `Constant(value='potato', kind=None)`},
		{&Constant{Value: py.String("potato")}, `Constant(value='potato', kind=None)`},
		{&Constant{Value: py.Bytes("potato")}, `Constant(value=b'potato', kind=None)`},
		{&BinOp{Left: &Constant{Value: py.String("one")}, Op: Add, Right: &Constant{Value: py.String("two")}},
			`BinOp(left=Constant(value='one', kind=None), op=Add(), right=Constant(value='two', kind=None))`},
		{&Module{}, `Module(body=[], type_ignores=[])`},
		{&Module{Body: []Stmt{&Pass{}}}, `Module(body=[Pass()], type_ignores=[])`},
		{&Module{Body: []Stmt{&ExprStmt{Value: &Tuple{}}}}, `Module(body=[Expr(value=Tuple(elts=[], ctx=UnknownExprContext(0)))], type_ignores=[])`},
		{&Constant{Value: py.True}, `Constant(value=True, kind=None)`},
		{&Name{Id: Identifier("hello"), Ctx: Load}, `Name(id='hello', ctx=Load())`},
		{&ListComp{Elt: &Constant{Value: py.String("potato")}, Generators: []Comprehension{{
			Target: &Name{Id: Identifier("hello"), Ctx: Load},
		}}}, `ListComp(elt=Constant(value='potato', kind=None), generators=[comprehension(target=Name(id='hello', ctx=Load()), iter=None, ifs=[], is_async=0)])`},
	} {
		out := Dump(test.in)
		if out != test.out {
			t.Errorf("Dump(%#v) got\n%q\nexpected\n%q", test.in, out, test.out)
		}
	}
}
