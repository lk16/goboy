package gb

import (
	"github.com/Humpheh/goboy/pkg/bits"
)

// instAdd performs a ADD instruction on the A register and another value and stores the result in A.
// It will also update the CPU flags using the result of the operation.
func (cpu *CPU) instAdd(val2 byte, addCarry bool) {

	val1 := byte(cpu.a())

	carry := int16(bits.B(cpu.flagC() && addCarry))
	total := int16(val1) + int16(val2) + carry
	cpu.setA(byte(total))

	cpu.setFlagZ(byte(total) == 0)
	cpu.setFlagN(false)
	cpu.setFlagH((val2&0xF)+(val1&0xF)+byte(carry) > 0xF)
	cpu.setFlagC(total > 0xFF) // If result is greater than 255
}

// instSub performs a SUB instruction on the A register and another value and stores the result (A - value) in A.
// It will also update the CPU flags using the result of the operation.
func (cpu *CPU) instSub(val2 byte, addCarry bool) {

	val1 := byte(cpu.a())

	carry := int16(bits.B(cpu.flagC() && addCarry))
	dirtySum := int16(val1) - int16(val2) - carry
	total := byte(dirtySum)
	cpu.setA(total)

	cpu.setFlagZ(total == 0)
	cpu.setFlagN(true)
	cpu.setFlagH(int16(val1&0x0f)-int16(val2&0xF)-int16(carry) < 0)
	cpu.setFlagC(dirtySum < 0)
}

// instAnd performs an AND instruction on the A register and another value and stores the result in A.
// It will also update the CPU flags using the result of the operation.
func (cpu *CPU) instAnd(val2 byte) {

	total := cpu.a() & val2
	cpu.setA(total)

	cpu.setFlagZ(total == 0)
	cpu.setFlagN(false)
	cpu.setFlagH(true)
	cpu.setFlagC(false)
}

// instOr performs an OR instruction on the A register and another value and stores the result in A.
// It will also update the CPU flags using the result of the operation.
func (cpu *CPU) instOr(val2 byte) {
	total := cpu.a() | val2
	cpu.setA(total)

	cpu.setFlagZ(total == 0)
	cpu.setFlagN(false)
	cpu.setFlagH(false)
	cpu.setFlagC(false)
}

// instXor performs an XOR instruction on the A register and another value and stores the result in A.
// It will also update the CPU flags using the result of the operation.
func (cpu *CPU) instXor(val2 byte) {
	total := cpu.a() ^ val2
	cpu.setA(total)
	cpu.setFlagZ(total == 0)
	cpu.setFlagN(false)
	cpu.setFlagH(false)
	cpu.setFlagC(false)
}

// instCp performs a CP operation on some value and A.
// It will update the CPU flags using the result of the operation.
func (cpu *CPU) instCp(val1 byte) {
	val2 := cpu.a()
	total := val2 - val1
	cpu.setFlagZ(total == 0)
	cpu.setFlagN(true)
	cpu.setFlagH((val1 & 0x0f) > (val2 & 0x0f))
	cpu.setFlagC(val1 > val2)
}

// Perform an INC operation on a value and stores the result using the set function.
// It will update the CPU flags using the result of the operation.
func (cpu *CPU) instInc(set func(byte), org byte) {
	total := org + 1
	set(total)

	cpu.setFlagZ(total == 0)
	cpu.setFlagN(false)
	cpu.setFlagH(bits.HalfCarryAdd(byte(org), 1))
}

// Perform an DEC operation on a value and stores the result using the set function.
// It will update the CPU flags using the result of the operation.
func (cpu *CPU) instDec(set func(byte), org byte) {
	total := org - 1
	set(total)

	cpu.setFlagZ(total == 0)
	cpu.setFlagN(true)
	cpu.setFlagH(org&0x0F == 0)
}

// Perform a 16bit ADD operation on a value and stores the result using the set function.
// It will update the CPU flags using the result of the operation.
func (cpu *CPU) instAdd16(set func(uint16), val1 uint16, val2 uint16) {
	total := int32(val1) + int32(val2)
	set(uint16(total))
	cpu.setFlagN(false)
	cpu.setFlagH(int32(val1&0xFFF) > (total & 0xFFF))
	cpu.setFlagC(total > 0xFFFF)
}

// instAdd16Signed performs a signed 16bit ADD operation on a value and stores the result using the set function.
// It will update the CPU flags using the result of the operation.
func (cpu *CPU) instAdd16Signed(set func(uint16), val1 uint16, val2 int8) {
	total := uint16(int32(val1) + int32(val2))
	set(total)
	tmpVal := val1 ^ uint16(val2) ^ total
	cpu.setFlagZ(false)
	cpu.setFlagN(false)
	cpu.setFlagH((tmpVal & 0x10) == 0x10)
	cpu.setFlagC((tmpVal & 0x100) == 0x100)
}

// instInc16 performs a 16 bit INC operation on a value and sore the result using the set function.
func (cpu *CPU) instInc16(set func(uint16), org uint16) {
	set(org + 1)
}

// instDec16 performs a 16 bit INC operation on a value and sore the result using the set function.
func (cpu *CPU) instDec16(set func(uint16), org uint16) {
	set(org - 1)
}

// instJump performs a JUMP operation by setting the PC to the value.
func (cpu *CPU) instJump(next uint16) {
	cpu.PC = next
}

// instCall performs a CALL operation by pushing the current PC to the stack and jumping to the next address.
func (cpu *CPU) instCall(mem *Memory, next uint16) {
	cpu.pushStack(mem, cpu.PC)
	cpu.PC = next
}

// instRet performs a RET operation by setting the PC to the next value popped off the stack.
func (cpu *CPU) instRet(mem *Memory) {
	cpu.PC = cpu.popStack(mem)
}

// pushStack pushes a 16 bit value onto the stack and decrements SP.
func (cpu *CPU) pushStack(mem *Memory, address uint16) {
	sp := cpu.sp()
	mem.Write(sp-1, byte(uint16(address&0xFF00)>>8))
	mem.Write(sp-2, byte(address&0xFF))
	cpu.setSp(sp - 2)
}

// popStack pops a 16 bit value off the stack and increments SP.
func (cpu *CPU) popStack(mem *Memory) uint16 {
	sp := cpu.sp()
	byte1 := uint16(mem.Read(sp))
	byte2 := uint16(mem.Read(sp+1)) << 8
	cpu.setSp(sp + 2)
	return byte1 | byte2
}
