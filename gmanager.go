package main

import (
	"flag"
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"key"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
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
	srcpath  = flag.String("s", "", "-s <目录或者文件路径> 要拷贝的文件或者目录")
	dstpath  = flag.String("d", "", "-d <目录或文件> 拷贝到目标机器上文件或目录")
)

//设置客户端的连接登录认证方式，并取消服务端验证
func ConntMth() (client *ssh.Client, err error) {
	var config = &ssh.ClientConfig{
		User:            *username,
		Auth:            []ssh.AuthMethod{ssh.Password(*password)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         30 * time.Second,
	}
	addr := fmt.Sprintf("%s:%s", *ip, *port)
	if client, err = ssh.Dial("tcp", addr, config); err != nil {
		return client, err
	}
	return client, nil
}

//远程登录进行交互式操作
func SshActive(sshclient *ssh.Client) error {
	session, err := sshclient.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}
	if err := session.RequestPty("xterm", 25, 80, modes); err != nil {
		return err
	}
	if err := session.Shell(); err != nil {
		return err
	}
	if err := session.Wait(); err != nil {
		return err
	}
	return nil

}

//远程命令操作
func SshCmd(sshclient *ssh.Client, command string) error {
	session, err := sshclient.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	if err := session.Run(command); err != nil {
		return err
	}
	return nil
}

//上传文件
func SshCopyPath(sshclient *ssh.Client, srcpath, dstpath string) error {
	sftpclient, err := sftp.NewClient(sshclient)
	if err != nil {
		return err
	}
	defer sftpclient.Close()

	buf := make([]byte, 1024)
	str := srcpath

	err = filepath.Walk(srcpath, func(pth string, f os.FileInfo, err error) error {
		if runtime.GOOS == "windows" {
			sarr := strings.Split(pth, "\\")
			pth = strings.Join(sarr, "/")
		}
		sli := strings.Split(pth, str)

		if f == nil {
			return err
		}

		if f.IsDir() {
			dstDir := path.Join(dstpath, path.Base(str)+sli[1])
			fmt.Println(dstDir)
			if err = sftpclient.Mkdir(dstDir); err != nil {
				return nil
			}
			return nil
		}

		srcFile, e := os.Open(pth)
		if e != nil {
			return e
		}
		defer srcFile.Close()

		dstFile, err := sftpclient.Create(path.Join(dstpath, path.Base(str)+sli[1]))
		if err != nil {
			return err
		}
		defer dstFile.Close()

		for {
			n, _ := srcFile.Read(buf)
			if n == 0 {
				break
			}
			dstFile.Write(buf)
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

//下载文件
func SshGetFile(sshclient *ssh.Client, srcpath, dstpath string) error {
	sftpclient, err := sftp.NewClient(sshclient)
	if err != nil {
		return nil
	}
	sftpclient.Close()

	buf := make([]byte, 1024)
	//str := srcpath

	err = filepath.Walk(dstpath, func(pth string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}

		if f.IsDir() {
			localdir := path.Base(pth)
			if err := os.Mkdir(path.Join(srcpath, localdir), 0644); err != nil {
				return err
			}
			return nil
		}

		srcFile, e := sftpclient.Open(pth)
		if e != nil {
			return e
		}
		defer srcFile.Close()

		srcfilename := path.Base(pth)
		dstFile, err := os.Create(path.Join(srcpath, srcfilename))
		if err != nil {
			return err
		}
		defer dstFile.Close()

		for {
			n, _ := srcFile.Read(buf)
			if n == 0 {
				break
			}
			dstFile.Write(buf)
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func main() {
	ce := func(err error, msg string) {
		if err != nil {
			log.Fatalf("%s error: %v", msg, err)
		}
	}

	flag.Parse()
	if *ip == "" {
		log.Fatal("i不能为空")
	}

	sshclient, err := ConntMth()
	ce(err, "ConntMth")

	switch *op_type {
	case "ssh":
		err := SshActive(sshclient)
		ce(err, "SshActive")

	case "cmd":
		err := SshCmd(sshclient, *cmds)
		ce(err, "SshCmd")

	case "copy":
		if runtime.GOOS == "windows" {
			s := strings.Split(*dstpath, "Git")
			*dstpath = s[len(s)-1]
		}
		err := SshCopyPath(sshclient, *srcpath, *dstpath)
		ce(err, "Sshcopypath")
		log.Println("上传完成")

	case "get":
		if runtime.GOOS == "windows" {
			s := strings.Split(*dstpath, "Git")
			*dstpath = s[len(s)-1]
		}
		err = SshGetFile(sshclient, *srcpath, *dstpath)
		ce(err, "SshGetFile")
		log.Println("下载完成")

	default:
		log.Fatalf("没有该模块: %s", *op_type)
	}

}
