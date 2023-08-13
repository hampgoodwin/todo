# todo

A simple todo applications

[![Go Report Card](https://goreportcard.com/badge/github.com/hampgoodwin/todo)](https://goreportcard.com/report/github.com/hampgoodwin/todo) [![golangci-lint](https://github.com/hampgoodwin/todo/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/hampgoodwin/todo/actions/workflows/golangci-lint.yml)
[![buf-lint](https://github.com/hampgoodwin/todo/actions/workflows/buf-lint.yml/badge.svg)](https://github.com/hampgoodwin/todo/actions/workflows/buf-lint.yml)

## Quickstart

```sh
make up && make run
# Optionally, to view the todo created events, run the wiretap in another shell window
make runwiretap
```

This will start the underlying database and start the todo application server.

To test start up your favorite gRPC client (or gui. I like postman and now deprecated bloomRPC).

Either source the proto files from this repository to fill in your grpc tooling, or make this request:

```json
{
  "create_to_dos": [
    {
      "message": "yubblebubble",
      "details": "gotta do the yubble bubble",
      "due_date": {
        "seconds": 20,
        "nanos": 10
      },
      "priority": 3,
      "level_of_effort": 3
    }
  ]
}
```

This should be called against the CreateToDos endpoint. Feel free to modify the message or details. Also, modifying the priority and level of effort should work. However, they are enums and have limits. If an incorrect value is used, it will default to `unspecified` so as to not cause an error. This can be modified.

The ListToDos endpoint example is below

```json
{
  "ids": [
    {valid_ksuid}
  ],
  "page_size": 10,
  "page_token": ""
}
```

There is rudimentary pagination on this service. If you create a few to dos, then performa  list with a low page size, you'll get back a next page token. You can use that in the subsequent request. You may, of course, omit the `ids` filter, as it is optional.

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
