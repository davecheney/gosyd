func (p *PSW) testAndSetZero(v int) {
        if v == 0 {
                *p |= FLAGZ
        }
}

func SWAB(c *cpu, i INST) {
        l := i.L()
        d := i.D()
        da := c.aget(d, l)
        val := c.memread(da, l)
        val = ((val >> 8) | (val << 8)) & 0xFFFF
        c.PS &= 0xFFF0
        c.PS.testAndSetZero(val & 0xff)
        c.PS.testAndSetNeg(val & 0x80)
        c.memwrite(da, l, val)
}

