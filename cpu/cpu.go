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
