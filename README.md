# Overview
This is a simple Go SDK wrapping the Lord of the Rings API https://the-one-api.dev/. This iteration covers only the Movie endpoints:

- /movie
- /movie/{id}
- /movie/{id}/quote


# Installation

This SDK requires go 1.19 and expects to be in your [GOPATH](https://golangr.com/what-is-gopath/)

```
go get github.com/chrypnotoad/chrisaf-sdk/sdk
```

Import it into your project with
```
import "github.com/chrypnotoad/chrisaf-sdk/"
```