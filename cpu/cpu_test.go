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
