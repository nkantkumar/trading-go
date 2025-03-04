package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"trading-go/internal/order"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) Save(o *order.Order) error {
	query := `INSERT INTO orders (id, user_id, symbol, quantity, price, type, timestamp) 
              VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.db.Exec(query, o.ID, o.UserID, o.Symbol, o.Quantity, o.Price, o.Type, o.Timestamp)
	return err
}

func (r *PostgresRepository) GetAll() ([]order.Order, error) {
	rows, err := r.db.Query(`SELECT id, user_id, symbol, quantity, price, type, timestamp FROM orders`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []order.Order
	for rows.Next() {
		var o order.Order
		err := rows.Scan(&o.ID, &o.UserID, &o.Symbol, &o.Quantity, &o.Price, &o.Type, &o.Timestamp)
		if err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}
	return orders, nil
}
