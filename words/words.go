package words

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)



type WordsList struct{
	words []string
}

func (w WordsList) NextWord() string {
	rand.Seed(time.Now().Unix())
	return w.words[rand.Intn(len(w.words))]
}

func WordsListFromFile(listName string) WordsList {
	readFile, err := os.Open(listName)
	if err != nil {
		fmt.Println(err)
	}
	words := []string{}
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
 	loop:
	for fileScanner.Scan() {
		split := strings.Fields(fileScanner.Text())[0]
		count := 0
		for _, rune := range split {
			count+=1
			if !unicode.IsLetter(rune) {
				continue loop	
			}
		}
		if count != 5 {
			continue loop
		}
		words = append(words, split)
	}
	return WordsList{
		words: words,
	}
}
