// Package pidfile implements pidfile creation.
//
// On Linux and Darwin, this package uses flock to take an exclusive lock of the pidfile.
// If another process tries to lock the file, an error will be returned. If the previous
// process dies and does not unlock the file or delete it, the operating system cleans up
// the lock automatically and this package will adopt the file and lock it.
//
// This package relies on the operating system to handle the locking. Some filesystems,
// such as NFS, may not implement flock in a reliable way. You should not use this package
// on such operating systems.
//
// The file locking is only an advisory. If the file is manually opened by a call to
// os.Open, the file will still open normally.
package pidfile
