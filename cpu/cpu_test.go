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
