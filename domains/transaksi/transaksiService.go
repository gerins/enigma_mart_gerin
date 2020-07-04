package transaction

import (
	"database/sql"
)

type TransactionService struct {
	TransactionRepo TransactionRepository
}

type TransactionServiceInterface interface {
	GetTransactions() (*[]Transaction, error)
	GetTransactionByID(id string) (*Transaction, error)
	HandlePOSTTransaction(d Transaction) (*Transaction, error)
	HandleUPDATETransaction(id string, data Transaction) (*Transaction, error)
	HandleDELETETransaction(id string) (*Transaction, error)
	GetTransactionsMontly(month string) (*[]Transaction, error)
	GetTransactionsDaily(daily string) (*[]Transaction, error)
}

func NewTransactionService(db *sql.DB) TransactionServiceInterface {
	return TransactionService{NewTransactionRepo(db)}
}

func (s TransactionService) GetTransactions() (*[]Transaction, error) {
	Transaction, err := s.TransactionRepo.HandleGETAllTransaction()
	if err != nil {
		return nil, err
	}

	return Transaction, nil
}

func (s TransactionService) GetTransactionByID(id string) (*Transaction, error) {
	Transaction, err := s.TransactionRepo.HandleGETTransaction(id, "A")
	if err != nil {
		return nil, err
	}
	return Transaction, nil
}

func (s TransactionService) HandlePOSTTransaction(d Transaction) (*Transaction, error) {
	result, err := s.TransactionRepo.HandlePOSTTransaction(d)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s TransactionService) HandleUPDATETransaction(id string, data Transaction) (*Transaction, error) {
	result, err := s.TransactionRepo.HandleUPDATETransaction(id, data)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s TransactionService) HandleDELETETransaction(id string) (*Transaction, error) {
	result, err := s.TransactionRepo.HandleDELETETransaction(id)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (s TransactionService) GetTransactionsMontly(month string) (*[]Transaction, error) {
	Transaction, err := s.TransactionRepo.HandleGETAllTransactionMontly(month)
	if err != nil {
		return nil, err
	}

	return Transaction, nil
}

func (s TransactionService) GetTransactionsDaily(daily string) (*[]Transaction, error) {
	Transaction, err := s.TransactionRepo.HandleGETAllTransactionDaily(daily)
	if err != nil {
		return nil, err
	}

	return Transaction, nil
}
