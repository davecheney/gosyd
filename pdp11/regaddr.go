package pdp11

// regaddr is a handle to an operand which can be
// read or written by an instruction.
type regaddr uint16

func (r regaddr) register() bool { return r&0177770 == 0170000 }
func (r regaddr) address() bool  { return !r.register() }

func (k *cpu) aget(v, l uint8) regaddr

func (k *cpu) memread16(a regaddr) uint16 {
	if a.register() {
		r := uint8(a & 7)
		return uint16(k.R[r&7])
	}
	return k.read16(uint16(a))
}
