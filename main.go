package main

import (
	"os"
	"os/exec"
	"path"
	"strings"
)

// Usage
// > gofetch github.com/nickcarenza/gotfetch#master

// Using gofetch on your package with no arugments will gofetch all extenral dependencies with the correct path
// You can use versioned import paths
// import "github.com/sirupsen/logrus.master

// Get current branch or tag name
// git symbolic-ref -q --short HEAD || git describe --tags --exact-match
// http://stackoverflow.com/a/18660163/863924

// remove matching files in src/ and pkg/
// go get with -d flag
// Checkout desired tag or branch
// go install

// List all external dependencies of package
// go list -f "{{range .Imports}}{{.}} {{end}}" | xargs go list -f "{{if not .Standard}}{{.ImportPath}}{{end}}"

func main() {
	if len(os.Args) < 1 || len(os.Args) > 2 {
		os.Stderr.WriteString("Incorrect number of arguments. Expected 1 or 2" + "\n")
		os.Stdout.WriteString(help() + "\n")
		os.Exit(1)
	}
	importPath := os.Args[1]
	importDir, pkg := path.Split(importPath)
	var targetRef string
	if strings.Contains(pkg, ".") {
		parts := strings.Split(file, ".")
		pkg = strings.Join(parts[:len(parts)-1], ".")
		targetRef = parts[len(parts)-1]
	}
	baseImportPath := path.Join(importDir, pkg)
	_, err := goGetPkg(baseImportPath)
	if err != nil {
		os.Stderr.Write(res)
		os.Exit(1)
	}
	repositoryPath := path.Join(os.Getenv("GOPATH"), "src", baseImportPath)
	currentRef, err := getCurrentRef(repositoryPath)
	if targetRef == "" {
		targetRef = currentRef
	}
	// TODO make sure we treat unversioned deps correctly
	// TODO we should probably make sure we ginore standard packages or warn
	if string(currentRef) != targetRef {
		res, err := checkoutRef(targetRef)
		if err != nil {
			os.Stderr.Write(res)
			os.Exit(1)
		}
	}
	versionedImportPath := path.Join(repositoryPath, targetRef)
	err = moveVersionedImport(repositoryPath, versionedImportPath)
	goInstall(versionedImportPath)
}

func help() string {
	return "\t Usage: gofetch repository-url#ref destination"
}

func checkoutRef(ref string) (res []byte, err error) {
	cmd := exec.Command("git", "checkout", ref)
	res, err = cmd.Output()
	return
}

func moveVersionedImport(from, to string) (err error) {
}

func goInstall(repositoryPath) (res []byte, err error) {
	cmd := exec.Command("go", "install")
	cmd.Dir = repositoryPath
	res, err = cmd.Output()
	return
}

func goGetPkg() (res []byte, err error) {
	cmd := exec.Command("go", "get", "-d", path.Join(importDir, pkg))
	res, err = cmd.Output()
	return
}

func getCurrentRef(repositoryPath) (res []byte, err error) {
	cmd := exec.Command("git symbolic-ref -q --short HEAD || git describe --tags --exact-match")
	cmd.Dir = repositoryPath
	res, err = cmd.Output()
	return
}
