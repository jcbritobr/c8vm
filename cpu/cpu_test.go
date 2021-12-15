package cpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldCreateNewCpuUnit(t *testing.T) {
	cpu := NewCPU()
	assert.NotNil(t, cpu)
}

func TestShouldLoadROM(t *testing.T) {
	cpu := NewCPU()
	n, err := cpu.LoadROM("testdata/ParticleDemo.ch8")
	if err != nil {
		assert.Fail(t, "LoadROM fail with %v", err)
	}
	assert.Equal(t, 353, n, "353 should be read as the demo has 353 bytes")
	for i := 0x50; i < 0x200; i++ {
		assert.Equal(t, uint8(0), cpu.memory[i], "First 512 bytes must be zero, as this is vm area")
	}
}

func TestLoadROMShouldFailWithWrongFilename(t *testing.T) {
	cpu := NewCPU()
	_, err := cpu.LoadROM("DummyRom")
	if err == nil {
		assert.Fail(t, "Error cant be nil")
	}
}

func TestCPUShouldReset(t *testing.T) {
	cpu := NewCPU()
	_, err := cpu.LoadROM("testdata/ParticleDemo.ch8")
	if err != nil {
		assert.Fail(t, "LoadROM failed with %v", err)
	}
	cpu.i = 42
	cpu.Reset()
	ncpu := NewCPU()
	assert.Equal(t, ncpu, cpu, "Cpu needs to be same as ncpu(new) after reset")
}

func TestShouldReturnFromSubrotine(t *testing.T) {
	cpu := NewCPU()
	cpu.stack[cpu.sp] = 0x30
	cpu.sp = cpu.sp + 1
	cpu.memory[0x200] = 0x00
	cpu.memory[0x201] = 0xee
	cpu.RunCycle()
	assert.Equal(t, uint16(0x30), cpu.pc)
	assert.Equal(t, uint16(0x00), cpu.sp)
}

func TestShouldClearDisplay(t *testing.T) {
	cpu := NewCPU()
	cpu.memory[0x200] = 0x00
	cpu.memory[0x201] = 0xe0
	cpu.display[0][0] = 0x1
	cpu.display[9][23] = 0x1
	cpu.RunCycle()
	for x := 0x00; x < 0x20; x++ {
		for y := 0x00; y < 0x40; y++ {
			assert.Equal(t, byte(0), cpu.display[x][y])
		}
	}
}

func TestShouldJumpToNNN(t *testing.T) {
	cpu := NewCPU()
	cpu.memory[0x200] = 0x10
	cpu.memory[0x201] = 0xff
	cpu.RunCycle()
	assert.Equal(t, uint16(0xff), cpu.pc)
}

func TestShouldCallAddr(t *testing.T) {
	cpu := NewCPU()
	cpu.memory[0x200] = 0x26
	cpu.memory[0x201] = 0x93
	cpu.RunCycle()
	assert.Equal(t, uint16(0x01), cpu.sp)
	assert.Equal(t, uint16(0x202), cpu.stack[cpu.sp-1])
	assert.Equal(t, uint16(0x693), cpu.pc)
}

func TestShouldSkipIfVxIsNNIsTrue(t *testing.T) {
	cpu := NewCPU()
	cpu.memory[0x200] = 0x3b
	cpu.memory[0x201] = 0x54
	cpu.v[0xb] = 0x54
	cpu.RunCycle()
	assert.Equal(t, uint16(0x204), cpu.pc)
}

func TestShouldNotSkipIfVxIsNNIsFalse(t *testing.T) {
	cpu := NewCPU()
	cpu.memory[0x200] = 0x31
	cpu.memory[0x201] = 0x54
	cpu.v[0x1] = 0x95
	cpu.RunCycle()
	assert.Equal(t, uint16(0x202), cpu.pc)
}
