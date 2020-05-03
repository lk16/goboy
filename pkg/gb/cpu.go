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

const (
	// these are register indexes
	regBC = 0
	regDE = 1
	regHL = 2
	regAF = 3
	regSP = 4
)

// CPU contains the registers used for program execution and
// provides methods for setting flags.
type CPU struct {
	reg [5]register
	PC  uint16

	Divider int
}

// Init CPU and its registers to the initial values.
func (cpu *CPU) Init(cgb bool) {
	cpu.PC = 0x100
	if cgb {
		cpu.reg[regAF].Set(0x1180)
	} else {
		cpu.reg[regAF].Set(0x01B0)
	}
	cpu.reg[regBC].Set(0x0000)
	cpu.reg[regDE].Set(0xFF56)
	cpu.reg[regHL].Set(0x000D)
	cpu.reg[regSP].Set(0xFFFE)

	cpu.reg[regAF].mask = 0xFFF0
	cpu.reg[regBC].mask = 0xFFFF
	cpu.reg[regDE].mask = 0xFFFF
	cpu.reg[regHL].mask = 0xFFFF
	cpu.reg[regSP].mask = 0xFFFF
}

// Internally set the value of a flag on the flag register.
func (cpu *CPU) setFlag(index int, on bool) {
	mask := uint(1 << index)
	if on {
		cpu.reg[regAF].value |= mask
	} else {
		cpu.reg[regAF].value &^= mask
	}
	cpu.reg[regAF].updateMask()
}

func (cpu *CPU) getFlag(index int) bool {
	mask := uint(1 << index)
	return cpu.reg[regAF].value&mask != 0
}

func (cpu *CPU) setFlagC(on bool) { cpu.setFlag(4, on) }
func (cpu *CPU) setFlagH(on bool) { cpu.setFlag(5, on) }
func (cpu *CPU) setFlagN(on bool) { cpu.setFlag(6, on) }
func (cpu *CPU) setFlagZ(on bool) { cpu.setFlag(7, on) }

func (cpu *CPU) flagC() bool { return cpu.getFlag(4) }
func (cpu *CPU) flagH() bool { return cpu.getFlag(5) }
func (cpu *CPU) flagN() bool { return cpu.getFlag(6) }
func (cpu *CPU) flagZ() bool { return cpu.getFlag(7) }

func (cpu *CPU) setA(val byte) { cpu.reg[regAF].SetHi(val) }
func (cpu *CPU) setB(val byte) { cpu.reg[regBC].SetHi(val) }
func (cpu *CPU) setC(val byte) { cpu.reg[regBC].SetLo(val) }
func (cpu *CPU) setD(val byte) { cpu.reg[regDE].SetHi(val) }
func (cpu *CPU) setE(val byte) { cpu.reg[regDE].SetLo(val) }
func (cpu *CPU) setF(val byte) { cpu.reg[regAF].SetLo(val) }
func (cpu *CPU) setH(val byte) { cpu.reg[regHL].SetHi(val) }
func (cpu *CPU) setL(val byte) { cpu.reg[regHL].SetLo(val) }

func (cpu *CPU) a() byte    { return cpu.reg[regAF].Hi() }
func (cpu *CPU) b() byte    { return cpu.reg[regBC].Hi() }
func (cpu *CPU) c() byte    { return cpu.reg[regBC].Lo() }
func (cpu *CPU) d() byte    { return cpu.reg[regDE].Hi() }
func (cpu *CPU) e() byte    { return cpu.reg[regDE].Lo() }
func (cpu *CPU) f() byte    { return cpu.reg[regAF].Lo() }
func (cpu *CPU) h() byte    { return cpu.reg[regHL].Hi() }
func (cpu *CPU) l() byte    { return cpu.reg[regHL].Lo() }
func (cpu *CPU) spHi() byte { return cpu.reg[regSP].Hi() }
func (cpu *CPU) spLo() byte { return cpu.reg[regSP].Lo() }

func (cpu *CPU) af() uint16 { return cpu.reg[regAF].HiLo() }
func (cpu *CPU) bc() uint16 { return cpu.reg[regBC].HiLo() }
func (cpu *CPU) de() uint16 { return cpu.reg[regDE].HiLo() }
func (cpu *CPU) hl() uint16 { return cpu.reg[regHL].HiLo() }
func (cpu *CPU) sp() uint16 { return cpu.reg[regSP].HiLo() }

func (cpu *CPU) setAf(val uint16) { cpu.reg[regAF].Set(val) }
func (cpu *CPU) setBc(val uint16) { cpu.reg[regBC].Set(val) }
func (cpu *CPU) setDe(val uint16) { cpu.reg[regDE].Set(val) }
func (cpu *CPU) setHl(val uint16) { cpu.reg[regHL].Set(val) }
func (cpu *CPU) setSp(val uint16) { cpu.reg[regSP].Set(val) }
