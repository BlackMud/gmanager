package main

import (
	"fmt"
	"os"
	//"path/filepath"
)

func main() {

	filenm := "./README.md"
	//dirnm := "./src"

	fs, err := os.Stat(filenm)
	if err != nil {
		fmt.Printf("err :%v\n", err)
		return
	}

	if fs.IsDir() {
		return
	}

	fd, err := os.Open()
	if err != nil {
		fmt.Println(err)
	}
	defer fd.Close()

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
