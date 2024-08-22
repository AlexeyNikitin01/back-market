package pg

import (
	// "context"
	"fmt"
	"database/sql"

	"task5/internal/config"

	_ "github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func DB(cfg *config.Config) (*sql.DB, error) {
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.Psql.User,
		cfg.Psql.Pass,
		cfg.Psql.Host,
		cfg.Psql.Port,
		cfg.Psql.Dbname,
		cfg.Psql.SSLmode,
	)
	db, err :=  sql.Open("pgx", url)
	// c, err := pgx.Connect(context.TODO(), url)
	if err != nil {
		return nil, err
	}
	
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("no connect to postgres") 
	}

	return db, nil
}
