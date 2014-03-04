package pdp11

func (k *cpu) step() {
	k.pc = uint16(k.R[7])
	ia := k.mmu.decode(k.pc, false, k.curuser)
	k.R[7] += 2
	instr := INST(k.unibus.read16(ia))
	if printState {
		k.printstate()
	}
	if timeInstr {
		Runtime += k.timing(ia)
	}
	switch instr & 0070000 {
	case 0010000: // MOV
		MOV(k, instr)
		return
	case 0020000: // CMP
		CMP(k, instr)
		return
	// more instructions ...
	}
	panic(trap{intINVAL, "invalid instruction"})
}
