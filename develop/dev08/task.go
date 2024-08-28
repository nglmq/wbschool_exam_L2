/*Взаимодействие с ОС

Необходимо реализовать свой собственный UNIX-шелл-утилиту
с поддержкой ряда простейших команд:

- cd <args> - смена директории (в качестве аргумента могут быть то-то и то)
- pwd - показать путь до текущего каталога
- echo <args> - вывод аргумента в STDOUT
- kill <args> - "убить" процесс, переданный в качесте аргумента (пример: такой-то пример)
- ps - выводит общую информацию по запущенным процессам
в формате *такой-то формат*

Так же требуется поддерживать функционал fork/exec-команд

*Шелл — это обычная консольная программа, которая будучи запущенной,
в интерактивном сеансе выводит некое приглашение в STDOUT
и ожидает ввода пользователя через STDIN. Дождавшись ввода,
обрабатывает команду согласно своей логике и при необходимости
выводит результат на экран. Интерактивный сеанс поддерживается до тех пор,
пока не будет введена команда выхода (например \quit).
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">>>>> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")

		args := strings.Fields(input)
		if len(args) == 0 {
			continue
		}

		switch args[0] {
		case "cd":
			if len(args) < 2 {
				fmt.Println("cd: missing argument")
			} else {
				changeDirectory(args[1])
			}
		case "pwd":
			printWorkingDirectory()
		case "echo":
			echoArgs(args[1:])
		case "kill":
			if len(args) < 2 {
				fmt.Println("kill: missing argument")
			} else {
				killProcess(args[1])
			}
		case "ps":
			listProcesses()
		case "\\quit":
			fmt.Println("Exiting shell...")
			os.Exit(0)
		default:
			executeCommand(args)
		}
	}
}

func changeDirectory(path string) {
	err := os.Chdir(path)
	if err != nil {
		fmt.Println("cd:", err)
	}
}

func printWorkingDirectory() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("pwd:", err)
	} else {
		fmt.Println(dir)
	}
}

func echoArgs(args []string) {
	fmt.Println(strings.Join(args, " "))
}

func killProcess(pidStr string) {
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		fmt.Println("kill: invalid PID")
		return
	}
	process, err := os.FindProcess(pid)
	if err != nil {
		fmt.Println("kill: cant find process with this pid")
		return
	}
	if err := process.Kill(); err != nil {
		fmt.Println("kill:", err)
		return
	}
}

func listProcesses() {
	cmd := exec.Command("ps")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("ps:", err)
	}
}

func executeCommand(args []string) {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		fmt.Println(args[0], ":", err)
	}
}
