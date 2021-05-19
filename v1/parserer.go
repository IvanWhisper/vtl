package v1

import (
	"bufio"
	"context"
	"fmt"
	mlog "github.com/IvanWhisper/michelangelo/log"
	"github.com/IvanWhisper/vtl/expr"
	"github.com/IvanWhisper/vtl/v1/parser"
	"github.com/IvanWhisper/vtl/v1/scanner"
	"io"
	"os"
	"strings"
	"unicode"
)

func ScanFile(ctx context.Context, filename string) []string {
	fi, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	mlog.InfoCtx(fmt.Sprintf("%s is open", filename), ctx)
	defer fi.Close()
	r := bufio.NewReader(fi)
	result := make([]string, 0)
	for {
		line, _, err := r.ReadLine()
		if err != nil && err != io.EOF {
			panic(err)
		}
		if err == io.EOF {
			break
		}
		if len(line) > 0 {
			if unicode.IsSpace(rune(line[0])) {
				result[len(result)-1] = result[len(result)-1] + string(line)
			} else {
				result = append(result, string(line))
			}
		}
	}
	return result
}

func Parse2Tree(ctx context.Context, expStr string) *expr.ExprItem {
	var s scanner.Scanner
	s.Init(strings.NewReader(expStr))
	s.Filename = "expStr"
	tokens := make([]*parser.Token, 0)
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		mlog.InfoCtx(fmt.Sprintf("%s: %s %d \n ", s.Position, s.TokenText(), tok), ctx)
		if s.TokenText() == "BOOL" {
		} else {
			tk := &parser.Token{
				Tok: s.TokenText(),
			}
			if tk.Tok == "+" || tk.Tok == "*" || tk.Tok == "(" || tk.Tok == ")" {
				tk.Type = parser.Operator
			} else {
				tk.Type = parser.Literal
			}
			tokens = append(tokens, tk)
		}
	}
	ast := parser.NewAST(tokens, expStr)
	if ast.Err != nil {
		mlog.ErrorCtx(ast.Err.Error(), ctx)
	}
	// AST builder
	ar := ast.ParseExpression()
	if ast.Err != nil {
		mlog.ErrorCtx(ast.Err.Error(), ctx)
	}
	return ar
}

//func demo(){
//	f:=make([]*expr.ExprItem,0)
//	expr.Flattening(ar,&f,0,0,func()int64{
//		count++
//		return count
//	})
//	if bs, err := json.Marshal(f); err == nil {
//		println("F-Json: ", string(bs))
//	} else {
//		println(err)
//	}
//
//	root:=expr.Tree(f)
//	if bs, err := json.Marshal(root); err == nil {
//		println("T-Json: ", string(bs))
//	} else {
//		println(err)
//	}
//
//	println("!")
//
//}
