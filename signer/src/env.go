package signer

import (
	"fmt"
	"os"
	"strconv"
)

type Env struct {
    HolderName, ServiceAddress string
	PublicCachingSecs int
}

func ReadEnv() (*Env, error) {
    result := Env{}

	result.HolderName = os.Getenv("HOLDER_NAME")
	if result.HolderName == "" {
	    return nil, fmt.Errorf("$HOLDER_NAME environment variable is empty or undefined")
	}

	result.ServiceAddress = os.Getenv("SERVICE_ADDRESS")
	if result.ServiceAddress == "" {
	    return nil, fmt.Errorf("$SERVICE_ADDRESS environment variable is empty or undefined")
	}

	pcs_str := os.Getenv("PUBLIC_CACHING_SECS")
	if pcs_str == "" {
	    return nil, fmt.Errorf("$PUBLIC_CACHING_SECS environment variable is empty or undefined")
	}

	var err error
	result.PublicCachingSecs, err = strconv.Atoi(pcs_str)
	if err != nil {
	    return nil, fmt.Errorf("$PUBLIC_CACHING_SECS environment variable is not a valid integer")
	}

	return &result, nil
}
