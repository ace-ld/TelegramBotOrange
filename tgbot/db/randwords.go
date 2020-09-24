package db

import (
	"math/rand"
	"time"
)

// return text about bot
func rangeDB(words []WordsDB) string {
	var slice = make([]WordsDB, 0)
	for _, valueMap := range words {
		slice = append(slice, valueMap)
	}

	return randWords(slice)
}

func randWords(words []WordsDB) string {
	rand.Seed(time.Now().UnixNano())
	var word = words[rand.Intn(len(words))]
	return word.text
}
