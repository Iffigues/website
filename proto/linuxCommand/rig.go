package linuxCommand

import (
	"golang.org/x/net/context"
)

func (s *Server) GetRig(ctx context.Context, in *Rig) (*Responses, error) {
	var t []string

	t = addBoolOptTable(t, in.Man, "-m")
	t = addBoolOptTable(t, in.Woman, "-f")
	t = addStringOptTable(t, in.Nbr, "-c", in.Nbr)
	return startExec("rig", t)
}
