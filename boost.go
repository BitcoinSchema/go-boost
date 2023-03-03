package boost

// Prefix is the protocol prefix for Boost Pow
const Prefix = "boostpow"

// Boost is Boost Pow data object
type Boost struct {
	Redeem BoostRedeem `json:"redeem" bson:"redeem"`
	Spend  BoostSpend  `json:"spend" bson:"spend"`
}

// BoostSpend is the parsed spend data. Version unknown when 0.
type BoostSpend struct {
	Hash           string  `json:"hash,omitempty" bson:"hash,omitempty"`
	Content        string  `json:"content,omitempty" bson:"content,omitempty"`
	Difficulty     float64 `json:"difficulty,omitempty" bson:"difficulty,omitempty"`
	Topic          *string `json:"topic,omitempty" bson:"topic,omitempty"`
	AdditionalData *string `json:"additional_data,omitempty" bson:"additional_data,omitempty"`
	Bits           uint64  `json:"bits" bson:"bits"`
	MetadataHash   string  `json:"metadata_hash,omitempty" bson:"metadata_hash,omitempty"`
	Time           uint64  `json:"time,omitempty" bson:"time,omitempty"`
	Nonce          uint32  `json:"nonce,omitempty" bson:"nonce,omitempty"`
	Category       uint32  `json:"category" bson:"category,omitempty"`
	MinerAddress   *string `json:"miner_address,omitempty" bson:"miner_address,omitempty"`
	Version        int32   `json:"version" bson:"version"`
}

type BoostRedeem struct {
	Signature       string `json:"signature" bson:"signature"`                 // 3044022100ac4003d62ddadbf0bff9cbe63d0f6ad740494ee7fcf5f296cfc056f52f087c7c021f2f9e2db03b141ce88edc1c10850a0831dea63edd6c6a8040d80e24737e6d4a41
	PubKey          string `json:"pubkey" bson:"pubkey"`                       // pubkey: 03097e9768554d40c0b5b18e44db2a15bbd137a373c39af46033049477bcbb79a4
	Nonce           uint32 `json:"nonce" bson:"nonce"`                         // nonce : 31497
	Timestamp       uint32 `json:"timestamp" bson:"timestamp"`                 // timestamp : 1677268580
	ExtraNonce2     string `json:"extra_nonce_2" bson:"extra_nonce_2"`         // extra_nonce_2 : "0f0445b186e64adc"
	ExtraNonce1     uint32 `json:"extra_nonce_1" bson:"extra_nonce_1"`         // extra_nonce_1 : 909479219
	MinerPubKeyHash string `json:"miner_pubkey_hash" bson:"miner_pubkey_hash"` // miner_pubkey_hash: 0xa3c10ac097a7da0009a786cc17edc1391a3bddf6
	Version         int32  `json:"version" bson:"version"`
}

// pattern output_script_pattern_no_asicboost = pattern{
// 0 - 	push{bytes{0x62, 0x6F, 0x6F, 0x73, 0x74, 0x70, 0x6F, 0x77}},
// 1 - OP_DROP,
// 2 -	optional{push_size{20, MinerAddress}},
// 2, 3 -	push_size{4, Category},
// 3, 4 -	push_size{32, Content},
// 4, 5 -	push_size{4, Target},
// 5, 6 -	push{x.Tag},
// 6, 7 - push_size{4, UserNonce},
// 	push{x.AdditionalData}, OP_CAT, OP_SWAP,
// 	// copy mining pool’s pubkey hash to alt stack. A copy remains on the stack.
// 	push{5}, OP_ROLL, OP_DUP, OP_TOALTSTACK, OP_CAT,
// 	// copy target and push to altstack.
// 	push{2}, OP_PICK, OP_TOALTSTACK,
// 	// check size of extra_nonce_1
// 	push{5}, OP_ROLL, OP_SIZE, push{4}, OP_EQUALVERIFY, OP_CAT,
// 	// check size of extra_nonce_2
// 	push{5}, OP_ROLL, OP_SIZE, push{8}, OP_EQUALVERIFY, OP_CAT,
// 	// create metadata document and hash it.
// 	OP_SWAP, OP_CAT, OP_HASH256,
// 	OP_SWAP, OP_TOALTSTACK, OP_CAT, OP_CAT,              // target to altstack.
// 	OP_SWAP, OP_SIZE, push{4}, OP_EQUALVERIFY, OP_CAT,   // check size of timestamp.
// 	OP_FROMALTSTACK, OP_CAT,                             // attach target
// 	// check size of nonce. Boost POW string is constructed.
// 	OP_SWAP, OP_SIZE, push{4}, OP_EQUALVERIFY, OP_CAT,
// 	// Take hash of work string and ensure that it is positive and minimally encoded.
// 	OP_HASH256, ensure_positive,
// 	// Get target, transform to expanded form, and ensure that it is positive and minimally encoded.
// 	OP_FROMALTSTACK, expand_target, ensure_positive,
// 	// check that the hash of the Boost POW string is less than the target
// 	OP_LESSTHAN, OP_VERIFY,
// 	// check that the given address matches the pubkey and check signature.
// 	OP_DUP, OP_HASH160, OP_FROMALTSTACK, OP_EQUALVERIFY, OP_CHECKSIG};

// hash: '0000000086915e291fe43f10bdd8232f65e6eb64628bbb4d128be3836c21b6cc',
// content: '00000000000000000000000000000000000000000048656c6c6f20776f726c64',
// bits: 486604799,
// difficulty: 1,
// metadataHash: "acd8278e84b037c47565df65a981d72fb09be5262e8783d4cf4e42633615962a",
// time: 1305200806,
// nonce: 3698479534,
// category: 1,
