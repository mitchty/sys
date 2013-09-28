package sys

import (
	// #include <termios.h>
	"C"
)

/*
Since Osx doesn't do TCGETS, use tcgetattr() instead.
TODO: check errno to see if we got EINTR or ENOTTY?
*/
func IsTty(fd uintptr) bool {

	//	var termios syscall.Termios
	var termios C.struct_termios
	err := int(C.tcgetattr(C.int(fd), &termios))
	return err == 0
}
