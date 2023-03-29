package ssh

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"

	"golang.org/x/crypto/ssh"
)

const (
	defaultKeyLength = 2048
	permPrivateFile  = 0o600
)

var ErrNoPEMData = errors.New("no PEM data")

type KeyPair struct {
	PrivateKey []byte
	PublicKey  []byte
}

func NewKeyPair(path string) (KeyPair, error) {
	var keyPair KeyPair

	privateKey, exist, err := readFileIfExist(path)
	if err != nil {
		return keyPair, fmt.Errorf("could not read file: %w", err)
	}

	if !exist {
		return generateKeyPair(path)
	}

	keyPair.PrivateKey = privateKey

	publicKey, exist, err := readFileIfExist(path)
	if err != nil {
		return keyPair, fmt.Errorf("could not read file: %w", err)
	}

	if !exist {
		return generatePublicKey(path, keyPair)
	}

	keyPair.PublicKey = publicKey

	return keyPair, nil
}

func generateKeyPair(path string) (KeyPair, error) {
	var keyPair KeyPair

	rsaKey, err := rsa.GenerateKey(rand.Reader, defaultKeyLength)
	if err != nil {
		return keyPair, fmt.Errorf("could not generate key: %w", err)
	}

	if err := rsaKey.Validate(); err != nil {
		return keyPair, fmt.Errorf("could not validate key: %w", err)
	}

	keyPair.PrivateKey = pem.EncodeToMemory(&pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   x509.MarshalPKCS1PrivateKey(rsaKey),
	})

	//nolint: nosnakecase // because these are predefined in the os package
	privateKeyFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, permPrivateFile)
	if err != nil {
		return keyPair, fmt.Errorf("could not write to file: %w", err)
	}

	defer privateKeyFile.Close()

	_, err = privateKeyFile.Write(keyPair.PrivateKey)
	if err != nil {
		return keyPair, fmt.Errorf("could not write to file: %w", err)
	}

	publicKey, err := ssh.NewPublicKey(&rsaKey.PublicKey)
	if err != nil {
		return keyPair, fmt.Errorf("could not generate public key: %w", err)
	}

	keyPair.PublicKey = ssh.MarshalAuthorizedKey(publicKey)

	publicKeyFile, err := os.Create(fmt.Sprintf("%s.pub", path))
	if err != nil {
		return keyPair, fmt.Errorf("could not write to file: %w", err)
	}

	defer publicKeyFile.Close()

	_, err = publicKeyFile.Write(keyPair.PublicKey)
	if err != nil {
		return keyPair, fmt.Errorf("could not write to file: %w", err)
	}

	return keyPair, nil
}

func generatePublicKey(path string, keyPair KeyPair) (KeyPair, error) {
	pemData, _ := pem.Decode(keyPair.PrivateKey)
	if pemData == nil {
		return keyPair, ErrNoPEMData
	}

	rsaKey, err := x509.ParsePKCS1PrivateKey(pemData.Bytes)
	if err != nil {
		return keyPair, fmt.Errorf("could not decrypt private key: %w", err)
	}

	publicKey, err := ssh.NewPublicKey(&rsaKey.PublicKey)
	if err != nil {
		return keyPair, fmt.Errorf("could not generate public key: %w", err)
	}

	keyPair.PublicKey = ssh.MarshalAuthorizedKey(publicKey)

	publicKeyFile, err := os.Create(fmt.Sprintf("%s.pub", path))
	if err != nil {
		return keyPair, fmt.Errorf("could not write to file: %w", err)
	}

	defer publicKeyFile.Close()

	_, err = publicKeyFile.Write(keyPair.PublicKey)
	if err != nil {
		return keyPair, fmt.Errorf("could not write to file: %w", err)
	}

	return keyPair, nil
}

func readFileIfExist(path string) ([]byte, bool, error) {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return []byte{}, false, nil
		}

		return []byte{}, true, fmt.Errorf("could not check file %s: %w", path, err)
	}

	contents, err := os.ReadFile(path)

	return contents, true, fmt.Errorf("could not read file %s: %w", path, err)
}
