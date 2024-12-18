package storage

import (
	"adsboard/internal/ads"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
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

func (p *Postgres) CreateAd(title, description string, price float64) (int, error) {
	query := `INSERT INTO ads (title, description, price) VALUES ($1, $2, $3) RETURNING id`
	var id int
	if err := p.db.QueryRow(query, title, description, price).Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (p *Postgres) GetAd(id int) (*ads.Ad, error) {
	query := `SELECT id, title, description, price FROM ads WHERE id = $1`
	rows, err := p.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	if !rows.Next() {
		return nil, nil
	}
	var ad ads.Ad
	if err := rows.Scan(&ad.ID, &ad.Title, &ad.Description, &ad.Price); err != nil {
		return nil, err
	}
	return ad, nil
}

func (p *Postgres) UpdateAd(ad ads.Ad) (bool, error) {
	query := `
			UPDATE ads 
			SET title = $1, description = $2, price = $3
			WHERE id = $4
			`
	result, err := p.db.Exec(query, ad.Title, ad.Description, ad.Price, ad.ID)
	if err != nil {
		log.Println("Failed to update Ad ID=%v", ad.ID)
		return false, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	if rowsAffected == 0 {
		log.Println("No ad found with ID=%v", ad.ID)
		return false, nil
	}
	return true, nil
}

func (p *Postgres) GetAllAd() ([]ads.Ad, error) {
	query := `SELECT id, title, description, price FROM ads`
	rows, err := p.db.Query(query)
	if err != nil {
		return nil, err
	}

	var adverts []ads.Ad

	for rows.Next() {
		var ad ads.Ad
		if err := rows.Scan(&ad.ID, &ad.Title, &ad.Description, &ad.Price); err != nil {
			return nil, err
		}
		adverts = append(adverts, ad)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return adverts, nil
}

// Ensure Postgres implements the Storage interface
var _ ads.Storage = (*Postgres)(nil)

// TODO: Implement other CRUD methods
