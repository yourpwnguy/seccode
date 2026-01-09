package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// Pattern for checking alphanumeric characters
var pattern = regexp.MustCompile(`^[A-Za-z0-9._-]+$`)

// validFileName checks if the filename only contains alphanumeric characters.
// It does by validating the filename against a regex,
// returns true if filename is alphanumeric
func validFileName(s string) bool {
	return pattern.Match([]byte(s))
}

func main() {
	// Example base directory
	// Take this by a faint of heart
	baseDir := "/var"

	// Example input we are trying
	// From an attacker's POV
	f := "../../../../../etc/rpc"

	// First check if the given input is alphanumeric
	if !validFileName(f) {
		log.Fatal("invalid filename")
	}

	// Append the given input to the base directory
	joinedPath := filepath.Join(baseDir, f)

	// Canocalize the full path, we are just normalizing it here to
	// get the resolved path. It handled symlinks bypasses.
	cleanPath, err := filepath.EvalSymlinks(joinedPath)
	if err != nil {
		log.Fatal(err)
	}

	// We do the same for the base directory
	baseClean, err := filepath.EvalSymlinks(baseDir)
	if err != nil {
		log.Fatal(err)
	}

	// We are doing the final check, if the given input starts with out base
	// directory. This is the important check.
	if !strings.HasPrefix(cleanPath, baseClean+string(os.PathSeparator)) {
		log.Fatal("path traversal detected")
	}

	// Here you might wanna start processing the file now
	fmt.Println("OK:", cleanPath)
}
