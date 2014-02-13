func (c *cpu) memread16(a uint32) uint16

func ADD(c *cpu, i INST) {
        sa := c.aget(i.S(), WORD)
        val1 := c.memread16(sa)
        da := c.aget(i.D(), WORD)
        val2 := c.memread16(da)
        val := val1 + val2
        c.PS &= 0xFFF0
        if val == 0 {
                c.PS |= FLAGZ
        }
        if val&0x8000 == 0x8000 {
                c.PS |= FLAGN
        }
        if val1+val2 >= 0xFFFF {
                c.PS |= FLAGC
        }
        c.memwrite16(da, val)
}
