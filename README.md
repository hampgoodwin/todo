# todo

A simple todo applications

[![Go Report Card](https://goreportcard.com/badge/github.com/hampgoodwin/todo)](https://goreportcard.com/report/github.com/hampgoodwin/todo) [![Coverage Status](https://coveralls.io/repos/github/hampgoodwin/GoLuca/badge.svg)](https://coveralls.io/github/hampgoodwin/GoLuca) [![golangci-lint](https://github.com/hampgoodwin/todo/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/hampgoodwin/todo/actions/workflows/golangci-lint.yml)
[![buf-lint](https://github.com/hampgoodwin/todo/actions/workflows/buf-lint.yml/badge.svg)](https://github.com/hampgoodwin/todo/actions/workflows/buf-lint.yml)

## Quickstart

```sh
make up && make run
# Optionally, to view the todo created events, run the wiretap in another shell window
make runwiretap
```
This will start the underlying database and start the todo application server

## Tooling

- [buf](https://buf.build/)
    - ```
      brew install bufbuild/buf/buf 
      ```
- [gofumpt](https://github.com/mvdan/gofumpt) for formatting.
- [golangci-lint](https://github.com/golangci/golangci-lint) for linting.
    - primarily ran as a container for local and ci flows, no installation is not really necessary.
- [colima](https://github.com/abiosoft/colima) for container runtimes.
    - colima is a free docker desktop alternative. However, this application _should_ work with whatever container runtime you use.
- [jaeger](https://www.jaegertracing.io/) for local trace collector and ui.
- [nats](https://nats.io/) JetStream for message bus.