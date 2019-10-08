<p align="center">
<img width="200" height="200" src="./logo.svg" alt="gnet">
<br /> <br />
<a title="Go Report Card" target="_blank" href="https://goreportcard.com/report/github.com/ImVexed/muon"><img src="https://goreportcard.com/badge/github.com/ImVexed/muon?style=flat-square"></a>
<img alt="GitHub code size in bytes" src="https://img.shields.io/github/languages/code-size/ImVexed/muon">
<br/>
<a target="_blank" href="https://gowalker.org/github.com/ImVexed/muon"><img src="https://img.shields.io/badge/api-reference-blue.svg?style=flat-square"></a>
</p>

----

`Muon` is a lightweight alternative to Electron written in Golang in about ~300 LoC, using Ultralight instead of Chromium. [Ultralight](https://tinyurl.com/2fcpre6) is a cross-platform WebKit rewrite using the GPU to target embeded desktop applications that resulted in a fast, lightweight, and low-memory HTML UI solution that blends the power of Chromium with the small footprint of Native UI.


# Features

- Full JS to Go interop
- GPU based rendering
- Cross-platform
- Hot-reloading
- Superior disk size + memory & cpu usage

Comparison with a "Hello, World!" React App

|      | Muon    | Electron |
|:----:|---------|----------|
| CPU  | 0.0%    | 1.2%     |
| MEM  | 26.0 MB | 201.7 MB |
| DISK | 42 MB   | 136 MB   |

# Example

From `examples/create-react-app/main.go`:
```go
package main

import (
  "github.com/ImVexed/muon"

  "cra-go/webfiles"
  "net/http"
)

func main() {
  // Any static asset packer of your liking (ex. fileb0x)
  fileHandler := http.FileServer(webfiles.HTTP)

  cfg := &muon.Config{
    Title:      "Hello, World!",
    Height:     500,
    Width:      500,
    Tilted:     true,
    Resizeable: true,
  }

  m := muon.New(cfg, fileHandler)

  // Expose our `add` function to the JS runtime
  m.Bind("add", add)

  // Show the Window and start the Runtime
  if err := m.Start(); err != nil {
    panic(err)
  }
}

// Muon automatically handles interop to and from the JS runtime
func add(a float64, b float64) float64 {
  return a + b
}
```

# FAQ

## Q: *How are JS types translated to Go types?*
- JS: `Boolean` Go: `bool`
- JS: `Number`  Go: `float64`
- JS: `String`  Go: `string`
- JS: `Object`  Go: `struct` via JSON

## Q: *I get `exit status 3221225781` when I try to run my program*
- Your program likely can't find the Ultralight libraries. Ensure they're either installed on the system, or, in the same folder as your program. Currently, Muon uses the 1.1 Ultralight pre-release that hasn't yet propogated to their main site and can only be downloaded from the [Ultralight](https://github.com/ultralight-ux/Ultralight#getting-the-latest-sdk) github repo.

## Q: *How do I get rid of the Console window on Windows?*
- Add `-ldflags -H=windowsgui` to either your `go build` or `go run` to get rid of the window.
