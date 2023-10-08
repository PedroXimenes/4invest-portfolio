package assets

import "github.com/PedroXimenes/4invest/internal/pkg/db"

func Get(id int64) (assets []Asset, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT portfolio_id, type, symbol, quantity, ideal_percentage FROM assets WHERE portfolio_id=$1`, id)
	if err != nil {
		return
	}

	for rows.Next() {
		var asset Asset
		err = rows.Scan(&asset.PortfolioID, &asset.Type, &asset.Symbol, &asset.Quantity, &asset.IdealPercentage)
		if err != nil {
			continue
		}
		assets = append(assets, asset)
	}
	return
}
