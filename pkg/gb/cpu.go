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

func (cpu *CPU) setFlagZ(on bool) {
	cpu.setFlag(7, on)
}

func (cpu *CPU) setFlagN(on bool) {
	cpu.setFlag(6, on)
}

func (cpu *CPU) setFlagH(on bool) {
	cpu.setFlag(5, on)
}

func (cpu *CPU) setFlagC(on bool) {
	cpu.setFlag(4, on)
}

func (cpu *CPU) flagZ() bool {
	return cpu.getFlag(7)
}

func (cpu *CPU) flagN() bool {
	return cpu.getFlag(6)
}

func (cpu *CPU) flagH() bool {
	return cpu.getFlag(5)
}

func (cpu *CPU) flagC() bool {
	return cpu.getFlag(4)
}
