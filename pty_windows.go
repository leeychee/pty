// +build windows

// Unsupport on windows. Just for build.

package pty

import (
	"os"
	"os/exec"
)

func Start(c *exec.Cmd) (pty *os.File, err error) {
	return nil, ErrUnsupported
}

func Getsize(t *os.File) (rows, cols int, err error) {
	return 0, 0, ErrUnsupported
}

func Setsize(t *os.File, rows uint16, cols uint16) error {
	return ErrUnsupported
}
