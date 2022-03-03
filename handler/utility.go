package handler

import (
	core "busha-challenge"
	"context"
	"encoding/json"
	"net/http"
)

func PopulateMovieDB(ctx context.Context, movieRepository core.MovieRepository) error {
	resp, err := http.Get("https://swapi.dev/api/films/")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var movies []*core.Movie
	err = json.NewDecoder(resp.Body).Decode(&movies)
	if err != nil {
		return err
	}
	return movieRepository.AddMany(ctx, movies)
}

func PopulateCharacterDB(ctx context.Context, characterRepository core.CharacterRepository) error {
	resp, err := http.Get("https://swapi.dev/api/people/")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var characters []*core.Character
	err = json.NewDecoder(resp.Body).Decode(&characters)
	if err != nil {
		return err
	}

	return characterRepository.AddMany(ctx, characters)
}
