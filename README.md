# envs

A CLI tool to manage multiple `.env` files. Switch between environments (dev, prod, staging) with a single command.

## How it works

`envs` keeps a `.env.sample` as a template and creates named environment files (`.env.dev`, `.env.prod`, etc). A config file (`.envs.json`) tracks which environment is active. When you switch environments, the active `.env.*` file is copied to `.env`.

## Usage

```bash
envs init                  # set up envs in the current directory
envs new dev               # create .env.dev from .env.sample
envs new prod --empty      # create .env.prod without copying sample
envs list                  # list all environments (* marks active)
envs use dev               # switch active environment to dev
```


## TODO

- [ ] Implement `use` command (switch active environment, copy to `.env`)
- [ ] Implement `current` command (show active environment)
- [ ] Implement `set` / `get` commands (manage individual env vars)
- [ ] Implement `validate` command (check env against sample)
- [ ] Add tests for `new` and `use` commands
