package repository

import (
	"context"
	"database/sql"
	"server-go/models"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type model struct {
}

type Interface interface {
	Save(ctx context.Context, data models.CurrencyRate) error
	GetCotes() ([]models.CurrencyRate, error)
}

func New() *model {
	return &model{}
}

func (tx model) Save(ctx context.Context, data models.CurrencyRate) error {
	ctxCancel, cancel := context.WithTimeout(ctx, 10*time.Millisecond)
	defer cancel()

	db, err := startConnection()
	if err != nil {
		return nil
	}
	defer db.Close()

	err = registrarCotacao(ctxCancel, db, data)
	if err != nil {
		return nil
	}

	return nil
}

func registrarCotacao(ctx context.Context, db *sql.DB, cotacao models.CurrencyRate) error {
	_, err := db.ExecContext(
		ctx, "INSERT INTO CURRENCY (code, codein, name, high, low, varBid, pctChange, bid, ask, timestamp, create_date) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		cotacao.Code, cotacao.Codein, cotacao.Name, cotacao.High, cotacao.Low, cotacao.VarBid, cotacao.PctChange, cotacao.Bid, cotacao.Ask, cotacao.Timestamp, cotacao.CreateDate)
	if err != nil {
		return err
	}
	return nil
}

func (tx model) GetCotes() ([]models.CurrencyRate, error) {
	var crs []models.CurrencyRate

	db, err := startConnection()
	if err != nil {
		return crs, nil
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM CURRENCY")
	if err != nil {
		return crs, err
	}
	defer rows.Close()

	var currencyRates []models.CurrencyRate

	for rows.Next() {
		var cr models.CurrencyRate
		err := rows.Scan(&cr.Code, &cr.Codein, &cr.Name, &cr.High, &cr.Low, &cr.VarBid, &cr.PctChange, &cr.Bid, &cr.Ask, &cr.Timestamp, &cr.CreateDate)
		if err != nil {
			return crs, err
		}
		currencyRates = append(currencyRates, cr)
	}

	return currencyRates, nil
}

func createTable(db *sql.DB) error {
	_, err := db.Exec(`
CREATE TABLE IF NOT EXISTS CURRENCY (
    code TEXT,
    codein TEXT,
    name TEXT,
    high TEXT,
    low TEXT,
    varBid TEXT,
    pctChange TEXT,
    bid TEXT,
    ask TEXT,
    timestamp TEXT,
    create_date TEXT
);
`)
	if err != nil {
		return err
	}
	return nil
}

func startConnection() (*sql.DB, error) {
	var db *sql.DB
	db, err := sql.Open("sqlite3", "/home/vinicius/Documents/Study/pos-grad/projects/client-server-api/server/app/repository/db.sqlite")
	if err != nil {
		return db, err
	}
	err = createTable(db)
	if err != nil {
		return db, err
	}
	return db, nil
}
