package db

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	//nolint:all
	_ "github.com/glebarez/go-sqlite"
)

const (
	AddressTable     = "addresses"
	TransactionTable = "transactions"
	PairTable        = "pairs"
)

type DB interface {
	CreateTables() error

	InsertIntoAddress(addr *Address) (*Address, error)
	InsertIntoTransaction(t *Transaction) (*Transaction, error)
	InsertIntoPair(key string, value string) (*Pair, error)

	UpdateAddressLabel(addr *Address) (*Address, error)

	GetAddressByID(id int) (*Address, error)
	GetAddressByAddress(address string) (*Address, error)
	GetAddressByPath(p string) (*Address, error)
	GetTransactionByID(id int) (*Transaction, error)
	GetPairByKey(key string) (*Pair, error)
	GetTotalRecords(tableName string) (int64, error)

	GetAllAddresses() ([]Address, error)
	GetAllTransactions() ([]Transaction, error)
	GetAllAddressesWithTotalRecords(pageIndex, pageSize int) ([]Address, int64, error)
	GetAllTransactionsWithTotalRecords(pageIndex, pageSize int) ([]Transaction, int64, error)
}

type db struct {
	*sql.DB
}

type Address struct {
	ID         int       `json:"id"`          // id of wallet
	Address    string    `json:"address"`     // Address in the wallet
	PublicKey  string    `json:"public_key"`  // Public key associated with the address
	Label      string    `json:"label"`       // Label for the address
	Path       string    `json:"path"`        // Path for the address
	IsImported bool      `json:"is_imported"` // imported for purpose
	CreatedAt  time.Time `json:"created_at"`
}

type Transaction struct {
	ID          int       `json:"id"`
	TxID        string    `json:"tx_id"`
	BlockHeight uint32    `json:"block_height"`
	BlockTime   uint32    `json:"block_time"`
	PayloadType string    `json:"payload_type"`
	Data        string    `json:"data"`
	Description string    `json:"description"`
	Amount      int64     `json:"amount"`
	Status      int       `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

type Pair struct {
	Key       string
	Value     string
	CreatedAt time.Time
}

func NewDB(path string) (DB, error) {
	dbInstance, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, ErrCouldNotOpenDatabase
	}

	return &db{
		dbInstance,
	}, nil
}

func (d *db) CreateTables() error {
	if err := d.createAddressTable(); err != nil {
		return err
	}

	if err := d.createTransactionTable(); err != nil {
		return err
	}

	return d.createPairTable()
}

func (d *db) createAddressTable() error {
	addressQuery := fmt.Sprintf("CREATE TABLE %s (id INTEGER PRIMARY KEY AUTOINCREMENT,"+
		" address VARCHAR, public_key VARCHAR, label VARCHAR, path VARCHAR, is_imported BOOLEAN, created_at TIMESTAMP)",
		AddressTable)

	_, err := d.ExecContext(context.Background(), addressQuery)
	if err != nil && !strings.Contains(err.Error(), "already exists") {
		return ErrCouldNotCreateTable
	}

	return nil
}

func (d *db) createTransactionTable() error {
	transactionQuery := fmt.Sprintf("CREATE TABLE %s (id INTEGER PRIMARY KEY AUTOINCREMENT,"+
		" tx_id VARCHAR, block_height INTEGER, block_time INTEGER, payload_type VARCHAR,"+
		" data VARCHAR, description VARCHAR, amount BIGINT,status INTEGER, created_at TIMESTAMP)", TransactionTable)
	_, err := d.ExecContext(context.Background(), transactionQuery)
	if err != nil && !strings.Contains(err.Error(), "already exists") {
		return ErrCouldNotCreateTable
	}

	return nil
}

func (d *db) createPairTable() error {
	pairQuery := fmt.Sprintf("CREATE TABLE %s (key VARCHAR PRIMARY KEY, value VARCHAR, created_at TIMESTAMP)", PairTable)
	_, err := d.ExecContext(context.Background(), pairQuery)
	if err != nil && !strings.Contains(err.Error(), "already exists") {
		return ErrCouldNotCreateTable
	}

	return nil
}

func (d *db) InsertIntoAddress(addr *Address) (*Address, error) {
	insertQuery := fmt.Sprintf("INSERT INTO %s (address, public_key, label, path, is_imported, created_at)"+
		" VALUES (?,?,?,?,?,?)", AddressTable)

	addr.CreatedAt = time.Now().UTC()
	r, err := d.ExecContext(context.Background(), insertQuery, addr.Address,
		addr.PublicKey, addr.Label, addr.Path, addr.IsImported, addr.CreatedAt)
	if err != nil {
		return nil, ErrCouldNotInsertRecordIntoTable
	}

	rowID, err := r.LastInsertId()
	if err != nil {
		return nil, ErrCouldNotInsertRecordIntoTable
	}

	return &Address{
		ID:         int(rowID),
		Address:    addr.Address,
		PublicKey:  addr.PublicKey,
		Label:      addr.Label,
		Path:       addr.Path,
		IsImported: addr.IsImported,
		CreatedAt:  addr.CreatedAt,
	}, nil
}

func (d *db) InsertIntoTransaction(t *Transaction) (*Transaction, error) {
	insertQuery := fmt.Sprintf("INSERT INTO %s (tx_id, block_height, block_time,"+
		" payload_type, data, description, amount, status, created_at) VALUES"+
		" (?,?,?,?,?,?,?,?,?)", TransactionTable)

	t.CreatedAt = time.Now().UTC()
	r, err := d.ExecContext(context.Background(), insertQuery, t.TxID, t.BlockHeight, t.BlockTime,
		t.PayloadType, t.Data, t.Description, t.Amount, t.Status, t.CreatedAt)
	if err != nil {
		return nil, ErrCouldNotInsertRecordIntoTable
	}

	rowID, err := r.LastInsertId()
	if err != nil {
		return nil, ErrCouldNotInsertRecordIntoTable
	}

	return &Transaction{
		ID:          int(rowID),
		TxID:        t.TxID,
		BlockHeight: t.BlockHeight,
		BlockTime:   t.BlockTime,
		PayloadType: t.PayloadType,
		Data:        t.Data,
		Description: t.Description,
		Amount:      t.Amount,
		Status:      t.Status,
		CreatedAt:   t.CreatedAt,
	}, nil
}

func (d *db) InsertIntoPair(key, value string) (*Pair, error) {
	createdAt := time.Now().UTC()
	insertQuery := fmt.Sprintf("INSERT INTO %s (key, value, created_at) VALUES (?,?,?)", PairTable)
	_, err := d.ExecContext(context.Background(), insertQuery, key, value, createdAt)
	if err != nil {
		return nil, ErrCouldNotInsertRecordIntoTable
	}

	return &Pair{
		Key:       key,
		Value:     value,
		CreatedAt: createdAt,
	}, nil
}

func (d *db) UpdateAddressLabel(addr *Address) (*Address, error) {
	insertQuery := fmt.Sprintf("UPDATE %s SET label = ? WHERE address = ?", AddressTable)

	r, err := d.ExecContext(context.Background(), insertQuery, addr.Label, addr.Address)
	if err != nil {
		return nil, ErrCouldNotUpdateRecordIntoTable
	}

	rowID, err := r.LastInsertId()
	if err != nil {
		return nil, ErrCouldNotUpdateRecordIntoTable
	}

	return &Address{
		ID:         int(rowID),
		Address:    addr.Address,
		PublicKey:  addr.PublicKey,
		Label:      addr.Label,
		Path:       addr.Path,
		IsImported: addr.IsImported,
		CreatedAt:  addr.CreatedAt,
	}, nil
}

func (d *db) GetAddressByID(id int) (*Address, error) {
	getQuery := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", AddressTable)
	row := d.QueryRowContext(context.Background(), getQuery, id)
	if row.Err() != nil {
		return nil, ErrCouldNotFindRecord
	}

	addr := &Address{}
	err := row.Scan(&addr.ID, &addr.Address, &addr.PublicKey, &addr.Label,
		&addr.Path, &addr.IsImported, &addr.CreatedAt)
	if err != nil {
		return nil, ErrCouldNotFindRecord
	}

	return addr, nil
}

func (d *db) GetAddressByAddress(address string) (*Address, error) {
	getQuery := fmt.Sprintf("SELECT * FROM %s WHERE address = ?", AddressTable)
	row := d.QueryRowContext(context.Background(), getQuery, address)
	if row.Err() != nil {
		return nil, ErrCouldNotFindRecord
	}

	addr := &Address{}
	err := row.Scan(&addr.ID, &addr.Address, &addr.PublicKey, &addr.Label,
		&addr.Path, &addr.IsImported, &addr.CreatedAt)
	if err != nil {
		return nil, ErrCouldNotFindRecord
	}

	return addr, nil
}

func (d *db) GetAddressByPath(p string) (*Address, error) {
	getQuery := fmt.Sprintf("SELECT * FROM %s WHERE path = ?", AddressTable)
	row := d.QueryRowContext(context.Background(), getQuery, p)
	if row.Err() != nil {
		return nil, ErrCouldNotFindRecord
	}

	addr := &Address{}
	err := row.Scan(&addr.ID, &addr.Address, &addr.PublicKey, &addr.Label,
		&addr.Path, &addr.IsImported, &addr.CreatedAt)
	if err != nil {
		return nil, ErrCouldNotFindRecord
	}

	return addr, nil
}

func (d *db) GetTransactionByID(id int) (*Transaction, error) {
	getQuery := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", TransactionTable)
	row := d.QueryRowContext(context.Background(), getQuery, id)
	if row.Err() != nil {
		return nil, ErrCouldNotFindRecord
	}

	t := &Transaction{}
	err := row.Scan(&t.ID, &t.TxID, &t.BlockHeight, &t.BlockTime, &t.PayloadType,
		&t.Data, &t.Description, &t.Amount, &t.Status, &t.CreatedAt)
	if err != nil {
		return nil, ErrCouldNotFindRecord
	}

	return t, nil
}

func (d *db) GetPairByKey(key string) (*Pair, error) {
	getQuery := fmt.Sprintf("SELECT * FROM %s WHERE key = ?", PairTable)
	row := d.QueryRowContext(context.Background(), getQuery, key)
	if row.Err() != nil {
		return nil, ErrCouldNotFindRecord
	}

	p := &Pair{}
	err := row.Scan(&p.Key, &p.Value, &p.CreatedAt)
	if err != nil {
		return nil, ErrCouldNotFindRecord
	}

	return p, nil
}

func (d *db) GetAllAddresses() ([]Address, error) {
	getAllQuery := fmt.Sprintf("SELECT * FROM %s ORDER BY id DESC", AddressTable)
	rows, err := d.QueryContext(context.Background(), getAllQuery)
	if err != nil || rows.Err() != nil {
		return nil, ErrCouldNotFindRecord
	}
	defer rows.Close()

	addrs := make([]Address, 0)
	for rows.Next() {
		addr := &Address{}
		err := rows.Scan(&addr.ID, &addr.Address, &addr.PublicKey, &addr.Label, &addr.Path, &addr.IsImported, &addr.CreatedAt)
		if err != nil {
			return nil, ErrCouldNotFindRecord
		}

		addrs = append(addrs, *addr)
	}

	return addrs, nil
}

func (d *db) GetAllTransactions() ([]Transaction, error) {
	getAllQuery := fmt.Sprintf("SELECT * FROM %s ORDER BY id DESC", TransactionTable)
	rows, err := d.QueryContext(context.Background(), getAllQuery)
	if err != nil || rows.Err() != nil {
		return nil, ErrCouldNotFindRecord
	}
	defer rows.Close()

	transactions := make([]Transaction, 0)
	for rows.Next() {
		t := &Transaction{}
		err := rows.Scan(&t.ID, &t.TxID, &t.BlockHeight, &t.BlockTime, &t.PayloadType,
			&t.Data, &t.Description, &t.Amount, &t.Status, &t.CreatedAt)
		if err != nil {
			return nil, ErrCouldNotFindRecord
		}

		transactions = append(transactions, *t)
	}

	return transactions, nil
}

func (d *db) GetAllAddressesWithTotalRecords(pageIndex, pageSize int) ([]Address, int64, error) {
	totalRecords, err := d.GetTotalRecords("addresses")
	if err != nil {
		return nil, 0, err
	}

	getAllQuery := fmt.Sprintf("SELECT * FROM %s ORDER BY id DESC LIMIT ? OFFSET ?", AddressTable)
	rows, err := d.QueryContext(context.Background(), getAllQuery, pageSize, calcOffset(pageIndex, pageSize))
	if err != nil || rows.Err() != nil {
		return nil, 0, ErrCouldNotFindRecord
	}
	defer rows.Close()

	addrs := make([]Address, 0, pageSize)
	for rows.Next() {
		addr := &Address{}
		err := rows.Scan(&addr.ID, &addr.Address, &addr.PublicKey, &addr.Label, &addr.Path, &addr.IsImported, &addr.CreatedAt)
		if err != nil {
			return nil, 0, ErrCouldNotFindRecord
		}

		addrs = append(addrs, *addr)
	}

	return addrs, totalRecords, nil
}

func (d *db) GetAllTransactionsWithTotalRecords(pageIndex, pageSize int) ([]Transaction, int64, error) {
	totalRecords, err := d.GetTotalRecords("transactions")
	if err != nil {
		return nil, 0, err
	}

	getAllQuery := fmt.Sprintf("SELECT * FROM %s ORDER BY id DESC LIMIT ? OFFSET ?", TransactionTable)
	rows, err := d.QueryContext(context.Background(), getAllQuery, pageSize, calcOffset(pageIndex, pageSize))
	if err != nil || rows.Err() != nil {
		return nil, 0, ErrCouldNotFindRecord
	}
	defer rows.Close()

	transactions := make([]Transaction, 0, pageSize)
	for rows.Next() {
		t := &Transaction{}
		err := rows.Scan(&t.ID, &t.TxID, &t.BlockHeight, &t.BlockTime, &t.PayloadType,
			&t.Data, &t.Description, &t.Amount, &t.Status, &t.CreatedAt)
		if err != nil {
			return nil, 0, ErrCouldNotFindRecord
		}

		transactions = append(transactions, *t)
	}

	return transactions, totalRecords, nil
}

func (d *db) GetTotalRecords(tableName string) (int64, error) {
	var totalRecords int64
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s", tableName)

	r := d.QueryRowContext(context.Background(), query)
	if r.Err() != nil {
		return totalRecords, ErrCouldNotFindTotalRecords
	}

	if err := r.Scan(&totalRecords); err != nil {
		return totalRecords, ErrCouldNotCreateTable
	}

	return totalRecords, nil
}

func calcOffset(pageIndex, pageSize int) int {
	return (pageIndex - 1) * pageSize
}