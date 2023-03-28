package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

// const shakespear = "shakespear.txt"
// const filenameShort = "shakespear_short.txt"
// const breshit = "breshit.txt"
// const breshit = "breshit_short.txt"
const breshit = "breshit_shorter.txt"

func main() {
	readAndExtractByLine(breshit)
}

func readAndExtractByLine(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Panicf("failed to open file: %s", err)
	}
	defer f.Close()

	var line string
	var chars, words []string
	var allChars []string
	var allWords []string

	// reNonLatinChar := regexp.MustCompile(`[\x{05D0}-\x{05EA}]`)
	// reNonLatinChar := regexp.MustCompile(`\p{Hebrew}`)
	reNonLatinWord := regexp.MustCompile(`\p{Hebrew}*`)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line = scanner.Text()
		line = replaceEndChars(line, "ך", "כ")
		line = replaceEndChars(line, "ם", "מ")
		line = replaceEndChars(line, "ן", "נ")
		line = replaceEndChars(line, "ף", "פ")
		line = replaceEndChars(line, "ץ", "צ")
		// line = replaceEndChars(line, "[", "")
		// line = replaceEndChars(line, "]", "")

		words = reNonLatinWord.FindAllString(line, -1)
		// chars = reNonLatinChar.FindAllString(line, -1)
		fmt.Println(line)
		// fmt.Println("Chars in the line:", len(chars))
		// fmt.Println(chars)
		// fmt.Println("Words in the line:", len(words))
		// fmt.Println(words)

		allChars = append(allChars, chars...)
		allWords = append(allWords, words...)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Chars in the file:", len(allChars))
	// fmt.Println(allChars)
	// fmt.Println("")
	fmt.Println("Words in the file:", len(allWords))
	fmt.Println(allWords)
	fmt.Println("")

	// ----------------------------------------
	var charStatsMap = make(map[string]int)
	for _, c := range allChars {
		if charStatsMap[c] == 0 {
			charStatsMap[c] = 1
		} else {
			charStatsMap[c]++
		}
	}

	// charStats := toKV(charStatsMap)
	// fmt.Println("CharStats:", charStats)
	// fmt.Println("Alef:", charStatsMap["א"])

	// sortedByKey := sortMapByKey(charStats)
	// sortedByValue := sortMapByValue(charStats)
	// fmt.Println("SortedByKey:", sortedByKey)
	// fmt.Println("SortedByValue:", sortedByValue)
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
