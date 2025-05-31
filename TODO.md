# TODO

## GENERAL

- [x] config port != live reload templ port
- [x] no plaintext passwords
- [x] session timeout
- [x] post mvc
- [x] rewrite all redirect and context
- [x] req_id to each request, also log it, and save to db (e.g. login -> save req_id to session)
- [x] csrf protection to post
- [x] change helpers to goat exported ones
- [x] fix js with import map
    - [x] generate importmap -> use external json file for config
    - [x] generate tsconfig paths -> use tsconfig's "extends" key to include?
- [x] goat ui
    - [x] dont use window.x = x, instead make event listeners for data attributes
- [x] better styling
- [x] page based scripts
- [x] page based css
- [x] combine goat-scaffhold repo with goat repo
- [x] slogger: default instead of export
- [ ] always use client side redirection, to avoid theme flicker -> do I need to do something with the `<head>` ?
- [ ] combine lytepage repo with goat repo
    - [x] basic static html rendering
    - [x] dynamic frontmatter -> `map[string]any`
    - [x] template for md rendering
    - [ ] special route handling (index, 404)
- [x] rewrite readme
- [ ] test!!!

## BARE

- no tasks

## BASIC AUTH

- [ ] updated_at
- [ ] email verification
- [ ] profile page to change name / email / password

## ADVANCED AUTH (these should be framework-provided functions)

- [ ] add google / github auth tepmlate
- [ ] add passkey
- [ ] add otp
- [ ] add magic link
