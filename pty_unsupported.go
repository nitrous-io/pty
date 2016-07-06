// +build !linux,!darwin,!freebsd,!dragonfly

package pty

import (
	"os"
)

func open() (pty, tty *os.File, err error) {
	return nil, nil, ErrUnsupported
}

func ptsname(f *os.File) (string, error) {
	return "", ErrUnsupported
}

func grantpt(f *os.File) error {
	return ErrUnsupported
}

func unlockpt(f *os.File) error {
	return ErrUnsupported
}

func ioctl(fd, cmd, ptr uintptr) error {
	return ErrUnsupported
}

func setsize(f *os.File, rows uint16, cols uint16) error {
	return ErrUnsupported
}
