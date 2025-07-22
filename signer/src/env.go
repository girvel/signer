package signer

import (
	"fmt"
	"os"
)

type Env struct {
    HolderName, ServiceAddress string
}

func ReadEnv() (*Env, error) {
    result := Env{}

	result.HolderName = os.Getenv("HOLDER_NAME")
	if result.HolderName == "" {
	    return nil, fmt.Errorf("$HOLDER_NAME environment variable is missing or undefined")
	}

	result.ServiceAddress = os.Getenv("SERVICE_ADDRESS")
	if result.ServiceAddress == "" {
	    return nil, fmt.Errorf("$SERVICE_ADDRESS environment variable is missing or undefined")
	}

	return &result, nil
}
