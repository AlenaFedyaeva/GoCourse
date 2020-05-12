package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

//СheckFile - reports whether the named file or directory exists.
func СheckFile(fname string) bool {
	if _, err := os.Stat(fname); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// copyFileContents - copies the contents of the file named src to the file named
// by dst. The file will be created if it does not already exist. If the
// destination file exists, all it's contents will be replaced by the contents
// of the source file.
func copyFileContents(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()
	return
}

//CpUtil - утилита копирования файла
func CpUtil() {
	// argsWithoutProg := os.Args[1:]
	// fmt.Println("args: ", argsWithoutProg)
	flag.Parse()
	tail := flag.Args()
	fmt.Println("tail:", tail)
	switch len(tail) {
	default:
		fmt.Println("cp: too many arguments")
		return
	case 0:
		fmt.Println("cp: missing file operand")
		return
	case 1:
		fmt.Println("cp: missing destination file operand after ", tail[0])
		return
	case 2:
		if !СheckFile(tail[0]) {
			fmt.Println("cp: ", tail[0], " No such file")
			return
		}
	}
	copyFileContents(tail[0], tail[1])
	if СheckFile(tail[1]) {
		fmt.Println("Файл ", tail[1], " успешно скопирован")
	}

}

func main() {

	CpUtil()
	fmt.Println("The end")

}
