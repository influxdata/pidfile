// +build linux darwin

package pidfile_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/influxdata/pidfile"
)

func TestPidfile(t *testing.T) {
	tmpdir, err := ioutil.TempDir(os.TempDir(), "pidfile")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpdir)

	fpath := filepath.Join(tmpdir, "file.pid")
	p, err := pidfile.New(fpath)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	// Check that the file exists.
	if _, err := os.Stat(fpath); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	// Attempt to double lock the file.
	if _, err := pidfile.New(fpath); err != pidfile.ErrAlreadyLocked {
		t.Fatalf("unexpected error: %s", err)
	}

	// Unlock the file.
	if err := p.Close(); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	// Attempt to lock the file again.
	if p, err := pidfile.New(fpath); err != nil {
		t.Fatalf("unexpected error: %s", err)
	} else {
		p.Close()
	}
}
