package linuxCommand


import (
	"strings"
	"bytes"
	"os/exec"
	"time"
	"os"
)

func logs(Command string, opt []string) {

	f, err := os.OpenFile("/log/linuxtool.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return
	}
	defer f.Close()
	ar := append([]string{Command}, opt...)
	tt := strings.Join(ar[:], " ")
	tt = tt + "\n"
	f.WriteString(tt)

}

func Exec(Command string, opt []string) (out, er bytes.Buffer, err error) {
	logs(Command, opt)
	path, errs := exec.LookPath(Command)
	if errs != nil {
		err = errs
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
		if errs := cmd.Process.Kill(); errs != nil {
			return out, er, errs
		}
		return
	case erro := <-done:
		if erro != nil {
			return out, er, erro
		}
	}
	return
}
