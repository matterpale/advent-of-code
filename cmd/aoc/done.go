package main

import (
	"github.com/matterpale/advent-of-code/domain"
	y15d01 "github.com/matterpale/advent-of-code/solutions/15/01"
	y21d01 "github.com/matterpale/advent-of-code/solutions/21/01"
	y21d02 "github.com/matterpale/advent-of-code/solutions/21/02"
	y22d01 "github.com/matterpale/advent-of-code/solutions/22/01"
	y22d02 "github.com/matterpale/advent-of-code/solutions/22/02"
	y22d03 "github.com/matterpale/advent-of-code/solutions/22/03"
	y22d04 "github.com/matterpale/advent-of-code/solutions/22/04"
	y22d05 "github.com/matterpale/advent-of-code/solutions/22/05"
	y22d06 "github.com/matterpale/advent-of-code/solutions/22/06"
	y22d07 "github.com/matterpale/advent-of-code/solutions/22/07"
	y22d08 "github.com/matterpale/advent-of-code/solutions/22/08"
	y22d09 "github.com/matterpale/advent-of-code/solutions/22/09"
	y22d10 "github.com/matterpale/advent-of-code/solutions/22/10"
	y22d11 "github.com/matterpale/advent-of-code/solutions/22/11"
	y22d12 "github.com/matterpale/advent-of-code/solutions/22/12"
	y22d13 "github.com/matterpale/advent-of-code/solutions/22/13"
	y23d01 "github.com/matterpale/advent-of-code/solutions/23/01"
	y23d02 "github.com/matterpale/advent-of-code/solutions/23/02"
	y23d03 "github.com/matterpale/advent-of-code/solutions/23/03"
	y23d04 "github.com/matterpale/advent-of-code/solutions/23/04"
)

var implemented = map[string]domain.Solver{
	// 2015
	"y15d01": &y15d01.Solver{},
	// 2021
	"y21d01": &y21d01.Solver{},
	"y21d02": &y21d02.Solver{},
	// 2022
	"y22d01": &y22d01.Solver{},
	"y22d02": &y22d02.Solver{},
	"y22d03": &y22d03.Solver{},
	"y22d04": &y22d04.Solver{},
	"y22d05": &y22d05.Solver{},
	"y22d06": &y22d06.Solver{},
	"y22d07": &y22d07.Solver{},
	"y22d08": &y22d08.Solver{},
	"y22d09": &y22d09.Solver{},
	"y22d10": &y22d10.Solver{},
	"y22d11": &y22d11.Solver{},
	"y22d12": &y22d12.Solver{},
	"y22d13": &y22d13.Solver{},
	// 2023
	"y23d01": &y23d01.Solver{},
	"y23d02": &y23d02.Solver{},
	"y23d03": &y23d03.Solver{},
	"y23d04": &y23d04.Solver{},
}

func initAllSolvers() map[string]domain.Solver {
	for _, solver := range implemented {
		solver.Init()
	}
	return implemented
}
