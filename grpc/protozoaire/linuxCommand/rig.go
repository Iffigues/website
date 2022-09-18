package linuxCommand

import (
	"golang.org/x/net/context"
)

func (s *Server) GetRig(ctx context.Context, in *Rig) (*Responses, error) {
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
	return startExec("rig", t)
}
