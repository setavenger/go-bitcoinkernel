# go-bitcoinkernel

⚠️🚧 This library is still under construction. ⚠️🚧

`go-bitcoinkernel` is a Go wrapper around libbitcoinkernel, a C++ library exposing Bitcoin Core's validation engine. It supports both validation of blocks and transaction outputs as well as reading block data.

## Building

The library statically compiles the Bitcoin Core libbitcoinkernel library as part of its build system. Currently it targets the kernelApi branch on the following fork: https://github.com/TheCharlatan/bitcoin/tree/kernelApi.

To build this library, the usual Bitcoin Core build requirements are needed:
- cmake
- A working C and C++ compiler
- Boost library
- Go 1.24 or later

Consult the Bitcoin Core documentation for the required dependencies. Once setup, run:

```bash
go build ./...
```

## Usage

```go
package main

import (
    "github.com/setavenger/go-bitcoinkernel"
)

func main() {
    // Example usage will be added as the library develops
}
```

## License

MIT License 