package vm

import (
	"sync"

	"github.com/coyove/bracket/base"
)

var lib_go = LibFunc{
	name: "go",
	args: 1,
	f: func(env *base.Env) base.Value {
		newEnv := base.NewEnv(env)
		cls := env.R0.Closure()
		args := env.R1.Array()

		if cls.ArgsCount() != len(args) {
			panic("not enough arguments to start")
		}

		for _, arg := range args {
			newEnv.Push(arg)
		}

		go Exec(newEnv, cls.Code())
		return base.NewValue()
	},
}

var lib_syncwaitgroupnew = LibFunc{
	name: "wait-group/new",
	args: 0,
	f: func(env *base.Env) base.Value {
		return base.NewGenericValue(&sync.WaitGroup{})
	},
}

var lib_syncwaitgroupadd = LibFunc{
	name: "wait-group/add",
	args: 2,
	f: func(env *base.Env) base.Value {
		wg := env.R0.Generic().(*sync.WaitGroup)
		wg.Add(int(env.R1.Number()))
		return base.NewValue()
	},
}

var lib_syncwaitgroupdone = LibFunc{
	name: "wait-group/done",
	args: 1,
	f: func(env *base.Env) base.Value {
		env.R0.Generic().(*sync.WaitGroup).Done()
		return base.NewValue()
	},
}

var lib_syncwaitgroupwait = LibFunc{
	name: "wait-group/wait",
	args: 1,
	f: func(env *base.Env) base.Value {
		wg := env.R0.Generic().(*sync.WaitGroup)
		wg.Wait()
		return base.NewValue()
	},
}

var lib_syncmutexnew = LibFunc{
	name: "mutex/new",
	args: 0,
	f: func(env *base.Env) base.Value {
		return base.NewGenericValue(&sync.Mutex{})
	},
}

var lib_syncmutexlock = LibFunc{
	name: "mutex/lock",
	args: 1,
	f: func(env *base.Env) base.Value {
		env.R0.Generic().(*sync.Mutex).Lock()
		return base.NewValue()
	},
}

var lib_syncmutexunlock = LibFunc{
	name: "mutex/unlock",
	args: 1,
	f: func(env *base.Env) base.Value {
		env.R0.Generic().(*sync.Mutex).Unlock()
		return base.NewValue()
	},
}
