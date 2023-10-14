package portfolio

import (
	"github.com/PedroXimenes/4invest/internal/pkg/db"
)

func Get(id int64) (portfolios []Portfolio, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT portfolio_id, user_id, name FROM portfolio WHERE user_id=$1`, id)
	if err != nil {
		return
	}

	for rows.Next() {
		var portfolio Portfolio
		err = rows.Scan(&portfolio.ID, &portfolio.UserID, &portfolio.Name)
		if err != nil {
			continue
		}
		portfolios = append(portfolios, portfolio)
	}

	return
}
