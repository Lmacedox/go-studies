package main

import (
	"database/sql"
	"fmt"
	"uuid"

	_ "github.com/go-sql-driver/mysql"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	// Abrindo conexão com banco de dados
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		fmt.Println("Erro ao estabelecer conexão com banco")
		panic(err)
	}

	fmt.Println("Conexão com BD estabelecida com sucesso!")
	defer db.Close()

	product := NewProduct("Notebook Inserido 1", 5500.90)
	err = insertProduct(db, product)
	if err != nil {
		panic(err)
	}

	product.Name = "Nome atualizado catatau?"

	err = updateProduct(db, product)
	if err != nil {
		panic(err)
	}

	// p, err := selectProduct(db, product.ID)
	// if err != nil {
	// 	panic(err)
	// }

	allProducts, err := selectAllProducts(db)
	if err != nil {
		panic(err)
	}

	fmt.Println("--- Lista de Produtos ---")
	for _, p := range allProducts {
		fmt.Printf("ID: %d | Nome: %s | Preço: %.2f\n", p.ID, p.Name, p.Price)
	}

}

func insertProduct(db *sql.DB, product *Product) error {
	// Statement contra SQLInjection
	stmt, err := db.Prepare("insert into products(id, name, price) values(?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	// No Exec passamos os valores para ele substituir os valores do ? no statement
	_, err = stmt.Exec(product.ID, product.Name, product.Price)

	if err != nil {
		return err
	}

	//se não houver erros retornamos nil
	return nil
}

func updateProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("update products set name = ?, price = ? where id = ?")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}

	return nil
}

func selectProduct(db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("select id, name, price from products where id = ?")
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	var p Product

	// Para buscar uma linha usamos o QueryRow
	// E o Scan vai setar os valores
	err = stmt.QueryRow(id).Scan(&p.ID, &p.Name, &p.Price)

	if err != nil {
		return nil, err
	}

	return &p, nil

}

func selectAllProducts(db *sql.DB) ([]Product, error) {
	// Fazendo a chamda sem o prepare, porque nao recebemos parametros entao nao pode ocorrer sqlInjection
	rows, err := db.Query("select * from products")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var products []Product

	for rows.Next() {
		var p Product
		err = rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	return products, nil
}
