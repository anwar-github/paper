package repository

import (
	"database/sql"
	"errors"
	"papper/database/mysql"
	"papper/model"
)

type Repository struct {
	DB *sql.DB
}

type Interface interface {
	Transaction() (*sql.Tx, error)
	Rollback(tx *sql.Tx) error
	Commit(tx *sql.Tx) error

	FindWalletForUpdateByUserCode(tx *sql.Tx, userCode string) (*model.Wallet, error)
	UpdateWalletBalanceByUserCode(tx *sql.Tx, balance int64, userCode string) error

	FindDisbursementByRequestID(tx *sql.Tx, requestID string) (*model.Disbursement, error)
	FindDisbursementByOrderID(tx *sql.Tx, orderID string) (*model.Disbursement, error)
	InsertDisbursement(tx *sql.Tx, disbursement *model.Disbursement) (*model.Disbursement, error)
}

func NewRepository(mysql *mysql.Mysql) Interface {
	return Repository{
		DB: mysql.DB,
	}
}

func (repository Repository) Transaction() (*sql.Tx, error) {
	tx, err := repository.DB.Begin()
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (repository Repository) Rollback(tx *sql.Tx) error {
	err := tx.Rollback()
	if err != nil {
		return err
	}

	return nil
}

func (repository Repository) Commit(tx *sql.Tx) error {
	err := tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (repository Repository) FindWalletForUpdateByUserCode(tx *sql.Tx, userCode string) (*model.Wallet, error) {
	var wallet model.Wallet
	row := tx.QueryRow(`SELECT balance, created_at, updated_at FROM wallet WHERE user_code = ? FOR UPDATE`, userCode)
	if err := row.Scan(&wallet.Balance, &wallet.CreatedAt, &wallet.UpdatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &wallet, nil
}

func (repository Repository) UpdateWalletBalanceByUserCode(tx *sql.Tx, balance int64, userCode string) error {
	if _, err := tx.Exec(
		`UPDATE wallet SET balance = ?, updated_at = CURRENT_TIMESTAMP WHERE user_code = ?`, balance, userCode); err != nil {
		return err
	}
	return nil
}

func (repository Repository) FindDisbursementByRequestID(tx *sql.Tx, requestID string) (*model.Disbursement, error) {
	var disbursement model.Disbursement
	row := tx.QueryRow(`SELECT request_id FROM disbursement WHERE request_id = ?`, requestID)

	if err := row.Scan(&disbursement.Request); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &disbursement, nil
}

func (repository Repository) FindDisbursementByOrderID(tx *sql.Tx, orderID string) (*model.Disbursement, error) {
	var disbursement model.Disbursement
	row := tx.QueryRow(`SELECT request_id, order_id, success, created_at FROM disbursement WHERE request_id = ?`, orderID)

	if err := row.Scan(&disbursement.Request, &disbursement.Order, &disbursement.Success, &disbursement.CreatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &disbursement, nil
}

func (repository Repository) InsertDisbursement(tx *sql.Tx, disbursement *model.Disbursement) (*model.Disbursement, error) {
	if _, err := tx.Query(
		`INSERT INTO disbursement (
                          order_id, 
                          request_id, 
                          user_code, 
                          provider,
                          account_number,  
                          success, 
                          amount, 
                          balance, 
                          created_at, 
                          updated_at)
		VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		disbursement.Order,
		disbursement.Request,
		disbursement.UserCode,
		disbursement.Provider,
		disbursement.AccountNumber,
		disbursement.Success,
		disbursement.Amount,
		disbursement.Balance,
		disbursement.CreatedAt,
		disbursement.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return disbursement, nil
}
