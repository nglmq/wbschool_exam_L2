/*
	Утилита sort
	================

	Отсортировать строки в файле по аналогии с консольной утилитой sort
	(man sort — смотрим описание и основные параметры):
	на входе подается файл с несортированными строками, на выходе — файл с отсортированными.

	Реализовать поддержку утилитой следующих ключей:

	-k — указание колонки для сортировки
	(слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел)
	-n — сортировать по числовому значению
	-r — сортировать в обратном порядке
	-u — не выводить повторяющиеся строки

	==================
	Дополнительно
	==================

	Реализовать поддержку утилитой следующих ключей:

	-M — сортировать по названию месяца
	-b — игнорировать хвостовые пробелы
	-c — проверять отсортированы ли данные
	-h — сортировать по числовому значению с учетом суффиксов
*/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type SortArgs struct {
	ColumnForSort       int
	SortByNumber        bool
	ReverseSort         bool
	DontPrintDuplicates bool
	Filename            string
}

func parseArgs() SortArgs {
	columnForSort := flag.Int("k", 0, "choose column for sort")
	sortByNumber := flag.Bool("n", false, "sort by number")
	reverseSort := flag.Bool("r", false, "reverse sort")
	dontPrintDuplicates := flag.Bool("u", false, "dont print duplicates")

	flag.Parse()

	flags := SortArgs{
		ColumnForSort:       *columnForSort,
		SortByNumber:        *sortByNumber,
		ReverseSort:         *reverseSort,
		DontPrintDuplicates: *dontPrintDuplicates,
		Filename:            flag.Arg(0),
	}

	return flags
}

func ParseFile(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func WriteToFile(filepath string, lines []string) error {
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := os.WriteFile(filepath, []byte(strings.Join(lines, "\n")), 0644); err != nil {
		return err
	}

	return nil
}

func SortStrings(filepath string, column int) ([]string, error) {
	lines, err := ParseFile(filepath)
	if err != nil {
		return nil, err
	}

	sort.Strings(lines)

	return lines, nil
}

func SortByNumber(filepath string, column int) ([]string, error) {
	lines, err := ParseFile(filepath)
	if err != nil {
		return nil, err
	}

	sort.Slice(lines, func(i, j int) bool {
		splitI := strings.Split(lines[i], " ")
		splitJ := strings.Split(lines[j], " ")

		if len(splitI) <= column || len(splitJ) <= column {
			return false
		}

		i, err := strconv.Atoi(splitI[column])
		if err != nil {
			return false
		}

		j, err = strconv.Atoi(splitJ[column])
		if err != nil {
			return false
		}

		return j > i
	})

	return lines, nil
}

func ReverseSortInt(filepath string, column int) ([]string, error) {
	lines, err := ParseFile(filepath)
	if err != nil {
		return nil, err
	}

	sort.Slice(lines, func(i, j int) bool {
		splitI := strings.Split(lines[i], " ")
		splitJ := strings.Split(lines[j], " ")

		if len(splitI) <= column || len(splitJ) <= column {
			return false
		}

		i, err := strconv.Atoi(splitI[column])
		if err != nil {
			return false
		}

		j, err = strconv.Atoi(splitJ[column])
		if err != nil {
			return false
		}

		return i > j
	})

	return lines, nil
}

func ReverseSortString(filepath string, column int) ([]string, error) {
	lines, err := ParseFile(filepath)
	if err != nil {
		return nil, err
	}

	sort.Slice(lines, func(i, j int) bool {
		return lines[i] > lines[j]
	})

	return lines, nil
}

func DeleteDuplicates(filepath string) ([]string, error) {
	lines, err := ParseFile(filepath)
	if err != nil {
		return nil, err
	}

	uniqueLines := make(map[string]bool, len(lines))

	for _, line := range lines {
		if _, ok := uniqueLines[line]; ok {
			uniqueLines[line] = false
		}
		uniqueLines[line] = true
	}

	fmt.Print(uniqueLines)

	var onlyUnique []string

	for line, isUnique := range uniqueLines {
		if isUnique {
			onlyUnique = append(onlyUnique, line)
		}
	}

	return onlyUnique, nil
}

func main() {
	flags := parseArgs()

	if flags.SortByNumber {
		lines, err := SortByNumber(flags.Filename, flags.ColumnForSort)
		if err != nil {
			fmt.Errorf("error sorting by number: %v", err)
		}

		err = WriteToFile(flags.Filename, lines)
		if err != nil {
			fmt.Errorf("error writing new file while sorting by number: %v", err)
		}
	}

	if !flags.SortByNumber {
		lines, err := SortStrings(flags.Filename, flags.ColumnForSort)
		if err != nil {
			fmt.Errorf("error sorting strings: %v", err)
		}

		err = WriteToFile(flags.Filename, lines)
		if err != nil {
			fmt.Errorf("error writing new file while sorting strings: %v", err)
		}
	}

	if flags.ReverseSort && flags.SortByNumber {
		lines, err := ReverseSortInt(flags.Filename, flags.ColumnForSort)
		if err != nil {
			fmt.Errorf("error reverse sorting ints: %v", err)
		}

		err = WriteToFile(flags.Filename, lines)
		if err != nil {
			fmt.Errorf("error writing new file while reverse sorting: %v", err)
		}
	}

	if flags.ReverseSort && !flags.SortByNumber {
		lines, err := ReverseSortString(flags.Filename, flags.ColumnForSort)
		if err != nil {
			fmt.Errorf("error reverse sorting strings: %v", err)
		}

		err = WriteToFile(flags.Filename, lines)
		if err != nil {
			fmt.Errorf("error writing to file while reverse sorting strings: %v", err)
		}
	}

	if flags.DontPrintDuplicates {
		lines, err := DeleteDuplicates(flags.Filename)
		if err != nil {
			fmt.Errorf("error deleting duplicates: %v", err)
		}

		err = WriteToFile(flags.Filename, lines)
		if err != nil {
			fmt.Errorf("error writing new file while deleting duplicates: %v", err)
		}
	}
}
