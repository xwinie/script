package parser

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
)

func (lex *Lexer) loadFile(path string, pos Token) *Node {
	if strings.Contains(lex.loop, path) {
		lex.Error(fmt.Sprintf("%s and %s are importing each other", lex.scanner.Pos.Source, path))
	}

	if n, ok := lex.cache[path]; ok {
		return n
	}

	code, err := ioutil.ReadFile(path)
	if err != nil {
		lex.Error(err.Error())
	}
	n, _, err := parse(bytes.NewReader(code), path, lex.cache, lex.loop+"?"+path)
	if err != nil {
		lex.Error(err.Error())
	}

	// now the required code is loaded, for naming scope we will wrap them into a closure
	cls := CompNode("func", "<a>", CompNode(), n).setPos0(pos)
	node := CompNode("call", cls, CompNode()).setPos0(pos)
	lex.cache[path] = node
	return node
}
