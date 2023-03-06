# go-boost

A transaction parser for boost (Bitcoin protocol). It takes a BOB formatted transaction and produces an easy to use BOOST strut with all protocol fields populated.

## Boost Object

```go
type Boost struct {
	Redeem BoostRedeem `json:"redeem" bson:"redeem"`
	Spend  BoostSpend  `json:"spend" bson:"spend"`
}

```

## Coverage

- [x] V1 Bounty Spend
- [x] V1 Bounty Redeem
- [x] V2 Contract Spend
- [x] V2 Contract Redeem

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

### V1 Contract Spend Example

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

More examples can be found in tests.
