package model

import (
	"database/sql"
	"fmt"
	"gomicroservice/config"
	"log"
)

type Product struct {
	ID 		int 	`json:"id"`
	SKU 	string 	`json:"sku"`
	Name 	string `json:"name_product"`
	Description string `json:"description"`
	Price   float32 `json:"price"`
}

const (
	table          = "product"
	//layoutDatetime = "2006-01-02 15:04:05"
)

type ListProducts []Product

func ShowProduct(id int64) (ListProducts, error)  {
	var listProduct ListProducts
	var product Product

	db, err := config.MYSQL()
	if err != nil {
		fmt.Printf("%v \n", err)
	}
	defer db.Close()

	sqlSatement := fmt.Sprintf("SELECT id, sku, name_product, description, price FROM %v WHERE id=?", table)
	rowQuery, err := db.Query(sqlSatement, id)
	if err != nil {
		fmt.Printf("%v \n", err)
	}
	defer rowQuery.Close()

	for rowQuery.Next() {
		switch err := rowQuery.Scan(&product.ID, &product.SKU,
			&product.Name, &product.Description, &product.Price); err {
		case sql.ErrNoRows:
			fmt.Printf("No rows returned \n", err)
		case nil:
			listProduct = append(listProduct, product)
		default:
			return nil, err
		}
	}
	return listProduct, nil
}

func GetProductAll() (ListProducts, error)  {
	var listProduct ListProducts
	var product Product

	db, err := config.MYSQL()
	if err != nil {
		fmt.Printf("%v \n", err)
	}
	defer db.Close()

	sqlSatement := fmt.Sprintf("SELECT id, sku, name_product, description, price FROM %v", table)
	rowQuery, err := db.Query(sqlSatement)
	if err != nil {
		fmt.Printf("%v \n", err)
	}
	defer rowQuery.Close()

	for rowQuery.Next() {
		switch err := rowQuery.Scan(&product.ID, &product.SKU,
			&product.Name, &product.Description, &product.Price); err {
		case sql.ErrNoRows:
			fmt.Printf("Error : %v \n", err)
		case nil:
			listProduct = append(listProduct, product)
		default:
			return nil, err
		}
	}
	return listProduct, nil
}

func PostProduct(product Product) error {
	db, err := config.MYSQL()
	if err != nil {
		log.Fatal("Error database connection", err)
	}
	defer db.Close()

	tx, err := db.Begin()
		if err != nil {
			log.Fatal("Error database transaction", err)
		}

	sqlStatement := fmt.Sprintf("INSERT INTO %v (sku, name_product, description, price) " +
		"VALUES('%v','%v','%v',%v)", table,
		product.SKU,
		product.Name,
		product.Description,
		product.Price)

	_, err = tx.Exec(sqlStatement)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Fatalf("update drivers: unable to rollback: %v", rollbackErr)
		}
		return err
	}

	_ = tx.Commit()

	return nil
}

func UpdateProduct(product Product) error {

	db, err := config.MYSQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	defer db.Close()

	tx, err := db.Begin()
		if err != nil {
			log.Fatal("Error database transaction", err)
		}

	sqlStatement := fmt.Sprintf("UPDATE %v set sku = '%s', name_product ='%s', description = '%s', price = %v where id = %d",
		table,
		product.SKU,
		product.Name,
		product.Description,
		product.Price,
		product.ID,
	)

	id := product.ID

	switch id {
		case 0:
			log.Printf("Wrong ID : %v \n", id)
			return nil
	default:
			_, err = tx.Exec(sqlStatement)
			if err != nil {
				if rollbackErr := tx.Rollback(); rollbackErr != nil {
					log.Fatalf("update drivers: unable to rollback: %v", rollbackErr)
				}
				return err
			}
			_ = tx.Commit()
			return nil
	}

}