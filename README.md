# hypixel

[![GoDoc](https://godoc.org/github.com/t1ra/hypixel?status.svg)](https://godoc.org/github.com/t1ra/hypixel)
[![Go Report Card](https://goreportcard.com/badge/github.com/t1ra/hypixel)](https://goreportcard.com/report/github.com/t1ra/hypixel)

**hypixel** is a [Hypixel API](https://github.com/HypixelDev/PublicAPI) library for *Go*.
It provides near full coverage of Hypixel's public API (see TODO below) in a lightweight package.

## Getting Started

* master

    * The latest stable release of Hypixel with the most stable API. You should probably use it for
production tooling.

* working

    * The development branch. There may be rapid, breaking API changes. Bugs are fixed faster, but
they may be more of them.

## Installing

To get the latest stable release

```
go get github.com/t1ra/hypixel
```

But if you want to use working

```
cd $GOPATH/src/github.com/t1ra/hypixel
git checkout develop
```

## Usage

1. Import the library into your project

```
import "github.com/t1ra/hypixel"
```

2. Create a new instance of the Hypixel struct. Creating a new instance requires an API key, which can
be retrieved from in-game by running `/uuid` (or `/uuid new`).

```
hypixel, err := Hypixel.New(KEY)
```

3. Follow documentation (or examples (TODO)) to construct your application.

## Documentation

As is the style of Go code, all structs and functions in Hypixel are documented in the source. This
allows Godoc to automatically generate very useful documentation.

* [![GoDoc](https://godoc.org/github.com/t1ra/hypixel?status.svg)](https://godoc.org/github.com/t1ra/hypixel)

There isn't any hand-written documentation just yet.

## Contributing

Contributions, whether in the form of Pull Requests or Issues are encouraged.
If you want to create a pull request, make sure you follow these guidelines:

* Open an issue describing the bug or feature.

* Fork the *develop* branch and make your changes.

* Follow predefined naming conventions and run `go lint`.

* Create a pull request stating your changes with a link to the issue.

## TODO

* https://github.com/HypixelDev/PublicAPI/tree/master/Documentation/methods/skyblock
* There are probably naming inconsistencies, especially in automatically generated structs.
* Examples