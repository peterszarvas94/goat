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
- [ ] slogger: default instead of export
- [ ] embed this repo into goat repo

- [ ] bump version script
- [ ] rewrite readme

## BARE

- no tasks

## BASIC AUTH

- [ ] updated_at
- [ ] email verification
- [ ] profile page to change name / email / password

## OAUTH

- [ ] add google / github auth tepmlate
