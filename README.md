# GoToTS

> Convert your Types from TS to structs in Go and visversa

## What it is doing for now.

```go
// Give the file to read from
goToTs := NewGotots("./examples/1_type_param.go")

err := goToTs.ConvertToTs("output/example1.d.ts")
```

- You can define the file it should read and then read it to and output file

> For now only the Go to TS is working.

