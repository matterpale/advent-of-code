package ropes

type Direction string

const (
	up    Direction = "U"
	down  Direction = "D"
	left  Direction = "L"
	right Direction = "R"
)

type Position [2]int
type Knot struct {
	position Position
}

func (k *Knot) move(d Direction) {
	switch d {
	case left:
		k.position[0]--
	case right:
		k.position[0]++
	case up:
		k.position[1]--
	case down:
		k.position[1]++
	default:
		panic("invalid direction")
	}
}

type Rope struct {
	Head, Tail *Knot
}

func NewRope() *Rope {
	return &Rope{&Knot{}, &Knot{}}
}

func (r *Rope) MoveHead(d Direction) {
	r.Head.move(d)
	r.motion()
}

func (r *Rope) motion() {
	r.tailCatchup(up)
	r.tailCatchup(down)
	r.tailCatchup(left)
	r.tailCatchup(right)
}

func (r *Rope) tailCatchup(d Direction) {
	side1, side2, x, y, sgn := parameterMagic(d)
	if r.Head.position[x]-r.Tail.position[x] == sgn*2 {
		r.Tail.move(d)
		if r.Head.position[y]-r.Tail.position[y] >= 1 {
			r.Tail.move(side1)
		}
		if r.Head.position[y]-r.Tail.position[y] <= -1 {
			r.Tail.move(side2)
		}
	}
}

func parameterMagic(d Direction) (side1, side2 Direction, x, y, sgn int) {
	side1, side2 = down, up
	x, y, sgn = 0, 1, 1
	if d == left {
		sgn = -1
	} else if d != right {
		side1, side2, x, y = right, left, y, x
		if d == up {
			sgn = -1
		} else if d != down {
			panic("invalid direction")
		}
	}
	return
}

func (r *Rope) TailPosition() Position {
	return r.Tail.position
}

type LongRope []*Rope

func tieRope(k *Knot) *Rope {
	return &Rope{k, &Knot{}}
}

func NewLongRope(knots uint8) *LongRope {
	lr := make(LongRope, knots-1)
	lr[0] = NewRope()
	for i := 1; i < len(lr); i++ {
		lr[i] = tieRope(lr[i-1].Tail)
	}
	return &lr
}

func (lr LongRope) MoveHead(d Direction) {
	lr[0].Head.move(d)
	lr.motion()
}

func (lr LongRope) TailPosition() Position {
	return lr[len(lr)-1].Tail.position
}

func (lr LongRope) motion() {
	for _, r := range lr {
		r.motion()
	}
}
