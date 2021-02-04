// Copyright 2018 The go-python Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Opcodes from opcodes.h
package vm

// Type for OpCodes
type OpCode byte

const (
	// Instruction opcodes for compiled code

	POP_TOP     OpCode = 1
	ROT_TWO     OpCode = 2
	ROT_THREE   OpCode = 3
	DUP_TOP     OpCode = 4
	DUP_TOP_TWO OpCode = 5
	ROT_FOUR    OpCode = 6
	NOP         OpCode = 9

	UNARY_POSITIVE OpCode = 10
	UNARY_NEGATIVE OpCode = 11
	UNARY_NOT      OpCode = 12

	UNARY_INVERT OpCode = 15

	BINARY_MATRIX_MULTIPLY  OpCode = 16
	INPLACE_MATRIX_MULTIPLY OpCode = 17

	BINARY_POWER OpCode = 19

	BINARY_MULTIPLY OpCode = 20

	BINARY_MODULO        OpCode = 22
	BINARY_ADD           OpCode = 23
	BINARY_SUBTRACT      OpCode = 24
	BINARY_SUBSCR        OpCode = 25
	BINARY_FLOOR_DIVIDE  OpCode = 26
	BINARY_TRUE_DIVIDE   OpCode = 27
	INPLACE_FLOOR_DIVIDE OpCode = 28
	INPLACE_TRUE_DIVIDE  OpCode = 29

	GET_AITER         OpCode = 50
	GET_ANEXT         OpCode = 51
	BEFORE_ASYNC_WITH OpCode = 52
	BEGIN_FINALLY     OpCode = 53
	END_ASYNC_FOR     OpCode = 54

	// STORE_MAP was removed
	STORE_MAP        OpCode = 54
	INPLACE_ADD      OpCode = 55
	INPLACE_SUBTRACT OpCode = 56
	INPLACE_MULTIPLY OpCode = 57

	INPLACE_MODULO OpCode = 59
	STORE_SUBSCR   OpCode = 60
	DELETE_SUBSCR  OpCode = 61

	BINARY_LSHIFT    OpCode = 62
	BINARY_RSHIFT    OpCode = 63
	BINARY_AND       OpCode = 64
	BINARY_XOR       OpCode = 65
	BINARY_OR        OpCode = 66
	INPLACE_POWER    OpCode = 67
	GET_ITER         OpCode = 68
	PRINT_EXPR       OpCode = 70
	LOAD_BUILD_CLASS OpCode = 71
	YIELD_FROM       OpCode = 72

	GET_YIELD_FROM_ITER OpCode = 69
	GET_AWAITABLE       OpCode = 73

	INPLACE_LSHIFT OpCode = 75
	INPLACE_RSHIFT OpCode = 76
	INPLACE_AND    OpCode = 77
	INPLACE_XOR    OpCode = 78
	INPLACE_OR     OpCode = 79
	BREAK_LOOP     OpCode = 80 // Removed
	WITH_CLEANUP   OpCode = 81 // Removed, replaced with WITH_CLEANUP_START

	WITH_CLEANUP_START  OpCode = 81
	WITH_CLEANUP_FINISH OpCode = 82

	RETURN_VALUE OpCode = 83
	IMPORT_STAR  OpCode = 84

	SETUP_ANNOTATIONS OpCode = 85

	YIELD_VALUE OpCode = 86
	POP_BLOCK   OpCode = 87
	END_FINALLY OpCode = 88
	POP_EXCEPT  OpCode = 89

	HAVE_ARGUMENT OpCode = 90 // OpCodes from here have an argument:

	STORE_NAME      OpCode = 90 // Index in name list
	DELETE_NAME     OpCode = 91 // ""
	UNPACK_SEQUENCE OpCode = 92 // Number of sequence items
	FOR_ITER        OpCode = 93
	UNPACK_EX       OpCode = 94 // Num items before variable part + (Num items after variable part << 8)

	STORE_ATTR    OpCode = 95 // Index in name list
	DELETE_ATTR   OpCode = 96 // ""
	STORE_GLOBAL  OpCode = 97 // ""
	DELETE_GLOBAL OpCode = 98 // ""

	LOAD_CONST  OpCode = 100 // Index in const list
	LOAD_NAME   OpCode = 101 // Index in name list
	BUILD_TUPLE OpCode = 102 // Number of tuple items
	BUILD_LIST  OpCode = 103 // Number of list items
	BUILD_SET   OpCode = 104 // Number of set items
	BUILD_MAP   OpCode = 105 // Always zero for now
	LOAD_ATTR   OpCode = 106 // Index in name list
	COMPARE_OP  OpCode = 107 // Comparison operator
	IMPORT_NAME OpCode = 108 // Index in name list
	IMPORT_FROM OpCode = 109 // Index in name list

	JUMP_FORWARD         OpCode = 110 // Number of bytes to skip
	JUMP_IF_FALSE_OR_POP OpCode = 111 // Target byte offset from beginning of code
	JUMP_IF_TRUE_OR_POP  OpCode = 112 // ""
	JUMP_ABSOLUTE        OpCode = 113 // ""
	POP_JUMP_IF_FALSE    OpCode = 114 // ""
	POP_JUMP_IF_TRUE     OpCode = 115 // ""

	LOAD_GLOBAL OpCode = 116 // Index in name list

	CONTINUE_LOOP OpCode = 119 // Removed
	SETUP_LOOP    OpCode = 120 // Removed
	SETUP_EXCEPT  OpCode = 121 // Removed
	SETUP_FINALLY OpCode = 122 // ""

	LOAD_FAST   OpCode = 124 // Local variable number
	STORE_FAST  OpCode = 125 // Local variable number
	DELETE_FAST OpCode = 126 // Local variable number

	RAISE_VARARGS OpCode = 130 // Number of raise arguments (1, 2 or 3)
	// CALL_FUNCTION_XXX opcodes defined below depend on this definition
	CALL_FUNCTION OpCode = 131 // #args + (#kwargs<<8)
	MAKE_FUNCTION OpCode = 132 // #defaults + #kwdefaults<<8 + #annotations<<16
	BUILD_SLICE   OpCode = 133 // Number of items

	MAKE_CLOSURE OpCode = 134 // Removed
	LOAD_CLOSURE OpCode = 135 // Load free variable from closure
	LOAD_DEREF   OpCode = 136 // Load and dereference from closure cell
	STORE_DEREF  OpCode = 137 // Store into cell
	DELETE_DEREF OpCode = 138 // Delete closure cell

	// The next 3 opcodes must be contiguous and satisfy
	// (CALL_FUNCTION_VAR - CALL_FUNCTION) & 3 == 1
	CALL_FUNCTION_VAR    OpCode = 140 // Removed
	CALL_FUNCTION_KW     OpCode = 141 // #args + (#kwargs<<8)
	CALL_FUNCTION_VAR_KW OpCode = 142 // Removed
	CALL_FUNCTION_EX     OpCode = 142

	SETUP_WITH OpCode = 143

	// Support for opargs more than 16 bits long
	EXTENDED_ARG OpCode = 144

	LIST_APPEND OpCode = 145
	SET_ADD     OpCode = 146
	MAP_ADD     OpCode = 147

	LOAD_CLASSDEREF OpCode = 148 // New in Python 3.4

	// Added after Python 3.4
	BUILD_LIST_UNPACK            OpCode = 149
	BUILD_MAP_UNPACK             OpCode = 150
	BUILD_MAP_UNPACK_WITH_CALL   OpCode = 151
	BUILD_TUPLE_UNPACK           OpCode = 152
	BUILD_SET_UNPACK             OpCode = 153
	SETUP_ASYNC_WITH             OpCode = 154
	FORMAT_VALUE                 OpCode = 155
	BUILD_CONST_KEY_MAP          OpCode = 156
	BUILD_STRING                 OpCode = 157
	BUILD_TUPLE_UNPACK_WITH_CALL OpCode = 158

	LOAD_METHOD  OpCode = 160
	CALL_METHOD  OpCode = 161
	CALL_FINALLY OpCode = 162
	POP_FINALLY  OpCode = 163
)

// Rich comparison opcodes
const (
	PyCmp_LT = iota
	PyCmp_LE
	PyCmp_EQ
	PyCmp_NE
	PyCmp_GT
	PyCmp_GE
	PyCmp_IN
	PyCmp_NOT_IN
	PyCmp_IS
	PyCmp_IS_NOT
	PyCmp_EXC_MATCH
	PyCmp_BAD
)

// If op has an argument
func (op OpCode) HAS_ARG() bool {
	return op >= HAVE_ARGUMENT
}
