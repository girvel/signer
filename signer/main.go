package main

import (
	signer "github.com/girvel/signer/src"
)

func main() {
    api := signer.CreateAPI()
    api.Run()
}
