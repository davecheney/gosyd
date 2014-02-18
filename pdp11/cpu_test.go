package pdp11

// START1 OMIT
type regs struct { R0, R1, R2, R3, R4, R5, R6, R7 int; PS psw }

type core map[uint18]uint16

type suite struct {
	name string
	regs
	core
	steps    int
	wantregs regs
}

var instrTests = []suite{
	{
		name:     "CLR R1",
		regs:     regs{R1: 0177777, R7: 001000, PS: 000017},
		core:     core{001000: 005001},
		steps:    1,
		wantregs: regs{R0: 0000000, R7: 001002, PS: 000004},
	},
}

// END1 OMIT
// START2 OMIT
func TestInstructions(t *testing.T) {
	for _, tt := range instrTests {
		instrTest(t, tt)
	}
}

func instrTest(t *testing.T, tt suite) {
	t.Log(tt.name)
	cpu := New()
	cpu.LoadMemory(tt.core)
	loadRegs(&cpu.cpu, tt.regs)
	for i := 0; i < tt.steps; i++ {
		cpu.Step()
	}
	checkRegs(t, &cpu.cpu, tt.wantregs)
}

// END2 OMIT
