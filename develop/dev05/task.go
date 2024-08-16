/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
)

type GrepArgs struct {
	StringsAfter      int
	StringsBefore     int
	StringsAround     int
	OnlyCount         bool
	IgnoreCase        bool
	WriteInvert       bool
	WriteExactString  bool
	WriteStringNumber bool
	Pattern           string
	Filename          string
}

func parseArgs() GrepArgs {
	stringsAfter := flag.Int("A", 0, "write strings after pattern")
	stringsBefore := flag.Int("B", 0, "write strings before pattern")
	stringsAround := flag.Int("C", 0, "write strings before and after pattern")
	onlyCount := flag.Bool("c", false, "write quantity of strings with pattern")
	ignoreCase := flag.Bool("i", false, "write strings ignoring case")
	writeInvert := flag.Bool("v", false, "write strings without pattern")
	writeExactString := flag.Bool("F", false, "write exact strings")
	writeStringNumber := flag.Bool("n", false, "write line number before string")

	flag.Parse()

	return GrepArgs{
		StringsAfter:      *stringsAfter,
		StringsBefore:     *stringsBefore,
		StringsAround:     *stringsAround,
		OnlyCount:         *onlyCount,
		IgnoreCase:        *ignoreCase,
		WriteInvert:       *writeInvert,
		WriteExactString:  *writeExactString,
		WriteStringNumber: *writeStringNumber,
		Pattern:           flag.Arg(0),
		Filename:          flag.Arg(1),
	}
}

// TODO: implement func for -A, -B, -C flags
func grep(args GrepArgs) ([]string, error) {
	file, err := os.Open(args.Filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var results []string
	scanner := bufio.NewScanner(file)
	lineNum := 0
	count := 0

	pattern := args.Pattern
	if args.IgnoreCase {
		pattern = "(?i)" + pattern
	}
	if !args.WriteExactString {
		pattern = ".*" + pattern + ".*"
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}

	for scanner.Scan() {
		line := scanner.Text()
		lineNum++
		match := re.MatchString(line)

		if args.WriteInvert {
			match = !match
		}

		if match {
			count++
			if args.OnlyCount {
				continue
			}
			if args.WriteStringNumber {
				line = fmt.Sprintf("%d:%s", lineNum, line)
			}
			results = append(results, line)
		}
	}

	if args.OnlyCount {
		results = []string{fmt.Sprintf("Count: %d", count)}
	}

	return results, scanner.Err()
}

func main() {
	args := parseArgs()

	results, err := grep(args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "grep: %v\n", err)
		os.Exit(1)
	}

	for _, result := range results {
		fmt.Println(result)
	}
}
