package pg

import (
	"context"
	"database/sql"
	"testing"

	// "github.com/jackc/pgx/v5"
)

// func TestGetPrds(t *testing.T) {
// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()
// 	conn, err := pgx.Connect(ctx, "postgres://postgres:postgres@localhost:5432/task_5")
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	prdRepo := NewProductRepository(conn)

// 	_, _, _, err = prdRepo.GetProducts(ctx, 0, 1000, 0, 100, 0, 10, "", "test1+string")
// 	if err != nil {
// 		t.Error(err)
// 	}
// }

func TestDelPrd(t *testing.T) {
	conn, err := sql.Open("pgx", "postgres://postgres:postgres@localhost:5432/task_5")
	if err != nil {
		t.Error(err)
	}
	prdRepo := NewProductRepository(conn)

	_, err = prdRepo.DeleteProductById(context.TODO(), 1)
	if err != nil {
		t.Error(err)
	}
}
