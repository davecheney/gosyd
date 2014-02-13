func MUL(c *cpu, i INST) {
        l := i.L()
        s := i.S()
        val1 := c.R[s&7]
        if val1&0x8000 == 0x8000 {
                val1 = -((0xFFFF ^ val1) + 1)
        }
        d := i.D()
        da := c.aget(d, l)
        val2 := int(c.memread(da, WORD))
        if val2&0x8000 == 0x8000 {
                val2 = -((0xFFFF ^ val2) + 1)
        }
        val := val1 * val2
        c.R[s&7] = val >> 16
        c.R[(s&7)|1] = val & 0xFFFF
        c.PS &= 0xFFF0
        c.PS.testAndSetNeg(val & 0x80000000)
        c.PS.testAndSetZero(val & 0xFFFFFFFF)
        if val < (1<<15) || val >= ((1<<15)-1) {
                c.PS |= FLAGC
        }
}
