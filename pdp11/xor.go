package pdp11

func XOR(c *cpu, i INST) {
        s := i.S()
        val1 := c.R[s&7]
        d := i.D()
        da := c.aget(d, WORD)
        val2 := c.memread(da, WORD)
        val := val1 ^ val2
        c.PS &= 0xFFF1
        c.PS.testAndSetZero(val)
        c.PS.testAndSetNeg(val & 0x8000)
        c.memwrite(da, WORD, val)
}
