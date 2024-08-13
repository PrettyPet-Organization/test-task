package products

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Product struct {
	Id      int     `json:"id"`
	Name    string  `json:"name"`
	Price   float64 `json:"price"`
	Country string  `json:"country"`
	Count   int     `json:"count"`
}

type ProductStore struct {
	db *sql.DB
}

func NewStore(maxOpenConn, maxIdleConn, connMaxIdleTime int) (*ProductStore, error) {
	db, err := sql.Open("sqlite3", "./store.db")
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenConn)
	db.SetMaxIdleConns(maxIdleConn)
	db.SetConnMaxIdleTime(time.Second * time.Duration(connMaxIdleTime))

	createDB := `CREATE TABLE IF NOT EXISTS products (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		price REAL,
		country TEXT,
		count INTEGER
	)`
	_, err = db.Exec(createDB)
	if err != nil {
		return nil, err
	}

	return &ProductStore{db}, nil
}

func (ps *ProductStore) Close() error {
	return ps.db.Close()
}

func (ps *ProductStore) AddProduct(name, country string, price float64, count int) (int, error) {
	res, err := ps.db.Exec("INSERT INTO products (name, country, price, count) VALUES (?,?,?,?)", name, country, price, count)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (ps *ProductStore) GetProduct(id int) (Product, error) {
	row := ps.db.QueryRow("SELECT id, name, price, country, count FROM products WHERE id = ?", id)

	var temp Product
	err := row.Scan(&temp.Id, &temp.Name, &temp.Price, &temp.Country, &temp.Count)
	if err == sql.ErrNoRows {
		return Product{}, fmt.Errorf("not found id:%d", id)
	} else if err != nil {
		return Product{}, err
	}

	return temp, nil
}

func (ps *ProductStore) GetAllProducts() ([]Product, error) {
	rows, err := ps.db.Query("SELECT id, name, price, country, count FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var temp Product

		err := rows.Scan(&temp.Id, &temp.Name, &temp.Price, &temp.Country, &temp.Count)
		if err != nil {
			return nil, err
		}

		products = append(products, temp)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (ps *ProductStore) DeleteProduct(id int) error {
	res, err := ps.db.Exec("DELETE FROM products WHERE id = ?", id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (ps *ProductStore) DeleteAllProducts() error {
	_, err := ps.db.Exec("DELETE FROM products")
	if err != nil {
		return err
	}
	return nil
}

func (ps *ProductStore) ChangeProduct(id int, name, country string, price float64, count int) error {
	res, err := ps.db.Exec("UPDATE products SET name = ?, price = ?, country = ?, count = ? WHERE id = ?", name, price, country, count, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (ps *ProductStore) BuyProduct(id, count int) error {
	res, err := ps.db.Exec("UPDATE products SET count = count - ? WHERE id = ?", count, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
