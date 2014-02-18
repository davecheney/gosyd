package pdp11

// START1 OMIT
type trap struct {
	num int
	msg string
}

func (t trap) String() string { return fmt.Sprintf("trap %06o occured: %s", t.num, t.msg) }

func (k *cpu) step() {
	pc := uint16(k.R[7])
	ia := k.mmu.decode(pc, false, k.curuser)
	k.R[7] += 2 // increment PC
	instr := INST(k.unibus.read16(ia))
	switch instr {
	// handle instruction and return.
	}
	panic(trap{INTINVAL, "invalid instruction"})
}

// END1 OMIT

// START2 OMIT
func (p *PDP1140) Step() {
	defer func() {
		t := recover()
		switch t := t.(type) {
		case trap:
			p.trapat(t.num, t.msg)
		case nil:
			// ignore
		default:
			panic(t)
		}
	}()
	p.step()
}

// END2 OMIT
