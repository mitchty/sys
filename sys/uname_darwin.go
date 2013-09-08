package sys

type Uname struct {
	Release, All, Machine, Processor, Nodename, Kernel string
}

/*
Return a Uname struct that makes sense for osx/darwin uname.
*/
func NewUname() *Uname {
	u := Uname{
		Release:   unameGet("-r"),
		All:       unameGet("-a"),
		Nodename:  unameGet("-n"),
		Machine:   unameGet("-m"),
		Processor: unameGet("-p"),
		Kernel:    unameGet("-s"),
	}
	return &u
}
