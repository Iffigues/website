package linuxCommand

import (
	"golang.org/x/net/context"
)

func (s *Server) GetFortune(ctx context.Context, in *Fortune)(*Responses, error) {
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
		t = addOptTab(t, "-m", in.M)
	}
	if in.N != "" {
		t = addOptTab(t, "-n", in.N)
	}
	for _, val := range in.Percent {
		t = addOptTab(t, val.Percent, val.Fortune)
	}
	stdout, stderr, erro := Exec("fortune", t)
	if erro != nil {
		return nil, erro
	}
	e.StdoutResponse, e.StderrResponse = stdout.String(), stderr.String()
	return e, nil
}
