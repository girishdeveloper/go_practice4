package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
)

func main() {
	fi, err := os.Lstat("/dev/null") // /usr/bin/php, /home/girish/Documents/go
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The file mode is:", fi.Mode())
	fmt.Printf("permissions %#o\n", fi.Mode().Perm())
	switch mode := fi.Mode(); {
	case mode.IsRegular():
		fmt.Println("regular file")
	case mode.IsDir():
		fmt.Println("directory")
	case mode&fs.ModeSymlink != 0:
		fmt.Println("symbolic link")
	case mode&fs.ModeNamedPipe != 0:
		fmt.Println("named pipe")
	case mode&fs.ModeCharDevice != 0:
		fmt.Println("character device")
	default:
		fmt.Println("find the file type")
	}
}
