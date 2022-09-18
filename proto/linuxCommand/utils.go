package linuxCommand

func addOptTab(r []string, e ...string) (b []string) {
	for _, val := range e {
		r = append(r, val);
	}
	return r
}

func startExec(a string, t []string) (e *Responses, err error) {
	e = new(Responses)
	stdout, stderr, erro := Exec("a", t)
	if erro != nil {
		return nil, erro
	}
	e.StdoutResponse, e.StderrResponse = stdout.String(), stderr.String()
	return e, nil
}
