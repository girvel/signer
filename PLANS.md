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

- [x] Review #1
  - [x] GH Actions: go tests, test docker compose up
  - [x] Missing env_file in docker-compose.yml
  - [x] Caching headers for /public
  - [x] Log internal errors
  - [x] All endpoints in plain text, keep both the dated text & its signature concatenated into a
        single text mass
  - [x] Redo swagger
  - [x] Redo README again

- [x] Review 2

- [x] Release v0.1

## v0.2: Front-end

Console-like UI for the thing

- [x] Hello world
- [x] Font
- [x] UI: two text input fields, two buttons
- [ ] Attach to the service
  - [x] Sign
  - [ ] Verify
- [ ] Color scheme
- [ ] Columns & arrow
- [ ] Indicate if text & signed text don't match
- [ ] Resolve TODOs
- [ ] Report service unavailability
- [ ] Copy button
- [ ] Ctrl+Enter while input textarea is active -> sign
- [ ] Ctrl+Enter while output textarea is active -> verify
- [ ] Timezone header?
- [ ] Checkbox to use timezone header?

- [ ] Thoughts:
  - [ ] docker health checks
  - [ ] update dating: replace /public reference with reference to the frontend
  - [ ] describe justification in the frontend page: assert that someone claimed something at a
        given date, no server-side storage, secured with modern cryptography
  - [ ] how to deploy JS?
  - [ ] made by girvel
  - [ ] display owner

## Deploy

- [ ] set up the server

## Authentication?
