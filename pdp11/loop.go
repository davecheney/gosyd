package pdp11

func (p *PDP1140) step() {
	if p.cpu.interrupts[0].vec > 0 && p.cpu.interrupts[0].pri >= ((int(p.cpu.PS)>>5)&7) {
		// handle interrupt
		return
	}
	p.cpu.step()
	clkcounter++
	if clkcounter >= 40000 {
		clkcounter = 0
		p.LKS |= (1 << 7)
		if p.LKS&(1<<6) != 0 {
			p.cpu.interrupt(INTCLOCK, 6)
		}
	}
	p.rk.Step()
	p.cons.Step()
}
