## v0.1: API

- [x] GET /public
  - [x] dummy endpoint
  - [x] connect private.pem, generate a public key
- [x] POST /sign
- [x] POST /verify
- [x] Extract src/crypto.go
- [x] Test /sign + /verify
- [x] `Content-Type: text/plain` for GET /public and POST /sign

- [x] Finalization 1
  - [x] Logging
  - [x] Swagger

- [x] Include signing date

The schema would become:

- [x] GET /public -> plaintext
- [x] POST /sign json(text) -> json(dated_text, signature)
- [x] POST /verify json(dated_text, signature) -> json()

- [ ] Finalization 2
  - [x] docker-compose
  - [x] Better README
  - [x] Do go library paths work with monorepos? (rename signer -> signer/signer)
  - [ ] TODOs
  - [ ] Review

- [ ] Maybe it's better to keep everything in plain text, so it can easily be saved into a text
      file; look up if there is a convention for that.

## v0.2: Front-end

Console-like UI for the thing
