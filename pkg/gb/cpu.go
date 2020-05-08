package gb

import "log"

const (
	maskFlagC = 1 << 4
	maskFlagH = 1 << 5
	maskFlagN = 1 << 6
	maskFlagZ = 1 << 7
)

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
	// registers
	reg [5]uint
	pc  uint

	Divider int
}

// Init CPU and its registers to the initial values.
func (cpu *CPU) Init(cgb bool) {
	cpu.pc = 0x100
	if cgb {
		cpu.setAf(0x1180)
	} else {
		cpu.setAf(0x01B0)
	}
	cpu.setBc(0x0000)
	cpu.setDe(0xFF56)
	cpu.setHl(0x000D)
	cpu.setSp(0xFFFE)
}

// Internally set the value of a flag on the flag register.
func (cpu *CPU) setFlag(mask uint, on bool) {
	if on {
		cpu.reg[regAF] |= mask
	} else {
		cpu.reg[regAF] &^= mask
	}
}

func (cpu *CPU) getFlag(index int) bool {
	mask := uint(1 << index)
	return cpu.reg[regAF]&mask != 0
}

func (cpu *CPU) setFlagC(on bool) { cpu.setFlag(maskFlagC, on) }
func (cpu *CPU) setFlagH(on bool) { cpu.setFlag(maskFlagH, on) }
func (cpu *CPU) setFlagN(on bool) { cpu.setFlag(maskFlagN, on) }
func (cpu *CPU) setFlagZ(on bool) { cpu.setFlag(maskFlagZ, on) }

func (cpu *CPU) flagC() bool { return cpu.reg[regAF]&maskFlagC != 0 }
func (cpu *CPU) flagH() bool { return cpu.reg[regAF]&maskFlagH != 0 }
func (cpu *CPU) flagN() bool { return cpu.reg[regAF]&maskFlagN != 0 }
func (cpu *CPU) flagZ() bool { return cpu.reg[regAF]&maskFlagZ != 0 }

func (cpu *CPU) carryBit() uint { return (cpu.reg[regAF] >> 4) & 0x1 }

// TODO move to test
func panicIfNotByte(val uint) {
	if val&^0xFF != 0 {
		log.Panicf("error: attempt to write more than 8 bits: 0x%x", val)
	}
}

// TODO move to test
func panicIfNotUint16(val uint) {
	if val&^0xFFFF != 0 {
		log.Panicf("error: attempt to write more than 16 bits: 0x%x", val)
	}
}

func (cpu *CPU) setA(val uint) {
	panicIfNotByte(val)
	cpu.reg[regAF] &= 0xFF
	cpu.reg[regAF] |= val << 8
}

func (cpu *CPU) setB(val uint) {
	panicIfNotByte(val)
	cpu.reg[regBC] &= 0xFF
	cpu.reg[regBC] |= val << 8
}

func (cpu *CPU) setC(val uint) {
	panicIfNotByte(val)
	cpu.reg[regBC] &= 0xFF00
	cpu.reg[regBC] |= val
}

func (cpu *CPU) setD(val uint) {
	panicIfNotByte(val)
	cpu.reg[regDE] &= 0xFF
	cpu.reg[regDE] |= val << 8
}

func (cpu *CPU) setE(val uint) {
	panicIfNotByte(val)
	cpu.reg[regDE] &= 0xFF00
	cpu.reg[regDE] |= val
}

func (cpu *CPU) setF(val uint) {
	panicIfNotByte(val)
	cpu.reg[regAF] &= 0xFF00
	cpu.reg[regAF] |= val
}

func (cpu *CPU) setH(val uint) {
	panicIfNotByte(val)
	cpu.reg[regHL] &= 0xFF
	cpu.reg[regHL] |= val << 8
}

func (cpu *CPU) setL(val uint) {
	panicIfNotByte(val)
	cpu.reg[regHL] &= 0xFF00
	cpu.reg[regHL] |= val
}

func (cpu *CPU) a() uint    { return (cpu.reg[regAF] >> 8) & 0xFF }
func (cpu *CPU) b() uint    { return cpu.reg[regBC] >> 8 }
func (cpu *CPU) c() uint    { return cpu.reg[regBC] & 0xFF }
func (cpu *CPU) d() uint    { return cpu.reg[regDE] >> 8 }
func (cpu *CPU) e() uint    { return cpu.reg[regDE] & 0xFF }
func (cpu *CPU) f() uint    { return cpu.reg[regAF] & 0xF0 } // mask out low F-nibble
func (cpu *CPU) h() uint    { return cpu.reg[regHL] >> 8 }
func (cpu *CPU) l() uint    { return cpu.reg[regHL] & 0xFF }
func (cpu *CPU) spHi() uint { return cpu.reg[regSP] >> 8 }
func (cpu *CPU) spLo() uint { return cpu.reg[regSP] & 0xFF }

func (cpu *CPU) af() uint { return cpu.reg[regAF] & 0xFFF0 } // mask out low F-nibble
func (cpu *CPU) bc() uint { return cpu.reg[regBC] }
func (cpu *CPU) de() uint { return cpu.reg[regDE] }
func (cpu *CPU) hl() uint { return cpu.reg[regHL] }
func (cpu *CPU) sp() uint { return cpu.reg[regSP] }

func (cpu *CPU) setAf(val uint) { panicIfNotUint16(val); cpu.reg[regAF] = val }
func (cpu *CPU) setBc(val uint) { panicIfNotUint16(val); cpu.reg[regBC] = val }
func (cpu *CPU) setDe(val uint) { panicIfNotUint16(val); cpu.reg[regDE] = val }
func (cpu *CPU) setHl(val uint) { panicIfNotUint16(val); cpu.reg[regHL] = val }
func (cpu *CPU) setSp(val uint) { panicIfNotUint16(val); cpu.reg[regSP] = val }
