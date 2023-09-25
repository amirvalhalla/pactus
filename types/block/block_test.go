package block_test

import (
	"encoding/hex"
	"testing"

	"github.com/fxamacker/cbor/v2"
	"github.com/pactus-project/pactus/crypto/hash"
	"github.com/pactus-project/pactus/types/block"
	"github.com/pactus-project/pactus/types/certificate"
	"github.com/pactus-project/pactus/types/tx"
	"github.com/pactus-project/pactus/types/tx/payload"
	"github.com/pactus-project/pactus/util"
	"github.com/pactus-project/pactus/util/simplemerkle"
	"github.com/pactus-project/pactus/util/testsuite"
	"github.com/stretchr/testify/assert"
)

func TestBasicCheck(t *testing.T) {
	ts := testsuite.NewTestSuite(t)

	t.Run("No transactions", func(t *testing.T) {
		b0 := ts.GenerateTestBlock()
		b := block.NewBlock(b0.Header(), b0.PrevCertificate(), block.Txs{})

		err := b.BasicCheck()
		assert.ErrorIs(t, err, block.BasicCheckError{
			Reason: "no subsidy transaction",
		})
	})

	t.Run("Without the previous certificate", func(t *testing.T) {
		b0 := ts.GenerateTestBlock()
		b := block.NewBlock(b0.Header(), nil, b0.Transactions())

		err := b.BasicCheck()
		assert.ErrorIs(t, err, block.BasicCheckError{
			Reason: "invalid genesis block hash",
		})
	})

	t.Run("Invalid certificate", func(t *testing.T) {
		b0 := ts.GenerateTestBlock()
		cert0 := b0.PrevCertificate()
		invCert := certificate.NewCertificate(0, 0, cert0.Committers(), cert0.Absentees(), cert0.Signature())
		b := block.NewBlock(b0.Header(), invCert, b0.Transactions())

		err := b.BasicCheck()
		assert.ErrorIs(t, err, block.BasicCheckError{
			Reason: "invalid certificate: certificate basic check failed: height is not positive: 0",
		})
	})

	//TODO fix me later
	//t.Run("Invalid transaction", func(t *testing.T) {
	//	b0 := ts.GenerateTestBlock()
	//	trxs0 := b0.Transactions()
	//	invalidValKey := ts.RandValKey()
	//	invalidValKey.Sign(trxs0[0].SignBytes())
	//	b := block.NewBlock(b0.Header(), b0.PrevCertificate(), trxs0)
	//
	//	err := b.BasicCheck()
	//	assert.ErrorIs(t, err, block.BasicCheckError{
	//		Reason: "invalid transaction: transaction basic check failed: invalid address: invalid address",
	//	})
	//})

	t.Run("Invalid state root hash", func(t *testing.T) {
		d := ts.DecodingHex(
			"01" + // Version
				"00000000" + // UnixTime
				"BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB" + // PrevBlockHash
				"0000000000000000000000000000000000000000000000000000000000000000" + // StateRoot
				"333333333333333333333333333333333333333333333333" + // SortitionSeed
				"333333333333333333333333333333333333333333333333" +
				"01AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" + // ProposerAddress
				"04030201" + // PrevCert: Height
				"0100" + // PrevCert: Round
				"0401020304" + // PrevCert: Committers
				"0102" + // PrevCert: Absentees
				"b53d79e156e9417e010fa21f2b2a96bee6be46fcd233295d" +
				"2f697cdb9e782b6112ac01c80d0d9d64c2320664c77fa2a6" + // PrevCert: Signature
				"01" + // Txs: Len
				"00" + // Tx[0]: Flags
				"01" + // Tx[0]: Version
				"a1b2c3d4" + // Tx[0]: Stamp
				"01000000" + // Tx[0]: LockTime
				"01" + // Tx[0]: Fee
				"00" + // Tx[0]: Memo
				"01" + // Tx[0]: PayloadType
				"00" + // Tx[0]: Sender (treasury)
				"012222222222222222222222222222222222222222" + // Tx[0]: Receiver
				"01") // Tx[0]: Amount

		b, _ := block.FromBytes(d)

		err := b.BasicCheck()
		assert.ErrorIs(t, err, block.BasicCheckError{
			Reason: "invalid state root: hash is zero",
		})
	})

	t.Run("Invalid previous block hash", func(t *testing.T) {
		d := ts.DecodingHex(
			"01" + // Version
				"00000000" + // UnixTime
				"0000000000000000000000000000000000000000000000000000000000000000" + // PrevBlockHash
				"BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB" + // StateRoot
				"333333333333333333333333333333333333333333333333" + // SortitionSeed
				"333333333333333333333333333333333333333333333333" +
				"01AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" + // ProposerAddress
				"04030201" + // PrevCert: Height
				"0100" + // PrevCert: Round
				"0401020304" + // PrevCert: Committers
				"0102" + // PrevCert: Absentees
				"b53d79e156e9417e010fa21f2b2a96bee6be46fcd233295d" +
				"2f697cdb9e782b6112ac01c80d0d9d64c2320664c77fa2a6" + // PrevCert: Signature
				"01" + // Txs: Len
				"00" + // Tx[0]: Flags
				"01" + // Tx[0]: Version
				"a1b2c3d4" + // Tx[0]: Stamp
				"01000000" + // Tx[0]: LockTime
				"01" + // Tx[0]: Fee
				"00" + // Tx[0]: Memo
				"01" + // Tx[0]: PayloadType
				"00" + // Tx[0]: Sender (treasury)
				"012222222222222222222222222222222222222222" + // Tx[0]: Receiver
				"01") // Tx[0]: Amount

		_, err := block.FromBytes(d)
		assert.ErrorIs(t, err, tx.InvalidPayloadTypeError{
			PayloadType: payload.Type(121),
		})
	})

	t.Run("Invalid proposer address (type is 2)", func(t *testing.T) {
		d := ts.DecodingHex(
			"01" + // Version
				"00000000" + // UnixTime
				"BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB" + // PrevBlockHash
				"BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB" + // StateRoot
				"333333333333333333333333333333333333333333333333" + // SortitionSeed
				"333333333333333333333333333333333333333333333333" +
				"02AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" + // ProposerAddress
				"04030201" + // PrevCert: Height
				"0100" + // PrevCert: Round
				"0401020304" + // PrevCert: Committers
				"0102" + // PrevCert: Absentees
				"b53d79e156e9417e010fa21f2b2a96bee6be46fcd233295d" +
				"2f697cdb9e782b6112ac01c80d0d9d64c2320664c77fa2a6" + // PrevCert: Signature
				"01" + // Txs: Len
				"00" + // Tx[0]: Flags
				"01" + // Tx[0]: Version
				"a1b2c3d4" + // Tx[0]: Stamp
				"01000000" + // Tx[0]: LockTime
				"01" + // Tx[0]: Fee
				"00" + // Tx[0]: Memo
				"01" + // Tx[0]: PayloadType
				"00" + // Tx[0]: Sender (treasury)
				"012222222222222222222222222222222222222222" + // Tx[0]: Receiver
				"01") // Tx[0]: Amount

		b, _ := block.FromBytes(d)
		err := b.BasicCheck()
		assert.ErrorIs(t, err, block.BasicCheckError{
			Reason: "invalid proposer address: pc1z42424242424242424242424242424242klpmq4",
		})
	})

	t.Run("Ok", func(t *testing.T) {
		d := ts.DecodingHex(
			"01" + // Version
				"00000000" + // UnixTime
				"BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB" + // PrevBlockHash
				"BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB" + // StateRoot
				"333333333333333333333333333333333333333333333333" + // SortitionSeed
				"333333333333333333333333333333333333333333333333" +
				"01AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" + // ProposerAddress
				"04030201" + // PrevCert: Height
				"0100" + // PrevCert: Round
				"0401020304" + // PrevCert: Committers
				"0102" + // PrevCert: Absentees
				"b53d79e156e9417e010fa21f2b2a96bee6be46fcd233295d" +
				"2f697cdb9e782b6112ac01c80d0d9d64c2320664c77fa2a6" + // PrevCert: Signature
				"01" + // Txs: Len
				"00" + // Tx[0]: Flags
				"01" + // Tx[0]: Version
				"a1b2c3d4" + // Tx[0]: Stamp
				"01000000" + // Tx[0]: LockTime
				"01" + // Tx[0]: Fee
				"00" + // Tx[0]: Memo
				"01" + // Tx[0]: PayloadType
				"00" + // Tx[0]: Sender (treasury)
				"022222222222222222222222222222222222222222" + // Tx[0]: Receiver
				"01") // Tx[0]: Amount

		b, _ := block.FromBytes(d)
		assert.NoError(t, b.BasicCheck())
		assert.Zero(t, b.Header().UnixTime())
		assert.Equal(t, b.Header().Version(), uint8(1))
	})
}

func TestCBORMarshaling(t *testing.T) {
	ts := testsuite.NewTestSuite(t)

	b1 := ts.GenerateTestBlock()
	bz1, err := cbor.Marshal(b1)
	assert.NoError(t, err)
	var b2 block.Block
	err = cbor.Unmarshal(bz1, &b2)
	assert.NoError(t, err)
	assert.NoError(t, b2.BasicCheck())
	assert.Equal(t, b1.Hash(), b2.Hash())

	assert.Equal(t, b1.Hash(), b2.Hash())

	err = cbor.Unmarshal([]byte{1}, &b2)
	assert.Error(t, err)
}

func TestEncodingBlock(t *testing.T) {
	ts := testsuite.NewTestSuite(t)

	blk := ts.GenerateTestBlock()
	length := blk.SerializeSize()

	for i := 0; i < length; i++ {
		w := util.NewFixedWriter(i)
		assert.Error(t, blk.Encode(w), "encode test %v failed", i)
	}
	w := util.NewFixedWriter(length)
	assert.NoError(t, blk.Encode(w))

	for i := 0; i < length; i++ {
		blk2 := new(block.Block)
		r := util.NewFixedReader(i, w.Bytes())
		assert.Error(t, blk2.Decode(r), "decode test %v failed", i)
	}

	blk2 := new(block.Block)
	r := util.NewFixedReader(length, w.Bytes())
	assert.NoError(t, blk2.Decode(r))
	assert.Equal(t, blk.Hash(), blk2.Hash())
	assert.Equal(t, blk.Header(), blk2.Header())
}

func TestTxFromBytes(t *testing.T) {
	ts := testsuite.NewTestSuite(t)

	blk := ts.GenerateTestBlock()
	bs, _ := blk.Bytes()
	_, err := block.FromBytes(bs)
	assert.NoError(t, err)
	_, err = blk.Bytes()
	assert.NoError(t, err)

	_, err = block.FromBytes([]byte{1})
	assert.Error(t, err)
}

func TestBlockHash(t *testing.T) {
	d, _ := hex.DecodeString(
		"01" + // Version
			"00000000" + // UnixTime
			"BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB" + // PrevBlockHash
			"BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB" + // StateRoot
			"333333333333333333333333333333333333333333333333" + // SortitionSeed
			"333333333333333333333333333333333333333333333333" +
			"01AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" + // ProposerAddress
			"04030201" + // PrevCert: Height
			"0100" + // PrevCert: Round
			"0401020304" + // PrevCert: Committers
			"0102" + // PrevCert: Absentees
			"b53d79e156e9417e010fa21f2b2a96bee6be46fcd233295d" +
			"2f697cdb9e782b6112ac01c80d0d9d64c2320664c77fa2a6" + // PrevCert: Signature
			"01" + // Txs: Len
			"00" + // Tx[0]: Flags
			"01" + // Tx[0]: Version
			"a1b2c3d4" + // Tx[0]: Stamp
			"01000000" + // Tx[0]: LockTime
			"01" + // Tx[0]: Fee
			"00" + // Tx[0]: Memo
			"01" + // Tx[0]: PayloadType
			"00" + // Tx[0]: Sender (treasury)
			"012222222222222222222222222222222222222222" + // Tx[0]: Receiver
			"01") // Tx[0]: Amount

	b, err := block.FromBytes(d)
	assert.NoError(t, err)
	assert.Equal(t, b.SerializeSize(), len(d))
	d2, _ := b.Bytes()
	assert.Equal(t, d, d2)

	headerSize := b.Header().SerializeSize()
	headerData := d[:headerSize]
	certSize := b.PrevCertificate().SerializeSize()
	certData := d[headerSize : headerSize+certSize]
	certHash := hash.CalcHash(certData)

	txHashes := make([]hash.Hash, 0)
	for _, trx := range b.Transactions() {
		txHashes = append(txHashes, trx.ID())
	}
	txRoot := simplemerkle.NewTreeFromHashes(txHashes).Root()

	hashData := headerData
	hashData = append(hashData, certHash.Bytes()...)
	hashData = append(hashData, txRoot.Bytes()...)
	hashData = append(hashData, util.Int32ToSlice(int32(b.Transactions().Len()))...)

	expected1 := hash.CalcHash(hashData)
	expected2, _ := hash.FromString("3f8364675a5a458eee7c594e92dce03223c87ee66107a6c11de0978b7c7c4bd3")
	assert.Equal(t, b.Hash(), expected1)
	assert.Equal(t, b.Hash(), expected2)
	assert.Equal(t, b.Stamp(), hash.Stamp{0x3f, 0x83, 0x64, 0x67})
}

func TestMakeBlock(t *testing.T) {
	ts := testsuite.NewTestSuite(t)

	b0 := ts.GenerateTestBlock()
	b1 := block.MakeBlock(1, b0.Header().Time(), b0.Transactions(),
		b0.Header().PrevBlockHash(),
		b0.Header().StateRoot(),
		b0.PrevCertificate(),
		b0.Header().SortitionSeed(),
		b0.Header().ProposerAddress())

	assert.Equal(t, b0.Hash(), b1.Hash())
}
