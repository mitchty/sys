package sys

import (
	"syscall"
	"unsafe"
	// #include <unistd.h>
	"C"
)

/*
Termios struct for using TCGETS to see if we have a tty.
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
