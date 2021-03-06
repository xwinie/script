package lib

import (
	"bytes"
	"fmt"

	"github.com/coyove/script"
)

type Set struct {
	m   map[interface{}]struct{}
	max int64
}

func init() {
	script.AddGlobalValue("set", func(env *script.Env) {
		s := &Set{m: map[interface{}]struct{}{}}
		s.max = env.Global.MaxStackSize - int64(len(*env.Global.Stack))
		if env.Global.MaxStackSize == 0 {
			s.max = 0
		} else {
			if s.max < 1 {
				panic("restricted")
			}
		}
		for _, e := range env.Stack() {
			s.Add(e.Interface())
		}
		env.A = script.Interface(s)
	},
		"set() => unique_set",
		"set(e1, e2, ..., en) => unique_set",
		"\tcreate a unique set, methods:",
		"\t\tunique_set.add(value) => ok_or_not",
		"\t\tunique_set.exists(value) => ok_or_not",
		"\t\tunique_set.size() => int",
		"\t\tunique_set.set() => n, e1, e2, ..., en",
	)
}

func (s *Set) Add(v interface{}) bool {
	if len(s.m) > int(s.max) && s.max > 0 {
		panic("set overflow")
	}
	_, exist := s.m[v]
	s.m[v] = struct{}{}
	return !exist
}

func (s *Set) Size() int {
	return len(s.m)
}

func (s *Set) Exists(v interface{}) bool {
	_, ok := s.m[v]
	return ok
}

func (s *Set) Set() []script.Value {
	x := make([]script.Value, 0, len(s.m))
	for k := range s.m {
		x = append(x, script.Interface(k))
	}
	return x
}

func (s *Set) String() string {
	p := bytes.NewBufferString("set(")
	for k := range s.m {
		p.WriteString(fmt.Sprint(k))
		p.WriteString(",")
	}
	if len(s.m) > 0 {
		p.Truncate(p.Len() - 1)
	}
	p.WriteString(")")
	return p.String()
}
