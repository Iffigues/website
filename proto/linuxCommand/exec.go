package linuxCommand


import (
	"fmt"
	"bytes"
	"os/exec"
	"time"
	"os"
)

func Exec(Command string, opt []string) (out, er bytes.Buffer, err error) {

	f, err := os.OpenFile("/log/linuxtool.log", os.O_WRONLY|os.O_APPEND, 0644)
	if err == nil {
		defer f.Close()
		var tt string = Command
		for _, e := range opt {
			tt = tt + " " + e
		}
		tt = tt + "\n"
		f.WriteString(tt)
	}

	path, errs := exec.LookPath(Command)
	if errs != nil {
		fmt.Println(errs)
		err = err
		return
	}
	cmd := exec.Command(path, opt...)
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
	fmt.Println(out, er, err)
	return
}
