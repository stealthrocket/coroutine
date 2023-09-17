package compiler

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"strconv"
)

func typeExpr(typ types.Type) ast.Expr {
	switch t := typ.(type) {
	case *types.Basic:
		return ast.NewIdent(t.String())
	case *types.Slice:
		return &ast.ArrayType{Elt: typeExpr(t.Elem())}
	case *types.Array:
		return &ast.ArrayType{
			Len: &ast.BasicLit{Kind: token.INT, Value: strconv.FormatInt(t.Len(), 10)},
			Elt: typeExpr(t.Elem()),
		}
	case *types.Map:
		return &ast.MapType{
			Key:   typeExpr(t.Key()),
			Value: typeExpr(t.Elem()),
		}
	case *types.Struct:
		fields := make([]*ast.Field, t.NumFields())
		for i := range fields {
			f := t.Field(i)
			fields[i] = &ast.Field{Type: typeExpr(f.Type())}
			if !f.Anonymous() {
				fields[i].Names = []*ast.Ident{ast.NewIdent(f.Name())}
			}
			if tag := t.Tag(i); tag != "" {
				panic("not implemented: struct tags")
			}
		}
		return &ast.StructType{Fields: &ast.FieldList{List: fields}}
	case *types.Pointer:
		return &ast.StarExpr{X: typeExpr(t.Elem())}
	case *types.Interface:
		if t.Empty() {
			return ast.NewIdent("any")
		}
	}
	panic(fmt.Sprintf("not implemented: %T", typ))
}