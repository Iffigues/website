package linuxCommand

import (
	"golang.org/x/net/context"
)

func (s *Server) GetBanner(ctx context.Context, in *Banner)(*Responses, error) {
	com := "banner"
	return startExec(com, in.Message)
}
