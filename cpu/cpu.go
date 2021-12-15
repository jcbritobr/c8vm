package cpu

import "os"

const (
	height = byte(0x20)
	width  = byte(0x40)
)

type CPU struct {
	pc            uint16
	memory        [4096]byte
	stack         [16]uint16
	sp            uint16
	v             [16]byte
	i             uint16
	delayTimer    byte
	soundTimer    byte
	display       [height][width]byte
	keys          [16]byte
	draw          bool
	inputFlag     bool
	inputRegister byte
}

func NewCPU() *CPU {
	cpu := &CPU{pc: 0x200}
	cpu.LoadFontSet()
	return cpu
}

func getFirstNibble(opcode uint16) uint16 {
	firstNibble := opcode & 0xf000
	return firstNibble
}

func getLastNibble(opcode uint16) uint16 {
	lastNibble := opcode & 0x000f
	return lastNibble
}

func (c *CPU) LoadROM(rom string) (int, error) {
	file, err := os.Open(rom)
	if err != nil {
		return 0, err
	}

	defer file.Close()
	memory := make([]byte, 3584)
	size, err := file.Read(memory)
	if err != nil {
		return 0, err
	}

	for index, data := range memory {
		c.memory[index+0x200] = data
	}

	return size, nil
}

func (c *CPU) LoadFontSet() {
	for i := 0x00; i < 0x50; i++ {
		c.memory[i] = fontset[i]
	}
}

func (c *CPU) Reset() {
	c.pc = 0x200
	c.delayTimer = 0
	c.soundTimer = 0
	c.i = 0
	c.sp = 0
	for i := 0; i < len(c.memory); i++ {
		c.memory[i] = 0
	}

	for i := 0; i < len(c.stack); i++ {
		c.stack[i] = 0
	}

	for i := 0; i < len(c.v); i++ {
		c.v[i] = 0
	}

	for i := 0; i < len(c.keys); i++ {
		c.keys[i] = 0
	}
	c.LoadFontSet()
	c.ClearDisplay()
}

func (c *CPU) fetchOpCode() uint16 {
	opCode := uint16(c.memory[c.pc])<<8 | uint16(c.memory[c.pc+1])
	return opCode
}

func (c *CPU) ClearDisplay() {
	for x := 0x00; x < 0x20; x++ {
		for y := 0x00; y < 0x40; y++ {
			c.display[x][y] = 0x00
		}
	}
}

func (c *CPU) Run() {
	c.RunCycle()
	if c.delayTimer > 0 {
		c.delayTimer = c.delayTimer - 1
	}

	if c.soundTimer > 0 {
		c.soundTimer = c.soundTimer - 1
	}
}

func (c *CPU) RunCycle() {
	opcode := c.fetchOpCode()
	c.pc = c.pc + 2
	switch getFirstNibble(opcode) {
	case 0x0000:
		switch getLastNibble(opcode) {
		case 0x0000:
			c.ClearDisplay()
		case 0x000e:
			c.pc = c.stack[c.sp-1]
			c.sp = c.sp - 1
		}
	case 0x1000:
		c.pc = opcode & 0x0fff
	case 0x2000:
		c.stack[c.sp] = c.pc
		c.sp = c.sp + 1
		c.pc = opcode & 0x0fff
	case 0x3000:
		nn := byte(opcode & 0x00ff)
		register := (opcode & 0x0f00) >> 8
		if c.v[register] == nn {
			c.pc = c.pc + 2
		}
	}
}
