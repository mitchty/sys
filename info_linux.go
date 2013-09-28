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
	vendor, releaseFile string
}

type thisSystem struct {
	vendor, releaseFile string
	releaseFileContents []string
}

var theLinuxIKnow thisSystem

/*
map of where release files per vendor type
*/
var linuxVendorReleaseFiles = []linuxVendorReleaseFile{
	{"sles", "/etc/SuSE-release"},
	{"arch", "/etc/arch-release"},
	{"redhat", "/etc/Redhat-release"},
	{"debian", "/etc/lsb-release"},
	{"ubuntu", "/etc/lsb-release"},
}

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
	for _, line := range releaseFileContents() {
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
Parse out Ubuntu derpitude from lsb-release.
verInfo equates to for 12.04:
major = 12
minor = 04
*/
func ubuntuInfo() (info verInfo) {
	for _, line := range releaseFileContents() {
		if strings.Contains(line, "DISTRIB_RELEASE=") {
			all := strings.Split(line, "=")[1]
			all = strings.Trim(all, "\n")
			ver := strings.Split(all, ".")
			info.major, info.minor = ver[0], ver[1]
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
	case "ubuntu":
		return fmt.Sprintf("%s-%s.%s", i.Vendor, i.Major, i.Minor)
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
	setupTheLinuxIKnow()
	switch vendor() {
	case "sles":
		info = slesInfo()
	case "ubuntu":
		info = ubuntuInfo()
	default:
		info = verInfo{major: "", minor: ""}
	}
	return
}

/*
Setup the struct that keeps track of the release filename
and its contents.

Bit of a giant method, but I'll refactor later.
*/
func setupTheLinuxIKnow() {
	// filename
	for _, i := range linuxVendorReleaseFiles {
		if _, err := os.Stat(i.releaseFile); err == nil {
			theLinuxIKnow.releaseFile = i.releaseFile
		}
	}
	// contents
	contents, err := ReadLines(theLinuxIKnow.releaseFile)
	if err != nil {
		log.Fatal(err)
		// panic?
	}
	theLinuxIKnow.releaseFileContents = contents
	// vendor
	theLinuxIKnow.vendor = "unknown_linux"
	if theLinuxIKnow.releaseFile == "/etc/lsb-release" {
		content := contents
		if strings.Contains(content[0], "Ubuntu") {
			theLinuxIKnow.vendor = "ubuntu"
		}
	} else {
		for _, i := range linuxVendorReleaseFiles {
			if i.releaseFile == theLinuxIKnow.releaseFile {
				theLinuxIKnow.vendor = i.vendor
				break
			}
		}
	}
	return
}

/*
Where's our release file?
*/
func releasefile() (name string) {
	name = theLinuxIKnow.releaseFile
	return
}

/*
Dump the contents of the file cached in the struct.
*/
func releaseFileContents() (contents []string) {
	contents = theLinuxIKnow.releaseFileContents
	return
}

/*
What vendor of linux are we running from the cached struct.
*/
func vendor() (vendor string) {
	vendor = theLinuxIKnow.vendor
	return
}
