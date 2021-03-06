package script

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"

	"github.com/coyove/script/parser"
)

func inst(op opCode, a, b uint16) uint32 {
	// 6 + 13 + 13
	return uint32(op)<<26 + uint32(a&0x1fff)<<13 + uint32(b&0x1fff)
}

func jmpInst(op opCode, dist int) uint32 {
	if dist < -(1<<23) || dist >= 1<<23 {
		panic("long jump")
	}
	// 6 + 26
	return uint32(op)<<26 + uint32(dist+1<<23)
}

func splitInst(op uint32) (op1 opCode, a, b uint16) {
	op1 = opCode(op >> 26)
	a = uint16(op>>13) & 0x1fff
	b = uint16(op) & 0x1fff
	return
}

type posVByte []byte

func (p *posVByte) append(idx uint32, line uint32) {
	v := func(v uint64) {
		*p = append(*p, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)
		n := binary.PutUvarint((*p)[len(*p)-10:], v)
		*p = (*p)[:len(*p)-10+n]
	}
	v(uint64(idx))
	v(uint64(line))

}

func (p posVByte) read(i int) (next int, idx, line uint32) {
	rd := p[i:]
	a, n := binary.Uvarint(rd)
	b, n2 := binary.Uvarint(rd[n:])
	if n == 0 || n2 == 0 {
		next = len(p) + 1
		return
	}
	return i + n + n2, uint32(a), uint32(b)
}

type packet struct {
	Code []uint32
	Pos  posVByte
}

func (b *packet) write(buf packet) {
	datalen := len(b.Code)
	b.Code = append(b.Code, buf.Code...)
	i := 0
	for i < len(buf.Pos) {
		var idx, line uint32
		i, idx, line = buf.Pos.read(i)
		b.Pos.append(idx+uint32(datalen), line)
	}
}

func (b *packet) writeInst(op opCode, opa, opb uint16) {
	b.Code = append(b.Code, inst(op, opa, opb))
}

func (b *packet) writeJmpInst(op opCode, d int) {
	b.Code = append(b.Code, jmpInst(op, d))
}

func (b *packet) writePos(p parser.Position) {
	if p.Line == 0 {
		// Debug Code, used to detect a null meta struct
		panicf("DEBUG: null line")
	}
	b.Pos.append(uint32(len(b.Code)), p.Line)
}

func (b *packet) truncateLast() {
	if len(b.Code) > 0 {
		b.Code = b.Code[:len(b.Code)-1]
	}
}

func (b *packet) Len() int {
	return len(b.Code)
}

var singleOp = map[opCode]string{
	OpConcat: parser.AConcat,
	OpAdd:    parser.AAdd,
	OpSub:    parser.ASub,
	OpMul:    parser.AMul,
	OpDiv:    parser.ADiv,
	OpIDiv:   parser.AIDiv,
	OpMod:    parser.AMod,
	OpEq:     parser.AEq,
	OpNeq:    parser.ANeq,
	OpLess:   parser.ALess,
	OpLessEq: parser.ALessEq,
	OpLoad:   parser.ALoad,
	OpStore:  parser.AStore,
	OpSlice:  parser.ASlice,
	OpPow:    parser.APow,
}

func pkPrettify(c *Func, p *Program, toplevel bool, tab int) string {
	sb := &bytes.Buffer{}
	spaces := strings.Repeat("        ", tab)
	spaces2 := ""
	if tab > 0 {
		spaces2 = strings.Repeat("        ", tab-1) + "+-------"
	}

	sb.WriteString(spaces2 + "+ START " + c.String() + "\n")

	readAddr := func(a uint16, rValue bool) string {
		if a == regA {
			return "$a"
		}

		suffix := ""
		if rValue {
			if a&0xfff == p.NilIndex {
				return "nil"
			}
			if a>>12 == 1 || toplevel {
				v := (*p.Stack)[a&0xfff]
				if !v.IsNil() {
					suffix = "`" + v.String()
				}
			}
		}

		if a>>12 == 1 {
			return fmt.Sprintf("g$%d", a&0xfff) + suffix
		}
		return fmt.Sprintf("$%d", a&0xfff) + suffix
	}

	oldpos := c.Code.Pos
	lastLine := uint32(0)

	for i, inst := range c.Code.Code {
		cursor := uint32(i) + 1
		bop, a, b := splitInst(inst)
		sb.WriteString(spaces)

		if len(c.Code.Pos) > 0 {
			next, op, line := c.Code.Pos.read(0)
			// log.Println(cursor, splitInst, unsafe.Pointer(&Pos))
			for uint32(cursor) > op {
				c.Code.Pos = c.Code.Pos[next:]
				if len(c.Code.Pos) == 0 {
					break
				}
				if next, op, line = c.Code.Pos.read(0); uint32(cursor) <= op {
					break
				}
			}

			if op == uint32(cursor) {
				x := "."
				if line != lastLine {
					x = strconv.Itoa(int(line))
					lastLine = line
				}
				sb.WriteString(fmt.Sprintf("|%-4s % 4d| ", x, cursor-1))
				c.Code.Pos = c.Code.Pos[next:]
			} else {
				sb.WriteString(fmt.Sprintf("|     % 4d| ", cursor-1))
			}
		} else {
			sb.WriteString(fmt.Sprintf("|$    % 4d| ", cursor-1))
		}

		switch bop {
		case OpSet:
			sb.WriteString(readAddr(a, false) + " = " + readAddr(b, true))
		case OpPopVClear:
			sb.WriteString("clear-v")
		case OpPopVAll:
			if a == 0 {
				sb.WriteString("$a = popv-all")
			} else {
				sb.WriteString("$a = popv-all-with-a")
			}
		case OpPopV:
			sb.WriteString("$a = pop-v")
		case OpPushV:
			sb.WriteString("pushv " + readAddr(a, true))
		case OpPush:
			sb.WriteString(fmt.Sprintf("push-%d %v", b, readAddr(a, true)))
		case OpMergeAV:
			sb.WriteString("merge-av")
		case OpRet:
			sb.WriteString("ret " + readAddr(a, true))
		case OpLoadFunc:
			sb.WriteString("$a = function:\n")
			cls := p.Functions[a]
			sb.WriteString(pkPrettify(cls, p, false, tab+1))
		case OpCall:
			if b == 1 {
				sb.WriteString("tail-call " + readAddr(a, true))
			} else {
				sb.WriteString("call " + readAddr(a, true))
			}
		case OpCallMap:
			sb.WriteString("callmap " + readAddr(a, true))
		case OpIf, OpIfNot, OpJmp:
			pos := int32(inst&0xffffff) - 1<<23
			pos2 := uint32(int32(cursor) + pos)
			switch bop {
			case OpIfNot:
				sb.WriteString("if not $a ")
			case OpIf:
				sb.WriteString("if $a ")
			}
			sb.WriteString(fmt.Sprintf("jmp %d to %d", pos, pos2))
		case OpInc:
			sb.WriteString("inc " + readAddr(a, false) + " " + readAddr(uint16(b), true))
		case OpLen:
			sb.WriteString("len " + readAddr(a, true))
		case OpNot:
			sb.WriteString("not " + readAddr(a, true))
		default:
			if bs, ok := singleOp[bop]; ok {
				sb.WriteString(bs + " " + readAddr(a, true) + " " + readAddr(b, true))
			} else {
				sb.WriteString(fmt.Sprintf("? %02x", bop))
			}
		}

		sb.WriteString("\n")
	}

	c.Code.Pos = oldpos

	sb.WriteString(spaces2 + "+ END " + c.String())
	return sb.String()
}
