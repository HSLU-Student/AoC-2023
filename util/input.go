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
