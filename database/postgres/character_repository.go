package postgres

import (
	core "busha-challenge"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
)

type characterRepository struct {
	client *pgx.Conn
}

func NewCharacterRepository(client *pgx.Conn) *characterRepository {
	return &characterRepository{client: client}
}

func (r *characterRepository) Get(ctx context.Context, movieID int, filterArg, sortArg, order *string) ([]*core.Character, error) {
	var characters []*core.Character

	query := fmt.Sprintf("SELECT * FROM characters WHERE %v IN movies_id", movieID)
	if filterArg != nil {
		query += fmt.Sprintf(" AND gender = %v", *filterArg)
	}
	if sortArg != nil {
		query += fmt.Sprintf("ORDER BY %v %v", *sortArg, *order)
	}

	rows, err := r.client.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		char := &core.Character{}
		rows.Scan(char.ID, char.Name, char.Gender, char.Height, char.MoviesID)
		characters = append(characters, char)
	}

	if rows.Err() != nil {
		return nil, err
	}
	return characters, nil
}

func (r *characterRepository) add(ctx context.Context, char *core.Character) error {
	_, err := r.client.Exec(ctx, "INSERT INTO characters (name, gender, height, movies_id) VALUES ($1, $2, $3, $4)", char.Name, char.Gender, char.Height, char.MoviesID)
	return err
}
