package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

//自定义命令行的参数变量
var (
	ip      = flag.String("i", "", "input host ip")
	op_type = flag.String("m", "", "input the module which you will chose and do sth...")
	src_fd  = flag.String("src", "", "input source file")
	dst_fd  = flag.String("dst", "", "input dest file or directory")
)

//复制文件或目录到目标主机上
//func scp(ip, src_fd, dst_fd string) (status bool) {
//
//}

//判断源文件或者目录是否存在
func PathExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, nil
}

func main() {
	flag.Parse()

	if (*ip == "") || (*op_type == "") {
		log.Fatal("i and m 不能为空")
	}

	ok, _ := PathExist(*src_fd)
	if !ok {
		log.Fatal("file or directory is not exists")
	}

	fmt.Printf("src file: %s\n", *src_fd)
	fmt.Println(flag.Args())
}
