# go-boost
> A transaction parser for boost (Bitcoin protocol)

[![Release](https://img.shields.io/github/release-pre/BitcoinSchema/go-boost.svg?logo=github&style=flat&v=3)](https://github.com/BitcoinSchema/go-boost/releases)
[![Build Status](https://img.shields.io/github/actions/workflow/status/BitcoinSchema/go-boost/run-tests.yml?branch=master&logo=github&v=3)](https://github.com/BitcoinSchema/go-boost/actions)
[![Go](https://img.shields.io/github/go-mod/go-version/BitcoinSchema/go-boost?v=3)](https://golang.org/)
<br>
[![Mergify Status](https://img.shields.io/endpoint.svg?url=https://api.mergify.com/v1/badges/BitcoinSchema/go-boost&style=flat&v=3)](https://mergify.io)
[![Sponsor](https://img.shields.io/badge/sponsor-BitcoinSchema-181717.svg?logo=github&style=flat&v=3)](https://github.com/sponsors/BitcoinSchema)
[![Donate](https://img.shields.io/badge/donate-bitcoin-ff9900.svg?logo=bitcoin&style=flat&v=3)](https://gobitcoinsv.com/#sponsor?utm_source=github&utm_medium=sponsor-link&utm_campaign=go-boost&utm_term=go-boost&utm_content=go-boost)
<br>
<br>

## Installation

**go-boost** requires a [supported release of Go](https://golang.org/doc/devel/release.html#policy).
```shell script
go get -u github.com/bitcoinschema/go-boost
```
<br>

## Documentation
View the generated [documentation](https://pkg.go.dev/github.com/bitcoinschema/go-boost)

[![GoDoc](https://godoc.org/github.com/bitcoinschema/go-boost?status.svg&style=flat)](https://pkg.go.dev/github.com/bitcoinschema/go-boost)

<br>

## Boost Object
It takes a BOB formatted transaction and produces an easy-to-use BOOST struct with all protocol fields populated.

```go
type Boost struct {
	Redeem Redeem `json:"redeem" bson:"redeem"`
	Spend  Spend  `json:"spend" bson:"spend"`
}
```
<br>

## Coverage
- [x] V1 Bounty Spend
- [x] V1 Bounty Redeem
- [x] V2 Contract Spend
- [x] V2 Contract Redeem

<br>

## Usage
Get the package

```bash
go get github.com/bitcoinschema/go-boost
```

Import the package
```
import (
  github.com/bitcoinschema/go-boost
)
```

Use to transform [BOB](https://github.com/bitcoinschema/go-bob) formatted transactions.

```
var b *boost.Boost
b, err = boostNewFromTape(&bobData.In[0].Tape[0])
```

V1 Contract Spend Example
```go
	// Get BOB data from string
	bobData, err := bob.NewFromRawTxString(boostTx)
	if err != nil {
		t.Fatalf("error occurred: %s", err.Error())
	}

	var b *Boost
	b, err = boost.NewFromTape(&bobData.Out[0].Tape[0])thing": "else"
}
```

boost.BoostSpend Result:
```json
{
  "Content": "7332808b5283f8acedcc6240a42f669cc3d305413201527852061fd5b283d0d8",
  "MinerAddress": "16nhPWCkbkR1bNACwPYULBWyvxQ5MCDZBo",
  "Category": 1111,
  "Difficulty": 0.01,
  "Topic": "theory",
  "AdditionalData": "this is the Boost whitepaper",
  "Nonce": 137,
  "Version": 1
}
```

_More examples can be found in tests._

<br>

<br/>

## Contributing

View the [contributing guidelines](.github/CONTRIBUTING.md) and follow the [code of conduct](.github/CODE_OF_CONDUCT.md).

### How can I help?
All kinds of contributions are welcome :raised_hands:!
The most basic way to show your support is to star :star2: the project, or to raise issues :speech_balloon:.
You can also support this project by [becoming a sponsor on GitHub](https://github.com/sponsors/BitcoinSchema) :clap:
or by making a [**bitcoin donation**](https://gobitcoinsv.com/#sponsor?utm_source=github&utm_medium=sponsor-link&utm_campaign=go-boost&utm_term=go-boost&utm_content=go-boost) to ensure this journey continues indefinitely! :rocket:

[![Stars](https://img.shields.io/github/stars/BitcoinSchema/go-boost?label=Please%20like%20us&style=social)](https://github.com/BitcoinSchema/go-boost/stargazers)

<br/>

## License

[![License](https://img.shields.io/github/license/BitcoinSchema/go-boost.svg?style=flat&v=3)](LICENSE)
