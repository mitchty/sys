package sys

import (
	// #include <unistd.h>
	"C"
)

/*
IsTty just returns a bool if the fd pointer passed happens to be a tty or not.

TODO: return errno as well?
*/
func IsTty(fd uintptr) bool {
	return int(C.isatty(C.int(fd))) != int(0)
}
