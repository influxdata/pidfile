// +build linux darwin

package pidfile

import (
	"os"
	"strconv"
	"syscall"
)

type Pidfile struct {
	file *os.File
}

// New writes a new file, locks the file, and then writes the pid to the file.
func New(fpath string) (*Pidfile, error) {
	file, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}

	if err := syscall.Flock(int(file.Fd()), syscall.LOCK_EX|syscall.LOCK_NB); err != nil {
		if err == syscall.EWOULDBLOCK {
			return nil, ErrAlreadyLocked
		}
		return nil, err
	}

	if err := file.Truncate(0); err != nil {
		return nil, err
	}
	file.WriteString(strconv.Itoa(os.Getpid()))
	return &Pidfile{file: file}, nil
}

// Close releases the pidfile by unlocking and removing it.
func (p *Pidfile) Close() error {
	syscall.Flock(int(p.file.Fd()), syscall.LOCK_UN)
	if err := p.file.Close(); err != nil {
		return err
	}
	return os.Remove(p.file.Name())
}
