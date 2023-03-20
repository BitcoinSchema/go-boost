package boost

import (
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"

	"github.com/bitcoinschema/go-bpu"
)

// FromTape builds a boost object from a *bpu.Tape
func (b *Boost) FromTape(tape *bpu.Tape) (err error) {
	if (len(tape.Cell) < 8 || tape.Cell[1].Ops == nil) && !isBoostInput(tape) {
		err = fmt.Errorf("invalid %s record len %d, %+v", b.Spend.Hash, len(tape.Cell), tape.Cell)
		return
	}
	var isContract = false
	if tape.Cell[0].S != nil && *tape.Cell[0].S == Prefix {

		// Boostpow output found
		if *tape.Cell[1].Ops != "OP_DROP" {
			err = fmt.Errorf("no OP_DROP %s", *tape.Cell[1].Ops)
			return
		}

		data, error := base64.StdEncoding.DecodeString(*tape.Cell[2].B)
		if err != nil {
			log.Fatal("error:", error)
			err = error
		}
		var baseIdx uint8
		if len(data) == 4 {
			// Bounty
			b.Spend.Category = binary.LittleEndian.Uint32(data)
			baseIdx = 3
		} else if len(data) == 20 {
			isContract = true
			// Contract
			baseIdx = 4
			// Miner address
			// If the miner pubkey hash is present, then the redeeming tx needs to have a signature
			// that matches the pubkey hash. only the miner who has the private key can mine the boost output
			dataStr := PubkeyHashToAddress(data)
			b.Spend.MinerAddress = &dataStr

			// Set the category
			dataCat, error := base64.StdEncoding.DecodeString(*tape.Cell[3].B)
			if err != nil {
				log.Fatal("error:", error)
				err = error
			}
			b.Spend.Category = binary.LittleEndian.Uint32(dataCat)

		}

		if tape.Cell[baseIdx].S == nil {
			err = fmt.Errorf("no content")
			return
		}

		b.Spend.Content = *tape.Cell[baseIdx].S

		// Calculate the difficulty from target
		dataTarget, errTarget := base64.StdEncoding.DecodeString(*tape.Cell[baseIdx+1].B)
		if errTarget != nil {
			log.Fatal("error target:", errTarget)
			err = errTarget
			return
		}

		ib := binary.BigEndian.Uint32(Reverse(dataTarget))
		// to compact size
		t := big.NewInt(int64(ib % 0x01000000))
		t.Mul(t, big.NewInt(2).Exp(big.NewInt(2), big.NewInt(8*(int64(ib/0x01000000)-3)), nil))

		// get difficulty
		flt, errDifficulty := targetToDifficulty(*tape.Cell[baseIdx+1].B)
		if errDifficulty != nil {
			log.Fatal("error difficulty:", errDifficulty)
			err = errDifficulty
			return
		}

		// set difficulty
		b.Spend.Difficulty = *flt

		if len(tape.Cell) > 20 && (tape.Cell[18].Ops != nil && *tape.Cell[18].Ops == "OP_5") || (tape.Cell[19].Ops != nil && *tape.Cell[19].Ops == "OP_5") {
			b.Spend.Version = 1
		} else if len(tape.Cell) > 20 && tape.Cell[18].Ops != nil && (*tape.Cell[18].Ops == "OP_6" || *tape.Cell[19].Ops == "OP_6") {
			b.Spend.Version = 2
		}

		// Set topic / tag
		if tape.Cell[baseIdx+2].B != nil {
			dataTagB, errDataTag := base64.StdEncoding.DecodeString(*tape.Cell[baseIdx+2].B)

			if errDataTag != nil {
				log.Fatal("error:", errDataTag)
				err = errDataTag
				return
			}
			if len(dataTagB) > 0 {
				topic := string(dataTagB)
				b.Spend.Topic = &topic
			}
		}

		if tape.Cell[baseIdx+3].B != nil {
			// Set nonce
			dataNonce, errNonce := base64.StdEncoding.DecodeString(*tape.Cell[baseIdx+3].B)

			if errNonce != nil {
				log.Fatal("error:", errNonce)
				err = errNonce
				return
			}

			b.Spend.Nonce = binary.LittleEndian.Uint32(dataNonce)
		}

		// Set additional data
		b.Spend.AdditionalData = tape.Cell[baseIdx+4].S

		return

	} else if isBoostInput(tape) {
		b.Redeem.Signature = *tape.Cell[0].H
		b.Redeem.PubKey = *tape.Cell[1].H

		// Set nonce
		dataNonce, errNonce := base64.StdEncoding.DecodeString(*tape.Cell[2].B)

		if errNonce != nil {
			log.Fatal("error:", errNonce)
			err = errNonce
			return
		}

		b.Redeem.Nonce = binary.LittleEndian.Uint32(dataNonce)

		// Set timestamp
		dataTimestamp, errTimestamp := base64.StdEncoding.DecodeString(*tape.Cell[3].B)

		if errTimestamp != nil {
			log.Fatal("error:", errTimestamp)
			err = errNonce
			return
		}
		b.Redeem.Timestamp = binary.LittleEndian.Uint32(dataTimestamp)

		// Set extra nonce 2
		b.Redeem.ExtraNonce2 = *tape.Cell[4].H

		// Set extra nonce 1
		dataNonce1, errNonce1 := base64.StdEncoding.DecodeString(*tape.Cell[5].B)

		if errNonce1 != nil {
			log.Fatal("error:", errNonce1)
			err = errNonce1
			return
		}

		b.Redeem.ExtraNonce1 = binary.BigEndian.Uint32(dataNonce1)

		// Set miner pubkey hash
		var dataMinerPkh []byte
		var errMinerPkh error
		if len(*tape.Cell[6].H) == 40 {
			b.Redeem.Version = 1
			// V1 bounty redeem
			dataMinerPkh, errMinerPkh = base64.StdEncoding.DecodeString(*tape.Cell[6].B)

		} else if len(tape.Cell) == 8 && len(*tape.Cell[6].H) == 8 {
			b.Redeem.Version = 2
			// V2 bounty redeem
			dataMinerPkh, errMinerPkh = base64.StdEncoding.DecodeString(*tape.Cell[7].B)

		} else if len(tape.Cell) == 7 && len(*tape.Cell[6].H) == 8 {
			b.Redeem.Version = 2
			// V2 contract redeem
			dataMinerPkh, errMinerPkh = base64.StdEncoding.DecodeString(*tape.Cell[6].B)
			isContract = true
		}

		if errMinerPkh != nil {
			log.Fatal("error:", errMinerPkh)
			err = errMinerPkh
			return
		}

		b.Redeem.MinerPubKeyHash = hex.EncodeToString(Reverse(dataMinerPkh[:]))

	}
	if b.Redeem.Version == 0 {
		b.Redeem.Version = redeemVersion(tape, isContract)
	}
	return
}

func isBoostInput(tape *bpu.Tape) bool {
	// v1 bounty redeem, v2 contract redeem = 7
	// v2 bounty redeem = 8
	return len(tape.Cell) == 7 || len(tape.Cell) == 8
}

func redeemVersion(tape *bpu.Tape, contract bool) int32 {
	// v1 bounty redeem, v2 contract redeem = 7
	// v2 bounty redeem = 8
	if len(tape.Cell) == 7 {
		if contract {
			return 2
		}
		return 1
	} else if len(tape.Cell) == 8 {
		return 2
	}
	return 0
}

// NewFromTape takes a bob.Tape and returns a BAP data structure
func NewFromTape(tape *bpu.Tape) (b *Boost, err error) {
	b = new(Boost)
	if tape == nil {
		err = fmt.Errorf("tape is nil %x", tape)
		return
	}
	err = b.FromTape(tape)
	return
}
