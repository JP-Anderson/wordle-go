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

func init() {
	rand.Seed(time.Now().Unix())
}

type WordsList struct{
	words []string
}

func (w WordsList) NextWord() string {
	return w.words[rand.Intn(len(w.words))]
}

func (w WordsList) Valid(s string) bool {
	l := 0
	r := len(w.words)-1
	for l<=r {
		mid := (l+r)/2
		if w.words[mid] == s {
			return true
		}
		if w.words[mid] > s {
			r = mid-1
		} else {
			l = mid+1
		} 
	}
	if l == len(w.words) || w.words[l] != s {
		return false
	}
	return true
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
		words = append(words, strings.ToLower(split))
	}
	return WordsList{
		words: words,
	}
}
