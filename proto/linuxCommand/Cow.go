package linuxCommand

import (
	"golang.org/x/net/context"
)

func (s *Server) GetCowFile(ctx context.Context, in *Empty)(*Responses, error) {
	return startExec("cowsay",[]string{"-l"})
}

func (s *Server) GetCow(ctx context.Context, in *Cow)(*Responses, error) {
	var t []string
	com := "cowsay"

	if in.Think {
		com = "cowthink"
	}
	if in.L {
		t = append(t, "-l")
	}
	if in.B {
		t = append(t, "-b")
	}
	if in.D {
		t = append(t, "-d")
	}
	if in.G {
		t = append(t, "-g")
	}
	if in.P {
		t = append(t, "-p")
	}
	if in.S {
		t = append(t, "-s")
	}
	if in.T {
		t = append(t, "-t")
	}
	if in.W {
		t = append(t, "-w")
	}
	if in.E != "" {
		addOptTab(t, "-e", in.E)
	}
	if in.TT != "" {
		addOptTab(t, "-T", in.TT)
	}
	if in.F != "" {
		addOptTab(t, "-f", in.F);
	}
	t = append(t, in.Message);
	return startExec(com, t)
}
