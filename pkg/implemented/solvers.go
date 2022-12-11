package implemented

import (
	"adventofcode/domain"
	y22d1 "adventofcode/pkg/22/01"
	y22d2 "adventofcode/pkg/22/02"
	y22d3 "adventofcode/pkg/22/03"
	y22d4 "adventofcode/pkg/22/04"
	y22d5 "adventofcode/pkg/22/05"
	y22d6 "adventofcode/pkg/22/06"
	y22d7 "adventofcode/pkg/22/07"
	y22d8 "adventofcode/pkg/22/08"
	y22d9 "adventofcode/pkg/22/09"
	y22d10 "adventofcode/pkg/22/10"
	y22d11 "adventofcode/pkg/22/11"
)

type Solvers map[string]domain.Solver

func GetSolvers() Solvers {
	return map[string]domain.Solver{

		// 2022
		"y22d01": (&y22d1.Solver{}).Init(),
		"y22d02": (&y22d2.Solver{}).Init(),
		"y22d03": (&y22d3.Solver{}).Init(),
		"y22d04": (&y22d4.Solver{}).Init(),
		"y22d05": (&y22d5.Solver{}).Init(),
		"y22d06": (&y22d6.Solver{}).Init(),
		"y22d07": (&y22d7.Solver{}).Init(),
		"y22d08": (&y22d8.Solver{}).Init(),
		"y22d09": (&y22d9.Solver{}).Init(),
		"y22d10": (&y22d10.Solver{}).Init(),
		"y22d11": (&y22d11.Solver{}).Init(),
	}
}
