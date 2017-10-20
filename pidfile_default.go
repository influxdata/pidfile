// +build !linux,!darwin

package pidfile

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Pidfile struct {
	fpath string
}

func New(fpath string) (*Pidfile, error) {
	pid := strconv.Itoa(os.Getpid())
	if err := ioutil.WriteFile(path, []byte(pid), 0644); err != nil {
		return nil, fmt.Errorf("write file: %s", err)
	}
	return &Pidfile{fpath: fpath}, nil
}

func (p *Pidfile) Close() error {
	return os.Remove(p.fpath)
}
