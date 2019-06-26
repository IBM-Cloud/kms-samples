# kms-encrypt-nonce

##Usage

1. Download and unzip the binary that is compatible with your operating system

2. Mark the file as executable using `chmod`

    ```
    chmod +x ./kms-encrypt-nonce
    ```

3. Run the script. The output will be json so it is pipe-able to a `.json` file

    ```
    ./kms-encrypt-nonce -key $KEY_MATERIAL -nonce $NONCE
    ```