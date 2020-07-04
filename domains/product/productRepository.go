package product

import (
	"database/sql"
	"errors"
	"log"
	"strconv"
)

type ProductRepo struct {
	db *sql.DB
}

type ProductRepository interface {
	HandleGETAllProduct() (*[]Product, error)
	HandleGETProduct(id, status string) (*Product, error)
	HandlePOSTProduct(d Product) (*Product, error)
	HandleUPDATEProduct(id string, data Product) (*Product, error)
	HandleDELETEProduct(id string) (*Product, error)
}

func NewProductRepo(db *sql.DB) ProductRepository {
	return ProductRepo{db}
}

// HandleGETAllProduct for GET all data from Product
func (p ProductRepo) HandleGETAllProduct() (*[]Product, error) {
	var d Product
	var AllProduct []Product

	result, err := p.db.Query("SELECT * FROM produk_idx WHERE status=?", "A")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for result.Next() {
		err := result.Scan(&d.ID, &d.ProductName, &d.CatID, &d.CategoryName, &d.Harga,
			&d.Created, &d.Updated, &d.Status)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		AllProduct = append(AllProduct, d)
	}
	return &AllProduct, nil
}

// HandleGETProduct for GET single data from Product
func (p ProductRepo) HandleGETProduct(id, status string) (*Product, error) {
	results := p.db.QueryRow("SELECT * FROM produk_idx WHERE id=? AND status=?", id, status)

	var d Product
	err := results.Scan(&d.ID, &d.ProductName, &d.CatID, &d.CategoryName, &d.Harga,
		&d.Created, &d.Updated, &d.Status)
	if err != nil {
		return nil, errors.New("Category ID Not Found")
	}

	return &d, nil
}

// HandlePOSTProduct will POST a new Product data
func (p ProductRepo) HandlePOSTProduct(d Product) (*Product, error) {
	tx, err := p.db.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	stmnt, _ := tx.Prepare(`INSERT INTO produk(produk_category_id,nama,harga) VALUES (?,?,?)`)
	defer stmnt.Close()

	result, err := stmnt.Exec(d.CatID, d.ProductName, d.Harga)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}
	lastInsertID, _ := result.LastInsertId()
	tx.Commit()
	return p.HandleGETProduct(strconv.Itoa(int(lastInsertID)), "A")
}

// HandleUPDATEProduct is used for UPDATE data Product
func (p ProductRepo) HandleUPDATEProduct(id string, data Product) (*Product, error) {
	tx, err := p.db.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	_, err = tx.Exec(`UPDATE produk SET produk_category_id=?,nama=?,harga=? WHERE id=?`,
		data.CatID, data.ProductName, data.Harga, id)

	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	checkAvaibility, err := p.HandleGETProduct(id, "A")
	if err != nil {
		return nil, err
	}
	return checkAvaibility, nil
}

// HandleDELETEProduct for DELETE single data from Product
func (p ProductRepo) HandleDELETEProduct(id string) (*Product, error) {
	tx, err := p.db.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	_, err = tx.Exec("UPDATE produk SET status=? WHERE id=?", "NA", id)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	checkAvaibility, err := p.HandleGETProduct(id, "NA")
	if err != nil {
		return nil, err
	}
	return checkAvaibility, nil
}
