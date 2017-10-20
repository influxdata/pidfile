package pidfile

import "errors"

// ErrAlreadyLocked is returned when the pidfile is already locked.
var ErrAlreadyLocked = errors.New("already locked")
