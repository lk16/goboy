package gb

// Register represents a GB CPU 16bit register which provides functions
// for setting and getting the higher and lower bytes.
type register struct {
	// The value of the register.
	value uint
}

// Hi gets the higher byte of the register.
func (reg *register) hi() byte {
	return byte(reg.value >> 8)
}

// Lo gets the lower byte of the register.
func (reg *register) lo() byte {
	return byte(reg.value & 0xFF)
}

// HiLo gets the 2 byte value of the register.
func (reg *register) hiLo() uint16 {
	return uint16(reg.value)
}

// setHi sets the higher byte of the register.
func (reg *register) sethi(val byte) {
	reg.value = uint(val)<<8 | (reg.value & 0xFF)
}

// setLog sets the lower byte of the register.
func (reg *register) setLo(val byte) {
	reg.value = uint(val) | (reg.value & 0xFF00)
}

// set the value of the register.
func (reg *register) set(val uint16) {
	reg.value = uint(val)
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
		cpu.reg[regAF].set(0x1180)
	} else {
		cpu.reg[regAF].set(0x01B0)
	}
	cpu.reg[regBC].set(0x0000)
	cpu.reg[regDE].set(0xFF56)
	cpu.reg[regHL].set(0x000D)
	cpu.reg[regSP].set(0xFFFE)
}

// Internally set the value of a flag on the flag register.
func (cpu *CPU) setFlag(index int, on bool) {
	mask := uint(1 << index)
	if on {
		cpu.reg[regAF].value |= mask
	} else {
		cpu.reg[regAF].value &^= mask
	}
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

func (cpu *CPU) setA(val byte) { cpu.reg[regAF].sethi(val) }
func (cpu *CPU) setB(val byte) { cpu.reg[regBC].sethi(val) }
func (cpu *CPU) setC(val byte) { cpu.reg[regBC].setLo(val) }
func (cpu *CPU) setD(val byte) { cpu.reg[regDE].sethi(val) }
func (cpu *CPU) setE(val byte) { cpu.reg[regDE].setLo(val) }
func (cpu *CPU) setF(val byte) { cpu.reg[regAF].setLo(val) }
func (cpu *CPU) setH(val byte) { cpu.reg[regHL].sethi(val) }
func (cpu *CPU) setL(val byte) { cpu.reg[regHL].setLo(val) }

func (cpu *CPU) a() byte    { return cpu.reg[regAF].hi() }
func (cpu *CPU) b() byte    { return cpu.reg[regBC].hi() }
func (cpu *CPU) c() byte    { return cpu.reg[regBC].lo() }
func (cpu *CPU) d() byte    { return cpu.reg[regDE].hi() }
func (cpu *CPU) e() byte    { return cpu.reg[regDE].lo() }
func (cpu *CPU) f() byte    { return cpu.reg[regAF].lo() & 0xF0 }
func (cpu *CPU) h() byte    { return cpu.reg[regHL].hi() }
func (cpu *CPU) l() byte    { return cpu.reg[regHL].lo() }
func (cpu *CPU) spHi() byte { return cpu.reg[regSP].hi() }
func (cpu *CPU) spLo() byte { return cpu.reg[regSP].lo() }

func (cpu *CPU) af() uint16 { return cpu.reg[regAF].hiLo() & 0xFFF0 }
func (cpu *CPU) bc() uint16 { return cpu.reg[regBC].hiLo() }
func (cpu *CPU) de() uint16 { return cpu.reg[regDE].hiLo() }
func (cpu *CPU) hl() uint16 { return cpu.reg[regHL].hiLo() }
func (cpu *CPU) sp() uint16 { return cpu.reg[regSP].hiLo() }

func (cpu *CPU) setAf(val uint16) { cpu.reg[regAF].set(val) }
func (cpu *CPU) setBc(val uint16) { cpu.reg[regBC].set(val) }
func (cpu *CPU) setDe(val uint16) { cpu.reg[regDE].set(val) }
func (cpu *CPU) setHl(val uint16) { cpu.reg[regHL].set(val) }
func (cpu *CPU) setSp(val uint16) { cpu.reg[regSP].set(val) }
