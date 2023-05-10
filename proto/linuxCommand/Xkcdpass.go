package linuxCommand

import (
	"golang.org/x/net/context"
)

func (s *Server) GetXkcdpass(ctx context.Context, in *Xkcdpass)(*Responses, error) {
	var t []string
	com := "xkcdpass"

	if in.Verbose {
		t = append(t, "-V")
	}
	if in.Min != "" {
		t = addOptTab(t, "---min", in.Min)
	}

	if in.Max != "" {
		t = addOptTab(t, "--Max", in.Max)
	}
	if in.Numwords != "" {
		t = addOptTab(t, "--numwords", in.Numwords)
	}

	if in.ValidChars != "" {
		t = addOptTab(t, "-v", in.ValidChars)
	}

	if in.Acrostic != "" {
		t = addOptTab(t, "-a", in.Acrostic)
	}

	if in.Count != "" {
		t = addOptTab(t, "-c", in.Count)
	}

	if in.Delimiter != "" {
		t = addOptTab(t, "-d", in.Delimiter)
	}

	if in.Separator != "" {
		t = addOptTab(t, "-s", in.Separator)
	}

	if in.Case != "" {
		t = addOptTab(t, "-C", in.Case)
	}

	return startExec(com, t)
}
