package sys

import (
	"syscall"
	"unsafe"
	// #include <unistd.h>
	"C"
)

/*
Isatty just returns a bool if the fd pointer passed happens to be a tty or not.

TODO: return errno as well?
*/
func Isatty(fd uintptr) bool {
	return int(C.isatty(C.int(fd))) != int(0)
}

/*
Termios struct for using TCGETS to see if we havea  tty.
*/
type Termios struct {
	Iflag, Oflag, Cflag, Lflag uint32
	Cc                         [20]byte
	Ispeed, Ospeed             uint32
}

/*
Fill up a Termios struct with a ioctl call to TCGETS to determine tty status.
*/
func IsTty(fd uintptr) bool {
	var termios Termios
	_, _, err := syscall.Syscall(syscall.SYS_IOCTL,
		fd,
		uintptr(syscall.TCGETS),
		uintptr(unsafe.Pointer(&termios)))
	return err == 0
}
