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

	t = addBoolOptTable(t, in.L, "-l")
	t = addBoolOptTable(t, in.B, "-b")
	t = addBoolOptTable(t, in.D, "-d")
	t = addBoolOptTable(t, in.G, "-g")
	t = addBoolOptTable(t, in.P, "-p")
	t = addBoolOptTable(t, in.S, "-s")
	t = addBoolOptTable(t, in.T, "-t")
	t = addBoolOptTable(t, in.W, "-w")
	t = addStringOptTable(t, in.E, "-e", in.E)
	t = addBoolOptTable(t, in.Y, "-y")
	t = addStringOptTable(t, in.TT, "-T", in.TT)
	t = addStringOptTable(t, in.F, "-f", in.F)
	t = append(t, in.Message);
	return startExec(com, t)
}
