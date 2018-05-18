package base

const REG_A = -1

const (
	// basic flat op
	OP_ASSERT = iota
	OP_NIL
	OP_TRUE
	OP_FALSE
	OP_SET
	OP_SET_NUM
	OP_SET_STR
	OP_STORE
	OP_LOAD
	OP_SAFE_STORE
	OP_SAFE_LOAD
	OP_ADD
	OP_SUB
	OP_MUL
	OP_DIV
	OP_MOD
	OP_EQ
	OP_NEQ
	OP_LESS
	OP_LESS_EQ
	OP_MORE
	OP_MORE_EQ
	OP_NOT
	OP_AND
	OP_OR
	OP_XOR
	OP_BIT_NOT
	OP_BIT_AND
	OP_BIT_OR
	OP_BIT_XOR
	OP_BIT_LSH
	OP_BIT_RSH

	// complex structure op
	OP_LIST
	OP_MAP

	// flow control op
	OP_IF
	OP_JMP
	OP_LAMBDA
	OP_CALL
	OP_R0
	OP_R0_NUM
	OP_R0_STR
	OP_R1
	OP_R1_NUM
	OP_R1_STR
	OP_R2
	OP_R2_NUM
	OP_R2_STR
	OP_R3
	OP_R3_NUM
	OP_R3_STR
	OP_PUSH
	OP_PUSH_NUM
	OP_PUSH_STR
	OP_RET
	OP_RET_NUM
	OP_RET_STR
	OP_YIELD
	OP_YIELD_NUM
	OP_YIELD_STR

	// special builtin op
	OP_WHO
	OP_DUP
	OP_ERROR
	OP_LEN
	OP_STACK

	OP_NOP = 0xFE
	OP_EOB = 0xFF
)
