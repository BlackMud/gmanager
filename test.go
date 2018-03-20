package main

import (
	"fmt"
	"os"
	//"path"
	"strings"
	//"path/filepath"
)

func main() {
	//filenm := "./README.md"
	dir := "./src"
	s, _ := os.Stat(dir)
	fmt.Println(s.Name())

	addr := "../usr/local/src..\\nihao\\wo"

	str := strings.Split(addr, "\\")
	fmt.Println(strings.Join(str, "/"))

	/*
		if dirstats.IsDir() {
			err = filepath.Walk(dirnm, func(path string, f os.FileInfo, err error) error {
				if f.IsDir() {
					fmt.Println("dir:", path)
					return nil
				}
				if f == nil {
					return err
				}
				fmt.Println("file:", path)
				return nil

			})
		}
		if err != nil {
			fmt.Printf("err : %v\n", err)
		}
	*/

}
