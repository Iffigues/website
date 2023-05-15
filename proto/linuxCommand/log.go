package linuxCommand

import (
	"strings"
	"os"
	"fmt"
)

func (c *Command)logs() {
	f, err := os.OpenFile("/log/linuxtool.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	fmt.Println("logs= ", c.command, c.args)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	ar := append([]string{c.command}, c.args...)
	fmt.Println(ar)
	tt := strings.Join(ar[:], " ")
	tt = tt + "\n"
	f.WriteString(tt)

}
