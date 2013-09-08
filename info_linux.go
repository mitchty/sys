package sys

import (
	"fmt"
	"log"
	"os"
	"strings"
)

/*
Store where the release file is, if at all.
*/
type linuxVendorReleaseFile struct {
	vendor      string
	releaseFile string
}

/*
map of where release files per vendor type
*/
var linuxVendorReleaseFiles = []linuxVendorReleaseFile{
	{"sles", "/etc/SuSE-release"},
	{"arch", "/etc/arch-release"},
	{"redhat", "/etc/Redhat-release"},
	{"unknown_linux", "/dev/null"},
}

/* NOT YET USED
type linuxVerInfo struct {
	arch string
	info verInfo
}

var linuxVerInfos = []func(){
	{"sles", func() verInfo { return slesInfo() }},
	{"unknown_linux", func() verInfo { return verInfo{major: "??", minor: "?"} }},
}*/

/*
Parse out SuSE Linux Enterprise Something release files.
verInfo equates to approximately, for say sles 11 service pack 1
major = 11
minor = 1

sp 2 is the same only
minor = 2

sp 3 the same
*/
func slesInfo() (info verInfo) {
	for _, line := range releasefileContents() {
		if strings.Contains(line, "PATCHLEVEL = ") {
			info.minor = strings.Trim(strings.Split(line, " = ")[1], "\n")
		}
		if strings.Contains(line, "VERSION = ") {
			info.major = strings.Trim(strings.Split(line, " = ")[1], "\n")
		}
	}
	return
}

/*
Stringify the output as vendor/major/minor

sles gets "special" treatment and output like so:
sles${major}u${minor}
*/
func (i *Info) String() string {
	switch vendor() {
	case "sles":
		return fmt.Sprintf("%s%su%s", i.Vendor, i.Major, i.Minor)
	default:
		return fmt.Sprintf("%s%s%s", i.Vendor, i.Major, i.Minor)
	}
}

/*
Dump out major/minor numbers.

And also have some defaults of nothing for major/minor when i've no idea what
i'm running on.
*/
func archInfo(arch string) (info verInfo) {
	switch vendor() {
	case "sles":
		info = slesInfo()
	default:
		info = verInfo{major: "", minor: ""}
	}
	return
}

/*
Where's our release file?
*/
func releasefile() (name string) {
	for _, i := range linuxVendorReleaseFiles {
		if vendor() == i.vendor {
			name = i.releaseFile
			break
		}
	}
	return
}

/*
Dump the contents of the file.
*/
func releasefileContents() (contents []string) {
	contents, err := ReadLines(releasefile())
	if err != nil {
		log.Fatal(err)
	}
	return
}

/*
What vendor of linux are we running.
*/
func vendor() (vendor string) {
	vendor = "unknown_linux"
	for _, i := range linuxVendorReleaseFiles {
		if _, err := os.Stat(i.releaseFile); err == nil {
			vendor = i.vendor
			break
		}
	}
	return
}
