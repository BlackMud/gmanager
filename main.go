package main

import (
	"flag"
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"os"
)

//自定义命令行的参数变量
var (
	ip      = flag.String("i", "", "input host ip")
	op_type = flag.String("m", "", "input the module which you will chose and do sth...")
	src_fd  = flag.String("src", "", "input source file")
	dst_fd  = flag.String("dst", "", "input dest file or directory")
	config  = &ssh.ClientConfig{User: "root", Auth: []ssh.AuthMethod{ssh.Password("root12300.")}}
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

	//if ok, _ := PathExist(*src_fd); !ok {
	//	log.Fatal("file or directory is not exists")
	//}

	//处理连接过程中的信息
	ce := func(err error, msg string) {
		if err != nil {
			log.Fatalf("%s error: %v", msg, err)
		}
	}

	//请求建立链接
	client, err := ssh.Dial("tcp", *ip, config)
	ce(err, "Dial")

	//创建连接回话
	session, err := client.NewSession()
	ce(err, "new session")
	defer session.Close()

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}

	err = session.RequestPty("xterm", 25, 80, modes)
	ce(err, "request pty")

	err = session.Shell()
	ce(err, "start shell")

	err = session.Wait()
	ce(err, "return")

}
