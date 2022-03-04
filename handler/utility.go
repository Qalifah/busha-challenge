package handler

import (
	core "busha-challenge"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

func PopulateMovieDB(ctx context.Context, movieRepository core.MovieRepository) error {
	resp, err := http.Get("https://swapi.dev/api/films/")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var rValue = struct {
		results []*core.Movie
	}{}
	err = json.NewDecoder(resp.Body).Decode(&rValue)
	if err != nil {
		return err
	}
	return movieRepository.AddMany(ctx, rValue.results)
}

func PopulateCharacterDB(ctx context.Context, characterRepository core.CharacterRepository) error {
	resp, err := http.Get("https://swapi.dev/api/people/")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// create response type
	rValue := struct {
		results []struct {
			name   string
			gender string
			height string
			films  []string
		}
	}{}

	err = json.NewDecoder(resp.Body).Decode(&rValue.results)
	if err != nil {
		return err
	}

	// transform response to fit the character model
	var characters []*core.Character
	for _, i := range rValue.results {
		char := &core.Character{
			Name:   i.name,
			Gender: i.gender,
			Height: i.height,
		}
		for _, j := range i.films {
			id, err := strconv.Atoi(j[len(j)-2 : len(j)-1])
			if err != nil {
				return err
			}
			char.MoviesID = append(char.MoviesID, id)
		}
		characters = append(characters, char)
	}
	return characterRepository.AddMany(ctx, characters)
}
