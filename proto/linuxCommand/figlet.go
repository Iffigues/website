package linuxCommand

import (
	"golang.org/x/net/context"
)

func (s *Server) GetFigletFile(ctx context.Context, in *Empty)(*Responses, error) {
	return startExec("figlist",[]string{})
}

func (s *Server) GetFiglet(ctx context.Context, in *Figlet) (*Responses, error) {
	var t []string
	if in.R {t = append(t, "-r")}
	if in.X {t = append(t, "-x")}
	if in.L {t = append(t, "-L")}
	if in.RR {t = append(t, "-R")}
	if in.XX {t = append(t, "-X")}
	if in.LL {t = append(t, "-l")}
	if in.C {t = append(t, "-c")}
	if in.P {t = append(t, "-P")}
	if in.N {t = append(t, "-n")}
	if in.O {t = append(t, "-o")}
	if in.W {t = append(t, "-W")}
	if in.K {t = append(t, "-k")}
	if in.S {t = append(t, "-S")}
	if in.SS {t = append(t, "-s")}
	if in.NN {t = append(t, "-N")}
	if in.E {t = append(t, "-E")}
	if in.D {t = append(t, "-D")}
	if in.T {t = append(t, "-T")}
	if in.NNN {t = append(t, "-N")}
	if in.WW != "" {t = addOptTab(t, "-w", in.WW);}
	if in.M != "" {t = addOptTab(t, "-m", in.M);}
	if in.F != "" {t = addOptTab(t, "-f", in.F);}
	if in.CC != "" {t = addOptTab(t, "-C", in.CC);}
	if in.Message != "" {t = addOptTab(t, in.Message);}
	return startExec("figlet", t)
}
