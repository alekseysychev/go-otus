package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

type wordStatistic struct {
	word  string //
	count int    //
}

func Top10(text string) []string {
	// словарь для пропуска элементов
	skipWords := map[string]bool{
		"-": true,
	}

	words := regexp.MustCompile(`[\p{L}\d-]+`).FindAllString(text, -1)

	// собираем уникальный словарь
	unique := map[string]int{}
	for _, word := range words {
		if _, ok := skipWords[word]; ok {
			continue
		}
		unique[strings.ToLower(word)]++
	}
	// делаем из словаря срез
	slice := make([]wordStatistic, 0, len(unique))

	for word, count := range unique {
		slice = append(slice, wordStatistic{word, count})
	}
	// сортируем по количеству вхождений
	sort.Slice(slice, func(i, j int) bool {
		return slice[i].count > slice[j].count
	})
	// отбираем первые 10 элементов
	result := []string{}
	for i := 0; i < 10 && i < len(slice)-1; i++ {
		result = append(result, slice[i].word)
	}
	return result
}
