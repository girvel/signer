package main

import (
	"log/slog"
	"os"

	signer "github.com/girvel/signer/src"
)

func must[T any](result T, err error) T {
    if err != nil {
        slog.Error(err.Error())
        panic("must() failed")
    }
    return result
}

func main() {
    slog.Info("signer started")

    privatePem := must(os.ReadFile("private.pem"))
    cryptographer := must(signer.CreateCryptographerRSA(privatePem))
    slog.Info("initialized cryptographer")

    api := signer.CreateAPI(cryptographer)
    slog.Info("initialized API; running it...")

    api.Run()
}
