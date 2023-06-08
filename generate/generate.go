package generate

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"github.com/spf13/cobra"
)

var bitSize = 2048

var Base64RSAPair = &cobra.Command{
	Use:   "b64rsa",
	Short: "Generate a base64 encoded RSA key pair",
	Long:  "First an asymmetric RSA key pair will be generated, then the base64 encoded versions will be printed",
	Run: func(cmd *cobra.Command, args []string) {
		prvKey, err := generateRSAPrvKey()
		if err != nil {
			panic(err)
		}
		prvPEM, pubPEM, err := generatePEMBlocks(prvKey)
		if err != nil {
			panic(err)
		}

		prvB64 := base64.StdEncoding.EncodeToString(prvPEM)
		pubB64 := base64.StdEncoding.EncodeToString(pubPEM)

		fmt.Println("-------- DO NOT COPY THIS LINE, BASE64 ENCODED PRIVATE KEY BELOW -------")
		fmt.Printf("%+v\n", prvB64)
		fmt.Println("-------- DO NOT COPY THIS LINE, BASE64 ENCODED PRIVATE KEY END -------")
		fmt.Println("-------- DO NOT COPY THIS LINE, BASE64 ENCODED PUBLIC KEY BELOW -------")
		fmt.Printf("%+v\n", pubB64)
		fmt.Println("-------- DO NOT COPY THIS LINE, BASE64 ENCODED PUBLIC KEY END -------")

	},
}

func generateRSAPrvKey() (*rsa.PrivateKey, error) {
	prvKey, err := rsa.GenerateKey(rand.Reader, bitSize)
	if err != nil {
		return nil, err
	}

	if err := prvKey.Validate(); err != nil {
		return nil, err
	}

	return prvKey, nil
}

func generatePEMBlocks(k *rsa.PrivateKey) ([]byte, []byte, error) {
	prvBytes, err := x509.MarshalPKCS8PrivateKey(k)
	if err != nil {
		return nil, nil, err
	}
	var prvPEMBlock = &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: prvBytes,
	}
	pubBytes, err := x509.MarshalPKIXPublicKey(&k.PublicKey)
	if err != nil {
		return nil, nil, err
	}
	var pubPEMBlock = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubBytes,
	}

	return pem.EncodeToMemory(prvPEMBlock), pem.EncodeToMemory(pubPEMBlock), nil
}
