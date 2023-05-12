package linuxCommand

import (
	"strings"
	"os"
)

func (C *Command)logs(Command string, opt []string) {
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
