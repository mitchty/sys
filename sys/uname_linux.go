package sys

type Uname struct {
	System, Release, All, Machine, Processor   string
	Information, Os, Nodename, Version, Kernel string
}

/*
Return a Uname struct that makes sense for gnu/uname
*/
func NewUname() *Uname {
	u := Uname{
		System:      unameGet("-o"),
		Release:     unameGet("-r"),
		All:         unameGet("-a"),
		Machine:     unameGet("-m"),
		Processor:   unameGet("-p"),
		Information: unameGet("-i"),
		Os:          unameGet("-o"),
		Nodename:    unameGet("-n"),
		Version:     unameGet("-v"),
		Kernel:      unameGet("-s"),
	}
	return &u
}
