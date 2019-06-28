# kms-encrypt-nonce

This script helps you run AES-GCM encryption on a nonce that is distributed by Key Protect. 

<!-- Add [Learn more]() link to tutorial -->

## Usage

1. Download and untar the binary that is compatible with your operating system.

2. Mark the file as executable using `chmod`.

    ```
    chmod +x ./kms-encrypt-nonce
    ```

3. Run the script to encrypt a nonce value with an AES symmetric key. 

    ```
    ./kms-encrypt-nonce -key $KEY_MATERIAL -nonce $NONCE
    ```

    The output displays the `encryptedNonce` and `iv` values that are used to verify a secure import request to the Key Protect service. The following snippet shows example values.

    ```
    {
        "encryptedNonce": "DVy/Dbk37X8gSVwRA5U6vrHdWQy8T2ej+riIVw==",
        "iv": "puQrzDX7gU1TcTTx"
    }