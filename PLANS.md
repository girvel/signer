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

- [x] Finalization 2
  - [x] docker-compose
  - [x] Better README
  - [x] Do go library paths work with monorepos? (rename signer -> signer/signer)
  - [x] TODOs
  - [x] Review script

- [ ] Review #1
  - [x] GH Actions: go tests, test docker compose up
  - [ ] Missing env_file in docker-compose.yml
  - [ ] Caching headers for /public
  - [ ] Log internal errors & don't show them to the outside; also do that for .Public()
  - [ ] All endpoints in plain text, keep both the dated text & its signature concatenated into a
        single text mass
  - [ ] Redo README again

- [ ] Release v0.1

## v0.2: Front-end

Console-like UI for the thing

- [ ] docker health checks
- [ ] update dating: replace /public reference with reference to the frontend

## Deploy

- [ ] set up the server
