package transaction

import (
	"database/sql"
	"enigma_mart_gerin/domains/product"
	"errors"
	"log"
	"strconv"
)

type TransactionRepo struct {
	db *sql.DB
}

type TransactionRepository interface {
	HandleGETAllTransaction() (*[]Transaction, error)
	HandleGETTransaction(id, status string) (*Transaction, error)
	HandlePOSTTransaction(d Transaction) (*Transaction, error)
	HandleUPDATETransaction(id string, data Transaction) (*Transaction, error)
	HandleDELETETransaction(id string) (*Transaction, error)
	HandleGETAllTransactionMontly(month string) (*[]Transaction, error)
	HandleGETAllTransactionDaily(daily string) (*[]Transaction, error)
}

func NewTransactionRepo(db *sql.DB) TransactionRepository {
	return TransactionRepo{db}
}

// HandleGETAllTransaction for GET all data from Transaction
func (p TransactionRepo) HandleGETAllTransaction() (*[]Transaction, error) {
	var d Transaction
	var AllTransaction []Transaction

	result, err := p.db.Query("SELECT * FROM transaksi_penjualan WHERE status=?", "A")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for result.Next() {
		err := result.Scan(&d.ID, &d.Created, &d.Updated, &d.Total, &d.Status)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		resultProduct, _ := p.db.Query(`SELECT nama,kategori,harga,kuantiti,total FROM transaksi_produk_idx WHERE trans_id=?`, d.ID)

		var soldItem SoldItems
		var allSoldItem []SoldItems
		for resultProduct.Next() {
			err := resultProduct.Scan(&soldItem.ProductName, &soldItem.Category, &soldItem.Harga, &soldItem.Quantity, &soldItem.Total)
			if err != nil {
				log.Println(err)
				return nil, err
			}
			allSoldItem = append(allSoldItem, soldItem)
		}
		d.SoldItems = allSoldItem
		AllTransaction = append(AllTransaction, d)
	}
	return &AllTransaction, nil
}

// HandleGETTransaction for GET single data from Transaction
func (p TransactionRepo) HandleGETTransaction(id, status string) (*Transaction, error) {
	results := p.db.QueryRow("SELECT * FROM transaksi_penjualan WHERE id=? AND status=?", id, status)

	var d Transaction
	err := results.Scan(&d.ID, &d.Created, &d.Updated, &d.Total, &d.Status)
	if err != nil {
		return nil, errors.New("Transaction ID Not Found")
	}

	resultProduct, _ := p.db.Query(`SELECT nama,kategori,harga,kuantiti,total FROM transaksi_produk_idx WHERE trans_id=?`, d.ID)
	var soldItem SoldItems

	for resultProduct.Next() {
		err := resultProduct.Scan(&soldItem.ProductName, &soldItem.Category, &soldItem.Harga, &soldItem.Quantity, &soldItem.Total)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		d.SoldItems = append(d.SoldItems, soldItem)
	}

	return &d, nil
}

// HandlePOSTTransaction will POST a new Transaction data
func (p TransactionRepo) HandlePOSTTransaction(d Transaction) (*Transaction, error) {
	tx, err := p.db.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	result, err := tx.Exec("INSERT INTO transaksi_penjualan(total_penjualan) VALUE(?)", "0")
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}
	lastInsertID, _ := result.LastInsertId()

	var Total int
	for _, value := range d.SoldItems {
		product, _ := product.NewProductRepo(p.db).HandleGETProduct(value.ProductName, "A")
		hargaProduct, _ := strconv.Atoi(product.Harga)
		quantitiProduct, _ := strconv.Atoi(value.Quantity)
		_, err = tx.Exec(`INSERT INTO transaksi_produk(transaksi_penjualan_id,produk_id,kuantiti,total) VALUE(?,?,?,?)`, lastInsertID, value.ProductName, value.Quantity, hargaProduct*quantitiProduct)
		if err != nil {
			log.Println(err)
			tx.Rollback()
			return nil, err
		}
		Total += hargaProduct * quantitiProduct
	}

	_, err = tx.Exec(`UPDATE transaksi_penjualan SET total_penjualan=? WHERE id=?`, Total, lastInsertID)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return p.HandleGETTransaction(strconv.Itoa(int(lastInsertID)), "A")
}

// HandleUPDATETransaction is used for UPDATE data Transaction
func (p TransactionRepo) HandleUPDATETransaction(id string, data Transaction) (*Transaction, error) {
	tx, err := p.db.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var Total int
	for _, value := range data.SoldItems {
		product, _ := product.NewProductRepo(p.db).HandleGETProduct(value.ProductName, "A")
		hargaProduct, _ := strconv.Atoi(product.Harga)
		quantitiProduct, _ := strconv.Atoi(value.Quantity)
		_, err = tx.Exec(`UPDATE transaksi_produk SET  kuantiti=?, total=? WHERE transaksi_penjualan_id=? AND produk_id=?`, value.Quantity, hargaProduct*quantitiProduct, id, value.ProductName)
		if err != nil {
			log.Println(err)
			tx.Rollback()
			return nil, err
		}
		Total += hargaProduct * quantitiProduct
	}

	_, err = tx.Exec(`UPDATE transaksi_penjualan SET total_penjualan=? WHERE id=?`, Total, id)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	checkAvaibility, err := p.HandleGETTransaction(id, "A")
	if err != nil {
		return nil, err
	}
	return checkAvaibility, nil
}

// HandleDELETETransaction for DELETE single data from Transaction
func (p TransactionRepo) HandleDELETETransaction(id string) (*Transaction, error) {
	tx, err := p.db.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	_, err = tx.Exec("UPDATE transaksi_penjualan SET status=? WHERE id=?", "NA", id)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	checkAvaibility, err := p.HandleGETTransaction(id, "NA")
	if err != nil {
		return nil, err
	}
	return checkAvaibility, nil
}

// HandleGETAllTransaction for GET all data this month from Transaction
func (p TransactionRepo) HandleGETAllTransactionMontly(month string) (*[]Transaction, error) {
	var d Transaction
	var AllTransaction []Transaction

	result, err := p.db.Query("SELECT * FROM transaksi_penjualan WHERE status=? AND month(created_at)=?", "A", month)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for result.Next() {
		err := result.Scan(&d.ID, &d.Created, &d.Updated, &d.Total, &d.Status)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		resultProduct, _ := p.db.Query(`SELECT nama,kategori,harga,kuantiti,total FROM transaksi_produk_idx WHERE trans_id=?`, d.ID)

		var soldItem SoldItems
		var allSoldItem []SoldItems
		for resultProduct.Next() {
			err := resultProduct.Scan(&soldItem.ProductName, &soldItem.Category, &soldItem.Harga, &soldItem.Quantity, &soldItem.Total)
			if err != nil {
				log.Println(err)
				return nil, err
			}
			allSoldItem = append(allSoldItem, soldItem)
		}
		d.SoldItems = allSoldItem
		AllTransaction = append(AllTransaction, d)
	}
	if len(AllTransaction) < 1 {
		return nil, errors.New("Transaction Not Found")
	}
	return &AllTransaction, nil
}

// HandleGETAllTransaction for GET all data this daily from Transaction
func (p TransactionRepo) HandleGETAllTransactionDaily(daily string) (*[]Transaction, error) {
	var d Transaction
	var AllTransaction []Transaction

	result, err := p.db.Query("SELECT * FROM transaksi_penjualan WHERE status=? AND day(created_at)=?", "A", daily)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for result.Next() {
		err := result.Scan(&d.ID, &d.Created, &d.Updated, &d.Total, &d.Status)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		resultProduct, _ := p.db.Query(`SELECT nama,kategori,harga,kuantiti,total FROM transaksi_produk_idx WHERE trans_id=?`, d.ID)

		var soldItem SoldItems
		var allSoldItem []SoldItems
		for resultProduct.Next() {
			err := resultProduct.Scan(&soldItem.ProductName, &soldItem.Category, &soldItem.Harga, &soldItem.Quantity, &soldItem.Total)
			if err != nil {
				log.Println(err)
				return nil, err
			}
			allSoldItem = append(allSoldItem, soldItem)
		}
		d.SoldItems = allSoldItem
		AllTransaction = append(AllTransaction, d)
	}
	if len(AllTransaction) < 1 {
		return nil, errors.New("Transaction Not Found")
	}
	return &AllTransaction, nil
}
