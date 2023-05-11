package linuxCommand

import ("fmt")

func addOptTab(r []string, e ...string) (b []string) {
	for _, val := range e {
		r = append(r, val);
	}
	return r
}

func startExec(a string, t []string) (e *Responses, err error) {
	fmt.Println(a, t)
	e = new(Responses)
	stdout, stderr, erro := Exec(a, t)
	if erro != nil {
		fmt.Println("hihi", erro, stderr)
		return nil, erro
	}
	e.StdoutResponse, e.StderrResponse = stdout.String(), stderr.String()
	fmt.Println("haha", string(e.StderrResponse))
	return e, nil
}
