# Signer

A service for signing data with RSA-PSS & SHA256.

## Verify the signature from command line:

```
openssl dgst -sha256 \
  -sigopt rsa_padding_mode:pss \
  -sigopt rsa_pss_saltlen:auto \
  -verify public.pem \
  -signature signature.bin \
  text.txt
```
