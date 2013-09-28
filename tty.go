package sys

import (
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
