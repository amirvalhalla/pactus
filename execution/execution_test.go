package execution

import (
	"testing"

	"github.com/pactus-project/pactus/crypto"
	"github.com/pactus-project/pactus/sandbox"
	"github.com/pactus-project/pactus/types/tx"
	"github.com/pactus-project/pactus/util/errors"
	"github.com/pactus-project/pactus/util/testsuite"
	"github.com/stretchr/testify/assert"
)

func TestLockTime(t *testing.T) {
	ts := testsuite.NewTestSuite(t)

	sb := sandbox.MockingSandbox(ts)
	sb.TestAcceptSortition = true
	exe := NewExecutor()
	rndValKey := ts.RandValKey()
	addr1 := rndValKey.Address()
	acc1 := sb.MakeNewAccount(addr1)
	acc1.AddToBalance(100 * 1e9)
	sb.UpdateAccount(addr1, acc1)
	rcvAddr := ts.RandAccAddress()
	_ = sb.TestStore.AddTestBlock(8642)

	t.Run("Future LockTime, Should returns error (+1)", func(t *testing.T) {
		lockTime := sb.CurrentHeight() + 1
		trx := tx.NewTransferTx(lockTime, addr1, rcvAddr, 1000, 1000, "expired-stamp")
		ts.HelperSignTransaction(rndValKey.PrivateKey(), trx)
		err := exe.Execute(trx, sb)
		assert.ErrorIs(t, err, FutureLockTimeError{LockTime: lockTime})
	})

	t.Run("Past LockTime, Should returns error (-8641)", func(t *testing.T) {
		lockTime := sb.CurrentHeight() - sb.TestParams.TransactionToLiveInterval - 1
		trx := tx.NewTransferTx(lockTime, addr1, rcvAddr, 1000, 1000, "expired-stamp")
		ts.HelperSignTransaction(rndValKey.PrivateKey(), trx)
		err := exe.Execute(trx, sb)
		assert.ErrorIs(t, err, PastLockTimeError{LockTime: lockTime})
	})

	t.Run("Transaction has valid LockTime (-8640)", func(t *testing.T) {
		lockTime := sb.CurrentHeight() - sb.TestParams.TransactionToLiveInterval
		trx := tx.NewTransferTx(lockTime, addr1, rcvAddr, 1000, 1000, "ok")
		ts.HelperSignTransaction(rndValKey.PrivateKey(), trx)
		err := exe.Execute(trx, sb)
		assert.NoError(t, err)
	})

	t.Run("Transaction has valid LockTime (0)", func(t *testing.T) {
		lockTime := sb.CurrentHeight()
		trx := tx.NewTransferTx(lockTime, addr1, rcvAddr, 1000, 1000, "ok")
		ts.HelperSignTransaction(rndValKey.PrivateKey(), trx)
		err := exe.Execute(trx, sb)
		assert.NoError(t, err)
	})

	t.Run("Subsidy transaction has invalid LockTime (+1)", func(t *testing.T) {
		lockTime := sb.CurrentHeight() + 1
		trx := tx.NewSubsidyTx(lockTime, rcvAddr, 1000,
			"invalid-lockTime")
		err := exe.Execute(trx, sb)
		assert.ErrorIs(t, err, FutureLockTimeError{LockTime: lockTime})
	})

	t.Run("Subsidy transaction has invalid LockTime (-1)", func(t *testing.T) {
		lockTime := sb.CurrentHeight() - 1
		trx := tx.NewSubsidyTx(lockTime, rcvAddr, 1000,
			"invalid-lockTime")
		err := exe.Execute(trx, sb)
		assert.ErrorIs(t, err, PastLockTimeError{LockTime: lockTime})
	})

	t.Run("Subsidy transaction has valid LockTime (0)", func(t *testing.T) {
		lockTime := sb.CurrentHeight()
		trx := tx.NewSubsidyTx(lockTime, rcvAddr, 1000, "ok")
		err := exe.Execute(trx, sb)
		assert.NoError(t, err)
	})

	t.Run("Sortition transaction has invalid LockTime (+1)", func(t *testing.T) {
		lockTime := sb.CurrentHeight() + 1
		proof := ts.RandProof()
		trx := tx.NewSortitionTx(lockTime, addr1, proof)
		ts.HelperSignTransaction(rndValKey.PrivateKey(), trx)
		err := exe.Execute(trx, sb)
		assert.ErrorIs(t, err, FutureLockTimeError{LockTime: lockTime})
	})

	t.Run("Sortition transaction has invalid LockTime (-8)", func(t *testing.T) {
		lockTime := sb.CurrentHeight() - sb.TestParams.SortitionInterval - 1
		proof := ts.RandProof()
		trx := tx.NewSortitionTx(lockTime, addr1, proof)
		ts.HelperSignTransaction(rndValKey.PrivateKey(), trx)
		err := exe.Execute(trx, sb)
		assert.ErrorIs(t, err, PastLockTimeError{LockTime: lockTime})
	})

	t.Run("Sortition transaction has valid LockTime (-7)", func(t *testing.T) {
		lockTime := sb.CurrentHeight() - sb.TestParams.SortitionInterval
		proof := ts.RandProof()
		val1 := sb.MakeNewValidator(rndValKey.PublicKey())
		val1.AddToStake(100 * 1e9)
		sb.UpdateValidator(val1)

		trx := tx.NewSortitionTx(lockTime, addr1, proof)
		ts.HelperSignTransaction(rndValKey.PrivateKey(), trx)
		err := exe.Execute(trx, sb)
		assert.NoError(t, err)
	})
}

func TestExecution(t *testing.T) {
	ts := testsuite.NewTestSuite(t)

	sb := sandbox.MockingSandbox(ts)
	exe := NewExecutor()

	rndValKey := ts.RandValKey()
	addr1 := rndValKey.Address()
	acc1 := sb.MakeNewAccount(addr1)
	acc1.AddToBalance(100 * 1e9)
	sb.UpdateAccount(addr1, acc1)
	rcvAddr := ts.RandAccAddress()
	_ = sb.TestStore.AddTestBlock(8642)
	lockTime := sb.CurrentHeight()

	t.Run("Invalid transaction, Should returns error", func(t *testing.T) {
		trx := tx.NewTransferTx(lockTime, ts.RandAccAddress(), rcvAddr, 1000, 1000, "invalid-tx")
		err := exe.Execute(trx, sb)
		assert.Equal(t, errors.Code(err), errors.ErrInvalidAddress)
	})

	t.Run("Invalid fee, Should returns error", func(t *testing.T) {
		trx := tx.NewTransferTx(lockTime, addr1, rcvAddr, 1000, 1, "invalid fee")
		ts.HelperSignTransaction(rndValKey.PrivateKey(), trx)
		err := exe.Execute(trx, sb)
		assert.Equal(t, errors.Code(err), errors.ErrInvalidFee)
	})

	t.Run("Invalid fee, Should returns error", func(t *testing.T) {
		trx := tx.NewTransferTx(lockTime, addr1, rcvAddr, 1000, 1001, "invalid fee")
		ts.HelperSignTransaction(rndValKey.PrivateKey(), trx)
		err := exe.Execute(trx, sb)
		assert.Equal(t, errors.Code(err), errors.ErrInvalidFee)
	})

	t.Run("Invalid fee (subsidy tx), Should returns error", func(t *testing.T) {
		trx := tx.NewTransferTx(lockTime, crypto.TreasuryAddress, rcvAddr, 1000, 1, "invalid fee")
		err := exe.Execute(trx, sb)
		assert.Equal(t, errors.Code(err), errors.ErrInvalidFee)
		assert.Error(t, exe.checkFee(trx, sb))
	})

	t.Run("Invalid fee (send tx), Should returns error", func(t *testing.T) {
		trx := tx.NewTransferTx(lockTime, addr1, rcvAddr, 1000, 0, "invalid fee")
		err := exe.Execute(trx, sb)
		assert.Equal(t, errors.Code(err), errors.ErrInvalidFee)
		assert.Error(t, exe.checkFee(trx, sb))
	})

	t.Run("Execution failed", func(t *testing.T) {
		proof := ts.RandProof()
		trx := tx.NewSortitionTx(lockTime, addr1, proof)
		ts.HelperSignTransaction(rndValKey.PrivateKey(), trx)
		err := exe.Execute(trx, sb)
		assert.Equal(t, errors.Code(err), errors.ErrInvalidAddress)
	})
}

func TestReplay(t *testing.T) {
	ts := testsuite.NewTestSuite(t)

	executor := NewExecutor()
	sb := sandbox.MockingSandbox(ts)
	rndValKey := ts.RandValKey()
	addr1 := rndValKey.Address()
	acc := sb.MakeNewAccount(addr1)
	acc.AddToBalance(1e9)
	sb.UpdateAccount(addr1, acc)
	lockTime := sb.CurrentHeight()

	trx := tx.NewTransferTx(lockTime,
		addr1, ts.RandAccAddress(), 10000, 1000, "")
	ts.HelperSignTransaction(rndValKey.PrivateKey(), trx)

	err := executor.Execute(trx, sb)
	assert.NoError(t, err)
	err = executor.Execute(trx, sb)
	assert.ErrorIs(t, err, TransactionCommittedError{
		ID: trx.ID(),
	})
}

func TestChecker(t *testing.T) {
	ts := testsuite.NewTestSuite(t)

	executor := NewExecutor()
	checker := NewChecker()
	sb := sandbox.MockingSandbox(ts)
	rndValKey := ts.RandValKey()
	addr1 := rndValKey.Address()
	acc := sb.MakeNewAccount(addr1)
	acc.AddToBalance(1e9)
	sb.UpdateAccount(addr1, acc)
	lockTime := sb.CurrentHeight() + 1

	trx := tx.NewTransferTx(lockTime,
		addr1, ts.RandAccAddress(), 10000, 1000, "")
	ts.HelperSignTransaction(rndValKey.PrivateKey(), trx)

	err := executor.Execute(trx, sb)
	assert.ErrorIs(t, err, FutureLockTimeError{LockTime: lockTime})
	err = checker.Execute(trx, sb)
	assert.NoError(t, err)
}

func TestFee(t *testing.T) {
	ts := testsuite.NewTestSuite(t)

	exe := NewChecker()
	sb := sandbox.MockingSandbox(ts)

	tests := []struct {
		amount          int64
		fee             int64
		expectedFee     int64
		expectedErrCode int
	}{
		{1, 1, sb.TestParams.MinimumFee, errors.ErrInvalidFee},
		{1, 1001, sb.TestParams.MinimumFee, errors.ErrInvalidFee},
		{1, 1000, sb.TestParams.MinimumFee, errors.ErrNone},

		{1 * 1e9, 1, 100000, errors.ErrInvalidFee},
		{1 * 1e9, 100001, 100000, errors.ErrInvalidFee},
		{1 * 1e9, 100000, 100000, errors.ErrNone},

		{1 * 1e12, 1, 1000000, errors.ErrInvalidFee},
		{1 * 1e12, 1000001, 1000000, errors.ErrInvalidFee},
		{1 * 1e12, 1000000, 1000000, errors.ErrNone},
	}

	sender := ts.RandAccAddress()
	receiver := ts.RandAccAddress()
	for i, test := range tests {
		trx := tx.NewTransferTx(sb.CurrentHeight()+1, sender, receiver, test.amount, test.fee,
			"testing fee")
		err := exe.checkFee(trx, sb)

		assert.Equal(t, errors.Code(err), test.expectedErrCode,
			"test %v failed. unexpected error", i)

		assert.Equal(t, CalculateFee(test.amount, sb.Params()), test.expectedFee,
			"test %v failed. invalid fee", i)
	}
}
