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

	if in.F3 != "" {
		t = addOptTab(t, "-F", in.F3)
	}

	if in.F4 != "" {
		t = addOptTab(t, "-F", in.F4)
	}

	if in.F5 != "" {
		t = addOptTab(t, "-F", in.F5)
	}

	if in.F6 != "" {
		t = addOptTab(t, "-F", in.F6)
	}

	if in.F7 != "" {
		t = addOptTab(t, "-F", in.F7)
	}


	if in.F8 != "" {
		t = addOptTab(t, "-F", in.F8)
	}

	if in.F9 != "" {
		t = addOptTab(t, "-F", in.F9)
	}


	if in.Message != "" {
		t = append(t, in.Message)
	}
	return startExec("toilet", t)
}
