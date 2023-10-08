package assets

import (
	"database/sql"
	"fmt"

	"github.com/PedroXimenes/4invest/internal/pkg/db"
)

func Insert(assets []*Asset) error {
	conn, err := db.OpenConnection()
	if err != nil {
		fmt.Printf("DB Connection error: %s\n", err)
		return err
	}
	defer conn.Close()

	query := `INSERT INTO assets (portfolio_id, type, symbol, quantity, max_value_limit, ideal_percentage) VALUES ($1, $2, $3, $4, $5, $6)`

	for _, asset := range assets {
		fmt.Printf("%+v\n", asset)
		err = conn.QueryRow(query, asset.PortfolioID, asset.Type, asset.Symbol, asset.Quantity, asset.MaxValueLimit, asset.IdealPercentage).Scan()
		if err != sql.ErrNoRows {
			fmt.Println("entrou aqui")
			return err
		}

	}

	return nil
}

// gcloud sql instances describe sincere-almanac-401411:us-central1:four-invest --format='value(connectionName)'

// ./../../Downloads/cloud-sql-proxy --port 6000 sincere-almanac-401411:us-central1:four-invest
//./../../Downloads/cloud-sql-proxy --port=9000 --run-connection-test sincere-almanac-401411:us-central1:four-invest
