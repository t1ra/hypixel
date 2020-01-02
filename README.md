# hypixel

[![GoDoc](https://img.shields.io/badge/Godoc-Reference-%2300ADD8?style=flat-square)](https://godoc.org/github.com/t1ra/hypixel)
[![Go Report Card](https://img.shields.io/badge/Go%20Report-A%2B-%2300ADD8?style=flat-square)](https://goreportcard.com/report/github.com/t1ra/hypixel)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/t1ra/hypixel?color=%2300ADD8&style=flat-square)
![License ISC](https://img.shields.io/badge/License-ISC-%2300ADD8?style=flat-square)

<img align="right" src="https://raw.githubusercontent.com/ashleymcnamara/gophers/master/Azure_Bit_Gopher.png" alt="Gaming Gopher" width="500px" height="auto">

**hypixel** is a [Hypixel API](https://github.com/HypixelDev/PublicAPI) library for Go.
It provides near full coverage of Hypixel's public API (see TODO below) in a lightweight package.

Some changes are made internally to try and keep the API consistent on the
library's end, but the Hypixel api is very inconsistent so check the documentation
if something looks like it should work, but doesn't.

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

3. Follow documentation and its examples to construct your application.

## Documentation

As is the style of Go code, all structs and functions in Hypixel are documented in the source. This
allows Godoc to automatically generate very useful documentation.

For documentation on methods

* [![GoDoc](https://godoc.org/github.com/t1ra/hypixel?status.svg)](https://godoc.org/github.com/t1ra/hypixel)

And for documentation on structs

* [![GoDoc](https://godoc.org/github.com/t1ra/hypixel/structs?status.svg)](https://godoc.org/github.com/t1ra/hypixel/structs)

There isn't any hand-written documentation just yet.

## Contributing

Contributions, whether in the form of Pull Requests or Issues are encouraged.
If you want to create a pull request, make sure you follow these guidelines:

* Open an issue describing the bug or feature.

* Fork the *develop* branch and make your changes.

* Follow predefined naming conventions and run `go lint`.

* Create and run a test and/or example (if applicable) for your changes.

* Create a pull request stating your changes with a link to the issue.

## Testing

Completing Hypixel tests requires you supply an API. To do so, join Hypixel
and type `/uuid new`, and add the key to a file called `api_key` in the root of
the repository.

You'll also have to modify the Example `ExampleKey()`'s output from my UUID (`d4acada6bc844dd384060bb77e207a7a`) to whatever your UUID is.

## TODO

* https://github.com/HypixelDev/PublicAPI/tree/master/Documentation/methods/skyblock
* There are probably naming inconsistencies, especially in automatically generated structs.

## Thanks

[Ashley McNamara](https://github.com/ashleymcnamara) for the Gopher above. Taken from https://github.com/ashleymcnamara/gophers.