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
- [x] Attach to the service
  - [x] Sign
  - [x] Verify
- [x] Reset output border color when text changes
- [x] Indicate if text & signed text don't match
- [x] Columns & arrow
- [x] Ctrl+Enter while input textarea is active -> sign
- [x] Ctrl+Enter while output textarea is active -> verify
- [x] Controller object
- [ ] Copy button
- [ ] Report service unavailability
- [ ] Resolve TODOs
- [ ] Review

- [ ] Thoughts:
  - [ ] docker compose for the backend only
  - [ ] docker health checks? Are they needed for a single-service setup?
  - [ ] update dating: replace /public reference with reference to the frontend
  - [x] describe justification in the frontend page: assert that someone claimed something at a
        given date, no server-side storage, secured with modern cryptography
  - [ ] how to deploy JS?
  - [x] made by girvel
  - [ ] display owner

- [ ] Maybe later:
  - [ ] Timezone header?
  - [ ] Checkbox to use timezone header?
  - [ ] Color scheme

## Deploy

- [ ] set up the server

## Authentication?
