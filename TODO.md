# TODO

## UI

- [ ] ditch goat ui (honestly was a nice try)
- [ ] daysi ui + tw typography (for makrdown) + some syntaxt highlighter
    - [ ] rewrite all page head to manually include global

## GENERAL

- [ ] always use client side redirection, to avoid theme flicker -> do I need to do something with the `<head>` ?
- [x] combine lytepage repo with goat repo
    - [x] basic static html rendering
    - [x] dynamic frontmatter -> `map[string]any`
    - [x] template for md rendering
    - [x] special route handling
        - [x] index
        - [x] 404 as "catch all"
        - [x] tag
        - [x] category
    - [x] options for formatting html
    - [ ] ssr for dev?
- [ ] test!!!
- [ ] do I need preflight? (not)

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

## MARKDOWN

- [x] add links to example pages into homepage
- [ ] styling
