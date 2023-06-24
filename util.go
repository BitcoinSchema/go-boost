package boost

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"math/big"
	"regexp"
	"strconv"

	"github.com/libsv/go-bk/base58"
)

// RemoveLBR removes linebreaks
func RemoveLBR(text string) string {
	re := regexp.MustCompile(`\x{000D}\x{000A}|[\x{000A}\x{000B}\x{000C}\x{000D}\x{0085}\x{2028}\x{2029}]`)
	return re.ReplaceAllString(text, ``)
}

// Reverse a byte array
func Reverse(input []byte) []byte {
	l := len(input)
	reversed := make([]byte, l)
	for i, n := range input {
		j := l - i - 1
		reversed[j] = n
	}
	return reversed
}

// PubkeyHashToAddress converts a pubkey hash to a base58check encoded address
func PubkeyHashToAddress(pubkeyHash []byte) string {
	// Add version byte to front (0x00 for mainnet addresses)
	versionedHash := append([]byte{0x00}, pubkeyHash...)

	// Double SHA-256 hash the versioned hash
	hash1 := sha256.Sum256(versionedHash)
	hash2 := sha256.Sum256(hash1[:])

	// Get the first 4 bytes of the double SHA-256 hash
	checksum := hash2[:4]

	// Append the checksum to the versioned hash
	addressBytes := append(versionedHash, checksum...)

	// Convert the address bytes to a Base58Check-encoded string
	addressString := base58.Encode(addressBytes)

	return addressString
}

func targetToDifficulty(target string) (*float64, error) {

	dataTarget, err := base64.StdEncoding.DecodeString(target)
	if err != nil {
		return nil, err
	}

	ib := binary.BigEndian.Uint32(Reverse(dataTarget))

	// to compact size
	t := big.NewInt(int64(ib % 0x01000000))
	t.Mul(t, big.NewInt(2).Exp(big.NewInt(2), big.NewInt(8*(int64(ib/0x01000000)-3)), nil))

	a := float64(0xFFFF0000000000000000000000000000000000000000000000000000) // genesis difficulty

	var b float64
	b, err = strconv.ParseFloat(t.String(), 64)
	if err != nil {
		emptyResult := 0.0
		return &emptyResult, err
	}
	result := a / b
	return &result, nil
}
