package days

import (
	"fmt"
	"strconv"
	"strings"
	"utils"
)

func One() {
	fmt.Printf("Day 1 Part 1: %d\n", oneA())
	fmt.Printf("Day 1 Part 2: %d\n", oneB())
}

func oneA() int {
	lines := utils.LoadLines("1.txt")
	newLines := []string{}
	total := 0
	var first int
	var last int
	for _, v := range lines {
		for _, k := range v {
			c := string(k)
			if i, err := strconv.Atoi(c); err == nil {
				if first == 0 {
					first = i
				} else {
					last = i
				}
			}
		}
		if last == 0 {
			last = first
		}
		temp := fmt.Sprintf("%d%d", first, last)
		newLines = append(newLines, temp)
		tempnum, _ := strconv.Atoi(temp)
		total += tempnum
		first = 0
		last = 0
	}
	return total
}

func oneB() int {
	lines := utils.LoadLines("1.txt")
	total := 0
	for _, l := range lines {
		list := []int{}
		l = strings.TrimSpace(l) + "aaaaa"
		for i := range l {
			if len(l) <= i+5 {
				break
			}
			num, err := strconv.Atoi(string(l[i]))
			var one = (l[i:i+3] == "one")
			var two = (l[i:i+3] == "two")
			var three = (l[i:i+5] == "three")
			var four = (l[i:i+4] == "four")
			var five = (l[i:i+4] == "five")
			var six = (l[i:i+3] == "six")
			var seven = (l[i:i+5] == "seven")
			var eight = (l[i:i+5] == "eight")
			var nine = (l[i:i+4] == "nine")
			if err == nil {
				list = append(list, num)
			}
			if one {
				list = append(list, 1)
			}
			if two {
				list = append(list, 2)
			}
			if three {
				list = append(list, 3)
			}
			if four {
				list = append(list, 4)
			}
			if five {
				list = append(list, 5)
			}
			if six {
				list = append(list, 6)
			}
			if seven {
				list = append(list, 7)
			}
			if eight {
				list = append(list, 8)
			}
			if nine {
				list = append(list, 9)
			}
		}

		temp := fmt.Sprintf("%d%d", list[0], list[len(list)-1])
		tempnum, _ := strconv.Atoi(temp)
		total += tempnum
	}
	return total
}
