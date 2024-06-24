package model

import (
	"context"
	"database/sql"
	"errors"

	database "flover-market/db"
)

/*
CREATE TABLE IF NOT EXISTS price_change
(
    product_id        INTEGER NOT NULL,
    date_price_change integer NOT NULL,
    new_price         REAL    NOT NULL,
    CONSTRAINT PK_PRICE_CHANGE PRIMARY KEY (product_id, date_price_change),
    FOREIGN KEY ([product_id]) REFERENCES "products" ([product_id]) ON DELETE NO ACTION ON UPDATE NO ACTION
);
*/

type PriceChange struct {
	ProductId       int
	DatePriceChange int
	NewPrice        int
}

func GetLastPriceChange(productId int) (PriceChange, error) {
	db := database.GetDB()

	var pc PriceChange
	err := db.QueryRow(context.Background(), "SELECT date_price_change, new_price FROM price_change WHERE product_id = $1 ORDER BY date_price_change DESC LIMIT 1", productId).Scan(&pc.DatePriceChange, &pc.NewPrice)
	if err != nil {
		return pc, err
	}
	if errors.Is(err, sql.ErrNoRows) {
		return pc, errors.New("no price changes")
	}
	return pc, nil
}

func AddPriceChange(priceChange PriceChange) error {
	db := database.GetDB()

	_, err := db.Exec(context.Background(), "INSERT INTO price_change (product_id, date_price_change, new_price) VALUES ($1, $2, $3)", priceChange.ProductId, priceChange.DatePriceChange, priceChange.NewPrice)
	if err != nil {
		return err
	}

	return nil
}

func DeletePriceChange(productId int, datePriceChange int) {
	// TODO
}

func GetPriceChanges(productId int) []PriceChange {
	// TODO
	return nil
}
