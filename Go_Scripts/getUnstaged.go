package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func getUnstagedFiles() ([]string, error) {
	cmd := exec.Command("git", "status", "--porcelain")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("error executing git status: %w", err)
	}

	lines := strings.Split(string(output), "\n")

	fmt.Printf("Output: ", lines)
	var unstagedFiles []string
	for _, line := range lines {
		if len(line) > 0 {
			unstagedFiles = append(unstagedFiles, strings.TrimSpace(line[3:]))
		}
	}
	return unstagedFiles, nil
}

func main() {
	unstagedFiles, err := getUnstagedFiles()
	fmt.Println("unstaged", unstagedFiles)
	if err != nil {
		log.Fatal(err)
	}

	if len(unstagedFiles) > 0 {
		fmt.Println("Unstaged files:")
		for _, file := range unstagedFiles {
			fmt.Println(file)
		}
	} else {
		fmt.Println("No unstaged files.")
	}
}
