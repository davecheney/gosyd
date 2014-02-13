package rk11

type RK11 struct {
	RKBA, RKDS, RKER, RKCS, RKWC uint16
	sector, surface, cylinder    int
}

func (r *RK11) Step() {
	pos := (r.cylinder*24 + r.surface*12 + r.sector) * 512
	if pos >= len(rkdisk) {
		panic(fmt.Sprintf("pos outside rkdisk length, pos: %v, len %v", pos, len(rkdisk)))
	}
	// ...
}
