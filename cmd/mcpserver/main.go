package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/example/testing-stuff/agent"
)

const generatedFile = "types/generated.go"

// addConstant parses generatedFile and appends the provided constant using
// the Go AST. The file is then formatted to keep lint tools happy.
func addConstant(c agent.Constant) error {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, generatedFile, nil, parser.ParseComments)
	if err != nil {
		return err
	}

	// locate the const declaration
	var constDecl *ast.GenDecl
	for _, d := range node.Decls {
		gd, ok := d.(*ast.GenDecl)
		if ok && gd.Tok == token.CONST {
			constDecl = gd
			break
		}
	}
	if constDecl == nil {
		return fmt.Errorf("no const block in %s", generatedFile)
	}

	spec := &ast.ValueSpec{
		Names:  []*ast.Ident{ast.NewIdent(c.Name)},
		Values: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: strconv.Quote(c.Value)}},
	}
	constDecl.Specs = append(constDecl.Specs, spec)

	var buf bytes.Buffer
	if err := format.Node(&buf, fset, node); err != nil {
		return err
	}
	return os.WriteFile(generatedFile, buf.Bytes(), 0644)
}

func generateHandler(w http.ResponseWriter, r *http.Request) {
	client := &agent.Client{}
	constant, err := client.FetchConstant("")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := addConstant(constant); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Added constant %s", constant.Name)
}

func main() {
	if _, err := os.Stat(generatedFile); os.IsNotExist(err) {
		log.Fatalf("%s not found", generatedFile)
	}
	http.HandleFunc("/generate", generateHandler)
	log.Println("Starting MCP server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
