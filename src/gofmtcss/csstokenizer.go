package gofmtcss

import (
	"io"
	"bufio"
)

const(
	EOF = -(iota + 1)
	Comment
	Declaration
	OpenBrace
	CloseBrace
)

type CssToken struct {
	Value string
	Token int
}

type CssScanner struct {
	src *bufio.Reader
	srcPos int
	srcEnd int
	
}

func (s *CssScanner) Init(src io.Reader) {
	s.src = bufio.NewReader(src)
}

func (s *CssScanner) Next() {
	
}