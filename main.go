package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fmarmol/basename/pkg/basename"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("usage: %v <glob_pattern> <target_pattern>\n", os.Args[0])
		os.Exit(1)
	}
	glob := os.Args[1]
	targetPattern := os.Args[2]

	matches, err := filepath.Glob(glob)
	if err != nil {
		panic(err)
	}
	for _, match := range matches {
		fileInfo := basename.ParseFile(match)
		newFile := targetPattern
		if strings.Contains(targetPattern, "$basename") {
			newFile = strings.ReplaceAll(newFile, "$basename", fileInfo.Basename)
		}
		if strings.Contains(targetPattern, "$ext") {
			newFile = strings.ReplaceAll(newFile, "$ext", fileInfo.Ext)
		}
		err = os.Rename(match, newFile)
		if err != nil {
			panic(err)
		}
	}
}
