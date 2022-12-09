package ropes

type Direction string

const (
	up    Direction = "U"
	down  Direction = "D"
	left  Direction = "L"
	right Direction = "R"
)

type Rope []*segment

func NewRope(knots uint8) *Rope {
	rope := make(Rope, knots-1)
	rope[0] = &segment{&knot{}, &knot{}}
	for i := 1; i < len(rope); i++ {
		rope[i] = &segment{rope[i-1].tail, &knot{}}
	}
	return &rope
}

func (r Rope) MoveHead(d Direction) {
	r[0].head.move(d)
	r.motion()
}

func (r Rope) motion() {
	for _, seg := range r {
		seg.motion()
	}
}

func (r Rope) TailPosition() Position {
	return r[len(r)-1].tail.position
}

type segment struct {
	head, tail *knot
}

func (r *segment) motion() {
	r.tailCatchup(up)
	r.tailCatchup(down)
	r.tailCatchup(left)
	r.tailCatchup(right)
}

func (r *segment) tailCatchup(d Direction) {
	side1, side2, x, y, sgn := parameterMagic(d)
	if r.head.position[x]-r.tail.position[x] == sgn*2 {
		r.tail.move(d)
		if r.head.position[y]-r.tail.position[y] >= 1 {
			r.tail.move(side1)
		}
		if r.head.position[y]-r.tail.position[y] <= -1 {
			r.tail.move(side2)
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
		}
	}
	return
}

type Position [2]int
type knot struct {
	position Position
}

func (k *knot) move(d Direction) {
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
