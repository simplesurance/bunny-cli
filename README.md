# bunny-cli
![CI](https://github.com/simplesurance/bunny-cli/actions/workflows/ci.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/simplesurance/bunny-cli)](https://goreportcard.com/report/github.com/simplesurance/bunny-cli)

bunny-cli is a command line tool to interact with the [Bunny.net
API](https://bunny.net).

The tool was created to manually test the
[github.com/simplesurance-bunny-go](https://github.com/simplesurance/bunny-go)
package and make experimenting with the bunny API more user-friendly.

## Getting Started

### Compilation

Run:

```sh
make build
```

### Usage

See the help output:
```sh
./bunny-cli --help
```

### Authentication

The Bunny.net API key can either be passed as parameter on the command line:

```sh
./bunny-cli --api-key API-KEY pullzone list
```

or via the `BUNNY_API_KEY` environment variable:

```sh
export BUNNY_API_KEY=API-KEY
./bunny-cli API-KEY pullzone list
```

## Status

The package is under initial development. \
Incompatible changes to the command line interface can happen any time.
