package gofmtcss

import (
	"testing"
	"bytes"
	"log"
)

func TestInit(t *testing.T) {
	inputString := "H1 { color:red; }"
	input := bytes.NewBufferString(inputString)
	tokenizer := new(CssScanner)
	tokenizer.Init(input)
}

func TestSingleCommentToken(t *testing.T) {
	inputString := "/* hello */"
	input := bytes.NewBufferString(inputString)
	tokenizer := new(CssScanner)
	tokenizer.Init(input)
	tokenizer.Next()
	if(len(tokenizer.tokens) != 1) {
		t.Errorf("Should only be one comment token. There were %v tokens", len(tokenizer.tokens))
	}
	log.Printf("First Token: %v", tokenizer.tokens[0].Value)
}

func TestTwoCommentTokens(t *testing.T) {
	inputString := "/* hello */   /* good bye */"
	input := bytes.NewBufferString(inputString)
	tokenizer := new(CssScanner)
	tokenizer.Init(input)
	tokenizer.Next()
	if(len(tokenizer.tokens) != 2) {
		t.Errorf("Should only be two comment tokens. There were %v tokens", len(tokenizer.tokens))
	}
	log.Printf("First Token: %v", tokenizer.tokens[0].Value)
	log.Printf("Second Token: %v", tokenizer.tokens[1].Value)
}