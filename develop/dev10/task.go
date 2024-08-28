/*
=== Утилита telnet ===

Реализовать простейший telnet-клиент.

Примеры вызовов:
go-telnet --timeout=10s host port
go-telnet mysite.ru 8080
go-telnet --timeout=3s 1.1.1.1 123

Требования:
1. Программа должна подключаться к указанному хосту
(ip или доменное имя + порт) по протоколу TCP. После подключения STDIN
программы должен записываться в сокет, а данные полученные и сокета
должны выводиться в STDOUT.

2. Опционально в программу можно передать таймаут на подключение к серверу
(через аргумент --timeout, по умолчанию 10s)

3. При нажатии Ctrl+D программа должна закрывать сокет и завершаться.
Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout
*/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func main() {
	// Parse command-line arguments
	timeoutFlag := flag.Duration("timeout", 10*time.Second, "connection timeout")
	flag.Parse()

	if flag.NArg() < 2 {
		fmt.Println("Usage: go-telnet [--timeout=10s] host port")
		os.Exit(1)
	}

	host := flag.Arg(0)
	port := flag.Arg(1)

	address := net.JoinHostPort(host, port)
	conn, err := net.DialTimeout("tcp", address, *timeoutFlag)
	if err != nil {
		fmt.Printf("Failed to connect to %s: %v\n", address, err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Printf("Connected to %s\n", address)

	done := make(chan struct{})

	go func() {
		io.Copy(os.Stdout, conn)
		done <- struct{}{}
	}()

	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			input, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					fmt.Println("Detected EOF, closing connection...")
					conn.Close()
					done <- struct{}{}
					return
				}
				fmt.Printf("Error reading from stdin: %v\n", err)
				done <- struct{}{}
				return
			}
			_, err = conn.Write([]byte(input))
			if err != nil {
				fmt.Printf("Error writing to connection: %v\n", err)
				done <- struct{}{}
				return
			}
		}
	}()

	<-done
	fmt.Println("\n\nConnection closed")
}
