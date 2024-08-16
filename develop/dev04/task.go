/*
Написать функцию поиска всех множеств анаграмм по словарю.

===========
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.
===========

Требования:
	Входные данные для функции: ссылка на массив, каждый элемент которого - слово на русском языке в кодировке utf8
	Выходные данные: ссылка на мапу множеств анаграмм

	Ключ - первое встретившееся в словаре слово из множества. Значение - ссылка на массив, каждый элемент которого,
	слово из множества.

	Массив должен быть отсортирован по возрастанию.
	Множества из одного элемента не должны попасть в результат.
	Все слова должны быть приведены к нижнему регистру.
	В результате каждое слово должно встречаться только один раз.
*/

package main

import (
	"fmt"
	"sort"
	"strings"
)

func findAnogram(words []string) map[string][]string {
	tempMap := make(map[string][]string)
	anogramMap := make(map[string][]string)

	for _, word := range words {
		strings.ToLower(word)
		sortedWord := sortChars(word)

		tempMap[sortedWord] = append(tempMap[sortedWord], word)
	}

	for _, group := range tempMap {
		if len(group) > 1 {
			sort.Strings(group)

			key := group[0]

			uniqueWords := onlyUnique(group)
			anogramMap[key] = uniqueWords
		}
	}

	return anogramMap
}

func sortChars(word string) string {
	runes := []rune(word)

	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	return string(runes)
}

func onlyUnique(words []string) []string {
	seen := make(map[string]bool)
	result := []string{}

	for _, word := range words {
		if !seen[word] {
			seen[word] = true
			result = append(result, word)
		}
	}

	return result
}

func main() {
	sl := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик","пятак","пятак","пятак","пятак","слиток","слиток","слиток","слиток"}
	
	fmt.Println(findAnogram(sl))
}
