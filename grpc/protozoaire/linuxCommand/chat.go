package linuxCommand

import (
	"golang.org/x/net/context"
)


func addOptTab(r []string, e ...string) {
	for _, val := range e {
		r = append(r, val);
	}
}

type Server struct {
	UnimplementedRigServiceServer
}

func (s *Server) GetRig(ctx context.Context, in *Rig) (*Responses, error) {
	e := new(Responses)
	var t []string
	if in.Man {
		t = append(t, "-m")
	}
	if in.Woman {
		t = append(t, "-f")
	}
	if in.Nbr != "" {
		addOptTab(t, "-c", in.Nbr)
	}
	print(in.Nbr)
	stdout, stderr, erro := Exec("rig", t)
	if erro != nil {
		return nil, erro
	}
	e.StdoutResponse, e.StderrResponse = stdout.String(), stderr.String()
	return e, nil
}
