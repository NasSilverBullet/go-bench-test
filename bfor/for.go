package bfor

var benchcount = 1000000

type BFor struct {
	Count int
}

func (b *BFor) getBenchCount() int {
	if b.Count > 0 {
		return b.Count
	}
	return benchcount
}

type Big struct {
	A string
	B string
	C string
	D string
	E string
	F string
	G string
	H string
	I string
	J string
	K string
	L string
	M string
	N string
	O string
	P string
	Q string
	R string
	S string
	T string
	U string
	V string
	W string
	X string
	Y string
	Z string
}

func (b *BFor) Gen() []*Big {
	c := b.getBenchCount()

	m := make([]*Big, c)
	for i := 0; i < c; i++ {
		m[i] = &Big{
			A: "A",
			B: "B",
			C: "C",
			D: "D",
			E: "E",
			F: "F",
			G: "G",
			H: "H",
			I: "I",
			K: "K",
			L: "L",
			M: "M",
			N: "N",
			O: "O",
			P: "P",
			Q: "Q",
			R: "R",
			S: "S",
			T: "T",
			U: "U",
			V: "V",
			W: "W",
			X: "X",
			Y: "Y",
			Z: "Z",
		}
	}

	return m
}

func (b *BFor) LoopWithCopy(bigs []*Big) {
	for _, big := range bigs {
		// var s string
		// s = big.A
		// s = big.B
		// s = big.C
		// s = big.D
		// s = big.E
		// s = big.F
		// s = big.G
		// s = big.H
		// s = big.I
		// s = big.K
		// s = big.L
		// s = big.M
		// s = big.N
		// s = big.O
		// s = big.P
		// s = big.Q
		// s = big.R
		// s = big.S
		// s = big.T
		// s = big.U
		// s = big.V
		// s = big.W
		// s = big.X
		// s = big.Y
		// s = big.Z
		// _ = s
		_ = big.A
	}
}

func (b *BFor) LoopMapAccess(bigs []*Big) {
	for i := range bigs {
		// var s string
		// s = bigs[i].A
		// s = bigs[i].B
		// s = bigs[i].C
		// s = bigs[i].D
		// s = bigs[i].E
		// s = bigs[i].F
		// s = bigs[i].G
		// s = bigs[i].H
		// s = bigs[i].I
		// s = bigs[i].K
		// s = bigs[i].L
		// s = bigs[i].M
		// s = bigs[i].N
		// s = bigs[i].O
		// s = bigs[i].P
		// s = bigs[i].Q
		// s = bigs[i].R
		// s = bigs[i].S
		// s = bigs[i].T
		// s = bigs[i].U
		// s = bigs[i].V
		// s = bigs[i].W
		// s = bigs[i].X
		// s = bigs[i].Y
		// s = bigs[i].Z
		_ = bigs[i].A
	}
}
