package linuxCommand

import (
	"golang.org/x/net/context"
)

func (s *Server) GetFortune(ctx context.Context, in *Fortune)(*Response, error) {
	e := new(Responses)
	var t []string

	if in.A {
		t = append(t, "-a")
	}
	if in.C {
		t = append(t, "-c")
	}
	if in.E {
		t = append(t, "-e")
	}
	if in.F {
		t = append(t, "-f")
	}
	if in.L {
		t = append(t, "-l")
	}
	if in.O {
		t = append(t, "-o")
	}
	if in.S {
		t = append(t, "-s")
	}
	if in.I {
		t = append(t, "-i")
	}
	if in.U {
		t = append(t, "-u")
	}
	if in.M != "" {
	}
	if in.N != "" {
	}
	for _, val := range in.Percent {
		
	}
	return
}
