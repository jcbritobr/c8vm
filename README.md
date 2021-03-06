# C8VM is a chip 8 virtual machine built using Go language.

### **Memory**
The Chip-8 language is capable of accessing up to 4KB (4,096 bytes) of RAM, from location 0x000 (0) to 0xFFF (4095). The first 512 bytes, from 0x000 to 0x1FF, are where the original interpreter was located, and should not be used by programs.
Most Chip-8 programs start at location 0x200 (512), but some begin at 0x600 (1536). Programs beginning at 0x600 are intended for the ETI 660 computer.

![Memory Map](images/chip8memorymap.png)

### **Registers**
Chip-8 has 16 general purpose 8-bit registers, usually referred to as Vx, where x is a hexadecimal digit (0 through F). There is also a 16-bit register called I. This register is generally used to store memory addresses, so only the lowest (rightmost) 12 bits are usually used.
The VF register should not be used by any program, as it is used as a flag by some instructions.
Chip-8 also has two special purpose 8-bit registers, for the delay and sound timers. When these registers are non-zero, they are automatically decremented at a rate of 60Hz.
There are also some "pseudo-registers" which are not accessable from Chip-8 programs. The program counter (PC) should be 16-bit, and is used to store the currently executing address. The stack pointer (SP) can be 8-bit, it is used to point to the topmost level of the stack.

The stack is an array of 16 16-bit values, used to store the address that the interpreter shoud return to when finished with a subroutine. Chip-8 allows for up to 16 levels of nested subroutines.




### **Instruction Set**
The following table contains all thirty-five instructions in the CHIP-8 instruction set. **NNN** refers to a hexadecimal memory address. **NN** refers to a hexadecimal byte. **N** refers to a hexadecimal nibble. **X** and **Y** refer to registers.

|Instruction |Description|
|------------|-----------|
|0NNN |Execute machine language subroutine at address NNN|
|00E0 |Clear the screen|
|00EE |Return from a subroutine|
|1NNN |Jump to address NNN|
|2NNN |Execute subroutine starting at address NNN|
|3XNN |Skip the following instruction if the value of register VX equals NN|
|4XNN |Skip the following instruction if the value of register VX is not equal to NN|
|5XY0 |Skip the following instruction if the value of register VX is equal to the value of register VY|
|6XNN |Store number NN in register VX|
|7XNN |Add the value NN to register VX|
|8XY0 |Store the value of register VY in register VX|
|8XY1 |Set VX to VX OR VY|
|8XY2 |Set VX to VX AND VY|
|8XY3 |Set VX to VX XOR VY|
|8XY4 |Add the value of register VY to register VX. Set VF to 01 if a carry occurs. Set VF to 00 if a carry does not occur|
|8XY5 |Subtract the value of register VY from register VX. Set VF to 00 if a borrow occurs. Set VF to 01 if a borrow does not occur|
|8XY6 |Store the value of register VY shifted right one bit in register VX??. Set register VF to the least significant bit prior to the shift. VY is unchanged|
|8XY7 |Set register VX to the value of VY minus VX. Set VF to 00 if a borrow occurs. Set VF to 01 if a borrow does not occur|
|8XYE |Store the value of register VY shifted left one bit in register VX??. Set register VF to the most significant bit prior to the shift. VY is unchanged|
|9XY0 |Skip the following instruction if the value of register VX is not equal to the value of register VY|
|ANNN |Store memory address NNN in register I|
|BNNN |Jump to address NNN + V0|
|CXNN |Set VX to a random number with a mask of NN|
|DXYN |Draw a sprite at position VX, VY with N bytes of sprite data starting at the address stored in I. Set VF to 01 if any set pixels are changed to unset, and 00 otherwise|
|EX9E |Skip the following instruction if the key corresponding to the hex value currently stored in register VX is pressed|
|EXA1 |Skip the following instruction if the key corresponding to the hex value currently stored in register VX is not pressed|
|FX07 |Store the current value of the delay timer in register VX|
|FX0A |Wait for a keypress and store the result in register VX|
|FX15 |Set the delay timer to the value of register VX|
|FX18 |Set the sound timer to the value of register VX|
|FX1E |Add the value stored in register VX to register I|
|FX29 |Set I to the memory address of the sprite data corresponding to the hexadecimal digit stored in register VX|
|FX33 |Store the binary-coded decimal equivalent of the value stored in register VX at addresses I, I + 1, and I + 2|
|FX55 |Store the values of registers V0 to VX inclusive in memory starting at address I. I is set to I + X + 1 after operation|
|FX65 |Fill registers V0 to VX inclusive with the values stored in memory starting at address I. I is set to I + X + 1 after operation|

### **References**
* [CHIP-8 Wiki](https://en.wikipedia.org/wiki/CHIP-8)
* [Cowgod's CHIP-8 reference](http://devernay.free.fr/hacks/chip8/C8TECH10.HTM)
* [Matthew Mikolay - Mastering CHIP 8](https://github.com/mattmikolay/chip-8/wiki/Mastering-CHIP%E2%80%908)
