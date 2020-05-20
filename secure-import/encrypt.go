/*
 * Copyright 2019, 2020 IBM Corp. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the “License”);
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *	https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an “AS IS” BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
)

// takes in key material (key) and nonce (value)
func encryptNonceWithCBC(key, value string) (string, string, error) {
	keyMat, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		panic(err)
	}

	nonce, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		panic(err)
	}

	block, err := aes.NewCipher(keyMat)
	if err != nil {
		panic(err)
	}

	// PKCS7 Padding
	paddingLength := aes.BlockSize - len(nonce)%aes.BlockSize
	paddingBytes := []byte{byte(paddingLength)}
	paddingText := bytes.Repeat(paddingBytes, paddingLength)
	nonce = append(nonce, paddingText...)

	// Generate an IV to achieve semantic security
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	cipherText := make([]byte, len(nonce))

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText, nonce)

	return base64.StdEncoding.EncodeToString(cipherText), base64.StdEncoding.EncodeToString(iv), nil
}

// takes in key material (key) and nonce (value)
func encryptNonce(key, value string) (string, string, error) {
	var cipherText []byte
	// base64 decode the input
	keyMat, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return "", "", fmt.Errorf("Failed to decode key material: %s", err)
	}
	nonce, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		return "", "", fmt.Errorf("Failed to decode nonce: %s", err)
	}
	// set up aes-gcm
	block, err := aes.NewCipher(keyMat)
	if err != nil {
		return "", "", err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", "", err
	}
	// create random iv for security to pass into aes-gcm
	iv := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err.Error())
	}
	cipherText = aesgcm.Seal(nil, iv, nonce, nil)

	// base64 encode so it's already prepared to be passed into a request
	return base64.StdEncoding.EncodeToString(cipherText), base64.StdEncoding.EncodeToString(iv), nil
}

func printEncryptedData(encryptedNonce, iv string) {

	// set up values to print as proper json
	encryptionValues := map[string]interface{}{
		"encryptedNonce": encryptedNonce,
		"iv":             iv,
	}
	js, err := json.MarshalIndent(encryptionValues, "", "	")
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%s\n", js)
}

func main() {
	var nonce, key, alg, encryptedNonce, iv string
	var err error
	flag.StringVar(&nonce, "nonce", "", "Nonce generated by Key Protect service")
	flag.StringVar(&key, "key", "", "Key material that you want to import into the Key Protect service")
	flag.StringVar(&alg, "alg", "GCM", "Algorithm to use: GCM or CBC, default is GCM")
	flag.Parse()
	if nonce == "" || key == "" {
		fmt.Println("ERROR: '-nonce' and '-key' must both be defined")
		os.Exit(1)
	}
	switch alg {
	case "GCM":
		encryptedNonce, iv, err := encryptNonce(key, nonce)
		if err != nil {
			fmt.Printf("ERROR: %s", err)
			os.Exit(1)
		}
		printEncryptedData(encryptedNonce, iv)
	case "CBC":
		encryptedNonce, iv, err = encryptNonceWithCBC(key, nonce)
		if err != nil {
			fmt.Printf("ERROR: %s", err)
			os.Exit(1)
		}
		printEncryptedData(encryptedNonce, iv)

	}
}