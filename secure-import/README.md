# kms-encrypt-nonce

A lightweight tool to encrypt nonce with AES-GCM or AES-CBC encryption on a nonce that is distributed by Key Protect or Hyper Protect Crypto Services.

<!-- Add [Learn more]() link to tutorial -->

## Usage

1. From [releases](https://github.com/IBM-Cloud/kms-samples/releases), download the binary that is compatible with your operating system.

2. Mark the file as executable using `chmod`.

    ```
    chmod +x ./kms-encrypt-nonce
    ```

3. Run the script to encrypt a nonce value with an AES symmetric key.
    
    **Note:** In case you need to use `CBC` encryption (used by Hyper Protect Crypto Services but not Key-Protect service), set the flag `-alg CBC`.

    ```
    ./kms-encrypt-nonce -key $KEY_MATERIAL -nonce $NONCE [-alg CBC]
    ```
    The output displays the `encryptedNonce` and `iv` values that are used to verify a secure import request to the Key Protect service. The following snippet shows example values.

    ```
    {
        "encryptedNonce": "DVy/Dbk37X8gSVwRA5U6vrHdWQy8T2ej+riIVw==",
        "iv": "puQrzDX7gU1TcTTx"
    }