package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

var ctx = context.Background()

func TestNewDB(t *testing.T) {
	someDB, err := NewDB(ctx, ":memory:")

	assert.Nil(t, err)
	assert.NotNil(t, someDB)
}

func TestInsert(t *testing.T) {
	t.Run("could not insert into address table", func(t *testing.T) {
		someDB, _ := NewDB(ctx, ":memory:")

		addr := &AddressInfo{
			Address:   "some-address",
			PublicKey: "some-public-key",
			Label:     "some-label",
			Path:      "some-path",
		}
		_, err := someDB.InsertAddressInfo(addr)
		assert.EqualError(t, ErrCouldNotInsertRecordIntoTable, err.Error())
	})

	t.Run("insert into address table", func(t *testing.T) {
		someDB, _ := NewDB(ctx, ":memory:")
		_ = someDB.CreateTables()

		addr := &AddressInfo{
			Address:   "some-address",
			PublicKey: "some-public-key",
			Label:     "some-label",
			Path:      "some-path",
		}
		actual, err := someDB.InsertAddressInfo(addr)

		assert.Nil(t, err)
		assert.Equal(t, addr.Address, actual.Address)
	})

	t.Run("could not insert into tranasction table", func(t *testing.T) {
		someDB, _ := NewDB(ctx, ":memory:")

		tr := &Transaction{
			TxID:        "some-txid",
			Address:     "some-address",
			BlockHeight: 4,
			BlockTime:   5,
			PayloadType: "something",
			Data:        "some-data",
			Description: "some-description",
			Amount:      50,
			Status:      1,
		}
		_, err := someDB.InsertTransaction(tr)
		assert.EqualError(t, ErrCouldNotInsertRecordIntoTable, err.Error())
	})

	t.Run("insert into tranasction table", func(t *testing.T) {
		someDB, _ := NewDB(ctx, ":memory:")
		_ = someDB.CreateTables()

		tr := &Transaction{
			TxID:        "some-txid",
			Address:     "some-address",
			BlockHeight: 4,
			BlockTime:   5,
			PayloadType: "something",
			Data:        "some-data",
			Description: "some-description",
			Amount:      50,
			Status:      1,
		}
		actual, err := someDB.InsertTransaction(tr)

		assert.Nil(t, err)
		assert.Equal(t, 1, actual.ID)
		assert.Equal(t, tr.BlockHeight, actual.BlockHeight)
	})

	t.Run("could not insert into pair table", func(t *testing.T) {
		someDB, _ := NewDB(ctx, ":memory:")

		key, value := "key", "value"
		err := someDB.SetValue(key, value)

		assert.EqualError(t, ErrCouldNotInsertRecordIntoTable, err.Error())
	})

	t.Run("insert into pair table", func(t *testing.T) {
		someDB, _ := NewDB(ctx, ":memory:")
		_ = someDB.CreateTables()

		key, value := "key", "value"
		err := someDB.SetValue(key, value)

		assert.Nil(t, err)
	})
}

func TestGetById(t *testing.T) {
	t.Run("could not get transaction by id", func(t *testing.T) {
		someDB, _ := NewDB(ctx, ":memory:")
		_ = someDB.CreateTables()

		tr := &Transaction{
			TxID:        "some-txid",
			Address:     "some-address",
			BlockHeight: 4,
			BlockTime:   5,
			PayloadType: "something",
			Data:        "some-data",
			Description: "some-description",
			Amount:      50,
			Status:      1,
		}
		_, _ = someDB.InsertTransaction(tr)

		actual, err := someDB.GetTransactionByID(10)

		assert.Nil(t, actual)
		assert.EqualError(t, ErrCouldNotFindRecord, err.Error())
	})

	t.Run("get transaction by id", func(t *testing.T) {
		someDB, _ := NewDB(ctx, ":memory:")
		_ = someDB.CreateTables()

		tr := &Transaction{
			TxID:        "some-txid",
			Address:     "some-address",
			BlockHeight: 4,
			BlockTime:   5,
			PayloadType: "something",
			Data:        "some-data",
			Description: "some-description",
			Amount:      50,
			Status:      1,
		}
		expected, _ := someDB.InsertTransaction(tr)

		actual, err := someDB.GetTransactionByID(expected.ID)

		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("could not get pair by key", func(t *testing.T) {
		someDB, _ := NewDB(ctx, ":memory:")
		_ = someDB.CreateTables()

		key, value := "key", "value"
		_ = someDB.SetValue(key, value)

		actual, err := someDB.GetValue("some-thing-wrong")

		assert.Equal(t, "", actual)
		assert.EqualError(t, ErrCouldNotFindRecord, err.Error())
	})

	t.Run("get pair by key", func(t *testing.T) {
		someDB, _ := NewDB(ctx, ":memory:")
		_ = someDB.CreateTables()

		key, value := "key", "value"
		_ = someDB.SetValue(key, value)

		actual, err := someDB.GetValue(key)

		assert.Nil(t, err)
		assert.Equal(t, value, actual)
	})
}

func TestAddress(t *testing.T) {
	t.Run("Could not get address by address", func(t *testing.T) {
		someDB, _ := NewDB(ctx, ":memory:")
		_ = someDB.CreateTables()

		addr := &AddressInfo{
			Address:   "some-pactus-addr",
			PublicKey: "some-public-key",
			Label:     "some-label",
			Path:      "some-path",
		}
		expected, _ := someDB.InsertAddressInfo(addr)

		expected.Address = "some-other-pactus-addr"
		actual, err := someDB.GetAddressInfoByAddress(expected.Address)

		assert.Nil(t, actual)
		assert.EqualError(t, ErrCouldNotFindRecord, err.Error())
	})

	t.Run("Get address by address", func(t *testing.T) {
		someDB, _ := NewDB(ctx, ":memory:")
		_ = someDB.CreateTables()

		addr := &AddressInfo{
			Address:   "some-pactus-addr",
			PublicKey: "some-public-key",
			Label:     "some-label",
			Path:      "some-path",
		}
		expected, _ := someDB.InsertAddressInfo(addr)

		actual, err := someDB.GetAddressInfoByAddress(expected.Address)
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("could not get address by path", func(t *testing.T) {
		someDB, _ := NewDB(ctx, ":memory:")
		_ = someDB.CreateTables()

		addr := &AddressInfo{
			Address:   "some-pactus-addr",
			PublicKey: "some-public-key",
			Label:     "some-label",
			Path:      "some-path",
		}
		expected, _ := someDB.InsertAddressInfo(addr)

		expected.Path = "some-other-path"
		actual, err := someDB.GetAddressByPath(expected.Path)

		assert.Nil(t, actual)
		assert.EqualError(t, ErrCouldNotFindRecord, err.Error())
	})

	t.Run("Get address by path", func(t *testing.T) {
		someDB, _ := NewDB(ctx, ":memory:")
		_ = someDB.CreateTables()

		addr := &AddressInfo{
			Address:   "some-pactus-addr",
			PublicKey: "some-public-key",
			Label:     "some-label",
			Path:      "some-path",
		}
		expected, _ := someDB.InsertAddressInfo(addr)

		actual, err := someDB.GetAddressByPath(expected.Path)
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("update label of address", func(t *testing.T) {
		someDB, _ := NewDB(ctx, ":memory:")
		_ = someDB.CreateTables()

		addr := &AddressInfo{
			Address:   "some-pactus-addr",
			PublicKey: "some-public-key",
			Label:     "some-label",
			Path:      "some-path",
		}
		addr, _ = someDB.InsertAddressInfo(addr)

		addr.Label = "some-other-lable"
		_ = someDB.UpdateAddressLabel(addr.Label, addr.Address)

		actual, err := someDB.GetAddressInfoByAddress(addr.Address)

		assert.Nil(t, err)
		assert.Equal(t, addr.Label, actual.Label)
	})
}

func TestTransaction(t *testing.T) {
	t.Run("could not get transaction by tx id", func(t *testing.T) {
		someDB, _ := NewDB(ctx, ":memory:")
		_ = someDB.CreateTables()

		tr := &Transaction{
			TxID:        "some-txid",
			Address:     "some-address",
			BlockHeight: 4,
			BlockTime:   5,
			PayloadType: "something",
			Data:        "some-data",
			Description: "some-description",
			Amount:      50,
			Status:      1,
		}
		_, _ = someDB.InsertTransaction(tr)

		actual, err := someDB.GetTransactionByTxID("unknown-txid")

		assert.Nil(t, actual)
		assert.EqualError(t, ErrCouldNotFindRecord, err.Error())
	})

	t.Run("get transaction by tx id", func(t *testing.T) {
		someDB, _ := NewDB(ctx, ":memory:")
		_ = someDB.CreateTables()

		tr := &Transaction{
			TxID:        "some-txid",
			Address:     "some-address",
			BlockHeight: 4,
			BlockTime:   5,
			PayloadType: "something",
			Data:        "some-data",
			Description: "some-description",
			Amount:      50,
			Status:      1,
		}
		expected, _ := someDB.InsertTransaction(tr)

		actual, err := someDB.GetTransactionByTxID(expected.TxID)

		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})
}

func TestGetAll(t *testing.T) {
	t.Run("get all addresses", func(t *testing.T) {
		someDB, _ := NewDB(ctx, ":memory:")
		_ = someDB.CreateTables()

		addr := &AddressInfo{
			PublicKey: "some-public-key",
			Label:     "some-label",
			Path:      "some-path",
		}

		addr.Address = "addr1"
		someInsertOne, _ := someDB.InsertAddressInfo(addr)
		addr.Address = "addr2"
		someInsertTwo, _ := someDB.InsertAddressInfo(addr)
		addr.Address = "addr3"
		someInsertThree, _ := someDB.InsertAddressInfo(addr)

		expected := make([]AddressInfo, 0, 3)
		expected = append(expected, *someInsertThree, *someInsertTwo, *someInsertOne)

		acutal, err := someDB.GetAllAddressInfos()

		assert.Nil(t, err)
		assert.Equal(t, expected, acutal)
	})

	t.Run("get all addresses with total records", func(t *testing.T) {
		someDB, _ := NewDB(ctx, ":memory:")
		_ = someDB.CreateTables()

		addr := &AddressInfo{
			PublicKey: "some-public-key",
			Label:     "some-label",
			Path:      "some-path",
		}

		addr.Address = "addr1"
		someInsertOne, _ := someDB.InsertAddressInfo(addr)
		addr.Address = "addr2"
		someInsertTwo, _ := someDB.InsertAddressInfo(addr)
		addr.Address = "addr3"
		someInsertThree, _ := someDB.InsertAddressInfo(addr)

		expected := make([]AddressInfo, 0, 3)
		expected = append(expected, *someInsertThree, *someInsertTwo, *someInsertOne)

		acutal, totalRecords, err := someDB.GetAllAddressInfosWithTotalRecords(1, 3)

		assert.Nil(t, err)
		assert.Equal(t, int64(3), totalRecords)
		assert.Equal(t, expected, acutal)
	})

	t.Run("get all transactions", func(t *testing.T) {
		someDB, _ := NewDB(ctx, ":memory:")
		_ = someDB.CreateTables()

		tr := &Transaction{
			TxID:        "some-txid",
			Address:     "some-address",
			BlockHeight: 4,
			BlockTime:   5,
			PayloadType: "something",
			Data:        "some-data",
			Description: "some-description",
			Amount:      50,
			Status:      1,
		}
		someInsertOne, _ := someDB.InsertTransaction(tr)
		someInsertTwo, _ := someDB.InsertTransaction(tr)
		someInsertThree, _ := someDB.InsertTransaction(tr)

		expected := make([]Transaction, 0, 3)
		expected = append(expected, *someInsertThree, *someInsertTwo, *someInsertOne)

		acutal, err := someDB.GetAllTransactions(EmptyQuery)

		assert.Nil(t, err)
		assert.Equal(t, expected, acutal)
	})

	t.Run("get all transactions with transaction status query option", func(t *testing.T) {
		someDB, _ := NewDB(ctx, ":memory:")
		_ = someDB.CreateTables()

		tr := &Transaction{
			TxID:        "some-txid",
			Address:     "some-address",
			BlockHeight: 4,
			BlockTime:   5,
			PayloadType: "something",
			Data:        "some-data",
			Description: "some-description",
			Amount:      50,
			Status:      int(Confirmed),
		}
		someInsertOne, _ := someDB.InsertTransaction(tr)
		someInsertTwo, _ := someDB.InsertTransaction(tr)

		tr.Status = int(Pending)
		_, _ = someDB.InsertTransaction(tr)

		expected := make([]Transaction, 0, 2)
		expected = append(expected, *someInsertTwo, *someInsertOne)

		acutal, err := someDB.GetAllTransactions(WithTransactionStatus(), Confirmed)

		assert.Nil(t, err)
		assert.Equal(t, expected, acutal)
	})

	t.Run("get all transactions with transaction address query option", func(t *testing.T) {
		someDB, _ := NewDB(ctx, ":memory:")
		_ = someDB.CreateTables()

		tr := &Transaction{
			TxID:        "some-txid",
			Address:     "some-address",
			BlockHeight: 4,
			BlockTime:   5,
			PayloadType: "something",
			Data:        "some-data",
			Description: "some-description",
			Amount:      50,
			Status:      int(Confirmed),
		}
		someInsertOne, _ := someDB.InsertTransaction(tr)
		someInsertTwo, _ := someDB.InsertTransaction(tr)

		tr.Address = "some-another-address"
		_, _ = someDB.InsertTransaction(tr)

		expected := make([]Transaction, 0, 2)
		expected = append(expected, *someInsertTwo, *someInsertOne)

		acutal, err := someDB.GetAllTransactions(WithTransactionAddr(), "some-address")

		assert.Nil(t, err)
		assert.Equal(t, expected, acutal)
	})

	t.Run("get all transactions with total records", func(t *testing.T) {
		someDB, _ := NewDB(ctx, ":memory:")
		_ = someDB.CreateTables()

		tr := &Transaction{
			TxID:        "some-txid",
			BlockHeight: 4,
			BlockTime:   5,
			PayloadType: "something",
			Data:        "some-data",
			Description: "some-description",
			Amount:      50,
			Status:      1,
		}
		someInsertOne, _ := someDB.InsertTransaction(tr)
		someInsertTwo, _ := someDB.InsertTransaction(tr)
		someInsertThree, _ := someDB.InsertTransaction(tr)

		expected := make([]Transaction, 0, 3)
		expected = append(expected, *someInsertThree, *someInsertTwo, *someInsertOne)

		acutal, totalRecords, err := someDB.GetAllTransactionsWithTotalRecords(1, 3, EmptyQuery)

		assert.Nil(t, err)
		assert.Equal(t, int64(3), totalRecords)
		assert.Equal(t, expected, acutal)
	})

	t.Run("get all transactions with total records and transaction status query option", func(t *testing.T) {
		someDB, _ := NewDB(ctx, ":memory:")
		_ = someDB.CreateTables()

		tr := &Transaction{
			TxID:        "some-txid",
			BlockHeight: 4,
			BlockTime:   5,
			PayloadType: "something",
			Data:        "some-data",
			Description: "some-description",
			Amount:      50,
			Status:      int(Confirmed),
		}
		_, _ = someDB.InsertTransaction(tr)
		someInsertTwo, _ := someDB.InsertTransaction(tr)

		tr.Status = int(Pending)
		_, _ = someDB.InsertTransaction(tr)

		expected := make([]Transaction, 0, 2)
		expected = append(expected, *someInsertTwo)

		acutal, totalRecords, err := someDB.GetAllTransactionsWithTotalRecords(1, 1,
			WithTransactionStatus(), Confirmed)

		assert.Nil(t, err)
		assert.Equal(t, int64(2), totalRecords)
		assert.Equal(t, expected, acutal)
	})
}

func TestTotalRecords(t *testing.T) {
	t.Run("could not find total records", func(t *testing.T) {
		someDB, _ := NewDB(ctx, ":memory:")
		_ = someDB.CreateTables()

		addr := &AddressInfo{
			Address:   "some-address",
			PublicKey: "some-public-key",
			Label:     "some-label",
			Path:      "some-path",
		}
		_, _ = someDB.InsertAddressInfo(addr)
		_, _ = someDB.InsertAddressInfo(addr)
		_, _ = someDB.InsertAddressInfo(addr)

		totalRecords, err := someDB.GetTotalRecords("some-table", EmptyQuery)

		assert.Equal(t, int64(0), totalRecords)
		assert.EqualError(t, ErrCouldNotFindTotalRecords, err.Error())
	})

	t.Run("ok without any query option", func(t *testing.T) {
		someDB, _ := NewDB(ctx, ":memory:")
		_ = someDB.CreateTables()

		addr := &AddressInfo{
			PublicKey: "some-public-key",
			Label:     "some-label",
			Path:      "some-path",
		}

		addr.Address = "addr1"
		_, _ = someDB.InsertAddressInfo(addr)
		addr.Address = "addr2"
		_, _ = someDB.InsertAddressInfo(addr)
		addr.Address = "addr3"
		_, _ = someDB.InsertAddressInfo(addr)

		totalRecords, err := someDB.GetTotalRecords(AddressTable, EmptyQuery)

		assert.Nil(t, err)
		assert.Equal(t, int64(3), totalRecords)
	})

	t.Run("ok with transaction status query option", func(t *testing.T) {
		someDB, _ := NewDB(ctx, ":memory:")
		_ = someDB.CreateTables()

		tr := &Transaction{
			TxID:        "some-txid",
			Address:     "some-address",
			BlockHeight: 4,
			BlockTime:   5,
			PayloadType: "something",
			Data:        "some-data",
			Description: "some-description",
			Amount:      50,
			Status:      int(Confirmed),
		}
		_, _ = someDB.InsertTransaction(tr)
		_, _ = someDB.InsertTransaction(tr)

		tr.Status = int(Pending)
		_, _ = someDB.InsertTransaction(tr)

		totalRecords, err := someDB.GetTotalRecords(TransactionTable, WithTransactionStatus(), Confirmed)

		assert.Nil(t, err)
		assert.Equal(t, int64(2), totalRecords)
	})

	t.Run("ok with transaction status and addr query option", func(t *testing.T) {
		someDB, _ := NewDB(ctx, ":memory:")
		_ = someDB.CreateTables()

		tr := &Transaction{
			TxID:        "some-txid",
			Address:     "some-address",
			BlockHeight: 4,
			BlockTime:   5,
			PayloadType: "something",
			Data:        "some-data",
			Description: "some-description",
			Amount:      50,
			Status:      int(Confirmed),
		}
		_, _ = someDB.InsertTransaction(tr)

		tr.Address = "some-another-address"
		_, _ = someDB.InsertTransaction(tr)

		tr.Status = int(Pending)
		tr.Address = "some-address"
		_, _ = someDB.InsertTransaction(tr)

		totalRecords, err := someDB.GetTotalRecords(TransactionTable,
			WithTransactionStatusAndAddr(), Confirmed, "some-address")

		assert.Nil(t, err)
		assert.Equal(t, int64(1), totalRecords)
	})

	t.Run("ok with transaction address query option", func(t *testing.T) {
		someDB, _ := NewDB(ctx, ":memory:")
		_ = someDB.CreateTables()

		tr := &Transaction{
			TxID:        "some-txid",
			Address:     "some-address",
			BlockHeight: 4,
			BlockTime:   5,
			PayloadType: "something",
			Data:        "some-data",
			Description: "some-description",
			Amount:      50,
			Status:      int(Confirmed),
		}
		_, _ = someDB.InsertTransaction(tr)

		tr.Address = "some-another-address"
		_, _ = someDB.InsertTransaction(tr)

		tr.Status = int(Pending)
		tr.Address = "some-address"
		_, _ = someDB.InsertTransaction(tr)

		totalRecords, err := someDB.GetTotalRecords(TransactionTable,
			WithTransactionAddr(), "some-address")

		assert.Nil(t, err)
		assert.Equal(t, int64(2), totalRecords)
	})
}