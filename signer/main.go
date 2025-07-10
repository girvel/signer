package main

import (
	signer "github.com/girvel/signer/src"
)

func main() {
    cryptographer, err := signer.CreateCryptographerRSA("private.pem")
    if err != nil {
        panic(err.Error())
    }

    api := signer.CreateAPI(cryptographer)
    api.Run()
}
