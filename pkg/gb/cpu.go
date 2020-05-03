package gb

// Register represents a GB CPU 16bit register which provides functions
// for setting and getting the higher and lower bytes.
type register struct {
	// The value of the register.
	value uint

	// A mask over the possible values in the register.
	// Only used for the AF register where lower bits of
	// F cannot be set.
	mask uint
}

// Hi gets the higher byte of the register.
func (reg *register) Hi() byte {
	return byte(reg.value >> 8)
}

// Lo gets the lower byte of the register.
func (reg *register) Lo() byte {
	return byte(reg.value & 0xFF)
}

// HiLo gets the 2 byte value of the register.
func (reg *register) HiLo() uint16 {
	return uint16(reg.value)
}

// SetHi sets the higher byte of the register.
func (reg *register) SetHi(val byte) {
	reg.value = uint(val)<<8 | (reg.value & 0xFF)
}

// SetLog sets the lower byte of the register.
func (reg *register) SetLo(val byte) {
	reg.value = uint(val) | (reg.value & 0xFF00)
	reg.updateMask()
}

// Set the value of the register.
func (reg *register) Set(val uint16) {
	reg.value = uint(val)
	reg.updateMask()
}

// Mask the value if one is set on this register.
func (reg *register) updateMask() {
	reg.value &= reg.mask
}

// CPU contains the registers used for program execution and
// provides methods for setting flags.
type CPU struct {
	AF register
	BC register
	DE register
	HL register

	PC uint16
	SP register

	Divider int
}

// Init CPU and its registers to the initial values.
func (cpu *CPU) Init(cgb bool) {
	cpu.PC = 0x100
	if cgb {
		cpu.AF.Set(0x1180)
	} else {
		cpu.AF.Set(0x01B0)
	}
	cpu.BC.Set(0x0000)
	cpu.DE.Set(0xFF56)
	cpu.HL.Set(0x000D)
	cpu.SP.Set(0xFFFE)

	cpu.AF.mask = 0xFFF0
	cpu.BC.mask = 0xFFFF
	cpu.DE.mask = 0xFFFF
	cpu.HL.mask = 0xFFFF
	cpu.SP.mask = 0xFFFF
}

// Internally set the value of a flag on the flag register.
func (cpu *CPU) setFlag(index int, on bool) {
	mask := uint(1 << index)
	if on {
		cpu.AF.value |= mask
	} else {
		cpu.AF.value &^= mask
	}
	cpu.AF.updateMask()
}

func (cpu *CPU) getFlag(index int) bool {
	mask := uint(1 << index)
	return cpu.AF.value&mask != 0
}

func (cpu *CPU) setFlagC(on bool) { cpu.setFlag(4, on) }
func (cpu *CPU) setFlagH(on bool) { cpu.setFlag(5, on) }
func (cpu *CPU) setFlagN(on bool) { cpu.setFlag(6, on) }
func (cpu *CPU) setFlagZ(on bool) { cpu.setFlag(7, on) }

func (cpu *CPU) flagC() bool { return cpu.getFlag(4) }
func (cpu *CPU) flagH() bool { return cpu.getFlag(5) }
func (cpu *CPU) flagN() bool { return cpu.getFlag(6) }
func (cpu *CPU) flagZ() bool { return cpu.getFlag(7) }

func (cpu *CPU) setA(val byte) { cpu.AF.SetHi(val) }
func (cpu *CPU) setB(val byte) { cpu.BC.SetHi(val) }
func (cpu *CPU) setC(val byte) { cpu.BC.SetLo(val) }
func (cpu *CPU) setD(val byte) { cpu.DE.SetHi(val) }
func (cpu *CPU) setE(val byte) { cpu.DE.SetLo(val) }
func (cpu *CPU) setF(val byte) { cpu.AF.SetLo(val) }
func (cpu *CPU) setH(val byte) { cpu.HL.SetHi(val) }
func (cpu *CPU) setL(val byte) { cpu.HL.SetLo(val) }

func (cpu *CPU) a() byte { return cpu.AF.Hi() }
func (cpu *CPU) b() byte { return cpu.BC.Hi() }
func (cpu *CPU) c() byte { return cpu.BC.Lo() }
func (cpu *CPU) d() byte { return cpu.DE.Hi() }
func (cpu *CPU) e() byte { return cpu.DE.Lo() }
func (cpu *CPU) f() byte { return cpu.AF.Lo() }
func (cpu *CPU) h() byte { return cpu.HL.Hi() }
func (cpu *CPU) l() byte { return cpu.HL.Lo() }

func (cpu *CPU) af() uint16 { return cpu.AF.HiLo() }
func (cpu *CPU) bc() uint16 { return cpu.BC.HiLo() }
func (cpu *CPU) de() uint16 { return cpu.DE.HiLo() }
func (cpu *CPU) hl() uint16 { return cpu.HL.HiLo() }

func (cpu *CPU) setAf(val uint16) { cpu.AF.Set(val) }
func (cpu *CPU) setBc(val uint16) { cpu.BC.Set(val) }
func (cpu *CPU) setDe(val uint16) { cpu.DE.Set(val) }
func (cpu *CPU) setHl(val uint16) { cpu.HL.Set(val) }
