package sys

import (
	"runtime"
)

type Info struct {
	os       string
	minor    string
	major    string
	vendor   string
	arch     string
	version  string
	nickname string
	fullname func()
}

type verInfo struct {
	major string
	minor string
}

func NewInfo() *Info {
	arch := sysarch()
	info := archInfo(arch)
	i := Info{
		os:      sysos(),
		vendor:  vendor(),
		arch:    arch,
		version: version(),
		major:   info.major,
		minor:   info.minor,
	}
	return &i
}

func sysos() (os string) { return runtime.GOOS }

func sysarch() (arch string) {
	switch runtime.GOARCH {
	case "amd64":
		arch = "x86_64"
	default:
		arch = runtime.GOARCH
	}
	return
}

func version() string { return "unknown_version" }
