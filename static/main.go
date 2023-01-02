package main

import  (
	"io/fs"
	"fmt"
	"embed"
)

var (
	//go:embed statics
	static   embed.FS
)

type files struct {
	path string
	files fs.DirEntry
}

func GetContent(static embed.FS, a string) (dir []fs.DirEntry, file []fs.DirEntry, err error) {
	e, ee := static.ReadDir(a)

	if ee != nil {
		err = ee
		return
	}
	for _, aa := range e {
		if aa.IsDir() {
			dir = append(dir, aa)
		} else {
			file = append(file, aa)
		}
	}
	return
}


func FindFile(static embed.FS, file string, o int)(path string) {
	return
}

func main() {
	a, aa , err := GetContent(static, "statics")
	fmt.Println(a, aa, err)
	for _, e := range a {
		fmt.Println(e)
	}

}
