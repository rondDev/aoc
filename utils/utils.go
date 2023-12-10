package utils

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// Load files, fileName should be formatted like "1.txt"
func LoadLines(fileName string) []string {
	lines := []string{}
	f, err := os.OpenFile("input/"+fileName, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return lines
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}

			return lines
		}
		line = strings.TrimSpace(line)
		lines = append(lines, line) // GET the line string
	}
	return lines
}
