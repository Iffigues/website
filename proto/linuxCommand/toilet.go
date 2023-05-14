package linuxCommand

import (
	"golang.org/x/net/context"
)

func (s *Server) GetToiletFFile(ctx context.Context, in *Empty)(*Responses, error) {
	return startExec("toilet",[]string{"-F", "list"})
}

func (s *Server) GetToiletEFile(ctx context.Context, in *Empty)(*Responses, error) {
	return startExec("toilet",[]string{"-E", "list"})
}

func (s *Server) GetToiletFFFile(ctx context.Context, in *Empty) (*Responses, error) {
	return startExec("figlist", []string{})
}

func (s *Server) GetToilet(ctx context.Context, in *Toilet) (*Responses, error) {
	var t []string

	t = addBoolOptTable(t, in.S, "-S")
	t = addBoolOptTable(t, in.SS, "-s")
	t = addBoolOptTable(t, in.K, "-k")
	t = addBoolOptTable(t, in.W, "-W")
	t = addBoolOptTable(t, in.O, "-o")
	t = addStringOptTable(t, in.F, "-f", in.F)
	t = addStringOptTable(t, in.FF, "-F", in.FF)
	t = addStringOptTable(t, in.FFF, "-F", in.FFF)
	t = addStringOptTable(t, in.E, "-E", in.E)

	for _, val := range in.F3 {
		t = addOptTable(t, "-F", val)
	}

	if in.Message != "" {
		t = append(t, in.Message)
	}
	return startExec("toilet", t)
}
