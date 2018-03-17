package main

import (
	"flag"
	//"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"net"
	"os"
	"time"
)

//自定义命令行的参数变量
var (
	ip       = flag.String("i", "", "-i <ip地址> 必须输入ip地址")
	port     = flag.String("P", "22", "-P <端口> 默认端口22")
	username = flag.String("u", "root", "-u <用户名> 默认用户:root")
	password = flag.String("p", "root12300.", "-p <密码> 登录用户密码")
	op_type  = flag.String("m", "", "-m <模块名> 输入要执行的模块")
	cmds     = flag.String("c", "", "-c <命令> 要执行的命令")
)

//设置客户端的连接登录认证方式，并取消服务端验证
func ConntMth(passwd string) (session *ssh.Session, err error) {
	var client *ssh.Client
	var config = &ssh.ClientConfig{
		User: *username,
		Auth: []ssh.AuthMethod{ssh.Password(passwd)},
		HostKeyCallback: func(hostname string, remote net.Addr, pubkey ssh.PublicKey) error {
			return nil
		},
		Timeout: 10 * time.Second,
	}

	if client, err = ssh.Dial("tcp", *ip+":"+*port, config); err != nil {
		return session, err
	}

	if session, err = client.NewSession(); err != nil {
		return session, err
	}

	return session, nil
}

//交互模式
func SshActive(sen *ssh.Session) error {
	sen.Stdout = os.Stdout
	sen.Stderr = os.Stderr
	sen.Stdin = os.Stdin

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}

	if err := sen.RequestPty("xterm", 25, 80, modes); err != nil {
		return err
	}

	if err := sen.Shell(); err != nil {
		return err
	}

	if err := sen.Wait(); err != nil {
		return err
	}
	return nil

}

//远程命令操作
func SshCmd(sen *ssh.Session, command string) error {
	sen.Stdout = os.Stdout
	sen.Stderr = os.Stderr

	if err := sen.Run(command); err != nil {
		return err
	}
	return nil
}

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
	ce := func(err error, msg string) {
		if err != nil {
			log.Fatalf("%s error: %v", msg, err)
		}
	}

	flag.Parse()
	if (*ip == "") || (*op_type == "") {
		log.Fatal("i and m 不能为空")
	}

	session, err := ConntMth(*password)
	ce(err, "ConntMth")
	defer session.Close()

	switch *op_type {
	case "ssh":
		err := SshActive(session)
		ce(err, "SshActive")
	case "cmd":
		err := SshCmd(session, *cmds)
		ce(err, "SshCmd")
	default:
		log.Fatalf("没有该模块: %s", *op_type)

	}

}
