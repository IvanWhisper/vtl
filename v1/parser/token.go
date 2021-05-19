package parser

type Token struct {
	// 原始字符
	Tok string
	// 类型，有 Literal、Operator 两种
	Type int

	Offset int
}
