package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	var (
		dirF = flag.String("dir", "", "The directory to clone into.")
	)
	flag.Parse()

	cloneURL := flag.Arg(0)
	relCloneDir, err := ExtractPath(cloneURL)
	if err != nil {
		return err
	}
	cloneDir := filepath.Join(*dirF, relCloneDir)
	parentDir := filepath.Dir(cloneDir)
	if _, err := os.Stat(cloneDir); os.IsNotExist(err) {
		if err := os.MkdirAll(parentDir, 0755); err != nil {
			return err
		}

		cmd := exec.Command("git", "clone", cloneURL)
		cmd.Dir = parentDir
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return err
		}
	}
	fmt.Println(cloneDir)
	return nil
}

var (
	httpR = regexp.MustCompile(`^[^:]+://(.+?)(?:\.git)?$`)
	sshR  = regexp.MustCompile(`^git@(.+)\.git$`)
)

func ExtractPath(cloneURL string) (string, error) {
	if m := httpR.FindStringSubmatch(cloneURL); len(m) == 2 {
		return m[1], nil
	}
	if m := sshR.FindStringSubmatch(cloneURL); len(m) == 2 {
		return strings.ReplaceAll(m[1], ":", "/"), nil
	}
	return "", fmt.Errorf("unknown clone url format: %q", cloneURL)
}
