package config

import (
	"fmt"
	"os"
	"strconv"
)

var (
	Port            int
	Environment     string
	DbPath          string
	CertificatePath string
	KeyPath         string
)

const (
	PORT = iota + 1
	DB_PATH
	ENV
	CERT_PATH
	KEY_PATH
)

func init() {
	var err error
	fmt.Println("loading config...")
	if len(os.Args) < 5 && os.Args[ENV] != "dev" {
		panic("not enough arguments")
	}

	Port, err = strconv.Atoi(os.Args[PORT])
	if err != nil {
		panic(err)
	}

	DbPath = os.Args[DB_PATH]
	Environment = os.Args[ENV]
	CertificatePath = os.Args[CERT_PATH]
	KeyPath = os.Args[KEY_PATH]
}
