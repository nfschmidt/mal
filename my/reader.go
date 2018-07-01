package main

import (
	"strconv"
	"strings"
)

type token = string

type tokenReader struct {
	position int
	tokens   []token
}

func newTokenReader(tokens []token) *tokenReader {
	return &tokenReader{
		position: 0,
		tokens:   tokens,
	}
}

func (tr *tokenReader) Peek() token {
	t := tr.tokens[tr.position]
	return t
}

func (tr *tokenReader) Next() bool {
	tr.position++
	return tr.position < len(tr.tokens)
}

func (tr *tokenReader) Token() token {
	t := tr.tokens[tr.position]
	return t
}

func readStr(str string) malType {
	tokens := tokenizer(str)
	tr := newTokenReader(tokens)
	return readForm(tr)
}

func tokenizer(str string) []token {
	str = strings.Replace(str, ",", " ", -1)
	str = strings.Replace(str, "(", " ( ", -1)
	str = strings.Replace(str, ")", " ) ", -1)

	return strings.Fields(str)
}

func readForm(tr *tokenReader) malType {
	if tr.Peek() == "(" {
		return readList(tr)
	}

	return readAtom(tr)
}

func readList(tr *tokenReader) malType {
	list := []malType{}
	for tr.Next() {
		if tr.Peek() == ")" {
			break
		}

		list = append(list, readForm(tr))
	}

	return list
}

func readAtom(tr *tokenReader) malType {
	token := tr.Peek()

	if strings.Index(token[:0], "0123456789") >= 0 {
		atom, err := strconv.Atoi(token)
		if err != nil {
		}
		return atom
	}

	return malSymbol{name: token}
}
