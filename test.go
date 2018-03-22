package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"runtime"

	"golang.org/x/crypto/ssh"
	//"golang.org/x/crypto/ssh/agent"
)

func main() {

	var pvkeyfile string
	usr, err := user.Current()
	if err != nil {
		log.Fatalf("failed to get keys home: %v", err)
	}

	if runtime.GOOS == "windows" {
		pvkeyfile = usr.HomeDir + "\\.ssh\\id_rsa"
		//fmt.Println(path.Base(pvkeyfile))
	} else {
		pvkeyfile = usr.HomeDir + "/.ssh/id_rsa"
	}

	if _, err = os.Stat(pvkeyfile); err != nil {
		log.Fatalf("file is not exists: %v", err)
	}

	b, e := ioutil.ReadFile(pvkeyfile)
	if e != nil {
		log.Fatalf("failed to read file: %v", e)
	}

	signer, err := ssh.ParsePrivateKey(b)
	if err != nil {
		log.Fatalf("unable to parse private key: %v", err)
	}
	fmt.Println(signer)
}
