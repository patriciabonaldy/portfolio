package storage

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var db *sql.DB

func OpenDatabase() error {
	var err error
	db, err = sql.Open("sqlite3", "./sqlite-database.db")
	if err != nil {
		return err
	}

	return db.Ping()
}

func CreateTables() {
	createTableSQL := `CREATE TABLE IF NOT EXISTS stock (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"price" TEXT,
		"date_created" TEXT
	  );`
	statement, err := db.Prepare(createTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
	log.Println("Stock table created")

	createTableSQL = `CREATE TABLE IF NOT EXISTS stock_history (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"stock_id" INTEGER,
		"price" TEXT,
		"date_history" TEXT
	  );`
	statement, err = db.Prepare(createTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
	log.Println("portfolio_history table created")

	createTableSQL = `CREATE TABLE IF NOT EXISTS portfolio (
		"id" INTEGER NOT NULL,
		"stock_id" INTEGER,
		"date_created" TEXT
	  );`
	statement, err = db.Prepare(createTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
	log.Println("portfolio table created")
}

type Storage struct {
}


func (s Storage) AddStock(price string, date string) {
	insertNoteSQL := `INSERT INTO stock(price, date_created) VALUES (?, ?)`
	statement, err := db.Prepare(insertNoteSQL)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = statement.Exec(price, date)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Inserted stock successfully")
}

func (s Storage) UpdateStock(stockID int, price string, date string) {
	insertNoteSQL := `update stock set price=?, date_created=? where id=?`
	_, err := db.Exec(insertNoteSQL, price, date, stockID)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("stock updated successfully")
}

func (s Storage) AddStockHistory(stockID int, price string, date string) {
	insertNoteSQL := `INSERT INTO stock_history(stock_id, price, date_history) VALUES (?, ?, ?)`
	statement, err := db.Prepare(insertNoteSQL)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = statement.Exec(stockID, price, date)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Inserted stock successfully")
}

func (s Storage) AddPortfolio(stockID int, date string) {
	insertNoteSQL := `INSERT INTO portfolio(id, stock_id, date_created) VALUES (1, ?, ?)`
	statement, err := db.Prepare(insertNoteSQL)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = statement.Exec(stockID, date)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Inserted stock successfully")
}

func (s Storage) GetStock(stockID int) stockRepo{
	var price, dateCreated string
	row, err := db.Query("SELECT * FROM stock where id=?", stockID)
	if err != nil {
		log.Fatal(err)
	}

	defer row.Close()

	for row.Next() {
		row.Scan(&stockID, &price, &dateCreated)
	}
	return stockRepo{
		ID:          stockID,
		Price:       price,
		DateCreated: dateCreated,
	}
}

func (s Storage) GetPortfolio() potfolioRepo {
	portf:= potfolioRepo{}
	row, err := db.Query("SELECT * FROM portfolio where id=1")
	if err != nil {
		log.Fatal(err)
	}

	defer row.Close()

	for row.Next() {
		var stockID int
		var dateCreated string
		row.Scan(&stockID, &dateCreated)
		portf= potfolioRepo{
			ID:          1,
			stockID:     stockID,
			DateCreated: dateCreated,
		}
	}

	return portf
}