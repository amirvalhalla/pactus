package capnp

import (
	"fmt"

	"github.com/pactus-project/pactus/crypto/hash"
	"github.com/pactus-project/pactus/types/tx"
)

func (zs *pactusServer) GetTransaction(args PactusServer_getTransaction) error {
	data, _ := args.Params.Id()
	h, err := hash.FromBytes(data)
	if err != nil {
		return fmt.Errorf("invalid transaction id: %v", err)
	}
	st := zs.state.StoredTx(h)
	if st == nil {
		return fmt.Errorf("transaction not found")
	}

	res, _ := args.Results.NewResult()
	if err := res.SetData(st.Data); err != nil {
		return err
	}
	return res.SetId(st.TxID.Bytes())
}

// Send broadcasts a raw transaction.
func (zs *pactusServer) SendRawTransaction(args PactusServer_sendRawTransaction) error {
	rawTx, _ := args.Params.RawTx()

	trx, err := tx.FromBytes(rawTx)
	if err != nil {
		return err
	}

	if err := trx.SanityCheck(); err != nil {
		return err
	}

	if err := zs.state.AddPendingTxAndBroadcast(trx); err != nil {
		return err
	}

	res, _ := args.Results.NewResult()
	if err := res.SetId(trx.ID().Bytes()); err != nil {
		return err
	}
	res.SetStatus(0)
	return nil
}
