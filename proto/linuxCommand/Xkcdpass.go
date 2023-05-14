package linuxCommand

import (
	"golang.org/x/net/context"
)

func (s *Server) GetXkcdpass(ctx context.Context, in *Xkcdpass)(*Responses, error) {
	var t []string
	com := "xkcdpass"

	t = addBoolOptTable(t, in.Verbose, "-V")
	t = addStringOptTable(t, in.Min, "--min", in.Min)
	t = addStringOptTable(t, in.Max, "--max", in.Max)
	t = addStringOptTable(t, in.Numwords, "-numwords", in.Numwords)
	t = addStringOptTable(t, in.ValidChars, "-v", in.ValidChars)
	t = addStringOptTable(t, in.Acrostic, "-a", in.Acrostic)
	t = addStringOptTable(t, in.Count, "-c", in.Count)
	t = addStringOptTable(t, in.Delimiter, "-d", in.Delimiter)
	t = addStringOptTable(t, in.Separator, "-s", in.Separator)
	t = addStringOptTable(t, in.Case, "-C", in.Case)
	return startExec(com, t)
}
