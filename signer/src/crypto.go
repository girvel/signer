package signer

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

type CryptographerRSA struct {
    key *rsa.PrivateKey
}

func CreateCryptographerRSA(privateKeyPath string) (*CryptographerRSA, error) {
    privatePem, err := os.ReadFile(privateKeyPath)
    if err != nil {
        return nil, err
    }

    block, _ := pem.Decode([]byte(privatePem))
    if block == nil {
        return nil, fmt.Errorf("Failed to decode the ./private.pem")
    }

    privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
    if err != nil {
        return nil, err
    }

    rsaPrivateKey, ok := privateKey.(*rsa.PrivateKey)
    if !ok {
        return nil, fmt.Errorf("Key in ./private.pem is not RSA")
    }

    return &CryptographerRSA{rsaPrivateKey}, nil
}

var pssOptions rsa.PSSOptions = rsa.PSSOptions{
    SaltLength: rsa.PSSSaltLengthAuto,
    Hash: crypto.SHA256,
}

func (c *CryptographerRSA) Sign(data string) ([]byte, error) {
    hashed := sha256.Sum256([]byte(data))
    return rsa.SignPSS(rand.Reader, c.key, crypto.SHA256, hashed[:], &pssOptions)
}

func (c *CryptographerRSA) Verify(data string, signature []byte) error {
    hashed := sha256.Sum256([]byte(data))

    return rsa.VerifyPSS(
        &c.key.PublicKey,
        crypto.SHA256,
        hashed[:],
        signature,
        &pssOptions,
    )
}

func (c *CryptographerRSA) Public() string {
    der, err := x509.MarshalPKIXPublicKey(&c.key.PublicKey)
    if err != nil {
        panic(err.Error())
    }

    pem := pem.EncodeToMemory(&pem.Block{
        Type: "RSA PUBLIC KEY",
        Bytes: der,
    })

    return string(pem)
}
