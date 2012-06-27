package gofmtcss

import (
	"io"
	"bufio"
	"log"
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
	tokens []CssToken
	currentToken *CssToken
}

func (s *CssScanner) Init(src io.Reader) {
	s.src = bufio.NewReader(src)
}

func (s *CssScanner) Next() {
	for {
		nextByte, err := s.src.ReadByte()
		if err != nil {
			log.Print(err)
			return
		}
		var peek []byte
		var nextNextByte byte
		switch {
			case nextByte == '/': 
				peek, _ = s.src.Peek(1)
				if(string(peek) == "*") {
					nextNextByte, _ = s.src.ReadByte()
					if(s.currentToken != nil) {
						s.tokens = append(s.tokens, *s.currentToken)
					}
					s.currentToken = &CssToken {Value: string(nextByte) + string(nextNextByte), Token: Comment}
					log.Print("Start comment")
					log.Printf("Test %v %v", s.currentToken.Value, s.currentToken.Token)
				}
			case s.currentToken != nil && s.currentToken.Token == Comment && nextByte == '*':
				peek, _ = s.src.Peek(1)
				if(string(peek) == "/") {
					nextNextByte, _ = s.src.ReadByte()
					s.currentToken.Value += string(nextByte) + string(nextNextByte)
					s.tokens = append(s.tokens, *s.currentToken)
					s.currentToken = nil
					log.Print("End comment")
				} else {
					s.currentToken.Value += string(nextByte)
					log.Print("Append to comment")
				}
			case s.currentToken != nil && s.currentToken.Token == Comment:
				s.currentToken.Value += string(nextByte)
				log.Print("Append to comment")
		}
	}
	/* bytes, err := s.src.Peek(2)
	if err != nil {
		log.Fatal(err)
		return
	}
	if string(bytes) == "/*" {
		comment, _ := s.src.ReadString('*')
		//TODO: Check err
		for loopByte byte; string(loopByte) == "/"; loopByte, _ = s.src.ReadByte() {
			comment += string(loopByte)
			log.Print("in loop")
			loopComment, _ := s.src.ReadString('*')
			//TODO: Check err
			comment += loopComment
		}
		endOfComment, _ := s.src.ReadByte()
		comment += string(endOfComment)
		s.tokens = append(s.tokens, CssToken {Value: comment, Token: Comment} )
	} */
}