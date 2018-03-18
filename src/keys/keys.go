package main

import (
	"fmt"
	"os"
	"os/user"
	//"bytes"
	"io/ioutil"
	"log"
	"runtime"
	//"crypto/rsa"
	//"crypto/x509"
	//"crypto/pem"
)

//type Keychain struct {
//	key *rsa.PrivateKey
//}

func GetUserPvKey() (string, error) {
	var pvkeyfile string
	usr, err := user.Current()
	if err != nil {
		log.Fatalf("failed to get keys home: %v", err)
	}
	if runtime.GOOS == "windows" {
		pvkeyfile = usr.HomeDir + "\\.ssh\\id_rsa"
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
	return string(b), nil
}

func main() {
	pvkey, err := GetUserPvKey()
	if err != nil {
		log.Fatalf("failed to getuserkey: %v", err)
	}
	fmt.Println(pvkey)

}
