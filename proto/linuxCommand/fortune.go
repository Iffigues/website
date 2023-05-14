package linuxCommand

import (
	"golang.org/x/net/context"
)

func (s *Server) GetFortuneFile(ctx context.Context, in *Empty)(*Responses, error) {
	return startExec("fortune",[]string{"-f"})
}

func (s *Server) GetFortune(ctx context.Context, in *Fortune)(*Responses, error) {
	var t []string

	t = addBoolOptTable(t, in.A, "-a")
	t = addBoolOptTable(t, in.C, "-c")
	t = addBoolOptTable(t, in.E, "-e")
	t = addBoolOptTable(t, in.F, "-f")
	t = addBoolOptTable(t, in.L, "-l")
	t = addBoolOptTable(t, in.O, "-o")
	t = addBoolOptTable(t, in.S, "-s")
	t = addBoolOptTable(t, in.I, "-i")
	t = addBoolOptTable(t, in.U, "-u")
	t = addStringOptTable(t, in.M, "-m", in.M)
	t = addStringOptTable(t, in.N, "-n", in.N)
	for _, val := range in.Percent {
		t = addOptTable(t, val.Percent + "%", val.Fortune)
	}
	return startExec("fortune", t)
}
