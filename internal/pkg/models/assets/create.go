package assets

import (
	"database/sql"

	"github.com/PedroXimenes/4invest/internal/pkg/db"
)

func Insert(assets []*Asset) ([]string, error) {
	errorAssets := []string{}
	conn, err := db.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	query := `INSERT INTO assets (portfolio_id, type, symbol, quantity, max_value_limit, ideal_percentage) VALUES ($1, $2, $3, $4, $5, $6)`

	for _, asset := range assets {
		err = conn.QueryRow(query, asset.PortfolioID, asset.Type, asset.Symbol, asset.Quantity, asset.MaxValueLimit, asset.IdealPercentage).Scan()
		if err != sql.ErrNoRows {
			if err.Error() == `ERROR: duplicate key value violates unique constraint "unique_portfolio_symbol" (SQLSTATE 23505)` {
				errorAssets = append(errorAssets, asset.Symbol)
				continue
			}
			return nil, err
		}
	}

	return errorAssets, nil
}
