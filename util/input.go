package util

import (
	"fmt"
	"os"
	"path"
	"strings"
)

func GetContent(day string) string {
	//gets working directory of caller (main.go)
	cdir, err := os.Getwd()

	//err if current working dir can not be fetched
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	path := path.Join(cdir, "data", fmt.Sprintf("%s.txt", day))
	return readContent(path)

}

func GetContentRoot(day string) string {
	path := path.Join("/workspaces/AoC-2023/data", fmt.Sprintf("%s.txt", day))
	return readContent(path)
}

func readContent(path string) string {
	raw, err := os.ReadFile(path)

	//err if file can't be read
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	return string(raw)
}

func SplitContentLine(input string) []string {
	return strings.Split(input, "\n")
}

func SplitContentRow(input string) []string {
	rows := SplitContentLine(input)
	tinput := []string{}

	//find longest string for transposition
	longest := 0
	for _, row := range rows {
		if len(row) > longest {
			longest = len(row)
		}
	}

	//get rows into slice
	for i := 0; i < longest; i++ {
		var str string

		for _, row := range rows {
			if len(row) > i {
				str += string(row[i])
			}
		}
		tinput = append(tinput, str)
	}

	return tinput
}
