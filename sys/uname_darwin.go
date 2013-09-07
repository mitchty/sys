package sys

type Uname struct {
	release   string
	all       string
	machine   string
	processor string
	Nodename  string
	kernel    string
}

/*
Return a Uname struct that makes sense for osx/darwin uname.
*/
func NewUname() *Uname {
	u := Uname{
		release:   unameGet("-r"),
		all:       unameGet("-a"),
		Nodename:  unameGet("-n"),
		machine:   unameGet("-m"),
		processor: unameGet("-p"),
		kernel:    unameGet("-s"),
	}
	return &u
}
