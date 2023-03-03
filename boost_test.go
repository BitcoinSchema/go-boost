package boost

import (
	"testing"

	"github.com/bitcoinschema/go-bob"
)

func TestV1BountySpend(t *testing.T) {

	// spend - c5c7248302683107aa91014fd955908a7c572296e803512e497ddf7d1f458bd3
	theoryStr := "theory"
	additionalDataStr := "this is the Boost whitepaper"

	var expected = &BoostSpend{
		Content:        "7332808b5283f8acedcc6240a42f669cc3d305413201527852061fd5b283d0d8",
		Difficulty:     0.01,
		Topic:          &theoryStr,
		AdditionalData: &additionalDataStr,
		Category:       1111,
		Nonce:          137,
		Version:        1,
	}

	boostTx := bountyV1SpendRaw

	// Get BOB data from string
	bobData, err := bob.NewFromRawTxString(boostTx)
	if err != nil {
		t.Fatalf("error occurred: %s", err.Error())
	}

	var b *Boost
	b, err = NewFromTape(&bobData.Out[0].Tape[0])
	if err != nil {
		t.Fatalf("error occurred: %s\n", err.Error())
	} else if b.Spend.Category != expected.Category {
		t.Fatalf("expected category: %d got: %d\n", expected.Category, b.Spend.Category)
	} else if b.Spend.Nonce != expected.Nonce {
		t.Fatalf("expected nonce: %d got: %d\n", expected.Nonce, b.Spend.Nonce)
	} else if b.Spend.Topic != nil && expected.Topic != nil && *b.Spend.Topic != *expected.Topic {
		t.Fatalf("expected topic: %s got: %s\n", *expected.Topic, *b.Spend.Topic)
	} else if b.Spend.AdditionalData != nil && expected.AdditionalData != nil && *b.Spend.AdditionalData != *expected.AdditionalData {
		t.Fatalf("expected additional data: %s got: %s\n", *expected.AdditionalData, *b.Spend.AdditionalData)
	} else if b.Spend.Difficulty != expected.Difficulty {
		t.Fatalf("expected target: %f got: %f\n", expected.Difficulty, b.Spend.Difficulty)
	} else if b.Spend.Version != expected.Version {
		t.Fatalf("expected version: %d got: %d\n", expected.Version, b.Spend.Version)
	}
}

func TestV1BountyRedeem(t *testing.T) {

	// redeem tx
	expectedTxid := "99bbbc28d39427bf530c05dc90db12e0953122fe5055afce6370d89a0085c28d"

	// expected values
	var expected = &BoostRedeem{
		Signature:       "3044022100ac4003d62ddadbf0bff9cbe63d0f6ad740494ee7fcf5f296cfc056f52f087c7c021f2f9e2db03b141ce88edc1c10850a0831dea63edd6c6a8040d80e24737e6d4a41",
		PubKey:          "03097e9768554d40c0b5b18e44db2a15bbd137a373c39af46033049477bcbb79a4",
		Nonce:           31497,
		Timestamp:       1677268580,
		ExtraNonce1:     909479219,
		ExtraNonce2:     "0f0445b186e64adc",
		MinerPubKeyHash: "a3c10ac097a7da0009a786cc17edc1391a3bddf6",
		Version:         1,
	}

	boostTx := bountyV1RedeemRaw

	// Get BOB data from string
	bobData, err := bob.NewFromRawTxString(boostTx)
	if err != nil {
		t.Fatalf("error occurred: %s", err.Error())
	}

	var b *Boost
	b, err = NewFromTape(&bobData.In[0].Tape[0])
	if err != nil {
		t.Fatalf("error occurred: %s\n", err.Error())
	} else if bobData.Tx.H != expectedTxid {
		t.Fatalf("expected txid: %s got: %s\n", expectedTxid, bobData.Tx.H)
	} else if b.Redeem.Signature != expected.Signature {
		t.Fatalf("expected signature: %s got: %s\n", expected.Signature, b.Redeem.Signature)
	} else if b.Redeem.Nonce != expected.Nonce {
		t.Fatalf("expected nonce: %d got: %d\n", expected.Nonce, b.Redeem.Nonce)
	} else if b.Redeem.PubKey != expected.PubKey {
		t.Fatalf("expected pubkey: %s got: %s\n", expected.PubKey, b.Redeem.PubKey)
	} else if b.Redeem.Timestamp != expected.Timestamp {
		t.Fatalf("expected timestamp: %d got: %d\n", expected.Timestamp, b.Redeem.Timestamp)
	} else if b.Redeem.ExtraNonce2 != expected.ExtraNonce2 {
		t.Fatalf("expected extra_nonce_2: %s got: %s\n", expected.ExtraNonce2, b.Redeem.ExtraNonce2)
	} else if b.Redeem.ExtraNonce1 != expected.ExtraNonce1 {
		t.Fatalf("expected extra_nonce_1: %d got: %d\n", expected.ExtraNonce1, b.Redeem.ExtraNonce1)
	} else if b.Redeem.MinerPubKeyHash != expected.MinerPubKeyHash {
		t.Fatalf("expected miner_pubkey_hash: %s got: %s\n", expected.MinerPubKeyHash, b.Redeem.MinerPubKeyHash)
	} else if b.Redeem.Version != expected.Version {
		t.Fatalf("expected version: %d got: %d\n", expected.Version, b.Redeem.Version)
	}
}

func TestV2BountySpend(t *testing.T) {

	// spend

	// expected values
	expectedTxid := "12aaef887e8348e83eac2937849de22a0bda8f7c2c819199bbcbb20b01722144"
	theoryStr := "theory"
	additionalDataStr := "this is the Boost whitepaper"
	var expected = &BoostSpend{
		Version:        2,
		Content:        "7332808b5283f8acedcc6240a42f669cc3d305413201527852061fd5b283d0d8",
		Category:       1111,
		Difficulty:     0.01,
		Nonce:          137,
		Topic:          &theoryStr,
		AdditionalData: &additionalDataStr,
	}

	boostTx := bountyV2SpendRaw

	// Get BOB data from string
	bobData, err := bob.NewFromRawTxString(boostTx)
	if err != nil {
		t.Fatalf("error occurred: %s", err.Error())
	}

	var b *Boost
	b, err = NewFromTape(&bobData.Out[0].Tape[0])
	if err != nil {
		t.Fatalf("error occurred: %s\n", err.Error())
	} else if expectedTxid != bobData.Tx.H {
		t.Fatalf("expected txid: %s got: %s\n", expectedTxid, bobData.Tx.H)
	} else if b.Spend.Category != expected.Category {
		t.Fatalf("expected category: %d got: %d\n", expected.Category, b.Spend.Category)
	} else if b.Spend.Nonce != expected.Nonce {
		t.Fatalf("expected nonce: %d got: %d\n", expected.Category, b.Spend.Category)
	} else if b.Spend.Topic != nil && expected.Topic != nil && *b.Spend.Topic != *expected.Topic {
		t.Fatalf("expected topic: %s got: %s\n", *expected.Topic, *b.Spend.Topic)
	} else if b.Spend.AdditionalData != nil && expected.AdditionalData != nil && *b.Spend.AdditionalData != *expected.AdditionalData {
		t.Fatalf("expected topic: %s got: %s\n", *expected.AdditionalData, *b.Spend.AdditionalData)
	} else if b.Spend.Version != expected.Version {
		t.Fatalf("expected version: %d got: %d\n", expected.Version, b.Spend.Version)
	} else if b.Spend.Difficulty != expected.Difficulty {
		t.Fatalf("expected target: %f got: %f\n", expected.Difficulty, b.Spend.Difficulty)
	}
}

func TestV2BountyRedeem(t *testing.T) {

	// redeem - 0f38e43dfc603296ef6883da389fc93815c0535bfba255b070a98bd6cc4da984
	// expected values
	//
	//	{
	//	  signature: 304502210081cac0bdfb713e8c6632ec8c7b6f1d070b19a43c3b06e05174f25dc9065c6e910220787dd9d0f58f79cda8b7f5b436eb2f8cd6d50dc5271e6216308c286406d4166141
	//	  pubkey: 03e0fd48907c0117600a6326aafe7d43adbc9421a4381bb6579f1ab4912cd25e37
	//	  nonce : 5267719
	//	  timestamp : 1677269436
	//	  extra_nonce_2 : "b4d8e1f74255bebc"
	//	  extra_nonce_1 : 2329617541
	//	  miner_pubkey_hash: 0x81bb8505a9999135a105e2f0290d55b1b70f7d3f
	//	}

	var expected = &BoostRedeem{
		Signature:       "304502210081cac0bdfb713e8c6632ec8c7b6f1d070b19a43c3b06e05174f25dc9065c6e910220787dd9d0f58f79cda8b7f5b436eb2f8cd6d50dc5271e6216308c286406d4166141",
		PubKey:          "03e0fd48907c0117600a6326aafe7d43adbc9421a4381bb6579f1ab4912cd25e37",
		Nonce:           5267719,
		Timestamp:       1677269436,
		ExtraNonce1:     2329617541,
		ExtraNonce2:     "b4d8e1f74255bebc",
		MinerPubKeyHash: "81bb8505a9999135a105e2f0290d55b1b70f7d3f",
		Version:         2,
	}

	boostTx := bountyV2RedeemRaw

	// Get BOB data from raw tx string
	bobData, err := bob.NewFromRawTxString(boostTx)
	if err != nil {
		t.Fatalf("error occurred: %s", err.Error())
	}

	var b *Boost
	b, err = NewFromTape(&bobData.In[0].Tape[0])
	if err != nil {
		t.Fatalf("error occurred: %s\n", err.Error())
	} else if b.Redeem.Signature != expected.Signature {
		t.Fatalf("expected signature: %s got: %s\n", expected.Signature, b.Redeem.Signature)
	} else if b.Redeem.Nonce != expected.Nonce {
		t.Fatalf("expected nonce: %d got: %d\n", expected.Nonce, b.Redeem.Nonce)
	} else if b.Redeem.PubKey != expected.PubKey {
		t.Fatalf("expected pubkey: %s got: %s\n", expected.PubKey, b.Redeem.PubKey)
	} else if b.Redeem.Timestamp != expected.Timestamp {
		t.Fatalf("expected timestamp: %d got: %d\n", expected.Timestamp, b.Redeem.Timestamp)
	} else if b.Redeem.ExtraNonce2 != expected.ExtraNonce2 {
		t.Fatalf("expected extra_nonce_2: %s got: %s\n", expected.ExtraNonce2, b.Redeem.ExtraNonce2)
	} else if b.Redeem.ExtraNonce1 != expected.ExtraNonce1 {
		t.Fatalf("expected extra_nonce_1: %d got: %d\n", expected.ExtraNonce1, b.Redeem.ExtraNonce1)
	} else if b.Redeem.MinerPubKeyHash != expected.MinerPubKeyHash {
		t.Fatalf("expected miner_pubkey_hash: %s got: %s\n", expected.MinerPubKeyHash, b.Redeem.MinerPubKeyHash)
	} else if b.Redeem.Version != expected.Version {
		t.Fatalf("expected version: %d got: %d\n", expected.Version, b.Redeem.Version)
	}
}

func TestV1ContractSpend(t *testing.T) {

	// spend

	// expected values
	expectedTxid := "ed122aa475c02ee049b342d9224bc140f015eee30b8411ad999c6a8378d9766e"
	theoryStr := "theory"
	additionalDataStr := "this is the Boost whitepaper"
	minerAddressStr := "16nhPWCkbkR1bNACwPYULBWyvxQ5MCDZBo"

	var expected = &BoostSpend{
		Content:        "7332808b5283f8acedcc6240a42f669cc3d305413201527852061fd5b283d0d8",
		MinerAddress:   &minerAddressStr,
		Category:       1111,
		Difficulty:     .01,
		Topic:          &theoryStr,
		AdditionalData: &additionalDataStr,
		Nonce:          137,
		Version:        1,
	}

	boostTx := contractV1SpendRaw

	// Get BOB data from string
	bobData, err := bob.NewFromRawTxString(boostTx)
	if err != nil {
		t.Fatalf("error occurred: %s", err.Error())
	}

	var b *Boost
	b, err = NewFromTape(&bobData.Out[0].Tape[0])
	if err != nil {
		t.Fatalf("error occurred: %s\n", err.Error())
	} else if expectedTxid != bobData.Tx.H {
		t.Fatalf("expected txid: %s got: %s\n", expectedTxid, bobData.Tx.H)
	} else if b.Spend.Category != expected.Category {
		t.Fatalf("expected category: %d got: %d\n", expected.Category, b.Spend.Category)
	} else if b.Spend.Nonce != expected.Nonce {
		t.Fatalf("expected nonce: %d got: %d\n", expected.Nonce, b.Spend.Nonce)
	} else if b.Spend.MinerAddress != nil && expected.MinerAddress != nil && *b.Spend.MinerAddress != *expected.MinerAddress {
		t.Fatalf("expected miner_address: %s got: %s\n", *expected.MinerAddress, *b.Spend.MinerAddress)
	} else if b.Spend.Topic != nil && expected.Topic != nil && *b.Spend.Topic != *expected.Topic {
		t.Fatalf("expected topic: %s got: %s\n", *expected.Topic, *b.Spend.Topic)
	} else if b.Spend.AdditionalData != nil && expected.AdditionalData != nil && *b.Spend.AdditionalData != *expected.AdditionalData {
		t.Fatalf("expected topic: %s got: %s\n", *expected.AdditionalData, *b.Spend.AdditionalData)
	} else if b.Spend.Version != expected.Version {
		t.Fatalf("expected version: %d got: %d\n", expected.Version, b.Spend.Version)
	} else if b.Spend.Difficulty != expected.Difficulty {
		t.Fatalf("expected target: %f got: %f\n", expected.Difficulty, b.Spend.Difficulty)
	}
}

func TestV2ContractRedeem(t *testing.T) {

	// redeem - a512c846b0154d23325f40ef87a088d747252fc2a179cca067a6026ee59c5ea6
	// expected values
	expectedTxid := "a512c846b0154d23325f40ef87a088d747252fc2a179cca067a6026ee59c5ea6"

	var expected = &BoostRedeem{
		Signature:   "304402201f94a12ace389cd389ef129dc9b68eb1a357ff6f71a508aa0b3accd90736007702206d316fce43e5ae24a6b07acc342e0f7a5c0d0366a2a00dee00acbb25b8f4f6a941",
		PubKey:      "03e0fd48907c0117600a6326aafe7d43adbc9421a4381bb6579f1ab4912cd25e37",
		Nonce:       3901135,
		Timestamp:   1677271659,
		ExtraNonce1: 3783406472,
		ExtraNonce2: "4f22be6e277ead90",
		Version:     2,
	}

	boostTx := contractV2RedeemRaw

	// Get BOB data from raw tx string
	bobData, err := bob.NewFromRawTxString(boostTx)
	if err != nil {
		t.Fatalf("error occurred: %s", err.Error())
	}

	var b *Boost
	b, err = NewFromTape(&bobData.In[0].Tape[0])
	if err != nil {
		t.Fatalf("error occurred: %s\n", err.Error())
	} else if expectedTxid != bobData.Tx.H {
		t.Fatalf("expected txid: %s got: %s\n", expectedTxid, bobData.Tx.H)
	} else if b.Redeem.Signature != expected.Signature {
		t.Fatalf("expected signature: %s got: %s\n", expected.Signature, b.Redeem.Signature)
	} else if b.Redeem.Nonce != expected.Nonce {
		t.Fatalf("expected nonce: %d got: %d\n", expected.Nonce, b.Redeem.Nonce)
	} else if b.Redeem.PubKey != expected.PubKey {
		t.Fatalf("expected pubkey: %s got: %s\n", expected.PubKey, b.Redeem.PubKey)
	} else if b.Redeem.Timestamp != expected.Timestamp {
		t.Fatalf("expected timestamp: %d got: %d\n", expected.Timestamp, b.Redeem.Timestamp)
	} else if b.Redeem.ExtraNonce2 != expected.ExtraNonce2 {
		t.Fatalf("expected extra_nonce_2: %s got: %s\n", expected.ExtraNonce2, b.Redeem.ExtraNonce2)
	} else if b.Redeem.ExtraNonce1 != expected.ExtraNonce1 {
		t.Fatalf("expected extra_nonce_1: %d got: %d\n", expected.ExtraNonce1, b.Redeem.ExtraNonce1)
	} else if b.Redeem.Version != expected.Version {
		t.Fatalf("expected version: %d got: %d\n", expected.Version, b.Redeem.Version)
	}
}
