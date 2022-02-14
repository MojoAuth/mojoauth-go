# MojoAuth Go SDK

## Documentation

* [Configuration](https://mojoauth.com/docs/) - Everything you need to begin using the MojoAuth SDK.

## Installation

To install, run:
`go get github.com/mojoauth/go-sdk`

Import the package:

`import "github.com/mojoauth/go-sdk"`

Install all package dependencies by running `go get ./...` in the root folder of this SDK.  

## Usage

Take a peek:

Before making any API calls, the MojoAuth API client must be initialized with your MojoAuth API key.

Sample code:

```
cfg := go-sdk.Config{
  ApiKey: "<Enter ApiKey>",
}

mojoclient, err := go_mojoauth.NewMojoAuth(&cfg)

if err != nil {
  errors = errors + err.(mojoerror.Error).OrigErr().Error()
  //		respCode = 500
}

```
