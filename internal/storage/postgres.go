package storage

import (
	"adsboard/internal/ads"
	"database/sql"
	_ "github.com/lib/pq"
)

type Postgres struct {
	db *sql.DB
}

func NewPostgres(dsn string) (*Postgres, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return &Postgres{db: db}, nil
}
func (p *Postgres) GetAllAd() ([]ads.Ad, error) {
	query := `SELECT title, description, price FROM ads`
	rows, err := p.db.Query(query)
	if err != nil {
		return nil, err
	}

	var adverts []ads.Ad

	for rows.Next() {
		var ad ads.Ad
		if err := rows.Scan(&ad.Title, &ad.Description, &ad.Price); err != nil {
			return nil, err
		}
		adverts = append(adverts, ad)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return adverts, nil
}

func (p *Postgres) CreateAd(title, description string, price float64) (int, error) {
	query := `INSERT INTO ads (title, description, price) VALUES ($1, $2, $3) RETURNING id`
	var id int
	if err := p.db.QueryRow(query, title, description, price).Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

// Ensure Postgres implements the Storage interface
var _ ads.Storage = (*Postgres)(nil)

// TODO: Implement other CRUD methods
