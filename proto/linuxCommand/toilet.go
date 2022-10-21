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
	if in.S {
		t = append(t, "-S")
	}
	if in.SS {
		t = append(t, "-s")
	}
	if in.K {
		t = append(t, "-k")
	}
	if in.W {
		t = append(t, "-W")
	}
	if in.O {
		t = append(t, "-o")
	}
	if in.F != "" {
		t = addOptTab(t, "-f", in.F)
	}

	if in.FF != "" {
		t = addOptTab(t, "-F", in.FF)
	}

	if in.FFF != "" {
		t = addOptTab(t, "-F", in.FFF)
	}

	if in.E != "" {
		t = addOptTab(t, "-E", in.E)
	}

	for _, val := range in.F3 {
		t = addOptTab(t, "-F", val)
	}

	if in.Message != "" {
		t = append(t, in.Message)
	}
	return startExec("toilet", t)
}
