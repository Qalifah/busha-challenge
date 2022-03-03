package postgres

import (
	"context"

	core "busha-challenge"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCharacterRepository_AddMany(t *testing.T) {
	data := []*core.Character{
		{Name: "Avatar", Gender: "male", Height: "7 ft", MoviesID: []int{1, 2, 3}},
		{Name: "Naruto", Gender: "female", Height: "96 cm", MoviesID: []int{4, 3, 6}},
	}

	ctx := context.Background()
	conn, err := createTestDB(ctx)
	require.NoError(t, err)
	defer conn.Close(ctx)

	charRepo := NewCharacterRepository(conn)
	require.NoError(t, charRepo.AddMany(ctx, data))
}

func TestCharacterRepository_Get(t *testing.T) {
	data := []*core.Character{
		{Name: "Avatar", Gender: "male", Height: "7 ft", MoviesID: []int{1, 2, 3}},
		{Name: "Naruto", Gender: "female", Height: "96 cm", MoviesID: []int{4, 3, 6}},
	}

	ctx := context.Background()
	conn, err := createTestDB(ctx)
	require.NoError(t, err)
	defer conn.Close(ctx)

	charRepo := NewCharacterRepository(conn)
	require.NoError(t, charRepo.AddMany(ctx, data))
	_, err = charRepo.Get(ctx, 3, nil, nil, nil)
	require.NoError(t, err)
}
