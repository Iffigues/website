package linuxCommand

import (
	"context"
	"time"
)

func addBoolOptTable(r []string, a bool, t string) (b []string) {
	if !a {
		return r
	}
	r = append(r, t)
	return r
}

func addStringOptTable(r []string, a string, e ...string) (b []string) {
	if a == "" {
		return r
	}

	for _, val := range e {
		r = append(r, val);
	}
	return r
}

func addOptTable(r []string, e ...string) (b []string) {
	for _, val := range e {
		r = append(r, val);
	}
	return r
}

func startExec(a string, t []string) (e *Responses, err error) {
	e = new(Responses)
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second * 60)
	defer cancel()
	c, err := NewCmd(a, t)
	if err != nil {
		return nil, err
	}
	stdout, stderr, err := c.Exec(ctx)
	if err != nil {
		return nil, err
	}
	e.StdoutResponse, e.StderrResponse = stdout.String(), stderr.String()
	return e, nil
}
