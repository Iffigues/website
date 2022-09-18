package linuxCommand


import (
	"fmt"
	"bytes"
	"os/exec"
	"time"
)

func Exec(Command string, opt []string) (out, er bytes.Buffer, err error) {
	fmt.Println(opt)
	cmd := exec.Command(Command, opt...)
	cmd.Stdout = &out
	cmd.Stderr  = &er
	err = cmd.Start()
	if err != nil {
		return out, er, err
	}
	done := make(chan error, 1)
	go func() {
	    done <- cmd.Wait()
	}()
	select {
	case <-time.After(60 * time.Second):
		if err := cmd.Process.Kill(); err != nil {
		}
		return
	case err := <-done:
		if err != nil {
		}
	}
	fmt.Println(out.String(), er.String());
	return
}
