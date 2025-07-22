package signer

import (
	"fmt"
	"os"
)

type Env struct {
    HolderName string
}

func ReadEnv() (*Env, error) {
    result := Env{os.Getenv("HOLDER_NAME")}
	if result.HolderName == "" {
	    return nil, fmt.Errorf("$HOLDER_NAME environment variable is missing or undefined")
	}
	return &result, nil
}
