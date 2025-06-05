# Testing Stuff

This repository demonstrates a small Go project with a stubbed LLM agent and a simple server that updates constants.  The `agent` package returns a deterministic constant.  The `mcpserver` binary exposes an endpoint that uses the agent to append constants to `types/generated.go` using the Go AST so that formatting remains correct.

Run the server:

```bash
go run ./cmd/mcpserver
```

Then visit `http://localhost:8080/generate` to append a constant to `types/generated.go`.

Run tests with:

```bash
go test ./...
```

