package portfolio

import (
	"github.com/PedroXimenes/4invest/internal/pkg/db"
)

func Insert(portfolio *Portfolio) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `INSERT INTO portfolio (name, user_id) VALUES ($1, $2) RETURNING portfolio_id`

	err = conn.QueryRow(sql, portfolio.Name, portfolio.UserID).Scan(&id)
	if err != nil {
		return
	}

	return
}
