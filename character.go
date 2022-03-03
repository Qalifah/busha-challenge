package core

import "context"

type Character struct {
	ID       int    `json:"-"`
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	Height   string `json:"height"`
	MoviesID []int  `json:"-"`
}

type CharacterRepository interface {
	Get(ctx context.Context, movieID int, filterArg, sortArg, order *string) ([]*Character, error)
	AddMany(ctx context.Context, chars []*Character) error
}
