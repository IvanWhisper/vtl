package parser

import "strings"

var (
	OperationCahr = []byte{
		'(',
		')',
		'+',
		'*',
		'=',
	}
)

const (
	// 字面量，e.g. 50
	Literal = iota
	// 操作符, e.g. + - * /
	Operator
)

// 对错误包装，进行可视化展示
func ErrPos(s string, pos int) string {
	r := strings.Repeat("-", len(s)) + "\n"
	s += "\n"
	for i := 0; i < pos; i++ {
		s += " "
	}
	s += "^\n"
	return r + s + r
}

// AST 生成器结构体
type AST struct {
	// 词法分析的结果
	Tokens []*Token
	// 源字符串
	source string
	// 当前分析器分析的 Token
	currTok *Token
	// 当前分析器的位置
	currIndex int
	// 错误收集
	Err error

	IdFunc        func() int64
	ProjectIdFunc func() int64
}
