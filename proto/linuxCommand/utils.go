package linuxCommand

import(
	"fmt"
)

func addOptTab(r []string, e ...string) (b []string) {
	for _, val := range e {
		r = append(r, val);
	}
	return r
}

func startExec(a string, t []string) (e *Responses, err error) {
	e = new(Responses)
	stdout, stderr, erro := Exec(a, t)
	fmt.Println("haha",stdout, err)
	if erro != nil {
		return nil, erro
	}
	e.StdoutResponse, e.StderrResponse = stdout.String(), stderr.String()
	return e, nil
}
