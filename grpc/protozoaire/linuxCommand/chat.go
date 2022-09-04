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

func (s *Server) GetRig(ctx context.Context, in *Rig) (*RigResponses, error) {
	e := new(RigResponses)
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
	res, _,erro := Exec("rig", t)
	println("hello", res.String())
	if erro != nil {
		return nil, erro
	}
	return e, nil
}
