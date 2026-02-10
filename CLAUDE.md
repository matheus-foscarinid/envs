# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

`envs` is a Go CLI tool for managing multiple `.env` files (e.g., `.env.dev`, `.env.prod`). It tracks the active environment in `.envx.json` and provides commands to initialize, list, and switch between environments. Built with Cobra CLI framework.

## Build & Run

```bash
go build -o envs .        # build binary
go run main.go [command]   # run without building
go test ./...              # run all tests
go test ./tests/ -v        # run integration tests (verbose)
go test ./tests/ -run TestInitCreatesAllFiles  # run a single test
go vet ./...               # check for issues
go mod tidy                # sync dependencies
```

## Architecture

Entry point (`main.go`) calls `cmd.Execute()`. Commands live in `cmd/` and delegate to internal packages for business logic.

- **cmd/**: Cobra command definitions (root, init, list)
- **internal/manager/**: core business logic (listing envs, resolving active env)
- **internal/config/**: config struct and read/write for `.envx.json`
- **internal/env/**: operations on `.env` files (load, write, copy)
- **internal/constants/**: shared file name constants (`SampleFile`, `DotEnvFile`, `ConfigFile`)

Data flows: CLI command → `cmd/` handler → `internal/manager/` → `internal/config/` + `internal/env/` → filesystem.

## Key Files

- `.envx.json`: runtime config tracking version and active environment name
- `.env.sample`: template for environment variables
- `.env`: the active environment file (generated, git-ignored)
