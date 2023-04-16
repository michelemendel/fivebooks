package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

var char2Gematria = map[string]int{
	"א": 1,
	"ב": 2,
	"ג": 3,
	"ד": 4,
	"ה": 5,
	"ו": 6,
	"ז": 7,
	"ח": 8,
	"ט": 9,
	"י": 10,
	// "ך": 20,
	"כ": 20,
	"ל": 30,
	// "ם": 40,
	"מ": 40,
	// "ן": 50,
	"נ": 50,
	"ס": 60,
	"ע": 70,
	// "ף": 80,
	"פ": 80,
	// "ץ": 90,
	"צ": 90,
	"ק": 100,
	"ר": 200,
	"ש": 300,
	"ת": 400,
}

type HebChar struct {
	Char     string
	Gematria int
	Unicode  int
	Count    int
	Pos      []int
}

const (
	nofHebrewChars = 27
	aleph          = 0x05D0
)

var HebChars = make(map[string]*HebChar)

func ppHebchars() {

	b, _ := json.MarshalIndent(HebChars, "", "  ")
	fmt.Println(string(b))
}

func init() {
	for i := 0; i < nofHebrewChars; i++ {
		uni := aleph + i
		ch := string(rune(uni))
		if char2Gematria[ch] != 0 { //We only want the basic Hebrew chars, not the final chars
			HebChars[ch] = &HebChar{Char: ch, Gematria: char2Gematria[ch], Unicode: uni, Count: 0, Pos: []int{}}
		}
	}
}

var (
	// reHebrewChar := regexp.MustCompile(`[\x{05D0}-\x{05EA}]`)
	reHebrewChar = regexp.MustCompile(`\p{Hebrew}`)
	// reHebrewWord = regexp.MustCompile(`\p{Hebrew}*`)
)

func readBookChars(filename string) map[string]*HebChar {
	return readBook(filename, reHebrewChar)
}

func readBook(filename string, reElem *regexp.Regexp) map[string]*HebChar {
	f := fileHandler(filename)
	defer f.Close()

	var line string
	var elems []string
	var lineNr = 1

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line = scanner.Text()
		line = replaceEndChars(line, "ך", "כ")
		line = replaceEndChars(line, "ם", "מ")
		line = replaceEndChars(line, "ן", "נ")
		line = replaceEndChars(line, "ף", "פ")
		line = replaceEndChars(line, "ץ", "צ")
		line = replaceEndChars(line, "[", "")
		line = replaceEndChars(line, "]", "")

		elems = reElem.FindAllString(line, -1)

		for idx, c := range elems {
			_ = idx
			if HebChars[c] == nil {
				fmt.Printf("Bad char (#%v#) at line,col: %d,%d\n", c, lineNr, idx)
				continue
			}

			if HebChars[c].Count == 0 {
				HebChars[c].Count = 1
			} else {
				HebChars[c].Count++
			}
			HebChars[c].Pos = append(HebChars[c].Pos, idx)
		}
		lineNr++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// ppHebchars()

	return HebChars
}

func fileHandler(filename string) *os.File {
	f, err := os.Open(filename)
	if err != nil {
		log.Panicf("failed to open file: %s", err)
	}
	return f
}

func reverse[S ~[]E, E any](s S) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func replaceEndChars(s, old, new string) string {
	return strings.ReplaceAll(s, old, new)
}

type kv struct {
	k string
	v int
}

func toKV(m map[string]int) []kv {
	var kvs []kv
	for k, v := range m {
		kvs = append(kvs, kv{k, v})
	}
	return kvs
}

func sortMapByKey(kvs []kv) []kv {
	sort.Slice(kvs, func(i, j int) bool {
		return kvs[i].k < kvs[j].k
	})
	return kvs
}

func sortMapByValue(kvs []kv) []kv {
	sort.Slice(kvs, func(i, j int) bool {
		return kvs[i].v < kvs[j].v
	})
	return kvs
}
