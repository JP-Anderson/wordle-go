package words

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadValidWords(t *testing.T) {
	wordsList := WordsListFromFile("dic.txt")
	assert.Equal(t, 14868, len(wordsList.words))
	for _, word := range wordsList.words {
		assert.Equal(t, 5, len(word))
		assert.NotNil(t, word)
	}
}

func TestRandomWord(t *testing.T) {
	wordsList := WordsListFromFile("dic.txt")
	word := wordsList.NextWord()
	assert.NotNil(t, word)
	assert.Len(t, word, 5)
}
