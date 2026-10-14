# Go language toolkit revision

Read `main.go` from top to bottom: grouped definitions explain the concepts, then `main()` demonstrates them in numbered sections.

From this directory:

```bash
go run .
gofmt -w main.go
```

The local `go.mod` lets normal package commands work. Run these commands inside this exercise, not the repository root.

## Revision checklist

- [ ] Structs, methods, zero values, whitespace validation.
- [ ] Typed errors, `errors.As`, `errors.Is`, `%v` versus `%w`.
- [ ] Implicit interfaces and compile-time implementation checks.
- [ ] Composition and delivery-error propagation.
- [ ] Injected clocks, ID generators, closures, and per-generator state.
- [ ] Slices, ordered batch results, continuing after errors, empty input.
- [ ] Value and pointer receivers, shared backing arrays, shallow copies.
- [ ] Missing map entries versus false, nil-map reads.
- [ ] Comparable structs and valid map keys.
- [ ] File append, open/write/close errors, `defer`, named return values.

`ConsoleSender` retains its original direct-call validation. The service validates before every delivery so other sender implementations can focus on delivery. The obsolete standalone `Notify` forwarding function has been removed to avoid confusing it with the service method.

`go run .` appends two valid notifications to `notifications.log` in the working directory. The ID counter is sequential and local to each generator. This exercise expects all service dependencies to be provided.

Finish by explaining each output and rebuilding the core service without copying. This completes the language exercise. Continue to session 02 in the root STUDY_PLAN.md for design essentials.
