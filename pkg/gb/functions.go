package gb

import (
	"github.com/Humpheh/goboy/pkg/bits"
)

// Perform a ADD instruction on the values and store the value using the set
// function. Will also update the CPU flags using the result of the operation.
func (gb *Gameboy) instAdd(set func(byte), val1 byte, val2 byte, addCarry bool) {
	carry := int16(bits.B(gb.CPU.flagC() && addCarry))
	total := int16(val1) + int16(val2) + carry
	set(byte(total))

	gb.CPU.setFlagZ(byte(total) == 0)
	gb.CPU.setFlagN(false)
	gb.CPU.setFlagH((val2&0xF)+(val1&0xF)+byte(carry) > 0xF)
	gb.CPU.setFlagC(total > 0xFF) // If result is greater than 255
}

// Perform a SUB operation on the values (val1 - val2) and store the result using
// the set function. Will also update the CPU flags using the result of the operation.
func (gb *Gameboy) instSub(set func(byte), val1 byte, val2 byte, addCarry bool) {
	carry := int16(bits.B(gb.CPU.flagC() && addCarry))
	dirtySum := int16(val1) - int16(val2) - carry
	total := byte(dirtySum)
	set(total)

	gb.CPU.setFlagZ(total == 0)
	gb.CPU.setFlagN(true)
	gb.CPU.setFlagH(int16(val1&0x0f)-int16(val2&0xF)-int16(carry) < 0)
	gb.CPU.setFlagC(dirtySum < 0)
}

// Perform a AND operation on two values and store the result using the set function.
// Will also update the CPU flags using the result of the operation.
func (gb *Gameboy) instAnd(set func(byte), val1 byte, val2 byte) {
	total := val1 & val2
	set(total)
	gb.CPU.setFlagZ(total == 0)
	gb.CPU.setFlagN(false)
	gb.CPU.setFlagH(true)
	gb.CPU.setFlagC(false)
}

// Perform an OR operation on two values and store the result using the set function.
// Will also update the CPU flags using the result of the operation.
func (gb *Gameboy) instOr(set func(byte), val1 byte, val2 byte) {
	total := val1 | val2
	set(total)
	gb.CPU.setFlagZ(total == 0)
	gb.CPU.setFlagN(false)
	gb.CPU.setFlagH(false)
	gb.CPU.setFlagC(false)
}

// Perform an XOR operation on two values and store the result using the set function.
// Will also update the CPU flags using the result of the operation.
func (gb *Gameboy) instXor(set func(byte), val1 byte, val2 byte) {
	total := val1 ^ val2
	set(total)
	gb.CPU.setFlagZ(total == 0)
	gb.CPU.setFlagN(false)
	gb.CPU.setFlagH(false)
	gb.CPU.setFlagC(false)
}

// Perform a CP operation on two values. Will set the flags from the result of the
// comparison.
func (gb *Gameboy) instCp(val1 byte, val2 byte) {
	total := val2 - val1
	gb.CPU.setFlagZ(total == 0)
	gb.CPU.setFlagN(true)
	gb.CPU.setFlagH((val1 & 0x0f) > (val2 & 0x0f))
	gb.CPU.setFlagC(val1 > val2)
}

// Perform an INC operation on a value and store the result using the set function.
// Will also update the CPU flags using the result of the operation.
func (gb *Gameboy) instInc(set func(byte), org byte) {
	total := org + 1
	set(total)

	gb.CPU.setFlagZ(total == 0)
	gb.CPU.setFlagN(false)
	gb.CPU.setFlagH(bits.HalfCarryAdd(byte(org), 1))
}

// Perform an DEC operation on a value and store the result using the set function.
// Will also update the CPU flags using the result of the operation.
func (gb *Gameboy) instDec(set func(byte), org byte) {
	total := org - 1
	set(total)

	gb.CPU.setFlagZ(total == 0)
	gb.CPU.setFlagN(true)
	gb.CPU.setFlagH(org&0x0F == 0)
}

// Perform a 16bit ADD operation on a value and store the result using the set function.
// Will also update the CPU flags using the result of the operation.
func (gb *Gameboy) instAdd16(set func(uint16), val1 uint16, val2 uint16) {
	total := int32(val1) + int32(val2)
	set(uint16(total))
	gb.CPU.setFlagN(false)
	gb.CPU.setFlagH(int32(val1&0xFFF) > (total & 0xFFF))
	gb.CPU.setFlagC(total > 0xFFFF)
}

// Perform a signed 16bit ADD operation on a value and store the result using the set
// function. Will also update the CPU flags using the result of the operation.
func (gb *Gameboy) instAdd16Signed(set func(uint16), val1 uint16, val2 int8) {
	total := uint16(int32(val1) + int32(val2))
	set(total)
	tmpVal := val1 ^ uint16(val2) ^ total
	gb.CPU.setFlagZ(false)
	gb.CPU.setFlagN(false)
	gb.CPU.setFlagH((tmpVal & 0x10) == 0x10)
	gb.CPU.setFlagC((tmpVal & 0x100) == 0x100)
}

// Perform a 16 bit INC operation on a value ans tore the result using the set function.
func (gb *Gameboy) instInc16(set func(uint16 uint16), org uint16) {
	set(org + 1)
}

// Perform a 16 bit DEC operation on a value ans tore the result using the set function.
func (gb *Gameboy) instDec16(set func(uint16 uint16), org uint16) {
	set(org - 1)
}

// Perform a JUMP operation by setting the PC to the value.
func (gb *Gameboy) instJump(next uint16) {
	gb.CPU.PC = next
}

// Perform a CALL operation by pushing the current PC to the stack and jumping to
// the next address.
func (gb *Gameboy) instCall(next uint16) {
	gb.pushStack(gb.CPU.PC)
	gb.CPU.PC = next
}

// Perform a RET operation by setting the PC to the next value popped off the stack.
func (gb *Gameboy) instRet() {
	gb.CPU.PC = gb.popStack()
}
