package y22d12

type Grid [][]byte

func (g *Grid) neighbours(position [2]int) [][2]int {
	ns := make([][2]int, 0, 4)
	for _, d := range directions(position) {
		if g.inside(d) && g.doable(position, d) {
			ns = append(ns, d)
		}
	}
	return ns
}

func (g *Grid) doable(a, b [2]int) bool {
	return int((*g)[a[0]][a[1]])-int((*g)[b[0]][b[1]]) <= 1
}

func (g *Grid) inside(position [2]int) bool {
	return position[0] >= 0 && position[0] < len(*g) && position[1] >= 0 && position[1] < len((*g)[0])
}

type vertex struct {
	position [2]int
	steps    int
}

func directions(position [2]int) [4][2]int {
	return [4][2]int{
		{position[0] + 1, position[1]},
		{position[0] - 1, position[1]},
		{position[0], position[1] + 1},
		{position[0], position[1] - 1}}
}
