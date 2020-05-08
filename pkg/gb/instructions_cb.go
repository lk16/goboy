package gb

func (gb *Gameboy) instRlc(val uint) uint {
	carry := val >> 7
	rot := (val<<1)&0xFF | carry

	gb.CPU.setFlagZ(rot == 0)
	gb.CPU.setFlagN(false)
	gb.CPU.setFlagH(false)
	gb.CPU.setFlagC(carry == 1)
	return rot
}

func (gb *Gameboy) instRl(val uint) uint {
	newCarry := val >> 7
	oldCarry := gb.CPU.carryBit()
	rot := (val<<1)&0xFF | oldCarry

	gb.CPU.setFlagZ(rot == 0)
	gb.CPU.setFlagN(false)
	gb.CPU.setFlagH(false)
	gb.CPU.setFlagC(newCarry == 1)
	return rot
}

func (gb *Gameboy) instRrc(val uint) uint {
	carry := val & 1
	rot := ((val >> 1) | (carry << 7)) & 0xFF

	gb.CPU.setFlagZ(rot == 0)
	gb.CPU.setFlagN(false)
	gb.CPU.setFlagH(false)
	gb.CPU.setFlagC(carry == 1)
	return rot
}

func (gb *Gameboy) instRr(val uint) uint {
	newCarry := val & 1
	oldCarry := gb.CPU.carryBit()
	rot := (val >> 1) | (oldCarry << 7)

	gb.CPU.setFlagZ(rot == 0)
	gb.CPU.setFlagN(false)
	gb.CPU.setFlagH(false)
	gb.CPU.setFlagC(newCarry == 1)
	return rot
}

func (gb *Gameboy) instSla(val uint) uint {
	carry := val >> 7
	rot := (val << 1) & 0xFF

	gb.CPU.setFlagZ(rot == 0)
	gb.CPU.setFlagN(false)
	gb.CPU.setFlagH(false)
	gb.CPU.setFlagC(carry == 1)
	return rot
}

func (gb *Gameboy) instSra(val uint) uint {
	rot := (val & 0x80) | (val >> 1)

	gb.CPU.setFlagZ(rot == 0)
	gb.CPU.setFlagN(false)
	gb.CPU.setFlagH(false)
	gb.CPU.setFlagC(val&1 == 1)
	return rot
}

func (gb *Gameboy) instSrl(val uint) uint {
	carry := val & 1
	rot := val >> 1

	gb.CPU.setFlagZ(rot == 0)
	gb.CPU.setFlagN(false)
	gb.CPU.setFlagH(false)
	gb.CPU.setFlagC(carry == 1)
	return rot
}

func (gb *Gameboy) instBit(bit, val uint) {
	gb.CPU.setFlagZ((val>>bit)&1 == 0)
	gb.CPU.setFlagN(false)
	gb.CPU.setFlagH(true)
}

func (gb *Gameboy) instSwap(val uint) uint {
	swapped := val<<4&0xF0 | val>>4

	gb.CPU.setFlagZ(swapped == 0)
	gb.CPU.setFlagN(false)
	gb.CPU.setFlagH(false)
	gb.CPU.setFlagC(false)
	return swapped
}
