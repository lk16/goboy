package gb

import "log"

// OpcodeCycles is the number of cpu cycles for each normal opcode.
var OpcodeCycles = []int{
	1, 3, 2, 2, 1, 1, 2, 1, 5, 2, 2, 2, 1, 1, 2, 1, // 0
	0, 3, 2, 2, 1, 1, 2, 1, 3, 2, 2, 2, 1, 1, 2, 1, // 1
	2, 3, 2, 2, 1, 1, 2, 1, 2, 2, 2, 2, 1, 1, 2, 1, // 2
	2, 3, 2, 2, 3, 3, 3, 1, 2, 2, 2, 2, 1, 1, 2, 1, // 3
	1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1, // 4
	1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1, // 5
	1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1, // 6
	2, 2, 2, 2, 2, 2, 0, 2, 1, 1, 1, 1, 1, 1, 2, 1, // 7
	1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1, // 8
	1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1, // 9
	1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1, // a
	1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1, // b
	2, 3, 3, 4, 3, 4, 2, 4, 2, 4, 3, 0, 3, 6, 2, 4, // c
	2, 3, 3, 0, 3, 4, 2, 4, 2, 4, 3, 0, 3, 0, 2, 4, // d
	3, 3, 2, 0, 0, 4, 2, 4, 4, 1, 4, 0, 0, 0, 2, 4, // e
	3, 3, 2, 1, 0, 4, 2, 4, 3, 2, 4, 1, 0, 0, 2, 4, // f
} //0  1  2  3  4  5  6  7  8  9  a  b  c  d  e  f

// CBOpcodeCycles is the number of cpu cycles for each CB opcode.
var CBOpcodeCycles = []int{
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2, // 0
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2, // 1
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2, // 2
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2, // 3
	2, 2, 2, 2, 2, 2, 3, 2, 2, 2, 2, 2, 2, 2, 3, 2, // 4
	2, 2, 2, 2, 2, 2, 3, 2, 2, 2, 2, 2, 2, 2, 3, 2, // 5
	2, 2, 2, 2, 2, 2, 3, 2, 2, 2, 2, 2, 2, 2, 3, 2, // 6
	2, 2, 2, 2, 2, 2, 3, 2, 2, 2, 2, 2, 2, 2, 3, 2, // 7
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2, // 8
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2, // 9
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2, // A
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2, // B
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2, // C
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2, // D
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2, // E
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2, // F
} //0  1  2  3  4  5  6  7  8  9  a  b  c  d  e  f

// executeNextOpcode gets the value at the current PC address, increments the PC,
// updates the CPU ticks and executes the opcode.
func (gb *Gameboy) executeNextOpcode() int {
	opcode := uint(gb.CPU.popPC(gb.Memory))
	gb.thisCpuTicks = OpcodeCycles[opcode] * 4
	executeInstruction(gb, gb.CPU, gb.Memory, opcode)
	return gb.thisCpuTicks
}

func executeInstruction(gb *Gameboy, cpu *CPU, mem *Memory, opcode uint) {

	opcode &= 0xFF

	if opcode == 0x76 {
		gb.halted = true
	}

	switch opcode & 0xC0 {
	case 0x40:
		var val byte

		switch opcode & 0x07 {
		case 0x00:
			val = cpu.b()
		case 0x01:
			val = cpu.c()
		case 0x02:
			val = cpu.d()
		case 0x03:
			val = cpu.e()
		case 0x04:
			val = cpu.h()
		case 0x05:
			val = cpu.l()
		case 0x06:
			val = mem.Read(cpu.hl())
		case 0x07:
			val = cpu.a()
		}

		switch opcode & 0x38 {
		case 0x00:
			cpu.setB(val)
		case 0x08:
			cpu.setC(val)
		case 0x10:
			cpu.setD(val)
		case 0x18:
			cpu.setE(val)
		case 0x20:
			cpu.setH(val)
		case 0x28:
			cpu.setL(val)
		case 0x30:
			mem.Write(cpu.hl(), val)
		case 0x38:
			cpu.setA(val)
		}
		return
	case 0x80:
		var val byte

		switch opcode & 0x07 {
		case 0x00:
			val = cpu.b()
		case 0x01:
			val = cpu.c()
		case 0x02:
			val = cpu.d()
		case 0x03:
			val = cpu.e()
		case 0x04:
			val = cpu.h()
		case 0x05:
			val = cpu.l()
		case 0x06:
			val = mem.Read(cpu.hl())
		case 0x07:
			val = cpu.a()
		}

		switch opcode & 0x38 {
		case 0x00:
			cpu.instAdd(val, false)
		case 0x08:
			cpu.instAdd(val, true)
		case 0x10:
			cpu.instSub(val, false)
		case 0x18:
			cpu.instSub(val, true)
		case 0x20:
			cpu.instAnd(val)
		case 0x28:
			cpu.instXor(val)
		case 0x30:
			cpu.instOr(val)
		case 0x38:
			cpu.instCp(val)
		}
		return

	}

	if opcode&0xCF == 0x02 {

		var addr uint16

		switch opcode {
		case 0x02:
			addr = cpu.bc()
		case 0x12:
			addr = cpu.de()
		case 0x22:
			fallthrough
		case 0x32:
			addr = cpu.hl()
		}

		mem.Write(addr, cpu.a())

		switch opcode {
		case 0x22:
			cpu.setHl(cpu.hl() + 1)
		case 0x32:
			cpu.setHl(cpu.hl() - 1)
		}
		return
	}

	if opcode&0xC7 == 0x06 {

		pc := cpu.popPC(mem)

		switch opcode {
		case 0x06:
			cpu.setB(pc)
		case 0x0E:
			cpu.setC(pc)
		case 0x16:
			cpu.setD(pc)
		case 0x1E:
			cpu.setE(pc)
		case 0x26:
			cpu.setH(pc)
		case 0x2E:
			cpu.setL(pc)
		case 0x3E:
			cpu.setA(pc)
		case 0x36:
			mem.Write(cpu.hl(), pc)
		}
		return
	}

	if opcode&0xCF == 0x0A {

		var val byte

		switch opcode {
		case 0x0A:
			val = mem.Read(cpu.bc())
		case 0x1A:
			val = mem.Read(cpu.de())
		case 0x2A:
			val = mem.Read(cpu.hl())
		case 0x3A:
			val = mem.Read(cpu.hl())
		}

		cpu.setA(val)

		switch opcode {
		case 0x3A:
			cpu.setHl(cpu.hl() - 1)
		case 0x2A:
			cpu.setHl(cpu.hl() + 1)
		}
		return
	}

	if opcode&0xC7 == 0xC6 {

		pc := cpu.popPC(mem)

		switch opcode {
		case 0xC6:
			cpu.instAdd(pc, false)
		case 0xCE:
			cpu.instAdd(pc, true)
		case 0xD6:
			cpu.instSub(pc, false)
		case 0xDE:
			cpu.instSub(pc, true)
		case 0xE6:
			cpu.instAnd(pc)
		case 0xEE:
			cpu.instXor(pc)
		case 0xF6:
			cpu.instOr(pc)
		case 0xFE:
			cpu.instCp(pc)
		}
		return
	}

	switch opcode {

	case 0xFA:
		// LD A,(nn)
		val := mem.Read(cpu.popPC16(mem))
		cpu.setA(val)
	case 0xEA:
		// LD (nn),A
		val := cpu.a()
		mem.Write(cpu.popPC16(mem), val)
	case 0xF2:
		// LD A,(C)
		val := 0xFF00 + uint16(cpu.c())
		cpu.setA(mem.Read(val))
	case 0xE2:
		// LD (C),A
		val := cpu.a()
		addr := 0xFF00 + uint16(cpu.c())
		mem.Write(addr, val)

	case 0xE0:
		// LD (0xFF00+n),A
		val := 0xFF00 + uint16(cpu.popPC(mem))
		mem.Write(val, cpu.a())
	case 0xF0:
		// LD A,(0xFF00+n)
		val := mem.ReadHighRam(0xFF00 + uint16(cpu.popPC(mem)))
		cpu.setA(val)
	// ========== 16-Bit Loads ===========
	case 0x01:
		// LD BC,nn
		val := cpu.popPC16(mem)
		cpu.setBc(val)
	case 0x11:
		// LD DE,nn
		val := cpu.popPC16(mem)
		cpu.setDe(val)
	case 0x21:
		// LD HL,nn
		val := cpu.popPC16(mem)
		cpu.setHl(val)
	case 0x31:
		// LD SP,nn
		val := cpu.popPC16(mem)
		cpu.setSp(val)
	case 0xF9:
		// LD SP,HL
		val := cpu.hl()
		cpu.setSp(val)
	case 0xF8:
		// LD HL,SP+n
		cpu.instAdd16Signed(cpu.setHl, cpu.sp(), int8(cpu.popPC(mem)))
	case 0x08:
		// LD (nn),SP
		address := cpu.popPC16(mem)
		mem.Write(address, cpu.spLo())
		mem.Write(address+1, cpu.spHi())
	case 0xF5:
		// PUSH AF
		cpu.pushStack(mem, cpu.af())
	case 0xC5:
		// PUSH BC
		cpu.pushStack(mem, cpu.bc())
	case 0xD5:
		// PUSH DE
		cpu.pushStack(mem, cpu.de())
	case 0xE5:
		// PUSH HL
		cpu.pushStack(mem, cpu.hl())
	case 0xF1:
		// POP AF
		cpu.setAf(cpu.popStack(mem))
	case 0xC1:
		// POP BC
		cpu.setBc(cpu.popStack(mem))
	case 0xD1:
		// POP DE
		cpu.setDe(cpu.popStack(mem))
	case 0xE1:
		// POP HL
		cpu.setHl(cpu.popStack(mem))
	// ========== 8-Bit ALU ===========

	case 0x3C:
		// INC A
		cpu.instInc(cpu.setA, cpu.a())
	case 0x04:
		// INC B
		cpu.instInc(cpu.setB, cpu.b())
	case 0x0C:
		// INC C
		cpu.instInc(cpu.setC, cpu.c())
	case 0x14:
		// INC D
		cpu.instInc(cpu.setD, cpu.d())
	case 0x1C:
		// INC E
		cpu.instInc(cpu.setE, cpu.e())
	case 0x24:
		// INC H
		cpu.instInc(cpu.setH, cpu.h())
	case 0x2C:
		// INC L
		cpu.instInc(cpu.setL, cpu.l())
	case 0x34:
		// INC (HL)
		addr := cpu.hl()
		cpu.instInc(func(val byte) { mem.Write(addr, val) }, mem.Read(addr))
	case 0x3D:
		// DEC A
		cpu.instDec(cpu.setA, cpu.a())
	case 0x05:
		// DEC B
		cpu.instDec(cpu.setB, cpu.b())
	case 0x0D:
		// DEC C
		cpu.instDec(cpu.setC, cpu.c())
	case 0x15:
		// DEC D
		cpu.instDec(cpu.setD, cpu.d())
	case 0x1D:
		// DEC E
		cpu.instDec(cpu.setE, cpu.e())
	case 0x25:
		// DEC H
		cpu.instDec(cpu.setH, cpu.h())
	case 0x2D:
		// DEC L
		cpu.instDec(cpu.setL, cpu.l())
	case 0x35:
		// DEC (HL)
		addr := cpu.hl()
		cpu.instDec(func(val byte) { mem.Write(addr, val) }, mem.Read(addr))
		// ========== 16-Bit ALU ===========
	case 0x09:
		// ADD HL,BC
		cpu.instAdd16(cpu.setHl, cpu.hl(), cpu.bc())
	case 0x19:
		// ADD HL,DE
		cpu.instAdd16(cpu.setHl, cpu.hl(), cpu.de())
	case 0x29:
		// ADD HL,HL
		cpu.instAdd16(cpu.setHl, cpu.hl(), cpu.hl())
	case 0x39:
		// ADD HL,SP
		cpu.instAdd16(cpu.setHl, cpu.hl(), cpu.sp())
	case 0xE8:
		// ADD SP,n
		cpu.instAdd16Signed(cpu.setSp, cpu.sp(), int8(cpu.popPC(mem)))
		cpu.setFlagZ(false)
	case 0x03:
		// INC BC
		cpu.instInc16(cpu.setBc, cpu.bc())
	case 0x13:
		// INC DE
		cpu.instInc16(cpu.setDe, cpu.de())
	case 0x23:
		// INC HL
		cpu.instInc16(cpu.setHl, cpu.hl())
	case 0x33:
		// INC SP
		cpu.instInc16(cpu.setSp, cpu.sp())
	case 0x0B:
		// DEC BC
		cpu.instDec16(cpu.setBc, cpu.bc())
	case 0x1B:
		// DEC DE
		cpu.instDec16(cpu.setDe, cpu.de())
	case 0x2B:
		// DEC HL
		cpu.instDec16(cpu.setHl, cpu.hl())
	case 0x3B:
		// DEC SP
		cpu.instDec16(cpu.setSp, cpu.sp())
	case 0x27:
		// DAA

		// When this instruction is executed, the A register is BCD
		// corrected using the contents of the flags. The exact process
		// is the following: if the least significant four bits of A
		// contain a non-BCD digit (i. e. it is greater than 9) or the
		// H flag is set, then 0x60 is added to the register. Then the
		// four most significant bits are checked. If this more significant
		// digit also happens to be greater than 9 or the C flag is set,
		// then 0x60 is added.
		if !cpu.flagN() {
			if cpu.flagC() || cpu.a() > 0x99 {
				cpu.setA(cpu.a() + 0x60)
				cpu.setFlagC(true)
			}
			if cpu.flagH() || cpu.a()&0xF > 0x9 {
				cpu.setA(cpu.a() + 0x06)
				cpu.setFlagH(false)
			}
		} else if cpu.flagC() && cpu.flagH() {
			cpu.setA(cpu.a() + 0x9A)
			cpu.setFlagH(false)
		} else if cpu.flagC() {
			cpu.setA(cpu.a() + 0xA0)
		} else if cpu.flagH() {
			cpu.setA(cpu.a() + 0xFA)
			cpu.setFlagH(false)
		}
		cpu.setFlagZ(cpu.a() == 0)
	case 0x2F:
		// CPL
		cpu.setA(0xFF ^ cpu.a())
		cpu.setFlagN(true)
		cpu.setFlagH(true)
	case 0x3F:
		// CCF
		cpu.setFlagN(false)
		cpu.setFlagH(false)
		cpu.setFlagC(!cpu.flagC())
	case 0x37:
		// SCF
		cpu.setFlagN(false)
		cpu.setFlagH(false)
		cpu.setFlagC(true)
	case 0x00:
		// NOP
	case 0x76:
		// HALT
		gb.halted = true
	case 0x10:
		// STOP
		gb.halted = true
		if gb.IsCGB() {
			// Handle switching to double speed mode
			gb.checkSpeedSwitch()
		}

		// Pop the next value as the STOP instruction is 2 bytes long. The second value
		// can be ignored, although generally it is expected to be 0x00 and any other
		// value is counted as a corrupted STOP instruction.
		cpu.popPC(mem)
	case 0xF3:
		// DI
		gb.interruptsOn = false
	case 0xFB:
		// EI
		gb.interruptsEnabling = true
	case 0x07:
		// RLCA
		value := cpu.a()
		result := byte(value<<1) | (value >> 7)
		cpu.setA(result)
		cpu.setFlagZ(false)
		cpu.setFlagN(false)
		cpu.setFlagH(false)
		cpu.setFlagC(value > 0x7F)
	case 0x17:
		// RLA
		value := cpu.a()
		var carry byte
		if cpu.flagC() {
			carry = 1
		}
		result := byte(value<<1) + carry
		cpu.setA(result)
		cpu.setFlagZ(false)
		cpu.setFlagN(false)
		cpu.setFlagH(false)
		cpu.setFlagC(value > 0x7F)
	case 0x0F:
		// RRCA
		value := cpu.a()
		result := byte(value>>1) | byte((value&1)<<7)
		cpu.setA(result)
		cpu.setFlagZ(false)
		cpu.setFlagN(false)
		cpu.setFlagH(false)
		cpu.setFlagC(result > 0x7F)
	case 0x1F:
		// RRA
		value := cpu.a()
		var carry byte
		if cpu.flagC() {
			carry = 0x80
		}
		result := byte(value>>1) | carry
		cpu.setA(result)
		cpu.setFlagZ(false)
		cpu.setFlagN(false)
		cpu.setFlagH(false)
		cpu.setFlagC((1 & value) == 1)
	case 0xC3:
		// JP nn
		cpu.instJump(cpu.popPC16(mem))
	case 0xC2:
		// JP NZ,nn
		next := cpu.popPC16(mem)
		if !cpu.flagZ() {
			cpu.instJump(next)
			gb.thisCpuTicks += 4
		}
	case 0xCA:
		// JP Z,nn
		next := cpu.popPC16(mem)
		if cpu.flagZ() {
			cpu.instJump(next)
			gb.thisCpuTicks += 4
		}
	case 0xD2:
		// JP NC,nn
		next := cpu.popPC16(mem)
		if !cpu.flagC() {
			cpu.instJump(next)
			gb.thisCpuTicks += 4
		}
	case 0xDA:
		// JP C,nn
		next := cpu.popPC16(mem)
		if cpu.flagC() {
			cpu.instJump(next)
			gb.thisCpuTicks += 4
		}
	case 0xE9:
		// JP HL
		cpu.instJump(cpu.hl())
	case 0x18:
		// JR n
		addr := int32(cpu.PC) + int32(int8(cpu.popPC(mem)))
		cpu.instJump(uint16(addr))
	case 0x20:
		// JR NZ,n
		next := int8(cpu.popPC(mem))
		if !cpu.flagZ() {
			addr := int32(cpu.PC) + int32(next)
			cpu.instJump(uint16(addr))
			gb.thisCpuTicks += 4
		}
	case 0x28:
		// JR Z,n
		next := int8(cpu.popPC(mem))
		if cpu.flagZ() {
			addr := int32(cpu.PC) + int32(next)
			cpu.instJump(uint16(addr))
			gb.thisCpuTicks += 4
		}
	case 0x30:
		// JR NC,n
		next := int8(cpu.popPC(mem))
		if !cpu.flagC() {
			addr := int32(cpu.PC) + int32(next)
			cpu.instJump(uint16(addr))
			gb.thisCpuTicks += 4
		}
	case 0x38:
		// JR C,n
		next := int8(cpu.popPC(mem))
		if cpu.flagC() {
			addr := int32(cpu.PC) + int32(next)
			cpu.instJump(uint16(addr))
			gb.thisCpuTicks += 4
		}
	case 0xCD:
		// CALL nn
		cpu.instCall(mem, cpu.popPC16(mem))
	case 0xC4:
		// CALL NZ,nn
		next := cpu.popPC16(mem)
		if !cpu.flagZ() {
			cpu.instCall(mem, next)
			gb.thisCpuTicks += 12
		}
	case 0xCC:
		// CALL Z,nn
		next := cpu.popPC16(mem)
		if cpu.flagZ() {
			cpu.instCall(mem, next)
			gb.thisCpuTicks += 12
		}
	case 0xD4:
		// CALL NC,nn
		next := cpu.popPC16(mem)
		if !cpu.flagC() {
			cpu.instCall(mem, next)
			gb.thisCpuTicks += 12
		}
	case 0xDC:
		// CALL C,nn
		next := cpu.popPC16(mem)
		if cpu.flagC() {
			cpu.instCall(mem, next)
			gb.thisCpuTicks += 12
		}
	case 0xC7:
		// RST 0x00
		cpu.instCall(mem, 0x0000)
	case 0xCF:
		// RST 0x08
		cpu.instCall(mem, 0x0008)
	case 0xD7:
		// RST 0x10
		cpu.instCall(mem, 0x0010)
	case 0xDF:
		// RST 0x18
		cpu.instCall(mem, 0x0018)
	case 0xE7:
		// RST 0x20
		cpu.instCall(mem, 0x0020)
	case 0xEF:
		// RST 0x28
		cpu.instCall(mem, 0x0028)
	case 0xF7:
		// RST 0x30
		cpu.instCall(mem, 0x0030)
	case 0xFF:
		// RST 0x38
		cpu.instCall(mem, 0x0038)
	case 0xC9:
		// RET
		cpu.instRet(mem)
	case 0xC0:
		// RET NZ
		if !cpu.flagZ() {
			cpu.instRet(mem)
			gb.thisCpuTicks += 12
		}
	case 0xC8:
		// RET Z
		if cpu.flagZ() {
			cpu.instRet(mem)
			gb.thisCpuTicks += 12
		}
	case 0xD0:
		// RET NC
		if !cpu.flagC() {
			cpu.instRet(mem)
			gb.thisCpuTicks += 12
		}
	case 0xD8:
		// RET C
		if cpu.flagC() {
			cpu.instRet(mem)
			gb.thisCpuTicks += 12
		}
	case 0xD9:
		// RETI
		cpu.instRet(mem)
		gb.interruptsEnabling = true
	case 0xCB:
		// CB
		nextInst := cpu.popPC(mem)
		gb.thisCpuTicks += CBOpcodeCycles[nextInst] * 4
		gb.cbInst[nextInst]()
	default:
		log.Panicf("unknown opcode 0x%x", opcode)
	}
}

// Read the value at the PC and increment the PC.
func (cpu *CPU) popPC(mem *Memory) byte {
	opcode := mem.Read(cpu.PC)
	cpu.PC++
	return opcode
}

// Read the next 16bit value at the PC.
func (cpu *CPU) popPC16(mem *Memory) uint16 {
	b1 := uint16(cpu.popPC(mem))
	b2 := uint16(cpu.popPC(mem))
	return b2<<8 | b1
}
