package main

import (
	"log/slog"
	"os"

	signer "github.com/girvel/signer/signer/src"
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

	env := must(signer.ReadEnv())

    api := signer.CreateAPI(cryptographer, env)
    slog.Info("initialized API; running it...")

    api.Run()
}
