# go-bitcoinkernel

⚠️🚧 This library is still under construction. ⚠️🚧

`go-bitcoinkernel` is a Go wrapper around libbitcoinkernel, a C++ library exposing Bitcoin Core's validation engine. It supports both validation of blocks and transaction outputs as well as reading block data.

Still ironing out the build process to make sure it works for other people. Try following the PRs steps to build the library and then linking. Currently still using install_name_tool to link the lib path where bitcoinkernel is installed. Please help if you know what the issue is.

To build: 
Follow the steps of the PR https://github.com/bitcoin/bitcoin/pull/30595 under "How can I review this PR?"
```bash
cmake -B build -DBUILD_KERNEL_LIB=ON -DBUILD_UTIL_CHAINSTATE=ON
```

I also had to do this (tip: use `-j n` to build with n number of cores to speed up the build process.)
```bash
cmake --build build
```

```bash
cmake --install build
```
This will install the build 

## Building

The library statically compiles the Bitcoin Core libbitcoinkernel library as part of its build system. Currently it targets the kernelApi branch on the following fork: https://github.com/TheCharlatan/bitcoin/tree/kernelApi.


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
