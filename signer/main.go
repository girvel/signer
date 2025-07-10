package main

import (
	"os"

	signer "github.com/girvel/signer/src"
)

func must[T any](result T, err error) T {
    if err != nil {
        panic(err.Error())  // TODO logging here
    }
    return result
}

func main() {
    privatePem := must(os.ReadFile("private.pem"))
    cryptographer := must(signer.CreateCryptographerRSA(privatePem))

    api := signer.CreateAPI(cryptographer)
    api.Run()
}
