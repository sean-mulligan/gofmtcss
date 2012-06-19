package gofmtcss

import (
	"testing"
	"bytes"
)

func TestInit(t *testing.T) {
	inputString := "H1 { color:red; }"
	input := bytes.NewBufferString(inputString)
	tokenizer := new(CssScanner)
	tokenizer.Init(input)
}