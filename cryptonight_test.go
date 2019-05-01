package cryptonight

import (
	"bytes"
	"encoding/hex"
	"testing"

	"gitlab.neji.vm.tc/marconi/go-ethereum/common/hexutil"
)

// Couple of helper functions to make tests below more readable.
func hashVariant1(input []byte) []byte {
	return hashCryptonight(input, 1, 0 /*block_height*/)
}

func hashVariant2(input []byte) []byte {
	return hashCryptonight(input, 2, 0 /*block_height*/)
}

func hashVariant4(input []byte, block_height uint64) []byte {
	return hashCryptonight(input, 4, block_height)
}

func TestHashVariant1(t *testing.T) {
	var input []byte
	var expected_hash []byte
	var actual_hash []byte

	// The following five test cases come from the Monero github
	// repository's test file, tests-slow-1.txt. They make sure we
	// correctly pulled in and compiled the cryptonight variant 1
	// implementation. The file is just some pairs of input and
	// expected output. For the curious, you can see how exactly
	// tests-slow-1.txt gets used by taking a look at Monero's
	// tests/hash/main.cpp file.
	input = hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000")
	expected_hash = hexutil.MustDecode("0xb5a7f63abb94d07d1a6445c36c07c7e8327fe61b1647e391b4c7edae5de57a3d")
	actual_hash = hashVariant1(input)
	if !bytes.Equal(actual_hash, expected_hash) {
		t.Error("Unexpected result: ", hex.EncodeToString(actual_hash), " versus ", hex.EncodeToString(expected_hash))
	}

	input = hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")
	expected_hash = hexutil.MustDecode("0x80563c40ed46575a9e44820d93ee095e2851aa22483fd67837118c6cd951ba61")
	actual_hash = hashVariant1(input)
	if !bytes.Equal(actual_hash, expected_hash) {
		t.Error("Unexpected result: ", hex.EncodeToString(actual_hash), " versus ", hex.EncodeToString(expected_hash))
	}

	input = hexutil.MustDecode("0x8519e039172b0d70e5ca7b3383d6b3167315a422747b73f019cf9528f0fde341fd0f2a63030ba6450525cf6de31837669af6f1df8131faf50aaab8d3a7405589")
	expected_hash = hexutil.MustDecode("0x5bb40c5880cef2f739bdb6aaaf16161eaae55530e7b10d7ea996b751a299e949")
	actual_hash = hashVariant1(input)
	if !bytes.Equal(actual_hash, expected_hash) {
		t.Error("Unexpected result: ", hex.EncodeToString(actual_hash), " versus ", hex.EncodeToString(expected_hash))
	}

	input = hexutil.MustDecode("0x37a636d7dafdf259b7287eddca2f58099e98619d2f99bdb8969d7b14498102cc065201c8be90bd777323f449848b215d2977c92c4c1c2da36ab46b2e389689ed97c18fec08cd3b03235c5e4c62a37ad88c7b67932495a71090e85dd4020a9300")
	expected_hash = hexutil.MustDecode("0x613e638505ba1fd05f428d5c9f8e08f8165614342dac419adc6a47dce257eb3e")
	actual_hash = hashVariant1(input)
	if !bytes.Equal(actual_hash, expected_hash) {
		t.Error("Unexpected result: ", hex.EncodeToString(actual_hash), " versus ", hex.EncodeToString(expected_hash))
	}

	input = hexutil.MustDecode("0x38274c97c45a172cfc97679870422e3a1ab0784960c60514d816271415c306ee3a3ed1a77e31f6a885c3cb")
	expected_hash = hexutil.MustDecode("0xed082e49dbd5bbe34a3726a0d1dad981146062b39d36d62c71eb1ed8ab49459b")
	actual_hash = hashVariant1(input)
	if !bytes.Equal(actual_hash, expected_hash) {
		t.Error("Unexpected result: ", hex.EncodeToString(actual_hash), " versus ", hex.EncodeToString(expected_hash))
	}
	// Holy shit it works. Endianness almost tricked me.
	input = hexutil.MustDecode("0x07000000000000a25dd004a5561f04de75b908d671ffa39960352d03d0682d28ce0890d29b9c96fd1e0000777777777777777777777777777777777777777777777777777777777777777777")
	expected_hash = hexutil.MustDecode("0xcbe504327063949398a053ae97c057f7ced7075b5cc9186c5377e48433190100")
	actual_hash = hashVariant1(input)
	if !bytes.Equal(actual_hash, expected_hash) {
		t.Error("Unexpected result: ", hex.EncodeToString(actual_hash), " versus ", hex.EncodeToString(expected_hash))
	}
}

func TestHashVariant1ForEthereum(t *testing.T) {
	var block_header_bytes []byte = hexutil.MustDecode("0xb34f93a7c65392053cbbf073e9ad3bc7a7c0c3a45bfa0795f954b53686849db8")
	var nonce uint64 = 0xc526c0a1000008dc
	var digest []byte
	var result []byte
	digest, result = HashVariant1ForEthereumHeader(block_header_bytes, nonce)
	expected_digest := hexutil.MustDecode("0x834d72ab9e78b9a60808b9a49866c6a452826f11eb4a8d3ac4b49c0faf740100")
	if !bytes.Equal(digest, expected_digest) {
		t.Error("Unexpected digest: ", hex.EncodeToString(digest), " versus ", hex.EncodeToString(expected_digest))
	}
	expected_result := hexutil.MustDecode("0x000174af0f9cb4c43a8d4aeb116f8252a4c66698a4b90808a6b9789eab724d83") // interpreted in little-endian
	if !bytes.Equal(result, expected_result) {
		t.Error("Unexpected result: ", hex.EncodeToString(result), " versus ", hex.EncodeToString(expected_result))
	}
}

func TestHashVariant2(t *testing.T) {
	var input []byte
	var expected_hash []byte
	var actual_hash []byte

	// The following test cases come from the Monero github repository's
	// test file, tests-slow-2.txt. They make sure we correctly pulled
	// in and compiled the cryptonight variant 2 implementation. The
	// file is just some pairs of input and expected output. For the
	// curious, you can see how exactly tests-slow-2.txt gets used by
	// taking a look at Monero's tests/hash/main.cpp file.
	input = hexutil.MustDecode("0x5468697320697320612074657374205468697320697320612074657374205468697320697320612074657374")
	expected_hash = hexutil.MustDecode("0x353fdc068fd47b03c04b9431e005e00b68c2168a3cc7335c8b9b308156591a4f")
	actual_hash = hashVariant2(input)
	if !bytes.Equal(actual_hash, expected_hash) {
		t.Error("Unexpected result: ", hex.EncodeToString(actual_hash), " versus ", hex.EncodeToString(expected_hash))
	}

	input = hexutil.MustDecode("0x4c6f72656d20697073756d20646f6c6f722073697420616d65742c20636f6e73656374657475722061646970697363696e67")
	expected_hash = hexutil.MustDecode("0x72f134fc50880c330fe65a2cb7896d59b2e708a0221c6a9da3f69b3a702d8682")
	actual_hash = hashVariant2(input)
	if !bytes.Equal(actual_hash, expected_hash) {
		t.Error("Unexpected result: ", hex.EncodeToString(actual_hash), " versus ", hex.EncodeToString(expected_hash))
	}

	input = hexutil.MustDecode("0x656c69742c2073656420646f20656975736d6f642074656d706f7220696e6369646964756e74207574206c61626f7265")
	expected_hash = hexutil.MustDecode("0x410919660ec540fc49d8695ff01f974226a2a28dbbac82949c12f541b9a62d2f")
	actual_hash = hashVariant2(input)
	if !bytes.Equal(actual_hash, expected_hash) {
		t.Error("Unexpected result: ", hex.EncodeToString(actual_hash), " versus ", hex.EncodeToString(expected_hash))
	}

	input = hexutil.MustDecode("0x657420646f6c6f7265206d61676e6120616c697175612e20557420656e696d206164206d696e696d2076656e69616d2c")
	expected_hash = hexutil.MustDecode("0x4472fecfeb371e8b7942ce0378c0ba5e6d0c6361b669c587807365c787ae652d")
	actual_hash = hashVariant2(input)
	if !bytes.Equal(actual_hash, expected_hash) {
		t.Error("Unexpected result: ", hex.EncodeToString(actual_hash), " versus ", hex.EncodeToString(expected_hash))
	}

	input = hexutil.MustDecode("0x71756973206e6f737472756420657865726369746174696f6e20756c6c616d636f206c61626f726973206e697369")
	expected_hash = hexutil.MustDecode("0x577568395203f1f1225f2982b637f7d5e61b47a0f546ba16d46020b471b74076")
	actual_hash = hashVariant2(input)
	if !bytes.Equal(actual_hash, expected_hash) {
		t.Error("Unexpected result: ", hex.EncodeToString(actual_hash), " versus ", hex.EncodeToString(expected_hash))
	}

	input = hexutil.MustDecode("0x757420616c697175697020657820656120636f6d6d6f646f20636f6e7365717561742e20447569732061757465")
	expected_hash = hexutil.MustDecode("0xf6fd7efe95a5c6c4bb46d9b429e3faf65b1ce439e116742d42b928e61de52385")
	actual_hash = hashVariant2(input)
	if !bytes.Equal(actual_hash, expected_hash) {
		t.Error("Unexpected result: ", hex.EncodeToString(actual_hash), " versus ", hex.EncodeToString(expected_hash))
	}

	input = hexutil.MustDecode("0x697275726520646f6c6f7220696e20726570726568656e646572697420696e20766f6c7570746174652076656c6974")
	expected_hash = hexutil.MustDecode("0x422f8cfe8060cf6c3d9fd66f68e3c9977adb683aea2788029308bbe9bc50d728")
	actual_hash = hashVariant2(input)
	if !bytes.Equal(actual_hash, expected_hash) {
		t.Error("Unexpected result: ", hex.EncodeToString(actual_hash), " versus ", hex.EncodeToString(expected_hash))
	}

	input = hexutil.MustDecode("0x657373652063696c6c756d20646f6c6f726520657520667567696174206e756c6c612070617269617475722e")
	expected_hash = hexutil.MustDecode("0x512e62c8c8c833cfbd9d361442cb00d63c0a3fd8964cfd2fedc17c7c25ec2d4b")
	actual_hash = hashVariant2(input)
	if !bytes.Equal(actual_hash, expected_hash) {
		t.Error("Unexpected result: ", hex.EncodeToString(actual_hash), " versus ", hex.EncodeToString(expected_hash))
	}

	input = hexutil.MustDecode("0x4578636570746575722073696e74206f6363616563617420637570696461746174206e6f6e2070726f6964656e742c")
	expected_hash = hexutil.MustDecode("0x12a794c1aa13d561c9c6111cee631ca9d0a321718d67d3416add9de1693ba41e")
	actual_hash = hashVariant2(input)
	if !bytes.Equal(actual_hash, expected_hash) {
		t.Error("Unexpected result: ", hex.EncodeToString(actual_hash), " versus ", hex.EncodeToString(expected_hash))
	}

	input = hexutil.MustDecode("0x73756e7420696e2063756c706120717569206f666669636961206465736572756e74206d6f6c6c697420616e696d20696420657374206c61626f72756d2e")
	expected_hash = hexutil.MustDecode("0x2659ff95fc74b6215c1dc741e85b7a9710101b30620212f80eb59c3c55993f9d")
	actual_hash = hashVariant2(input)
	if !bytes.Equal(actual_hash, expected_hash) {
		t.Error("Unexpected result: ", hex.EncodeToString(actual_hash), " versus ", hex.EncodeToString(expected_hash))
	}
}

func TestHashVariant2ForEthereum(t *testing.T) {
	var block_header_bytes []byte = hexutil.MustDecode("0xb34f93a7c65392053cbbf073e9ad3bc7a7c0c3a45bfa0795f954b53686849db8")
	var nonce uint64 = 0xc526c0a1000008dc
	var digest []byte
	var result []byte
	digest, result = HashVariant2ForEthereumHeader(block_header_bytes, nonce)
	expected_digest := hexutil.MustDecode("0xa0e217e26c0c5c409a9e7119a8a7b1c4faa5886a2c101116548ab323f11bc4c2")
	if !bytes.Equal(digest, expected_digest) {
		t.Error("Unexpected digest: ", hex.EncodeToString(digest), " versus ", hex.EncodeToString(expected_digest))
	}
	expected_result := hexutil.MustDecode("0xc2c41bf123b38a541611102c6a88a5fac4b1a7a819719e9a405c0c6ce217e2a0") // interpreted in little-endian
	if !bytes.Equal(result, expected_result) {
		t.Error("Unexpected result: ", hex.EncodeToString(result), " versus ", hex.EncodeToString(expected_result))
	}
}

func TestHashVariant4(t *testing.T) {
	var input []byte
	var expected_hash []byte
	var actual_hash []byte

	// The following test cases come from the Monero github repository's
	// test file, tests-slow-4.txt. They make sure we correctly pulled
	// in and compiled the cryptonight variant 4 implementation. The
	// file is just some pairs of input and expected output. For the
	// curious, you can see how exactly tests-slow-4.txt gets used by
	// taking a look at Monero's tests/hash/main.cpp file.
	input = hexutil.MustDecode("0x5468697320697320612074657374205468697320697320612074657374205468697320697320612074657374")
	expected_hash = hexutil.MustDecode("0xf759588ad57e758467295443a9bd71490abff8e9dad1b95b6bf2f5d0d78387bc")
	actual_hash = hashVariant4(input, 1806260 /*block_height*/)
	if !bytes.Equal(actual_hash, expected_hash) {
		t.Error("Unexpected result: ", hex.EncodeToString(actual_hash), " versus ", hex.EncodeToString(expected_hash))
	}

	input = hexutil.MustDecode("0x4c6f72656d20697073756d20646f6c6f722073697420616d65742c20636f6e73656374657475722061646970697363696e67")
	expected_hash = hexutil.MustDecode("0x5bb833deca2bdd7252a9ccd7b4ce0b6a4854515794b56c207262f7a5b9bdb566")
	actual_hash = hashVariant4(input, 1806261 /*block_height*/)
	if !bytes.Equal(actual_hash, expected_hash) {
		t.Error("Unexpected result: ", hex.EncodeToString(actual_hash), " versus ", hex.EncodeToString(expected_hash))
	}

	input = hexutil.MustDecode("0x656c69742c2073656420646f20656975736d6f642074656d706f7220696e6369646964756e74207574206c61626f7265")
	expected_hash = hexutil.MustDecode("0x1ee6728da60fbd8d7d55b2b1ade487a3cf52a2c3ac6f520db12c27d8921f6cab")
	actual_hash = hashVariant4(input, 1806262 /*block_height*/)
	if !bytes.Equal(actual_hash, expected_hash) {
		t.Error("Unexpected result: ", hex.EncodeToString(actual_hash), " versus ", hex.EncodeToString(expected_hash))
	}

	input = hexutil.MustDecode("0x657420646f6c6f7265206d61676e6120616c697175612e20557420656e696d206164206d696e696d2076656e69616d2c")
	expected_hash = hexutil.MustDecode("0x6969fe2ddfb758438d48049f302fc2108a4fcc93e37669170e6db4b0b9b4c4cb")
	actual_hash = hashVariant4(input, 1806263 /*block_height*/)
	if !bytes.Equal(actual_hash, expected_hash) {
		t.Error("Unexpected result: ", hex.EncodeToString(actual_hash), " versus ", hex.EncodeToString(expected_hash))
	}

	input = hexutil.MustDecode("0x71756973206e6f737472756420657865726369746174696f6e20756c6c616d636f206c61626f726973206e697369")
	expected_hash = hexutil.MustDecode("0x7f3048b4e90d0cbe7a57c0394f37338a01fae3adfdc0e5126d863a895eb04e02")
	actual_hash = hashVariant4(input, 1806264 /*block_height*/)
	if !bytes.Equal(actual_hash, expected_hash) {
		t.Error("Unexpected result: ", hex.EncodeToString(actual_hash), " versus ", hex.EncodeToString(expected_hash))
	}

	input = hexutil.MustDecode("0x757420616c697175697020657820656120636f6d6d6f646f20636f6e7365717561742e20447569732061757465")
	expected_hash = hexutil.MustDecode("0x1d290443a4b542af04a82f6b2494a6ee7f20f2754c58e0849032483a56e8e2ef")
	actual_hash = hashVariant4(input, 1806265 /*block_height*/)
	if !bytes.Equal(actual_hash, expected_hash) {
		t.Error("Unexpected result: ", hex.EncodeToString(actual_hash), " versus ", hex.EncodeToString(expected_hash))
	}

	input = hexutil.MustDecode("0x697275726520646f6c6f7220696e20726570726568656e646572697420696e20766f6c7570746174652076656c6974")
	expected_hash = hexutil.MustDecode("0xc43cc6567436a86afbd6aa9eaa7c276e9806830334b614b2bee23cc76634f6fd")
	actual_hash = hashVariant4(input, 1806266 /*block_height*/)
	if !bytes.Equal(actual_hash, expected_hash) {
		t.Error("Unexpected result: ", hex.EncodeToString(actual_hash), " versus ", hex.EncodeToString(expected_hash))
	}

	input = hexutil.MustDecode("0x657373652063696c6c756d20646f6c6f726520657520667567696174206e756c6c612070617269617475722e")
	expected_hash = hexutil.MustDecode("0x87be2479c0c4e8edfdfaa5603e93f4265b3f8224c1c5946feb424819d18990a4")
	actual_hash = hashVariant4(input, 1806267 /*block_height*/)
	if !bytes.Equal(actual_hash, expected_hash) {
		t.Error("Unexpected result: ", hex.EncodeToString(actual_hash), " versus ", hex.EncodeToString(expected_hash))
	}

	input = hexutil.MustDecode("0x4578636570746575722073696e74206f6363616563617420637570696461746174206e6f6e2070726f6964656e742c")
	expected_hash = hexutil.MustDecode("0xdd9d6a6d8e47465cceac0877ef889b93e7eba979557e3935d7f86dce11b070f3")
	actual_hash = hashVariant4(input, 1806268 /*block_height*/)
	if !bytes.Equal(actual_hash, expected_hash) {
		t.Error("Unexpected result: ", hex.EncodeToString(actual_hash), " versus ", hex.EncodeToString(expected_hash))
	}

	input = hexutil.MustDecode("0x73756e7420696e2063756c706120717569206f666669636961206465736572756e74206d6f6c6c697420616e696d20696420657374206c61626f72756d2e")
	expected_hash = hexutil.MustDecode("0x75c6f2ae49a20521de97285b431e717125847fb8935ed84a61e7f8d36a2c3d8e")
	actual_hash = hashVariant4(input, 1806269 /*block_height*/)
	if !bytes.Equal(actual_hash, expected_hash) {
		t.Error("Unexpected result: ", hex.EncodeToString(actual_hash), " versus ", hex.EncodeToString(expected_hash))
	}
}

func TestHashVariant4ForEthereum(t *testing.T) {
	var block_header_bytes []byte = hexutil.MustDecode("0xb34f93a7c65392053cbbf073e9ad3bc7a7c0c3a45bfa0795f954b53686849db8")
	var nonce uint64 = 0xc526c0a1000008dc
	var digest []byte
	var result []byte
	digest, result = HashVariant4ForEthereumHeader(block_header_bytes, nonce, 8111222 /*block_height*/)
	expected_digest := hexutil.MustDecode("0x1621e81c0910c8167e2c37da637e212e24dd6882f1e9c0e043d6eff0d284a2b8")
	if !bytes.Equal(digest, expected_digest) {
		t.Error("Unexpected digest: ", hex.EncodeToString(digest), " versus ", hex.EncodeToString(expected_digest))
	}
	expected_result := hexutil.MustDecode("0xb8a284d2f0efd643e0c0e9f18268dd242e217e63da372c7e16c810091ce82116") // interpreted in little-endian
	if !bytes.Equal(result, expected_result) {
		t.Error("Unexpected result: ", hex.EncodeToString(result), " versus ", hex.EncodeToString(expected_result))
	}
}
