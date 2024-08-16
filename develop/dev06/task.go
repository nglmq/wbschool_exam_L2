/*
Реализовать утилиту аналог консольной команды cut (man cut).
Утилита должна принимать строки через STDIN,
разбивать по разделителю (TAB) на колонки и выводить запрошенные.

Реализовать поддержку утилитой следующих ключей:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем
*/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var fields string
	var delimiter string
	var separatedOnly bool

	flag.StringVar(&fields, "f", "", "fields to select")
	flag.StringVar(&delimiter, "d", "\t", "delimiter")
	flag.BoolVar(&separatedOnly, "s", false, "only lines with the delimiter")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if separatedOnly && !strings.Contains(line, delimiter) {
			continue
		}

		columns := strings.Split(line, delimiter)
		fieldIndexes := parseFields(fields)

		for i, index := range fieldIndexes {
			if index < len(columns) {
				if i > 0 {
					fmt.Print(delimiter)
				}
				fmt.Print(columns[index])
			}
		}
		fmt.Println()
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error reading standard input: %s\n", err)
	}
}

func parseFields(fieldStr string) []int {
	var result []int
	fields := strings.Split(fieldStr, ",")
	for _, f := range fields {
		var index int
		fmt.Sscanf(f, "%d", &index)
		result = append(result, index-1)
	}
	return result
}
