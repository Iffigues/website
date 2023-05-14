package linuxCommand

import (
	"golang.org/x/net/context"
)

func (s *Server) GetFigletFile(ctx context.Context, in *Empty)(*Responses, error) {
	return startExec("figlist",[]string{})
}

func (s *Server) GetFiglet(ctx context.Context, in *Figlet) (*Responses, error) {
	var t []string

	t = addBoolOptTable(t, in.R, "-r")
	t = addBoolOptTable(t, in.X, "-x")
	t = addBoolOptTable(t, in.L, "-L")
	t = addBoolOptTable(t, in.RR, "-R")
	t = addBoolOptTable(t, in.XX, "-X")
	t = addBoolOptTable(t, in.LL, "-l")
	t = addBoolOptTable(t, in.C, "-c")
	t = addBoolOptTable(t, in.P, "-P")
	t = addBoolOptTable(t, in.N, "-n")
	t = addBoolOptTable(t, in.O, "-o")
	t = addBoolOptTable(t, in.W, "-W")
	t = addBoolOptTable(t, in.K, "-k")
	t = addBoolOptTable(t, in.S, "-S")
	t = addBoolOptTable(t, in.SS, "-s")
	t = addBoolOptTable(t, in.NN, "-N")
	t = addBoolOptTable(t, in.E, "-E")
	t = addBoolOptTable(t, in.D, "-D")
	t = addBoolOptTable(t, in.T, "-T")
	t = addBoolOptTable(t, in.NNN, "-N")
	t = addStringOptTable(t, in.WW, "-w", in.WW)
	t = addStringOptTable(t, in.M, "-m", in.M)
	t = addStringOptTable(t, in.F, "-f", in.F)
	t = addStringOptTable(t, in.CC, "-C", in.CC)
	t = addStringOptTable(t, in.Message, in.Message)
	return startExec("figlet", t)
}
