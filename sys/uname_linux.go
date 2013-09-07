package sys

type Uname struct {
	System      string
	release     string
	all         string
	machine     string
	processor   string
	information string
	os          string
	Nodename    string
	version     string
	kernel      string
}

/*
Return a Uname struct that makes sense for gnu/uname
*/
func NewUname() *Uname {
	u := Uname{
		System:      unameGet("-o"),
		release:     unameGet("-r"),
		all:         unameGet("-a"),
		machine:     unameGet("-m"),
		processor:   unameGet("-p"),
		information: unameGet("-i"),
		os:          unameGet("-o"),
		Nodename:    unameGet("-n"),
		version:     unameGet("-v"),
		kernel:      unameGet("-s"),
	}
	return &u
}
