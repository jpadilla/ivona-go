# Go IVONA [![Build Status](https://travis-ci.org/jpadilla/ivona-go.svg?branch=master)](https://travis-ci.org/jpadilla/ivona-go)

Go client library for [IVONA](http://www.ivona.com/us/).

## Supported API Calls

- CreateSpeech

## Versioning

Each revision of the binding is tagged and the version is updated accordingly.

Given Go's lack of built-in versioning, it is highly recommended you use a
[package management tool](https://code.google.com/p/go-wiki/wiki/PackageManagementTools) in order
to ensure a newer version of the binding does not affect backwards compatibility.

To see the list of past versions, run `git tag`. To manually get an older
version of the client, clone this repo, checkout the specific tag and build the
library:

```sh
git clone https://github.com/jpadilla/ivona-go.git
cd ivona-go
git checkout api_version_tag
make build
```

## Installation

```
go get github.com/jpadilla/ivona-go
```

## Documentation

For details on all the functionality in this library, see the [GoDoc](http://godoc.org/github.com/jpadilla/ivona-go) documentation.

## Example usage

```go
package main

import (
    "log"

    ivona "github.com/jpadilla/ivona-go"
)

func main() {
    client := ivona.New("IVONA_ACCESS_KEY", "IVONA_SECRET_KEY")
    options := ivona.NewSpeechOptions("Hello World")
    r, err := client.CreateSpeech(options)

    if err != nil {
        log.Fatal(err)
    }

    log.Printf("%v\n", len(r.Audio))
    log.Printf("%v\n", r.ContentType)
    log.Printf("%v\n", r.RequestID)
}

```
