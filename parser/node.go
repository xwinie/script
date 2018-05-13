package parser

import (
	"fmt"
	"io"
	"log"
	"math"
	"strconv"
	"strings"
)

type Position struct {
	Source string
	Line   int
	Column int
}

type Token struct {
	Type int
	Name string
	Str  string
	Pos  Position
}

func (self *Token) String() string {
	return fmt.Sprintf("<type:%v, str:%v>", self.Name, self.Str)
}

const (
	NTNumber = iota
	NTString
	NTAtom
	NTCompound
	NTAddr
)

type Node struct {
	Type     byte
	Value    interface{}
	Pos      Position
	Compound []*Node
	LibWH    bool
}

func NewCompoundNode(args ...interface{}) *Node {
	n := &Node{
		Type:     NTCompound,
		Compound: make([]*Node, 0),
	}
	for _, arg := range args {
		switch arg.(type) {
		case string:
			n.Compound = append(n.Compound, &Node{
				Type:  NTAtom,
				Value: arg.(string),
			})
		case *Node:
			if n.Pos.Source == "" {
				n.Pos = arg.(*Node).Pos
			}
			n.Compound = append(n.Compound, arg.(*Node))
		default:
			panic("shouldn't happen")
		}
	}
	return n
}

func NewAtomNode(tok Token) *Node {
	return &Node{
		Type:  NTAtom,
		Value: tok.Str,
		Pos:   tok.Pos,
	}
}

func NewStringNode(arg string) *Node {
	return &Node{
		Type:  NTString,
		Value: arg,
	}
}

func NewNumberNode(arg string) *Node {
	if arg[0] == '0' && len(arg) > 1 {
		var num uint64
		var err error
		switch arg[1] {
		case 'x', 'X':
			num, err = strconv.ParseUint(arg[2:], 16, 64)
		case 'b', 'B':
			num, err = strconv.ParseUint(arg[2:], 2, 64)
		case 'i', 'I':
			num, err = strconv.ParseUint(arg[2:], 16, 64)
			if err == nil {
				return &Node{
					Type:  NTNumber,
					Value: math.Float64frombits(uint64(num)),
				}
			}
		default:
			num, err = strconv.ParseUint(arg[1:], 8, 64)
		}
		if err == nil {
			return &Node{
				Type:  NTNumber,
				Value: float64(num),
			}
		}
	}

	num, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		log.Panicf("invalid number: %s, %+v", arg, err)
	}

	return &Node{
		Type:  NTNumber,
		Value: num,
	}
}

func (n *Node) Dump(w io.Writer) {
	switch n.Type {
	case NTNumber:
		io.WriteString(w, "<"+strconv.FormatFloat(n.Value.(float64), 'f', 9, 64)+">")
	case NTString:
		io.WriteString(w, strconv.Quote(n.Value.(string)))
	case NTAtom:
		io.WriteString(w, n.Value.(string))
	case NTCompound:
		io.WriteString(w, "[")
		for _, a := range n.Compound {
			a.Dump(w)
			io.WriteString(w, " ")
		}
		io.WriteString(w, "]")
	}
}

func (n *Node) String() string {
	pos := fmt.Sprintf("@%s:%d:%d", n.Pos.Source, n.Pos.Line, n.Pos.Column)
	switch n.Type {
	case NTNumber:
		return strconv.FormatFloat(n.Value.(float64), 'f', 9, 64) + pos
	case NTString:
		return strconv.Quote(n.Value.(string)) + pos
	case NTAtom:
		return n.Value.(string) + pos
	case NTCompound:
		buf := make([]string, len(n.Compound))
		for i, a := range n.Compound {
			buf[i] = a.String()
		}
		return "[" + strings.Join(buf, " ") + "]" + pos
	}
	panic("shouldn't happen")
}
