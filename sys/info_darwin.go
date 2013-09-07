package sys

import (
	"fmt"
	"strings"
)

/*
Stringified output like osx-10.8.4 etc...
*/
func (i *Info) String() string {
	return fmt.Sprintf("%s-%s.%s", i.vendor, i.major, i.minor)
}

/*
verInfo equates to approximately, for say 10.8.4
major = 10.8
minor = 4
*/
func archInfo(arch string) (info verInfo) {
	pv := trimmedOutput("sw_vers", "-productVersion")
	version := strings.Split(pv, ": ")[0]
	vs := strings.Split(version, ".")
	info.major = strings.Join(vs[:2], ".")
	info.minor = strings.Join(vs[2:], "")
	return
}

/*
"vendor" is osx here, whatever it made sense at the time for reasons, don't
judge me.
*/
func vendor() (vendor string) {
	vendor = "osx"
	return
}
