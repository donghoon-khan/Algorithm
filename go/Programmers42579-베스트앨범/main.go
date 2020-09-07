package main

import (
	"fmt"
	"sort"
)

type PlayStruct struct {
	idx   int
	score int
}

type GenreStruct struct {
	totalScore int
	plays      []PlayStruct
}

func solution(genres []string, plays []int) []int {
	var ret []int
	genresMap := map[string]GenreStruct{}

	for i := 0; i < len(genres); i++ {
		genre := genresMap[genres[i]]
		var play PlayStruct
		genre.totalScore += plays[i]
		play.idx = i
		play.score = plays[i]
		genre.plays = append(genre.plays, play)
		genresMap[genres[i]] = genre
	}

	sortGenresMap := map[int]string{}
	sortGenresKeys := []int{}
	for key, val := range genresMap {
		sortGenresMap[val.totalScore] = key
		sortGenresKeys = append(sortGenresKeys, val.totalScore)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sortGenresKeys)))

	for _, val := range sortGenresKeys {
		playByGenres := genresMap[sortGenresMap[val]].plays

		if len(playByGenres) < 2 {
			ret = append(ret, playByGenres[0].idx)
			continue
		}

		sort.Slice(playByGenres, func(i, j int) bool {

			if playByGenres[i].score == playByGenres[j].score {
				return playByGenres[i].idx < playByGenres[j].idx
			}
			return playByGenres[i].score > playByGenres[j].score
		})
		ret = append(ret, playByGenres[0].idx, playByGenres[1].idx)

	}
	return ret
}

func main() {
	genres := []string{"classic", "pop", "classic", "classic", "pop"}
	plays := []int{500, 600, 150, 800, 2500}

	fmt.Println(solution(genres, plays))
}
