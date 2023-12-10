package days

import (
	"fmt"
	"strconv"
	"strings"
	"utils"
)

func Three() {
	fmt.Printf("Day 3 Part 1: %d\n", threeA())
	fmt.Printf("Day 3 Part 2: %d\n", threeB())
}

func threeA() int {

	total := 0

	testInput :=
		`467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`
	lns := strings.Split(testInput, "\n")
	lns = utils.LoadLines("3.txt")
	for li, l := range lns {
		curnum := ""
		for i := 0; i < len(l); i++ {
			check := true
			if x, err := strconv.Atoi(string(l[i])); err == nil {
				curnum += strconv.Itoa(x)
				if i != len(l)-1 {
					check = false
				}
			} else {
			}
			if check {
				if curnum != "" {
					add := false
					var mini int
					if i-len(curnum)-1 < 0 {
						mini = i - len(curnum)
					} else {
						mini = i - len(curnum) - 1
					}
					for x := mini; x <= i; x++ {
						_, e := strconv.Atoi(string(l[x]))
						if e != nil && l[x] != '.' {
							add = true
						}
						if li-1 >= 0 && len(lns[li-1]) > x {
							_, ed := strconv.Atoi(string(lns[li-1][x]))
							if ed != nil && lns[li-1][x] != '.' {
								add = true
							}
						}
						if li+1 <= len(lns)-1 && len(lns[li+1]) > x {
							_, eu := strconv.Atoi(string(lns[li+1][x]))
							if eu != nil && lns[li+1][x] != '.' {
								add = true
							}
						}

					}
					if add {
						c, _ := strconv.Atoi(curnum)
						total += c
						curnum = ""
					}

					curnum = ""
					// fmt.Printf("i: %d - curnum: %s\n", i, curnum)
				}
			}

		}
	}
	return total
}

func threeB() int {

	return 1
}
