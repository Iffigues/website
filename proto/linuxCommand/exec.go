package linuxCommand

import (
	"bytes"
	"os/exec"
	"context"
)

func NewCmd(command string, arg []string) (c *Command, err error) {
	c =  &Command{
		command: command,
		args: arg,
	}
	c.path, err = exec.LookPath(command)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Command) Command(ctx context.Context) (cmd *exec.Cmd) {
	cmd = exec.CommandContext(ctx, c.path, c.args...)
	cmd.Stdout = &c.out
	cmd.Stderr = &c.err
	return cmd
}

func (c *Command)Exec(ctx context.Context) (out, er bytes.Buffer, err error) {
	cmd := c.Command(ctx)
	c.logs()
	err = cmd.Start()
	if err != nil {
		return c.out, c.err, err
	}
	done := make(chan error, 1)
	go func() {
	    done <- cmd.Wait()
	}()
	select {
	case <-ctx.Done():
		if errs := cmd.Process.Kill(); errs != nil {
			return c.out, c.err, errs
		}
		return
	case erro := <-done:
		if erro != nil {
			return c.out, c.err, erro
		}
	}
	return c.out, c.err, err
}
