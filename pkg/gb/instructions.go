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
	opcode := uint(gb.popPC())
	gb.thisCpuTicks = OpcodeCycles[opcode] * 4
	gb.executeInstruction(opcode)
	return gb.thisCpuTicks
}

func (gb *Gameboy) executeInstruction(opcode uint) {

	opcode &= 0xFF

	if opcode == 0x76 {
		// HALT
		gb.halted = true
		return
	}

	if opcode&0xC0 == 0x40 {

		var val byte

		switch opcode & 0x07 {
		case 0x00:
			val = gb.CPU.b()
		case 0x01:
			val = gb.CPU.c()
		case 0x02:
			val = gb.CPU.d()
		case 0x03:
			val = gb.CPU.e()
		case 0x04:
			val = gb.CPU.h()
		case 0x05:
			val = gb.CPU.l()
		case 0x06:
			val = gb.Memory.Read(gb.CPU.hl())
		case 0x07:
			val = gb.CPU.a()
		}

		switch opcode & 0x38 {
		case 0x00:
			gb.CPU.setB(val)
		case 0x08:
			gb.CPU.setC(val)
		case 0x10:
			gb.CPU.setD(val)
		case 0x18:
			gb.CPU.setE(val)
		case 0x20:
			gb.CPU.setH(val)
		case 0x28:
			gb.CPU.setL(val)
		case 0x30:
			gb.Memory.Write(gb.CPU.hl(), val)
		case 0x38:
			gb.CPU.setA(val)
		}
		return
	}

	switch opcode {
	case 0x06:
		// LD B, n
		gb.CPU.setB(gb.popPC())
	case 0x0E:
		// LD C, n
		gb.CPU.setC(gb.popPC())
	case 0x16:
		// LD D, n
		gb.CPU.setD(gb.popPC())
	case 0x1E:
		// LD E, n
		gb.CPU.setE(gb.popPC())
	case 0x26:
		// LD H, n
		gb.CPU.setH(gb.popPC())
	case 0x2E:
		// LD L, n
		gb.CPU.setL(gb.popPC())
	case 0x0A:
		// LD A,(BC)
		val := gb.Memory.Read(gb.CPU.bc())
		gb.CPU.setA(val)
	case 0x1A:
		// LD A,(DE)
		val := gb.Memory.Read(gb.CPU.de())
		gb.CPU.setA(val)
	case 0xFA:
		// LD A,(nn)
		val := gb.Memory.Read(gb.popPC16())
		gb.CPU.setA(val)
	case 0x3E:
		// LD A,(nn)
		val := gb.popPC()
		gb.CPU.setA(val)
	case 0x36:
		// LD (HL),n 36
		val := gb.popPC()
		gb.Memory.Write(gb.CPU.hl(), val)
	case 0x02:
		// LD (BC),A
		val := gb.CPU.a()
		gb.Memory.Write(gb.CPU.bc(), val)
	case 0x12:
		// LD (DE),A
		val := gb.CPU.a()
		gb.Memory.Write(gb.CPU.de(), val)
	case 0xEA:
		// LD (nn),A
		val := gb.CPU.a()
		gb.Memory.Write(gb.popPC16(), val)
	case 0xF2:
		// LD A,(C)
		val := 0xFF00 + uint16(gb.CPU.c())
		gb.CPU.setA(gb.Memory.Read(val))
	case 0xE2:
		// LD (C),A
		val := gb.CPU.a()
		mem := 0xFF00 + uint16(gb.CPU.c())
		gb.Memory.Write(mem, val)
	case 0x3A:
		// LDD A,(HL)
		val := gb.Memory.Read(gb.CPU.hl())
		gb.CPU.setA(val)
		gb.CPU.setHl(gb.CPU.hl() - 1)
	case 0x32:
		// LDD (HL),A
		val := gb.CPU.hl()
		gb.Memory.Write(val, gb.CPU.a())
		gb.CPU.setHl(gb.CPU.hl() - 1)
	case 0x2A:
		// LDI A,(HL)
		val := gb.Memory.Read(gb.CPU.hl())
		gb.CPU.setA(val)
		gb.CPU.setHl(gb.CPU.hl() + 1)
	case 0x22:
		// LDI (HL),A
		val := gb.CPU.hl()
		gb.Memory.Write(val, gb.CPU.a())
		gb.CPU.setHl(gb.CPU.hl() + 1)
	case 0xE0:
		// LD (0xFF00+n),A
		val := 0xFF00 + uint16(gb.popPC())
		gb.Memory.Write(val, gb.CPU.a())
	case 0xF0:
		// LD A,(0xFF00+n)
		val := gb.Memory.ReadHighRam(0xFF00 + uint16(gb.popPC()))
		gb.CPU.setA(val)
	// ========== 16-Bit Loads ===========
	case 0x01:
		// LD BC,nn
		val := gb.popPC16()
		gb.CPU.setBc(val)
	case 0x11:
		// LD DE,nn
		val := gb.popPC16()
		gb.CPU.setDe(val)
	case 0x21:
		// LD HL,nn
		val := gb.popPC16()
		gb.CPU.setHl(val)
	case 0x31:
		// LD SP,nn
		val := gb.popPC16()
		gb.CPU.setSp(val)
	case 0xF9:
		// LD SP,HL
		val := gb.CPU.hl()
		gb.CPU.setSp(val)
	case 0xF8:
		// LD HL,SP+n
		gb.instAdd16Signed(gb.CPU.setHl, gb.CPU.sp(), int8(gb.popPC()))
	case 0x08:
		// LD (nn),SP
		address := gb.popPC16()
		gb.Memory.Write(address, gb.CPU.spLo())
		gb.Memory.Write(address+1, gb.CPU.spHi())
	case 0xF5:
		// PUSH AF
		gb.pushStack(gb.CPU.af())
	case 0xC5:
		// PUSH BC
		gb.pushStack(gb.CPU.bc())
	case 0xD5:
		// PUSH DE
		gb.pushStack(gb.CPU.de())
	case 0xE5:
		// PUSH HL
		gb.pushStack(gb.CPU.hl())
	case 0xF1:
		// POP AF
		gb.CPU.setAf(gb.popStack())
	case 0xC1:
		// POP BC
		gb.CPU.setBc(gb.popStack())
	case 0xD1:
		// POP DE
		gb.CPU.setDe(gb.popStack())
	case 0xE1:
		// POP HL
		gb.CPU.setHl(gb.popStack())
	// ========== 8-Bit ALU ===========
	case 0x87:
		// ADD A,A
		gb.instAdd(gb.CPU.setA, gb.CPU.a(), gb.CPU.a(), false)
	case 0x80:
		// ADD A,B
		gb.instAdd(gb.CPU.setA, gb.CPU.b(), gb.CPU.a(), false)
	case 0x81:
		// ADD A,C
		gb.instAdd(gb.CPU.setA, gb.CPU.c(), gb.CPU.a(), false)
	case 0x82:
		// ADD A,D
		gb.instAdd(gb.CPU.setA, gb.CPU.d(), gb.CPU.a(), false)
	case 0x83:
		// ADD A,E
		gb.instAdd(gb.CPU.setA, gb.CPU.e(), gb.CPU.a(), false)
	case 0x84:
		// ADD A,H
		gb.instAdd(gb.CPU.setA, gb.CPU.h(), gb.CPU.a(), false)
	case 0x85:
		// ADD A,L
		gb.instAdd(gb.CPU.setA, gb.CPU.l(), gb.CPU.a(), false)
	case 0x86:
		// ADD A,(HL)
		gb.instAdd(gb.CPU.setA, gb.Memory.Read(gb.CPU.hl()), gb.CPU.a(), false)
	case 0xC6:
		// ADD A,#
		gb.instAdd(gb.CPU.setA, gb.popPC(), gb.CPU.a(), false)
	case 0x8F:
		// ADC A,A
		gb.instAdd(gb.CPU.setA, gb.CPU.a(), gb.CPU.a(), true)
	case 0x88:
		// ADC A,B
		gb.instAdd(gb.CPU.setA, gb.CPU.b(), gb.CPU.a(), true)
	case 0x89:
		// ADC A,C
		gb.instAdd(gb.CPU.setA, gb.CPU.c(), gb.CPU.a(), true)
	case 0x8A:
		// ADC A,D
		gb.instAdd(gb.CPU.setA, gb.CPU.d(), gb.CPU.a(), true)
	case 0x8B:
		// ADC A,E
		gb.instAdd(gb.CPU.setA, gb.CPU.e(), gb.CPU.a(), true)
	case 0x8C:
		// ADC A,H
		gb.instAdd(gb.CPU.setA, gb.CPU.h(), gb.CPU.a(), true)
	case 0x8D:
		// ADC A,L
		gb.instAdd(gb.CPU.setA, gb.CPU.l(), gb.CPU.a(), true)
	case 0x8E:
		// ADC A,(HL)
		gb.instAdd(gb.CPU.setA, gb.Memory.Read(gb.CPU.hl()), gb.CPU.a(), true)
	case 0xCE:
		// ADC A,#
		gb.instAdd(gb.CPU.setA, gb.popPC(), gb.CPU.a(), true)
	case 0x97:
		// SUB A,A
		gb.instSub(gb.CPU.setA, gb.CPU.a(), gb.CPU.a(), false)
	case 0x90:
		// SUB A,B
		gb.instSub(gb.CPU.setA, gb.CPU.a(), gb.CPU.b(), false)
	case 0x91:
		// SUB A,C
		gb.instSub(gb.CPU.setA, gb.CPU.a(), gb.CPU.c(), false)
	case 0x92:
		// SUB A,D
		gb.instSub(gb.CPU.setA, gb.CPU.a(), gb.CPU.d(), false)
	case 0x93:
		// SUB A,E
		gb.instSub(gb.CPU.setA, gb.CPU.a(), gb.CPU.e(), false)
	case 0x94:
		// SUB A,H
		gb.instSub(gb.CPU.setA, gb.CPU.a(), gb.CPU.h(), false)
	case 0x95:
		// SUB A,L
		gb.instSub(gb.CPU.setA, gb.CPU.a(), gb.CPU.l(), false)
	case 0x96:
		// SUB A,(HL)
		gb.instSub(gb.CPU.setA, gb.CPU.a(), gb.Memory.Read(gb.CPU.hl()), false)
	case 0xD6:
		// SUB A,#
		gb.instSub(gb.CPU.setA, gb.CPU.a(), gb.popPC(), false)
	case 0x9F:
		// SBC A,A
		gb.instSub(gb.CPU.setA, gb.CPU.a(), gb.CPU.a(), true)
	case 0x98:
		// SBC A,B
		gb.instSub(gb.CPU.setA, gb.CPU.a(), gb.CPU.b(), true)
	case 0x99:
		// SBC A,C
		gb.instSub(gb.CPU.setA, gb.CPU.a(), gb.CPU.c(), true)
	case 0x9A:
		// SBC A,D
		gb.instSub(gb.CPU.setA, gb.CPU.a(), gb.CPU.d(), true)
	case 0x9B:
		// SBC A,E
		gb.instSub(gb.CPU.setA, gb.CPU.a(), gb.CPU.e(), true)
	case 0x9C:
		// SBC A,H
		gb.instSub(gb.CPU.setA, gb.CPU.a(), gb.CPU.h(), true)
	case 0x9D:
		// SBC A,L
		gb.instSub(gb.CPU.setA, gb.CPU.a(), gb.CPU.l(), true)
	case 0x9E:
		// SBC A,(HL)
		gb.instSub(gb.CPU.setA, gb.CPU.a(), gb.Memory.Read(gb.CPU.hl()), true)
	case 0xDE:
		// SBC A,#
		gb.instSub(gb.CPU.setA, gb.CPU.a(), gb.popPC(), true)
	case 0xA7:
		// AND A,A
		gb.instAnd(gb.CPU.setA, gb.CPU.a(), gb.CPU.a())
	case 0xA0:
		// AND A,B
		gb.instAnd(gb.CPU.setA, gb.CPU.b(), gb.CPU.a())
	case 0xA1:
		// AND A,C
		gb.instAnd(gb.CPU.setA, gb.CPU.c(), gb.CPU.a())
	case 0xA2:
		// AND A,D
		gb.instAnd(gb.CPU.setA, gb.CPU.d(), gb.CPU.a())
	case 0xA3:
		// AND A,E
		gb.instAnd(gb.CPU.setA, gb.CPU.e(), gb.CPU.a())
	case 0xA4:
		// AND A,H
		gb.instAnd(gb.CPU.setA, gb.CPU.h(), gb.CPU.a())
	case 0xA5:
		// AND A,L
		gb.instAnd(gb.CPU.setA, gb.CPU.l(), gb.CPU.a())
	case 0xA6:
		// AND A,(HL)
		gb.instAnd(gb.CPU.setA, gb.Memory.Read(gb.CPU.hl()), gb.CPU.a())
	case 0xE6:
		// AND A,#
		gb.instAnd(gb.CPU.setA, gb.popPC(), gb.CPU.a())
	case 0xB7:
		// OR A,A
		gb.instOr(gb.CPU.setA, gb.CPU.a(), gb.CPU.a())
	case 0xB0:
		// OR A,B
		gb.instOr(gb.CPU.setA, gb.CPU.b(), gb.CPU.a())
	case 0xB1:
		// OR A,C
		gb.instOr(gb.CPU.setA, gb.CPU.c(), gb.CPU.a())
	case 0xB2:
		// OR A,D
		gb.instOr(gb.CPU.setA, gb.CPU.d(), gb.CPU.a())
	case 0xB3:
		// OR A,E
		gb.instOr(gb.CPU.setA, gb.CPU.e(), gb.CPU.a())
	case 0xB4:
		// OR A,H
		gb.instOr(gb.CPU.setA, gb.CPU.h(), gb.CPU.a())
	case 0xB5:
		// OR A,L
		gb.instOr(gb.CPU.setA, gb.CPU.l(), gb.CPU.a())
	case 0xB6:
		// OR A,(HL)
		gb.instOr(gb.CPU.setA, gb.Memory.Read(gb.CPU.hl()), gb.CPU.a())
	case 0xF6:
		// OR A,#
		gb.instOr(gb.CPU.setA, gb.popPC(), gb.CPU.a())
	case 0xAF:
		// XOR A,A
		gb.instXor(gb.CPU.setA, gb.CPU.a(), gb.CPU.a())
	case 0xA8:
		// XOR A,B
		gb.instXor(gb.CPU.setA, gb.CPU.b(), gb.CPU.a())
	case 0xA9:
		// XOR A,C
		gb.instXor(gb.CPU.setA, gb.CPU.c(), gb.CPU.a())
	case 0xAA:
		// XOR A,D
		gb.instXor(gb.CPU.setA, gb.CPU.d(), gb.CPU.a())
	case 0xAB:
		// XOR A,E
		gb.instXor(gb.CPU.setA, gb.CPU.e(), gb.CPU.a())
	case 0xAC:
		// XOR A,H
		gb.instXor(gb.CPU.setA, gb.CPU.h(), gb.CPU.a())
	case 0xAD:
		// XOR A,L
		gb.instXor(gb.CPU.setA, gb.CPU.l(), gb.CPU.a())
	case 0xAE:
		// XOR A,(HL)
		gb.instXor(gb.CPU.setA, gb.Memory.Read(gb.CPU.hl()), gb.CPU.a())
	case 0xEE:
		// XOR A,#
		gb.instXor(gb.CPU.setA, gb.popPC(), gb.CPU.a())
	case 0xBF:
		// CP A,A
		gb.instCp(gb.CPU.a(), gb.CPU.a())
	case 0xB8:
		// CP A,B
		gb.instCp(gb.CPU.b(), gb.CPU.a())
	case 0xB9:
		// CP A,C
		gb.instCp(gb.CPU.c(), gb.CPU.a())
	case 0xBA:
		// CP A,D
		gb.instCp(gb.CPU.d(), gb.CPU.a())
	case 0xBB:
		// CP A,E
		gb.instCp(gb.CPU.e(), gb.CPU.a())
	case 0xBC:
		// CP A,H
		gb.instCp(gb.CPU.h(), gb.CPU.a())
	case 0xBD:
		// CP A,L
		gb.instCp(gb.CPU.l(), gb.CPU.a())
	case 0xBE:
		// CP A,(HL)
		gb.instCp(gb.Memory.Read(gb.CPU.hl()), gb.CPU.a())
	case 0xFE:
		// CP A,#
		gb.instCp(gb.popPC(), gb.CPU.a())
	case 0x3C:
		// INC A
		gb.instInc(gb.CPU.setA, gb.CPU.a())
	case 0x04:
		// INC B
		gb.instInc(gb.CPU.setB, gb.CPU.b())
	case 0x0C:
		// INC C
		gb.instInc(gb.CPU.setC, gb.CPU.c())
	case 0x14:
		// INC D
		gb.instInc(gb.CPU.setD, gb.CPU.d())
	case 0x1C:
		// INC E
		gb.instInc(gb.CPU.setE, gb.CPU.e())
	case 0x24:
		// INC H
		gb.instInc(gb.CPU.setH, gb.CPU.h())
	case 0x2C:
		// INC L
		gb.instInc(gb.CPU.setL, gb.CPU.l())
	case 0x34:
		// INC (HL)
		addr := gb.CPU.hl()
		gb.instInc(func(val byte) { gb.Memory.Write(addr, val) }, gb.Memory.Read(addr))
	case 0x3D:
		// DEC A
		gb.instDec(gb.CPU.setA, gb.CPU.a())
	case 0x05:
		// DEC B
		gb.instDec(gb.CPU.setB, gb.CPU.b())
	case 0x0D:
		// DEC C
		gb.instDec(gb.CPU.setC, gb.CPU.c())
	case 0x15:
		// DEC D
		gb.instDec(gb.CPU.setD, gb.CPU.d())
	case 0x1D:
		// DEC E
		gb.instDec(gb.CPU.setE, gb.CPU.e())
	case 0x25:
		// DEC H
		gb.instDec(gb.CPU.setH, gb.CPU.h())
	case 0x2D:
		// DEC L
		gb.instDec(gb.CPU.setL, gb.CPU.l())
	case 0x35:
		// DEC (HL)
		addr := gb.CPU.hl()
		gb.instDec(func(val byte) { gb.Memory.Write(addr, val) }, gb.Memory.Read(addr))
		// ========== 16-Bit ALU ===========
	case 0x09:
		// ADD HL,BC
		gb.instAdd16(gb.CPU.setHl, gb.CPU.hl(), gb.CPU.bc())
	case 0x19:
		// ADD HL,DE
		gb.instAdd16(gb.CPU.setHl, gb.CPU.hl(), gb.CPU.de())
	case 0x29:
		// ADD HL,HL
		gb.instAdd16(gb.CPU.setHl, gb.CPU.hl(), gb.CPU.hl())
	case 0x39:
		// ADD HL,SP
		gb.instAdd16(gb.CPU.setHl, gb.CPU.hl(), gb.CPU.sp())
	case 0xE8:
		// ADD SP,n
		gb.instAdd16Signed(gb.CPU.setSp, gb.CPU.sp(), int8(gb.popPC()))
		gb.CPU.setFlagZ(false)
	case 0x03:
		// INC BC
		gb.instInc16(gb.CPU.setBc, gb.CPU.bc())
	case 0x13:
		// INC DE
		gb.instInc16(gb.CPU.setDe, gb.CPU.de())
	case 0x23:
		// INC HL
		gb.instInc16(gb.CPU.setHl, gb.CPU.hl())
	case 0x33:
		// INC SP
		gb.instInc16(gb.CPU.setSp, gb.CPU.sp())
	case 0x0B:
		// DEC BC
		gb.instDec16(gb.CPU.setBc, gb.CPU.bc())
	case 0x1B:
		// DEC DE
		gb.instDec16(gb.CPU.setDe, gb.CPU.de())
	case 0x2B:
		// DEC HL
		gb.instDec16(gb.CPU.setHl, gb.CPU.hl())
	case 0x3B:
		// DEC SP
		gb.instDec16(gb.CPU.setSp, gb.CPU.sp())
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
		if !gb.CPU.flagN() {
			if gb.CPU.flagC() || gb.CPU.a() > 0x99 {
				gb.CPU.setA(gb.CPU.a() + 0x60)
				gb.CPU.setFlagC(true)
			}
			if gb.CPU.flagH() || gb.CPU.a()&0xF > 0x9 {
				gb.CPU.setA(gb.CPU.a() + 0x06)
				gb.CPU.setFlagH(false)
			}
		} else if gb.CPU.flagC() && gb.CPU.flagH() {
			gb.CPU.setA(gb.CPU.a() + 0x9A)
			gb.CPU.setFlagH(false)
		} else if gb.CPU.flagC() {
			gb.CPU.setA(gb.CPU.a() + 0xA0)
		} else if gb.CPU.flagH() {
			gb.CPU.setA(gb.CPU.a() + 0xFA)
			gb.CPU.setFlagH(false)
		}
		gb.CPU.setFlagZ(gb.CPU.a() == 0)
	case 0x2F:
		// CPL
		gb.CPU.setA(0xFF ^ gb.CPU.a())
		gb.CPU.setFlagN(true)
		gb.CPU.setFlagH(true)
	case 0x3F:
		// CCF
		gb.CPU.setFlagN(false)
		gb.CPU.setFlagH(false)
		gb.CPU.setFlagC(!gb.CPU.flagC())
	case 0x37:
		// SCF
		gb.CPU.setFlagN(false)
		gb.CPU.setFlagH(false)
		gb.CPU.setFlagC(true)
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
		gb.popPC()
	case 0xF3:
		// DI
		gb.interruptsOn = false
	case 0xFB:
		// EI
		gb.interruptsEnabling = true
	case 0x07:
		// RLCA
		value := gb.CPU.a()
		result := byte(value<<1) | (value >> 7)
		gb.CPU.setA(result)
		gb.CPU.setFlagZ(false)
		gb.CPU.setFlagN(false)
		gb.CPU.setFlagH(false)
		gb.CPU.setFlagC(value > 0x7F)
	case 0x17:
		// RLA
		value := gb.CPU.a()
		var carry byte
		if gb.CPU.flagC() {
			carry = 1
		}
		result := byte(value<<1) + carry
		gb.CPU.setA(result)
		gb.CPU.setFlagZ(false)
		gb.CPU.setFlagN(false)
		gb.CPU.setFlagH(false)
		gb.CPU.setFlagC(value > 0x7F)
	case 0x0F:
		// RRCA
		value := gb.CPU.a()
		result := byte(value>>1) | byte((value&1)<<7)
		gb.CPU.setA(result)
		gb.CPU.setFlagZ(false)
		gb.CPU.setFlagN(false)
		gb.CPU.setFlagH(false)
		gb.CPU.setFlagC(result > 0x7F)
	case 0x1F:
		// RRA
		value := gb.CPU.a()
		var carry byte
		if gb.CPU.flagC() {
			carry = 0x80
		}
		result := byte(value>>1) | carry
		gb.CPU.setA(result)
		gb.CPU.setFlagZ(false)
		gb.CPU.setFlagN(false)
		gb.CPU.setFlagH(false)
		gb.CPU.setFlagC((1 & value) == 1)
	case 0xC3:
		// JP nn
		gb.instJump(gb.popPC16())
	case 0xC2:
		// JP NZ,nn
		next := gb.popPC16()
		if !gb.CPU.flagZ() {
			gb.instJump(next)
			gb.thisCpuTicks += 4
		}
	case 0xCA:
		// JP Z,nn
		next := gb.popPC16()
		if gb.CPU.flagZ() {
			gb.instJump(next)
			gb.thisCpuTicks += 4
		}
	case 0xD2:
		// JP NC,nn
		next := gb.popPC16()
		if !gb.CPU.flagC() {
			gb.instJump(next)
			gb.thisCpuTicks += 4
		}
	case 0xDA:
		// JP C,nn
		next := gb.popPC16()
		if gb.CPU.flagC() {
			gb.instJump(next)
			gb.thisCpuTicks += 4
		}
	case 0xE9:
		// JP HL
		gb.instJump(gb.CPU.hl())
	case 0x18:
		// JR n
		addr := int32(gb.CPU.PC) + int32(int8(gb.popPC()))
		gb.instJump(uint16(addr))
	case 0x20:
		// JR NZ,n
		next := int8(gb.popPC())
		if !gb.CPU.flagZ() {
			addr := int32(gb.CPU.PC) + int32(next)
			gb.instJump(uint16(addr))
			gb.thisCpuTicks += 4
		}
	case 0x28:
		// JR Z,n
		next := int8(gb.popPC())
		if gb.CPU.flagZ() {
			addr := int32(gb.CPU.PC) + int32(next)
			gb.instJump(uint16(addr))
			gb.thisCpuTicks += 4
		}
	case 0x30:
		// JR NC,n
		next := int8(gb.popPC())
		if !gb.CPU.flagC() {
			addr := int32(gb.CPU.PC) + int32(next)
			gb.instJump(uint16(addr))
			gb.thisCpuTicks += 4
		}
	case 0x38:
		// JR C,n
		next := int8(gb.popPC())
		if gb.CPU.flagC() {
			addr := int32(gb.CPU.PC) + int32(next)
			gb.instJump(uint16(addr))
			gb.thisCpuTicks += 4
		}
	case 0xCD:
		// CALL nn
		gb.instCall(gb.popPC16())
	case 0xC4:
		// CALL NZ,nn
		next := gb.popPC16()
		if !gb.CPU.flagZ() {
			gb.instCall(next)
			gb.thisCpuTicks += 12
		}
	case 0xCC:
		// CALL Z,nn
		next := gb.popPC16()
		if gb.CPU.flagZ() {
			gb.instCall(next)
			gb.thisCpuTicks += 12
		}
	case 0xD4:
		// CALL NC,nn
		next := gb.popPC16()
		if !gb.CPU.flagC() {
			gb.instCall(next)
			gb.thisCpuTicks += 12
		}
	case 0xDC:
		// CALL C,nn
		next := gb.popPC16()
		if gb.CPU.flagC() {
			gb.instCall(next)
			gb.thisCpuTicks += 12
		}
	case 0xC7:
		// RST 0x00
		gb.instCall(0x0000)
	case 0xCF:
		// RST 0x08
		gb.instCall(0x0008)
	case 0xD7:
		// RST 0x10
		gb.instCall(0x0010)
	case 0xDF:
		// RST 0x18
		gb.instCall(0x0018)
	case 0xE7:
		// RST 0x20
		gb.instCall(0x0020)
	case 0xEF:
		// RST 0x28
		gb.instCall(0x0028)
	case 0xF7:
		// RST 0x30
		gb.instCall(0x0030)
	case 0xFF:
		// RST 0x38
		gb.instCall(0x0038)
	case 0xC9:
		// RET
		gb.instRet()
	case 0xC0:
		// RET NZ
		if !gb.CPU.flagZ() {
			gb.instRet()
			gb.thisCpuTicks += 12
		}
	case 0xC8:
		// RET Z
		if gb.CPU.flagZ() {
			gb.instRet()
			gb.thisCpuTicks += 12
		}
	case 0xD0:
		// RET NC
		if !gb.CPU.flagC() {
			gb.instRet()
			gb.thisCpuTicks += 12
		}
	case 0xD8:
		// RET C
		if gb.CPU.flagC() {
			gb.instRet()
			gb.thisCpuTicks += 12
		}
	case 0xD9:
		// RETI
		gb.instRet()
		gb.interruptsEnabling = true
	case 0xCB:
		// CB
		nextInst := gb.popPC()
		gb.thisCpuTicks += CBOpcodeCycles[nextInst] * 4
		gb.cbInst[nextInst]()
	default:
		log.Panicf("unknown opcode 0x%x", opcode)
	}
}

// Read the value at the PC and increment the PC.
func (gb *Gameboy) popPC() byte {
	opcode := gb.Memory.Read(gb.CPU.PC)
	gb.CPU.PC++
	return opcode
}

// Read the next 16bit value at the PC.
func (gb *Gameboy) popPC16() uint16 {
	b1 := uint16(gb.popPC())
	b2 := uint16(gb.popPC())
	return b2<<8 | b1
}
