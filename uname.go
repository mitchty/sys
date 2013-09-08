package sys

import (
	"os/exec"
	"strings"
)

/*
These helper functions should probably be abstracted away, I think, somewhere
I DON"T KNOW GO, err bo, whatever, I don't know what i'm doing.
*/
func trimmedOutput(command string, arg ...string) (output string) {
	out, _ := exec.Command(command, arg...).CombinedOutput()
	output = strings.Trim(string(out), "\n")
	return
}

/*
Cause, lazy
*/
func unameGet(arg ...string) (output string) {
	output = trimmedOutput("uname", arg...)
	return
}
