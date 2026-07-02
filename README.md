keyring-keyctl
==============
[![CI](https://github.com/lox/keyring-keyctl/actions/workflows/test.yml/badge.svg?branch=master)](https://github.com/lox/keyring-keyctl/actions/workflows/test.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/lox/keyring-keyctl.svg)](https://pkg.go.dev/github.com/lox/keyring-keyctl)

Linux keyctl provider for [`github.com/lox/keyring/v2`](https://github.com/lox/keyring).

## Usage

```bash
go get github.com/lox/keyring-keyctl
```

```go
import (
	"context"

	"github.com/lox/keyring/v2"
	keyctl "github.com/lox/keyring-keyctl"
)

ctx := context.Background()

ring, err := keyring.Open(ctx,
	keyring.WithServiceName("example"),
	keyring.WithProvider(keyctl.Provider()),
)
```

`keyctl.Provider` accepts `Scope` and `Perm` options. On non-Linux platforms,
it returns `keyring.ErrUnavailable` during open.
