# Signer

A service, allowing to assign a date to some text with cryptographic verification. Uses RSA-PSS & SHA256.

## Concept

Takes any text, like `John Doe bets the Moon is fake`, and transforms it into a signed and dated claim:

```
John Doe bets the Moon is fake

Signed 2025-07-22 13:37 by girvel

Mb27yOo4ffqcTmMhhX/2AW4KRLgwsetgTWPWZifzVSY1HIObYjYrX3yMMabBdKl4BtyOW7D432IVgvx7F+Hpv5pAS236p9VYKwzpXzzkMziR4Q32z8yYqLsYT3o4tClDEqqLoxZYuXUWA781nnjculthEz6OK9lXYWXu+hvdJDcXCnlBqQ5x4cGPzsG/2bvPpWX7KYaiQJLaceICxvZetcKy5m9LvX86zxrVPI7o1nQRSsedfFS1d7GZRoKM7ircKAJnSdFT3ICNMTRJd57rtnoUT1GH6jLc4tlHiFVdH+ns4YCHpUvOIf501/6v86hAfCAqDyylK8rVTXiQtUkcNw=
```

And then in 2026 the second wave of lunar expeditions hits, and people discover that the Moon was faked by the American government, John Doe can present this pair and show, that it's indeed him, who predicted this as early as July 2025. With this service, it becomes cryptographically verifiable fact, dependent only on the reputation of `girvel` as a signer.

## Endpoints

All endpoints accept and respond with plain text.

- `GET /public`: see the public key
- `POST /sign`: sign the provided text
- `POST /verify`: verify the provided signed text

## Launch

Via docker compose:

1. Create `.env` file (see `.env.example`)
2. Put your private RSA key into `signer/private.pem` (in PEM PKCS#8 format)

```bash
docker compose up
```

3. Go to `localhost:8080/swagger/index.html` to see docs.

Also the signer service itself can be launched via `go run .` inside the folder.

## Verify the signature from command line:

(In the example, dated text goes into `text.txt`, signature into `signature.bin`, public key into `public.pem`)

```
openssl dgst -sha256 \
  -sigopt rsa_padding_mode:pss \
  -sigopt rsa_pss_saltlen:auto \
  -verify public.pem \
  -signature signature.bin \
  text.txt
```
