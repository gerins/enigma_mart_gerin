package category

import (
	"database/sql"
	"errors"
	"log"
	"strconv"
)

type CategoryRepo struct {
	db *sql.DB
}

type CategoryRepository interface {
	HandleGETAllCategory() (*[]Category, error)
	HandleGETCategory(id, status string) (*Category, error)
	HandlePOSTCategory(d Category) (*Category, error)
	HandleUPDATECategory(id string, data Category) (*Category, error)
	HandleDELETECategory(id string) (*Category, error)
}

func NewCategoryRepo(db *sql.DB) CategoryRepository {
	return CategoryRepo{db}
}

// HandleGETAllCategory for GET all data from Category
func (p CategoryRepo) HandleGETAllCategory() (*[]Category, error) {
	var d Category
	var AllCategory []Category

	result, err := p.db.Query("SELECT * FROM produk_category WHERE status=?", "A")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for result.Next() {
		err := result.Scan(&d.ID, &d.CategoryName, &d.Created, &d.Updated, &d.Status)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		resultProduct, _ := p.db.Query("SELECT nama FROM produk WHERE status=? AND produk_category_id=?", "A", d.ID)

		var productName string
		var allProductName []string
		for resultProduct.Next() {
			err := resultProduct.Scan(&productName)
			if err != nil {
				log.Println(err)
				return nil, err
			}
			allProductName = append(allProductName, productName)
		}
		d.Produk = allProductName
		AllCategory = append(AllCategory, d)
	}
	return &AllCategory, nil
}

// HandleGETCategory for GET single data from Category
func (p CategoryRepo) HandleGETCategory(id, status string) (*Category, error) {
	results := p.db.QueryRow("SELECT * FROM produk_category WHERE id=? AND status=?", id, status)

	var d Category
	err := results.Scan(&d.ID, &d.CategoryName, &d.Created, &d.Updated, &d.Status)
	if err != nil {
		return nil, errors.New("Category ID Not Found")
	}

	resultProduct, _ := p.db.Query("SELECT nama FROM produk WHERE status=? AND produk_category_id=?", "A", d.ID)

	for resultProduct.Next() {
		var productName string
		err := resultProduct.Scan(&productName)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		d.Produk = append(d.Produk, productName)
	}

	return &d, nil
}

// HandlePOSTCategory will POST a new Category data
func (p CategoryRepo) HandlePOSTCategory(d Category) (*Category, error) {
	tx, err := p.db.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	stmnt, _ := tx.Prepare(`INSERT INTO produk_category(nama) VALUES (?)`)
	defer stmnt.Close()

	result, err := stmnt.Exec(d.CategoryName)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}
	lastInsertID, _ := result.LastInsertId()
	tx.Commit()
	return p.HandleGETCategory(strconv.Itoa(int(lastInsertID)), "A")
}

// HandleUPDATECategory is used for UPDATE data Category
func (p CategoryRepo) HandleUPDATECategory(id string, data Category) (*Category, error) {
	tx, err := p.db.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	_, err = tx.Exec(`UPDATE produk_category SET nama=? WHERE id=?`, data.CategoryName, id)

	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	checkAvaibility, err := p.HandleGETCategory(id, "A")
	if err != nil {
		return nil, err
	}
	return checkAvaibility, nil
}

// HandleDELETECategory for DELETE single data from Category
func (p CategoryRepo) HandleDELETECategory(id string) (*Category, error) {
	tx, err := p.db.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	_, err = tx.Exec("UPDATE produk_category SET status=? WHERE id=?", "NA", id)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	checkAvaibility, err := p.HandleGETCategory(id, "NA")
	if err != nil {
		return nil, err
	}
	return checkAvaibility, nil
}
