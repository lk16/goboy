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
	opcode := int(gb.popPC())
	gb.thisCpuTicks = OpcodeCycles[opcode] * 4
	gb.executeInstruction(opcode)
	return gb.thisCpuTicks
}

func (gb *Gameboy) executeInstruction(opcode int) {
	switch opcode {
	case 0x06:
		// LD B, n
		gb.CPU.BC.SetHi(gb.popPC())
	case 0x0E:
		// LD C, n
		gb.CPU.BC.SetLo(gb.popPC())
	case 0x16:
		// LD D, n
		gb.CPU.DE.SetHi(gb.popPC())
	case 0x1E:
		// LD E, n
		gb.CPU.DE.SetLo(gb.popPC())
	case 0x26:
		// LD H, n
		gb.CPU.HL.SetHi(gb.popPC())
	case 0x2E:
		// LD L, n
		gb.CPU.HL.SetLo(gb.popPC())
	case 0x7F:
		// LD A,A
		gb.CPU.AF.SetHi(gb.CPU.AF.Hi())
	case 0x78:
		// LD A,B
		gb.CPU.AF.SetHi(gb.CPU.BC.Hi())
	case 0x79:
		// LD A,C
		gb.CPU.AF.SetHi(gb.CPU.BC.Lo())
	case 0x7A:
		// LD A,D
		gb.CPU.AF.SetHi(gb.CPU.DE.Hi())
	case 0x7B:
		// LD A,E
		gb.CPU.AF.SetHi(gb.CPU.DE.Lo())
	case 0x7C:
		// LD A,H
		gb.CPU.AF.SetHi(gb.CPU.HL.Hi())
	case 0x7D:
		// LD A,L
		gb.CPU.AF.SetHi(gb.CPU.HL.Lo())
	case 0x0A:
		// LD A,(BC)
		val := gb.Memory.Read(gb.CPU.BC.HiLo())
		gb.CPU.AF.SetHi(val)
	case 0x1A:
		// LD A,(DE)
		val := gb.Memory.Read(gb.CPU.DE.HiLo())
		gb.CPU.AF.SetHi(val)
	case 0x7E:
		// LD A,(HL)
		val := gb.Memory.Read(gb.CPU.HL.HiLo())
		gb.CPU.AF.SetHi(val)
	case 0xFA:
		// LD A,(nn)
		val := gb.Memory.Read(gb.popPC16())
		gb.CPU.AF.SetHi(val)
	case 0x3E:
		// LD A,(nn)
		val := gb.popPC()
		gb.CPU.AF.SetHi(val)
	case 0x47:
		// LD B,A
		gb.CPU.BC.SetHi(gb.CPU.AF.Hi())
	case 0x40:
		// LD B,B
		gb.CPU.BC.SetHi(gb.CPU.BC.Hi())
	case 0x41:
		// LD B,C
		gb.CPU.BC.SetHi(gb.CPU.BC.Lo())
	case 0x42:
		// LD B,D
		gb.CPU.BC.SetHi(gb.CPU.DE.Hi())
	case 0x43:
		// LD B,E
		gb.CPU.BC.SetHi(gb.CPU.DE.Lo())
	case 0x44:
		// LD B,H
		gb.CPU.BC.SetHi(gb.CPU.HL.Hi())
	case 0x45:
		// LD B,L
		gb.CPU.BC.SetHi(gb.CPU.HL.Lo())
	case 0x46:
		// LD B,(HL)
		val := gb.Memory.Read(gb.CPU.HL.HiLo())
		gb.CPU.BC.SetHi(val)
	case 0x4F:
		// LD C,A
		gb.CPU.BC.SetLo(gb.CPU.AF.Hi())
	case 0x48:
		// LD C,B
		gb.CPU.BC.SetLo(gb.CPU.BC.Hi())
	case 0x49:
		// LD C,C
		gb.CPU.BC.SetLo(gb.CPU.BC.Lo())
	case 0x4A:
		// LD C,D
		gb.CPU.BC.SetLo(gb.CPU.DE.Hi())
	case 0x4B:
		// LD C,E
		gb.CPU.BC.SetLo(gb.CPU.DE.Lo())
	case 0x4C:
		// LD C,H
		gb.CPU.BC.SetLo(gb.CPU.HL.Hi())
	case 0x4D:
		// LD C,L
		gb.CPU.BC.SetLo(gb.CPU.HL.Lo())
	case 0x4E:
		// LD C,(HL)
		val := gb.Memory.Read(gb.CPU.HL.HiLo())
		gb.CPU.BC.SetLo(val)
	case 0x57:
		// LD D,A
		gb.CPU.DE.SetHi(gb.CPU.AF.Hi())
	case 0x50:
		// LD D,B
		gb.CPU.DE.SetHi(gb.CPU.BC.Hi())
	case 0x51:
		// LD D,C
		gb.CPU.DE.SetHi(gb.CPU.BC.Lo())
	case 0x52:
		// LD D,D
		gb.CPU.DE.SetHi(gb.CPU.DE.Hi())
	case 0x53:
		// LD D,E
		gb.CPU.DE.SetHi(gb.CPU.DE.Lo())
	case 0x54:
		// LD D,H
		gb.CPU.DE.SetHi(gb.CPU.HL.Hi())
	case 0x55:
		// LD D,L
		gb.CPU.DE.SetHi(gb.CPU.HL.Lo())
	case 0x56:
		// LD D,(HL)
		val := gb.Memory.Read(gb.CPU.HL.HiLo())
		gb.CPU.DE.SetHi(val)
	case 0x5F:
		// LD E,A
		gb.CPU.DE.SetLo(gb.CPU.AF.Hi())
	case 0x58:
		// LD E,B
		gb.CPU.DE.SetLo(gb.CPU.BC.Hi())
	case 0x59:
		// LD E,C
		gb.CPU.DE.SetLo(gb.CPU.BC.Lo())
	case 0x5A:
		// LD E,D
		gb.CPU.DE.SetLo(gb.CPU.DE.Hi())
	case 0x5B:
		// LD E,E
		gb.CPU.DE.SetLo(gb.CPU.DE.Lo())
	case 0x5C:
		// LD E,H
		gb.CPU.DE.SetLo(gb.CPU.HL.Hi())
	case 0x5D:
		// LD E,L
		gb.CPU.DE.SetLo(gb.CPU.HL.Lo())
	case 0x5E:
		// LD E,(HL)
		val := gb.Memory.Read(gb.CPU.HL.HiLo())
		gb.CPU.DE.SetLo(val)
	case 0x67:
		// LD H,A
		gb.CPU.HL.SetHi(gb.CPU.AF.Hi())
	case 0x60:
		// LD H,B
		gb.CPU.HL.SetHi(gb.CPU.BC.Hi())
	case 0x61:
		// LD H,C
		gb.CPU.HL.SetHi(gb.CPU.BC.Lo())
	case 0x62:
		// LD H,D
		gb.CPU.HL.SetHi(gb.CPU.DE.Hi())
	case 0x63:
		// LD H,E
		gb.CPU.HL.SetHi(gb.CPU.DE.Lo())
	case 0x64:
		// LD H,H
		gb.CPU.HL.SetHi(gb.CPU.HL.Hi())
	case 0x65:
		// LD H,L
		gb.CPU.HL.SetHi(gb.CPU.HL.Lo())
	case 0x66:
		// LD H,(HL)
		val := gb.Memory.Read(gb.CPU.HL.HiLo())
		gb.CPU.HL.SetHi(val)
	case 0x6F:
		// LD L,A
		gb.CPU.HL.SetLo(gb.CPU.AF.Hi())
	case 0x68:
		// LD L,B
		gb.CPU.HL.SetLo(gb.CPU.BC.Hi())
	case 0x69:
		// LD L,C
		gb.CPU.HL.SetLo(gb.CPU.BC.Lo())
	case 0x6A:
		// LD L,D
		gb.CPU.HL.SetLo(gb.CPU.DE.Hi())
	case 0x6B:
		// LD L,E
		gb.CPU.HL.SetLo(gb.CPU.DE.Lo())
	case 0x6C:
		// LD L,H
		gb.CPU.HL.SetLo(gb.CPU.HL.Hi())
	case 0x6D:
		// LD L,L
		gb.CPU.HL.SetLo(gb.CPU.HL.Lo())
	case 0x6E:
		// LD L,(HL)
		val := gb.Memory.Read(gb.CPU.HL.HiLo())
		gb.CPU.HL.SetLo(val)
	case 0x77:
		// LD (HL),A
		val := gb.CPU.AF.Hi()
		gb.Memory.Write(gb.CPU.HL.HiLo(), val)
	case 0x70:
		// LD (HL),B
		val := gb.CPU.BC.Hi()
		gb.Memory.Write(gb.CPU.HL.HiLo(), val)
	case 0x71:
		// LD (HL),C
		val := gb.CPU.BC.Lo()
		gb.Memory.Write(gb.CPU.HL.HiLo(), val)
	case 0x72:
		// LD (HL),D
		val := gb.CPU.DE.Hi()
		gb.Memory.Write(gb.CPU.HL.HiLo(), val)
	case 0x73:
		// LD (HL),E
		val := gb.CPU.DE.Lo()
		gb.Memory.Write(gb.CPU.HL.HiLo(), val)
	case 0x74:
		// LD (HL),H
		val := gb.CPU.HL.Hi()
		gb.Memory.Write(gb.CPU.HL.HiLo(), val)
	case 0x75:
		// LD (HL),L
		val := gb.CPU.HL.Lo()
		gb.Memory.Write(gb.CPU.HL.HiLo(), val)
	case 0x36:
		// LD (HL),n 36
		val := gb.popPC()
		gb.Memory.Write(gb.CPU.HL.HiLo(), val)
	case 0x02:
		// LD (BC),A
		val := gb.CPU.AF.Hi()
		gb.Memory.Write(gb.CPU.BC.HiLo(), val)
	case 0x12:
		// LD (DE),A
		val := gb.CPU.AF.Hi()
		gb.Memory.Write(gb.CPU.DE.HiLo(), val)
	case 0xEA:
		// LD (nn),A
		val := gb.CPU.AF.Hi()
		gb.Memory.Write(gb.popPC16(), val)
	case 0xF2:
		// LD A,(C)
		val := 0xFF00 + uint16(gb.CPU.BC.Lo())
		gb.CPU.AF.SetHi(gb.Memory.Read(val))
	case 0xE2:
		// LD (C),A
		val := gb.CPU.AF.Hi()
		mem := 0xFF00 + uint16(gb.CPU.BC.Lo())
		gb.Memory.Write(mem, val)
	case 0x3A:
		// LDD A,(HL)
		val := gb.Memory.Read(gb.CPU.HL.HiLo())
		gb.CPU.AF.SetHi(val)
		gb.CPU.HL.Set(gb.CPU.HL.HiLo() - 1)
	case 0x32:
		// LDD (HL),A
		val := gb.CPU.HL.HiLo()
		gb.Memory.Write(val, gb.CPU.AF.Hi())
		gb.CPU.HL.Set(gb.CPU.HL.HiLo() - 1)
	case 0x2A:
		// LDI A,(HL)
		val := gb.Memory.Read(gb.CPU.HL.HiLo())
		gb.CPU.AF.SetHi(val)
		gb.CPU.HL.Set(gb.CPU.HL.HiLo() + 1)
	case 0x22:
		// LDI (HL),A
		val := gb.CPU.HL.HiLo()
		gb.Memory.Write(val, gb.CPU.AF.Hi())
		gb.CPU.HL.Set(gb.CPU.HL.HiLo() + 1)
	case 0xE0:
		// LD (0xFF00+n),A
		val := 0xFF00 + uint16(gb.popPC())
		gb.Memory.Write(val, gb.CPU.AF.Hi())
	case 0xF0:
		// LD A,(0xFF00+n)
		val := gb.Memory.ReadHighRam(0xFF00 + uint16(gb.popPC()))
		gb.CPU.AF.SetHi(val)
	// ========== 16-Bit Loads ===========
	case 0x01:
		// LD BC,nn
		val := gb.popPC16()
		gb.CPU.BC.Set(val)
	case 0x11:
		// LD DE,nn
		val := gb.popPC16()
		gb.CPU.DE.Set(val)
	case 0x21:
		// LD HL,nn
		val := gb.popPC16()
		gb.CPU.HL.Set(val)
	case 0x31:
		// LD SP,nn
		val := gb.popPC16()
		gb.CPU.SP.Set(val)
	case 0xF9:
		// LD SP,HL
		val := gb.CPU.HL
		gb.CPU.SP = val
	case 0xF8:
		// LD HL,SP+n
		gb.instAdd16Signed(gb.CPU.HL.Set, gb.CPU.SP.HiLo(), int8(gb.popPC()))
	case 0x08:
		// LD (nn),SP
		address := gb.popPC16()
		gb.Memory.Write(address, gb.CPU.SP.Lo())
		gb.Memory.Write(address+1, gb.CPU.SP.Hi())
	case 0xF5:
		// PUSH AF
		gb.pushStack(gb.CPU.AF.HiLo())
	case 0xC5:
		// PUSH BC
		gb.pushStack(gb.CPU.BC.HiLo())
	case 0xD5:
		// PUSH DE
		gb.pushStack(gb.CPU.DE.HiLo())
	case 0xE5:
		// PUSH HL
		gb.pushStack(gb.CPU.HL.HiLo())
	case 0xF1:
		// POP AF
		gb.CPU.AF.Set(gb.popStack())
	case 0xC1:
		// POP BC
		gb.CPU.BC.Set(gb.popStack())
	case 0xD1:
		// POP DE
		gb.CPU.DE.Set(gb.popStack())
	case 0xE1:
		// POP HL
		gb.CPU.HL.Set(gb.popStack())
	// ========== 8-Bit ALU ===========
	case 0x87:
		// ADD A,A
		gb.instAdd(gb.CPU.AF.SetHi, gb.CPU.AF.Hi(), gb.CPU.AF.Hi(), false)
	case 0x80:
		// ADD A,B
		gb.instAdd(gb.CPU.AF.SetHi, gb.CPU.BC.Hi(), gb.CPU.AF.Hi(), false)
	case 0x81:
		// ADD A,C
		gb.instAdd(gb.CPU.AF.SetHi, gb.CPU.BC.Lo(), gb.CPU.AF.Hi(), false)
	case 0x82:
		// ADD A,D
		gb.instAdd(gb.CPU.AF.SetHi, gb.CPU.DE.Hi(), gb.CPU.AF.Hi(), false)
	case 0x83:
		// ADD A,E
		gb.instAdd(gb.CPU.AF.SetHi, gb.CPU.DE.Lo(), gb.CPU.AF.Hi(), false)
	case 0x84:
		// ADD A,H
		gb.instAdd(gb.CPU.AF.SetHi, gb.CPU.HL.Hi(), gb.CPU.AF.Hi(), false)
	case 0x85:
		// ADD A,L
		gb.instAdd(gb.CPU.AF.SetHi, gb.CPU.HL.Lo(), gb.CPU.AF.Hi(), false)
	case 0x86:
		// ADD A,(HL)
		gb.instAdd(gb.CPU.AF.SetHi, gb.Memory.Read(gb.CPU.HL.HiLo()), gb.CPU.AF.Hi(), false)
	case 0xC6:
		// ADD A,#
		gb.instAdd(gb.CPU.AF.SetHi, gb.popPC(), gb.CPU.AF.Hi(), false)
	case 0x8F:
		// ADC A,A
		gb.instAdd(gb.CPU.AF.SetHi, gb.CPU.AF.Hi(), gb.CPU.AF.Hi(), true)
	case 0x88:
		// ADC A,B
		gb.instAdd(gb.CPU.AF.SetHi, gb.CPU.BC.Hi(), gb.CPU.AF.Hi(), true)
	case 0x89:
		// ADC A,C
		gb.instAdd(gb.CPU.AF.SetHi, gb.CPU.BC.Lo(), gb.CPU.AF.Hi(), true)
	case 0x8A:
		// ADC A,D
		gb.instAdd(gb.CPU.AF.SetHi, gb.CPU.DE.Hi(), gb.CPU.AF.Hi(), true)
	case 0x8B:
		// ADC A,E
		gb.instAdd(gb.CPU.AF.SetHi, gb.CPU.DE.Lo(), gb.CPU.AF.Hi(), true)
	case 0x8C:
		// ADC A,H
		gb.instAdd(gb.CPU.AF.SetHi, gb.CPU.HL.Hi(), gb.CPU.AF.Hi(), true)
	case 0x8D:
		// ADC A,L
		gb.instAdd(gb.CPU.AF.SetHi, gb.CPU.HL.Lo(), gb.CPU.AF.Hi(), true)
	case 0x8E:
		// ADC A,(HL)
		gb.instAdd(gb.CPU.AF.SetHi, gb.Memory.Read(gb.CPU.HL.HiLo()), gb.CPU.AF.Hi(), true)
	case 0xCE:
		// ADC A,#
		gb.instAdd(gb.CPU.AF.SetHi, gb.popPC(), gb.CPU.AF.Hi(), true)
	case 0x97:
		// SUB A,A
		gb.instSub(gb.CPU.AF.SetHi, gb.CPU.AF.Hi(), gb.CPU.AF.Hi(), false)
	case 0x90:
		// SUB A,B
		gb.instSub(gb.CPU.AF.SetHi, gb.CPU.AF.Hi(), gb.CPU.BC.Hi(), false)
	case 0x91:
		// SUB A,C
		gb.instSub(gb.CPU.AF.SetHi, gb.CPU.AF.Hi(), gb.CPU.BC.Lo(), false)
	case 0x92:
		// SUB A,D
		gb.instSub(gb.CPU.AF.SetHi, gb.CPU.AF.Hi(), gb.CPU.DE.Hi(), false)
	case 0x93:
		// SUB A,E
		gb.instSub(gb.CPU.AF.SetHi, gb.CPU.AF.Hi(), gb.CPU.DE.Lo(), false)
	case 0x94:
		// SUB A,H
		gb.instSub(gb.CPU.AF.SetHi, gb.CPU.AF.Hi(), gb.CPU.HL.Hi(), false)
	case 0x95:
		// SUB A,L
		gb.instSub(gb.CPU.AF.SetHi, gb.CPU.AF.Hi(), gb.CPU.HL.Lo(), false)
	case 0x96:
		// SUB A,(HL)
		gb.instSub(gb.CPU.AF.SetHi, gb.CPU.AF.Hi(), gb.Memory.Read(gb.CPU.HL.HiLo()), false)
	case 0xD6:
		// SUB A,#
		gb.instSub(gb.CPU.AF.SetHi, gb.CPU.AF.Hi(), gb.popPC(), false)
	case 0x9F:
		// SBC A,A
		gb.instSub(gb.CPU.AF.SetHi, gb.CPU.AF.Hi(), gb.CPU.AF.Hi(), true)
	case 0x98:
		// SBC A,B
		gb.instSub(gb.CPU.AF.SetHi, gb.CPU.AF.Hi(), gb.CPU.BC.Hi(), true)
	case 0x99:
		// SBC A,C
		gb.instSub(gb.CPU.AF.SetHi, gb.CPU.AF.Hi(), gb.CPU.BC.Lo(), true)
	case 0x9A:
		// SBC A,D
		gb.instSub(gb.CPU.AF.SetHi, gb.CPU.AF.Hi(), gb.CPU.DE.Hi(), true)
	case 0x9B:
		// SBC A,E
		gb.instSub(gb.CPU.AF.SetHi, gb.CPU.AF.Hi(), gb.CPU.DE.Lo(), true)
	case 0x9C:
		// SBC A,H
		gb.instSub(gb.CPU.AF.SetHi, gb.CPU.AF.Hi(), gb.CPU.HL.Hi(), true)
	case 0x9D:
		// SBC A,L
		gb.instSub(gb.CPU.AF.SetHi, gb.CPU.AF.Hi(), gb.CPU.HL.Lo(), true)
	case 0x9E:
		// SBC A,(HL)
		gb.instSub(gb.CPU.AF.SetHi, gb.CPU.AF.Hi(), gb.Memory.Read(gb.CPU.HL.HiLo()), true)
	case 0xDE:
		// SBC A,#
		gb.instSub(gb.CPU.AF.SetHi, gb.CPU.AF.Hi(), gb.popPC(), true)
	case 0xA7:
		// AND A,A
		gb.instAnd(gb.CPU.AF.SetHi, gb.CPU.AF.Hi(), gb.CPU.AF.Hi())
	case 0xA0:
		// AND A,B
		gb.instAnd(gb.CPU.AF.SetHi, gb.CPU.BC.Hi(), gb.CPU.AF.Hi())
	case 0xA1:
		// AND A,C
		gb.instAnd(gb.CPU.AF.SetHi, gb.CPU.BC.Lo(), gb.CPU.AF.Hi())
	case 0xA2:
		// AND A,D
		gb.instAnd(gb.CPU.AF.SetHi, gb.CPU.DE.Hi(), gb.CPU.AF.Hi())
	case 0xA3:
		// AND A,E
		gb.instAnd(gb.CPU.AF.SetHi, gb.CPU.DE.Lo(), gb.CPU.AF.Hi())
	case 0xA4:
		// AND A,H
		gb.instAnd(gb.CPU.AF.SetHi, gb.CPU.HL.Hi(), gb.CPU.AF.Hi())
	case 0xA5:
		// AND A,L
		gb.instAnd(gb.CPU.AF.SetHi, gb.CPU.HL.Lo(), gb.CPU.AF.Hi())
	case 0xA6:
		// AND A,(HL)
		gb.instAnd(gb.CPU.AF.SetHi, gb.Memory.Read(gb.CPU.HL.HiLo()), gb.CPU.AF.Hi())
	case 0xE6:
		// AND A,#
		gb.instAnd(gb.CPU.AF.SetHi, gb.popPC(), gb.CPU.AF.Hi())
	case 0xB7:
		// OR A,A
		gb.instOr(gb.CPU.AF.SetHi, gb.CPU.AF.Hi(), gb.CPU.AF.Hi())
	case 0xB0:
		// OR A,B
		gb.instOr(gb.CPU.AF.SetHi, gb.CPU.BC.Hi(), gb.CPU.AF.Hi())
	case 0xB1:
		// OR A,C
		gb.instOr(gb.CPU.AF.SetHi, gb.CPU.BC.Lo(), gb.CPU.AF.Hi())
	case 0xB2:
		// OR A,D
		gb.instOr(gb.CPU.AF.SetHi, gb.CPU.DE.Hi(), gb.CPU.AF.Hi())
	case 0xB3:
		// OR A,E
		gb.instOr(gb.CPU.AF.SetHi, gb.CPU.DE.Lo(), gb.CPU.AF.Hi())
	case 0xB4:
		// OR A,H
		gb.instOr(gb.CPU.AF.SetHi, gb.CPU.HL.Hi(), gb.CPU.AF.Hi())
	case 0xB5:
		// OR A,L
		gb.instOr(gb.CPU.AF.SetHi, gb.CPU.HL.Lo(), gb.CPU.AF.Hi())
	case 0xB6:
		// OR A,(HL)
		gb.instOr(gb.CPU.AF.SetHi, gb.Memory.Read(gb.CPU.HL.HiLo()), gb.CPU.AF.Hi())
	case 0xF6:
		// OR A,#
		gb.instOr(gb.CPU.AF.SetHi, gb.popPC(), gb.CPU.AF.Hi())
	case 0xAF:
		// XOR A,A
		gb.instXor(gb.CPU.AF.SetHi, gb.CPU.AF.Hi(), gb.CPU.AF.Hi())
	case 0xA8:
		// XOR A,B
		gb.instXor(gb.CPU.AF.SetHi, gb.CPU.BC.Hi(), gb.CPU.AF.Hi())
	case 0xA9:
		// XOR A,C
		gb.instXor(gb.CPU.AF.SetHi, gb.CPU.BC.Lo(), gb.CPU.AF.Hi())
	case 0xAA:
		// XOR A,D
		gb.instXor(gb.CPU.AF.SetHi, gb.CPU.DE.Hi(), gb.CPU.AF.Hi())
	case 0xAB:
		// XOR A,E
		gb.instXor(gb.CPU.AF.SetHi, gb.CPU.DE.Lo(), gb.CPU.AF.Hi())
	case 0xAC:
		// XOR A,H
		gb.instXor(gb.CPU.AF.SetHi, gb.CPU.HL.Hi(), gb.CPU.AF.Hi())
	case 0xAD:
		// XOR A,L
		gb.instXor(gb.CPU.AF.SetHi, gb.CPU.HL.Lo(), gb.CPU.AF.Hi())
	case 0xAE:
		// XOR A,(HL)
		gb.instXor(gb.CPU.AF.SetHi, gb.Memory.Read(gb.CPU.HL.HiLo()), gb.CPU.AF.Hi())
	case 0xEE:
		// XOR A,#
		gb.instXor(gb.CPU.AF.SetHi, gb.popPC(), gb.CPU.AF.Hi())
	case 0xBF:
		// CP A,A
		gb.instCp(gb.CPU.AF.Hi(), gb.CPU.AF.Hi())
	case 0xB8:
		// CP A,B
		gb.instCp(gb.CPU.BC.Hi(), gb.CPU.AF.Hi())
	case 0xB9:
		// CP A,C
		gb.instCp(gb.CPU.BC.Lo(), gb.CPU.AF.Hi())
	case 0xBA:
		// CP A,D
		gb.instCp(gb.CPU.DE.Hi(), gb.CPU.AF.Hi())
	case 0xBB:
		// CP A,E
		gb.instCp(gb.CPU.DE.Lo(), gb.CPU.AF.Hi())
	case 0xBC:
		// CP A,H
		gb.instCp(gb.CPU.HL.Hi(), gb.CPU.AF.Hi())
	case 0xBD:
		// CP A,L
		gb.instCp(gb.CPU.HL.Lo(), gb.CPU.AF.Hi())
	case 0xBE:
		// CP A,(HL)
		gb.instCp(gb.Memory.Read(gb.CPU.HL.HiLo()), gb.CPU.AF.Hi())
	case 0xFE:
		// CP A,#
		gb.instCp(gb.popPC(), gb.CPU.AF.Hi())
	case 0x3C:
		// INC A
		gb.instInc(gb.CPU.AF.SetHi, gb.CPU.AF.Hi())
	case 0x04:
		// INC B
		gb.instInc(gb.CPU.BC.SetHi, gb.CPU.BC.Hi())
	case 0x0C:
		// INC C
		gb.instInc(gb.CPU.BC.SetLo, gb.CPU.BC.Lo())
	case 0x14:
		// INC D
		gb.instInc(gb.CPU.DE.SetHi, gb.CPU.DE.Hi())
	case 0x1C:
		// INC E
		gb.instInc(gb.CPU.DE.SetLo, gb.CPU.DE.Lo())
	case 0x24:
		// INC H
		gb.instInc(gb.CPU.HL.SetHi, gb.CPU.HL.Hi())
	case 0x2C:
		// INC L
		gb.instInc(gb.CPU.HL.SetLo, gb.CPU.HL.Lo())
	case 0x34:
		// INC (HL)
		addr := gb.CPU.HL.HiLo()
		gb.instInc(func(val byte) { gb.Memory.Write(addr, val) }, gb.Memory.Read(addr))
	case 0x3D:
		// DEC A
		gb.instDec(gb.CPU.AF.SetHi, gb.CPU.AF.Hi())
	case 0x05:
		// DEC B
		gb.instDec(gb.CPU.BC.SetHi, gb.CPU.BC.Hi())
	case 0x0D:
		// DEC C
		gb.instDec(gb.CPU.BC.SetLo, gb.CPU.BC.Lo())
	case 0x15:
		// DEC D
		gb.instDec(gb.CPU.DE.SetHi, gb.CPU.DE.Hi())
	case 0x1D:
		// DEC E
		gb.instDec(gb.CPU.DE.SetLo, gb.CPU.DE.Lo())
	case 0x25:
		// DEC H
		gb.instDec(gb.CPU.HL.SetHi, gb.CPU.HL.Hi())
	case 0x2D:
		// DEC L
		gb.instDec(gb.CPU.HL.SetLo, gb.CPU.HL.Lo())
	case 0x35:
		// DEC (HL)
		addr := gb.CPU.HL.HiLo()
		gb.instDec(func(val byte) { gb.Memory.Write(addr, val) }, gb.Memory.Read(addr))
		// ========== 16-Bit ALU ===========
	case 0x09:
		// ADD HL,BC
		gb.instAdd16(gb.CPU.HL.Set, gb.CPU.HL.HiLo(), gb.CPU.BC.HiLo())
	case 0x19:
		// ADD HL,DE
		gb.instAdd16(gb.CPU.HL.Set, gb.CPU.HL.HiLo(), gb.CPU.DE.HiLo())
	case 0x29:
		// ADD HL,HL
		gb.instAdd16(gb.CPU.HL.Set, gb.CPU.HL.HiLo(), gb.CPU.HL.HiLo())
	case 0x39:
		// ADD HL,SP
		gb.instAdd16(gb.CPU.HL.Set, gb.CPU.HL.HiLo(), gb.CPU.SP.HiLo())
	case 0xE8:
		// ADD SP,n
		gb.instAdd16Signed(gb.CPU.SP.Set, gb.CPU.SP.HiLo(), int8(gb.popPC()))
		gb.CPU.setFlagZ(false)
	case 0x03:
		// INC BC
		gb.instInc16(gb.CPU.BC.Set, gb.CPU.BC.HiLo())
	case 0x13:
		// INC DE
		gb.instInc16(gb.CPU.DE.Set, gb.CPU.DE.HiLo())
	case 0x23:
		// INC HL
		gb.instInc16(gb.CPU.HL.Set, gb.CPU.HL.HiLo())
	case 0x33:
		// INC SP
		gb.instInc16(gb.CPU.SP.Set, gb.CPU.SP.HiLo())
	case 0x0B:
		// DEC BC
		gb.instDec16(gb.CPU.BC.Set, gb.CPU.BC.HiLo())
	case 0x1B:
		// DEC DE
		gb.instDec16(gb.CPU.DE.Set, gb.CPU.DE.HiLo())
	case 0x2B:
		// DEC HL
		gb.instDec16(gb.CPU.HL.Set, gb.CPU.HL.HiLo())
	case 0x3B:
		// DEC SP
		gb.instDec16(gb.CPU.SP.Set, gb.CPU.SP.HiLo())
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
			if gb.CPU.flagC() || gb.CPU.AF.Hi() > 0x99 {
				gb.CPU.AF.SetHi(gb.CPU.AF.Hi() + 0x60)
				gb.CPU.setFlagC(true)
			}
			if gb.CPU.flagH() || gb.CPU.AF.Hi()&0xF > 0x9 {
				gb.CPU.AF.SetHi(gb.CPU.AF.Hi() + 0x06)
				gb.CPU.setFlagH(false)
			}
		} else if gb.CPU.flagC() && gb.CPU.flagH() {
			gb.CPU.AF.SetHi(gb.CPU.AF.Hi() + 0x9A)
			gb.CPU.setFlagH(false)
		} else if gb.CPU.flagC() {
			gb.CPU.AF.SetHi(gb.CPU.AF.Hi() + 0xA0)
		} else if gb.CPU.flagH() {
			gb.CPU.AF.SetHi(gb.CPU.AF.Hi() + 0xFA)
			gb.CPU.setFlagH(false)
		}
		gb.CPU.setFlagZ(gb.CPU.AF.Hi() == 0)
	case 0x2F:
		// CPL
		gb.CPU.AF.SetHi(0xFF ^ gb.CPU.AF.Hi())
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
		value := gb.CPU.AF.Hi()
		result := byte(value<<1) | (value >> 7)
		gb.CPU.AF.SetHi(result)
		gb.CPU.setFlagZ(false)
		gb.CPU.setFlagN(false)
		gb.CPU.setFlagH(false)
		gb.CPU.setFlagC(value > 0x7F)
	case 0x17:
		// RLA
		value := gb.CPU.AF.Hi()
		var carry byte
		if gb.CPU.flagC() {
			carry = 1
		}
		result := byte(value<<1) + carry
		gb.CPU.AF.SetHi(result)
		gb.CPU.setFlagZ(false)
		gb.CPU.setFlagN(false)
		gb.CPU.setFlagH(false)
		gb.CPU.setFlagC(value > 0x7F)
	case 0x0F:
		// RRCA
		value := gb.CPU.AF.Hi()
		result := byte(value>>1) | byte((value&1)<<7)
		gb.CPU.AF.SetHi(result)
		gb.CPU.setFlagZ(false)
		gb.CPU.setFlagN(false)
		gb.CPU.setFlagH(false)
		gb.CPU.setFlagC(result > 0x7F)
	case 0x1F:
		// RRA
		value := gb.CPU.AF.Hi()
		var carry byte
		if gb.CPU.flagC() {
			carry = 0x80
		}
		result := byte(value>>1) | carry
		gb.CPU.AF.SetHi(result)
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
		gb.instJump(gb.CPU.HL.HiLo())
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
