package pdp11

// uint18 represents a unibus 18 bit physical address
type uint18 uint32

func (m *KT11) decode(a uint16, w, user bool) (addr uint18)

func (c *cpu) read16(a uint16) uint16 {
	addr := c.mmu.decode(a, false, c.curuser)
	return c.unibus.read16(addr)
}
