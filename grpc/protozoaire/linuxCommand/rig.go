package linuxCommand

import (
	"golang.org/x/net/context"
)

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
		t = addOptTab(t, "-c", in.Nbr)
	}
	stdout, stderr, erro := Exec("rig", t)
	if erro != nil {
		return nil, erro
	}
	e.StdoutResponse, e.StderrResponse = stdout.String(), stderr.String()
	return e, nil
}
