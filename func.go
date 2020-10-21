package potatolang

import (
	"bytes"
	"fmt"
	"strconv"
)

const (
	FuncYield = 1 << iota
	FuncVararg
)

type Func struct {
	packet
	Name       string
	ConstTable []Value
	NumParam   byte
	options    byte
	stackSize  uint16
	yCursor    uint32
	yEnv       Env
	native     func(env *Env)
}

// Native creates a golang-native function
func Native(f func(env *Env)) Value {
	return Fun(&Func{native: f})
}

func (c *Func) setOpt(flag bool, opt byte) {
	if flag {
		c.options |= opt
	}
}

func (c *Func) Is(opt byte) bool { return (c.options & opt) > 0 }

func (c *Func) String() string {
	if c.native != nil {
		return "<native>"
	}

	p := bytes.Buffer{}
	if c.Name != "" {
		p.WriteString(c.Name)
	} else {
		p.WriteString("function")
	}

	p.WriteString("$")
	p.WriteString(strconv.Itoa(len(c.ConstTable)))
	p.WriteString("(")
	for i := 0; i < int(c.NumParam); i++ {
		p.WriteString("a")
		p.WriteString(strconv.Itoa(i))
		p.WriteString(",")
	}

	if c.Is(FuncVararg) {
		p.WriteString("...")
	} else {
		if c.NumParam > 0 {
			p.Truncate(p.Len() - 1)
		}
	}
	p.WriteString(")")

	if c.yEnv.stack != nil {
		p.WriteString(fmt.Sprintf("@%x", c.yCursor))
	}

	return p.String()
}

func (c *Func) PrettyString() string {
	if c.native != nil {
		return "[native Code]"
	}
	return pkPrettify(c, 0)
}

// exec executes the closure with the given Env
func (c *Func) exec(newEnv Env) (Value, []Value) {
	if c.native == nil {
		if c.yEnv.stack != nil {
			newEnv = c.yEnv
		}

		v, vb, np, yield := execCursorLoop(newEnv, c, c.yCursor)
		if yield {
			c.yCursor = np
			c.yEnv = newEnv
		} else {
			c.yCursor = 0
			c.yEnv = Env{}
		}
		return v, vb
	}

	c.native(&newEnv)
	return newEnv.A, newEnv.V
}

func (c *Func) Call(a ...Value) (Value, []Value) {
	var newEnv Env
	var varg []Value
	if c.yEnv.stack == nil {
		for i := range a {
			if i >= int(c.NumParam) {
				varg = append(varg, a[i])
			}
			newEnv.Push(a[i])
		}
		if c.native == nil && c.Is(FuncVararg) {
			newEnv.grow(int(c.NumParam) + 1)
			newEnv._set(uint16(c.NumParam), unpackedStack(&unpacked{a: varg}))
		}
	}
	return c.exec(newEnv)
}

// Terminate will try to stop the execution, when called the closure (along with duplicates) become invalid immediately
func (c *Func) Terminate() {
	const Stop = uint32(OpEOB) << 26
	for i := range c.Code {
		c.Code[i] = Stop
	}
}