package sys

import (
	"runtime"
)

type Info struct {
	Os, Minor, Major, Vendor, Arch, Version string
}

type verInfo struct {
	major string
	minor string
}

func NewInfo() *Info {
	arch := sysarch()
	info := archInfo(arch)
	i := Info{
		Os:      sysos(),
		Vendor:  vendor(),
		Arch:    arch,
		Version: version(),
		Major:   info.major,
		Minor:   info.minor,
	}
	return &i
}

func sysarch() (arch string) {
	arch = runtime.GOARCH
	if arch == "amd64" {
		arch = "x86_64"
	}
	return
}

func version() string { return "unknown_version" }

func sysos() (os string) { return runtime.GOOS }
