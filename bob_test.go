package boost

import (
	"strings"
	"testing"

	"github.com/bitcoinschema/go-bob"
)

// this tx has 2 boost outputs

// "BOOST": [
// 	{
// 		"content": "1d0be6ce78582478caa329e65a808d75d9e05302bd1025d3621a43b79da8ea18",
// 		"diff": 1,
// 		"category": "00000000",
// 		"tag": "6b73697223000000000000000000000000000000",
// 		"additionalData": "0000000000000000000000000000000000000000000000000000000000000000",
// 		"userNonce": "00000000",
// 		"useGeneralPurposeBits": false
// 	},
// 	{
// 		"content": "3332a9f3b0db1fcbd2ec312a1c45d07dadf35f2df090d713ba576fe7f34ad94b",
// 		"diff": 1,
// 		"category": "00000000",
// 		"tag": "65636e616e696623000000000000000000000000",
// 		"additionalData": "0000000000000000000000000000000000000000000000000000000000000000",
// 		"userNonce": "00000000",
// 		"useGeneralPurposeBits": false
// 	}
// ],

// TestFromTape will test the method NewFromTape()
func TestNewFromTape(t *testing.T) {

	// TODO: Move this to go-bob
	// Clean up the bob string
	boostTx := RemoveLBR(sampleValidBoostTx)
	boostTx = strings.ReplaceAll(boostTx, ",\"a\":false", "")

	// Get BOB data from string
	bobData, err := bob.NewFromString(boostTx)
	if err != nil {
		t.Fatalf("error occurred: %s", err.Error())
	}

	// Get from tape - instance 1
	var b *Boost
	b, err = NewFromTape(&bobData.Out[4].Tape[0])
	if err != nil {
		t.Fatalf("error occurred: %s\n", err.Error())
	} else if b.Spend.Category != 0 {
		t.Fatalf("expected: %d got: %d\n", 0, b.Spend.Category)
	} else if b.Spend.Difficulty != 1.0 {
		t.Fatalf("expected: %f got: %f\n", 1.0, b.Spend.Difficulty)
	} else if b.Spend.Topic != nil && *b.Spend.Topic != "#risk" {
		t.Fatalf("expected: %s got: %s\n", "#risk", *b.Spend.Topic)
	} else if b.Spend.Nonce != 0 {
		t.Fatalf("expected: %d got: %d\n", 0, b.Spend.Nonce)
	} else if b.Spend.Content != "what?" {
		t.Fatalf("expected: %s got: %x\n", "what?", b.Spend.Content)
	}

	// Get from tape - instance 2
	var b2 *Boost
	b2, err = NewFromTape(&bobData.Out[6].Tape[0])
	if err != nil {
		t.Fatalf("error occurred: %s\n", err.Error())
	} else if b2.Spend.Category != 0 {
		t.Fatalf("expected: %d got: %d\n", 0, b2.Spend.Category)
	} else if b2.Spend.Difficulty != 486604799 {
		t.Fatalf("expected: %d got: %f\n", 4866047999, b2.Spend.Difficulty)
	} else if b2.Spend.Topic != nil && *b.Spend.Topic != "what3?" {
		t.Fatalf("expected: %s got: %s\n", "what3?", *b2.Spend.Topic)
	} else if b2.Spend.Nonce != 0 {
		t.Fatalf("expected: %d got: %d\n", 0, b2.Spend.Nonce)
	} else if b2.Spend.Content != "what4?" {
		t.Fatalf("expected: %s got: %x\n", "what4?", b2.Spend.Content)
	}

	// Wrong tape
	_, err = NewFromTape(&bobData.Out[0].Tape[0])
	if err == nil {
		t.Fatalf("error should have occurred")
	}

}
