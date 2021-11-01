package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io"
	"os"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/gobuffalo/packr/v2"
)

var pb2 *packr.Box = packr.New("pb2", "./test")
var (
	_sd_           = "_serviceDesc"
	grpc_msg       = map[string]string{}
	grpc_packename = "main"
)

func readpb(pbfilepath string) map[string]string {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, pbfilepath, nil, parser.ParseComments)
	if err != nil {
		fmt.Println("readpb", err)
	}
	grpc_packename = node.Name.Name
	ast.Inspect(node, func(n ast.Node) bool {
		// Find Return Statements

		ret, ok := n.(*ast.GenDecl)

		if ok {
			for _, ss := range ret.Specs {

				if vs, ok := ss.(*ast.ValueSpec); ok && len(vs.Names) > 0 {
					if strings.Contains(vs.Names[0].Name, "_serviceDesc") {
						_sd_ = vs.Names[0].Name
					}

				}
				if sp, ok := ss.(*ast.TypeSpec); ok {

					if unicode.IsUpper(rune(sp.Name.Name[0])) {
						i := 0
						if ty, ok := sp.Type.(*ast.StructType); ok {
							for _, fd := range ty.Fields.List {
								if len(fd.Names) > 0 {
									if fd.Names[0].Name == "state" || fd.Names[0].Name == "sizeCache" || fd.Names[0].Name == "unknownFields" {
										i++
									}
								}

							}
						}
						if i >= 3 {
							// fmt.Printf("%v found on line %d:\n", sp.Name.Name, fset.Position(ret.Pos()).Line)
							grpc_msg[fmt.Sprintf("new%v", sp.Name.Name)] = sp.Name.Name
						}

					}
				}
			}

			// printer.Fprint(os.Stdout, fset, ret)
			return true
		}
		return true
	})
	return grpc_msg
}

func mkpb2(f func(ast.Node)) ast.Node {
	fset := token.NewFileSet()
	src, _ := pb2.FindString("pb.interface.go")
	node, err := parser.ParseFile(fset, "", src, parser.ParseComments)
	if err != nil {
		fmt.Println("读取错误", err)
	}
	ast.Inspect(node, func(n ast.Node) bool {
		// Find Return Statements
		ret, ok := n.(*ast.ValueSpec)
		// fmt.Printf("%T\n", n)
		if ok {

			for _, n := range ret.Names {

				if n.Name == "_sd_" {
					ret.Values = []ast.Expr{ast.NewIdent(_sd_)}
				}
				if n.Name == "grpc_msg" {
					// fmt.Println(n.Name, n.Obj.Data, n.Obj.Decl, n.Obj.Type)
					f(ret)

				}

			}
		}
		return true
	})
	node.Name = ast.NewIdent(grpc_packename)
	return node
}

func analyExpr(expr ast.Expr) {
	if comp, ok := expr.(*ast.CompositeLit); ok {
		for _, elt := range comp.Elts {
			analyExpr(elt)
		}
	} else {

		switch expr.(type) {
		case *ast.KeyValueExpr:
			fmt.Println("KeyValueExpr")
			analyExpr(expr.(*ast.KeyValueExpr).Value)
		case *ast.FuncType:
			fmt.Println("FuncType")
			fmt.Println(expr.(*ast.FuncType))
		case *ast.FuncLit:
			fmt.Println("FuncLit", expr.(*ast.FuncLit))
			if expr.(*ast.FuncLit).Type != nil {
				if expr.(*ast.FuncLit).Type.Params != nil {
					for _, param := range expr.(*ast.FuncLit).Type.Params.List {
						fmt.Println(param.Names)
					}
				}
				if expr.(*ast.FuncLit).Type.Results != nil {
					for _, result := range expr.(*ast.FuncLit).Type.Results.List {
						fmt.Println("xxx", result.Type.(*ast.InterfaceType))
						// analyExpr(result.Type)
					}
				}
			}

			for _, lis := range expr.(*ast.FuncLit).Body.List {
				analyStmt(lis)
			}
		case *ast.CallExpr:
			fmt.Println("CallExpr")

			analyExpr(expr.(*ast.CallExpr).Fun)
			for _, arg := range expr.(*ast.CallExpr).Args {
				analyExpr(arg)
			}
		case *ast.Ident:
			fmt.Println("Ident")
			fmt.Println(expr.(*ast.Ident), expr.(*ast.Ident).Obj)
		}
	}
}

func analyStmt(stmt ast.Stmt) {
	switch stmt.(type) {
	case *ast.ReturnStmt:
		fmt.Println("ReturnStmt")
		for _, result := range stmt.(*ast.ReturnStmt).Results {
			analyExpr(result)
		}
	}
}

func main() {

	var err error
	var filedir string
	if len(os.Args) > 1 {
		filedir, _ = filepath.Split(filepath.FromSlash(os.Args[1]))
		fmt.Println("read file:", filedir)
		_, err := os.Stat(os.Args[0])
		if err != nil {
			fmt.Println("输入可用的pb.go文件,错误信息:", err)
			return
		}

	} else {
		fmt.Println("请输入可用pb.go文件")

	}

	readpb(os.Args[1])

	gmt_set_func := func(msg_type string) ast.Expr {
		/*new(HelloRequest)*/
		_GMT_New_Type := &ast.CallExpr{
			Fun:  ast.NewIdent("new"),
			Args: []ast.Expr{ast.NewIdent(msg_type)},
		}
		/*return new(HelloRequest)*/
		_GMT_Func_Return := &ast.ReturnStmt{
			Results: []ast.Expr{_GMT_New_Type},
		}
		/*func() interface {} {
		        return new(HelloRequest)
		}*/
		_GMT_Func := &ast.FuncLit{
			Type: &ast.FuncType{
				Params: &ast.FieldList{},
				Results: &ast.FieldList{
					List: []*ast.Field{{Type: &ast.InterfaceType{
						Methods: &ast.FieldList{},
					}}},
				},
			},
			Body: &ast.BlockStmt{
				List: []ast.Stmt{_GMT_Func_Return},
			},
		}
		return _GMT_Func
	}

	msg_buff_new := func(newname string, value ast.Expr) ast.Expr {
		/*
			{Name: newHelloRequest, MsgType: func() interface {} {
				        return new(HelloRequest)
				}}
		*/
		msg_type := &ast.CompositeLit{
			Elts: []ast.Expr{
				&ast.KeyValueExpr{
					Key: ast.NewIdent("Name"),
					Value: &ast.BasicLit{
						Kind:  token.STRING,
						Value: fmt.Sprintf("\"%v\"", newname),
					},
				},
				&ast.KeyValueExpr{
					Key:   ast.NewIdent("MsgType"),
					Value: value,
				},
			},
		}
		return msg_type
	}
	pack(msg_buff_new, gmt_set_func)
	msg_type_add := mkpb2(func(expr ast.Node) {
		if v, ok := expr.(*ast.ValueSpec); ok {
			if comp, ok := v.Values[0].(*ast.CompositeLit); ok {
				for msg_new, msg_ty := range grpc_msg {
					comp.Elts = append(comp.Elts, msg_buff_new(msg_new, gmt_set_func(msg_ty)))
				}
			}
		}
	})

	pb2newfile := filepath.Join(filedir, "pb.interface.go")

	pf, err := os.Create(pb2newfile)
	if err != nil {
		fmt.Println("创建", pb2newfile, "文件失败", err)
	}
	defer pf.Close()
	// fmt.Println(err, printer.Fprint(f, token.NewFileSet(), mv2))
	writeGoFile(pf, msg_type_add)
	fmt.Println("cread pb.interface.go successful")
}

func writeGoFile(wr io.Writer, node interface{}) error {
	// 输出Go代码
	header := "// Code generated by tools of liuli`s  DO NOT EDIT.\n"

	wr.Write([]byte(header))
	err := printer.Fprint(wr, token.NewFileSet(), node)
	if err != nil {
		return err
	}

	return nil
}
func pack(newfunc func(newname string, value ast.Expr) ast.Expr, f func(msg_type string) ast.Expr) {
	msg_buff := ast.NewIdent("msg_buff")
	msg_buff.Obj = &ast.Object{
		Name: "msg_buff",
	}
	msg := &ast.CompositeLit{
		Type: &ast.ArrayType{
			Elt: msg_buff,
		},
		Elts: []ast.Expr{}}
	for n, v := range grpc_msg {
		msg.Elts = append(msg.Elts, newfunc(n, f(v)))
	}

}

func test(msg ast.Expr) {
	msg_struct := &ast.StructType{
		Fields: &ast.FieldList{
			List: []*ast.Field{
				{
					Names: []*ast.Ident{ast.NewIdent("Name")},
					Type:  ast.NewIdent("string"),
				},
				{
					Names: []*ast.Ident{ast.NewIdent("MsgType")},
					Type: &ast.FuncType{
						Params: &ast.FieldList{},
						Results: &ast.FieldList{
							List: []*ast.Field{
								{
									Type: &ast.InterfaceType{
										Methods: &ast.FieldList{},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	m := ast.NewIdent("grpc_msg")
	m.Obj = &ast.Object{
		Name: "grpc_msg",
		Kind: ast.Var,
		Decl: msg,
	}
	// mv := &ast.ValueSpec{
	// 	Names:  []*ast.Ident{m},
	// 	Values: []ast.Expr{msg},
	// }
	fmt.Println(msg_struct)
}
