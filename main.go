package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Fprintln(os.Stderr, "Incorrect numer of arguments. Expected 2 or 3")
		fmt.Fprintln(os.Stdout, help())
		os.Exit(1)
	}
	repositoryUrl := os.Args[1]
	reference := "master"
	if strings.Contains(repositoryUrl, "#") {
		parts := strings.SplitN(repositoryUrl, "#", 2)
		repositoryUrl = parts[0]
		reference = parts[1]
	}
	dst := os.Args[2]
	cmd := exec.Command("git", "clone", repositoryUrl, "--branch", reference, dst)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func help() string {
	return "\t Usage: gofetch repository-url#ref destination"
}
