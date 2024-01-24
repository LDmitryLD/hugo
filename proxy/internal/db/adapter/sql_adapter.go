package adapter

import (
	"fmt"
	"log"
	"projects/LDmitryLD/hugoproxy/proxy/internal/models"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type SQLAdapter struct {
	db         *sqlx.DB
	sqlBuilder sq.StatementBuilderType
}

func NewSQLAdapter(db *sqlx.DB) *SQLAdapter {
	builder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	return &SQLAdapter{
		db:         db,
		sqlBuilder: builder,
	}
}

func (s *SQLAdapter) Select(query string) (models.Address, error) {

	rows, err := s.db.Query(`
		SELECT address.lat, address.lon
		FROM history_search_address
		JOIN search_history ON search_history.id = history_search_address.search_id
		JOIN address ON history_search_address.address_id = address.id
		WHERE levenshtein(search_history.query, $1) <= LENGHT($1) * 0.3;
		`, query)
	if err != nil {
		return models.Address{}, err
	}
	defer rows.Close()

	if !rows.Next() {
		log.Println("совпадений нет")
		return models.Address{}, fmt.Errorf("совпадений нет")
	}

	var address models.Address

	if err = rows.Scan(&address.Lat, &address.Lon); err != nil {
		return models.Address{}, err
	}
	log.Printf("Lat %s Lon %s", address.Lat, address.Lon)
	return address, nil

	// for rows.Next() {
	// 	err := rows.Scan(&address.Lat, &address.Lon)
	// 	if err != nil {
	// 		return models.Address{}, err
	// 	}
	// 	log.Printf("Lat: %d, Lon: %d", address.Lat, address.Lon)
	// }

	//return address, nil
}

func (s *SQLAdapter) Insert(query string, lat string, lon string) error {
	_, err := s.db.Exec(`INSERT INTO search_history (query) VALUES ($1)`, query)
	if err != nil {
		return err
	}

	_, err = s.db.Exec(`INSERT INTO address (lat, lon) VALUES ($1, $2)`, lat, lon)
	if err != nil {
		return err
	}

	// var searchID int
	// err = s.db.QueryRow(`SELECT id FROM search_history ORDER BY id DESC LIMIT 1`).Scan(&searchID)
	// if err != nil {
	// 	return err
	// }

	// var addressID int
	// err = s.db.QueryRow(`SELECT id FROM address ORDER BY id DESC LIMIT 1`).Scan(&addressID)
	// if err != nil {
	// 	return err
	// }

	// _, err = s.db.Exec(`INSERT INTO history_search_address (search_id, address_id) VALUES ($1, $2)`, searchID, addressID)
	// if err != nil {
	// 	return err
	// }

	return nil

}
