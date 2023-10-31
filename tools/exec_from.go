//go:build ignore
// +build ignore

package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
)

func init() {
	log.SetOutput(os.Stderr)
	log.SetFlags(0)
}

func main() {
	dir := flag.String("dir", "", "directory to run from")
	flag.Parse()

	if n := len(flag.Args()); n < 1 {
		log.Fatalf("Expected at least 1 argument, got %d", n)
	}

	cmd := exec.Command(flag.Arg(0), flag.Args()[1:]...)
	cmd.Dir = *dir
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		switch err := err.(type) {
		case *exec.ExitError:
			os.Exit(err.ExitCode())
		default:
			log.Fatalf("Unknown error received: %s", err)
		}
	}
}
