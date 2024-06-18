package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
	"os/signal"
    "strings"
    "sync"
	"syscall"
	"time"
)

func t23_1() {
    var wg sync.WaitGroup
    numberChan := make(chan int)
    squareChan := make(chan int)
    productChan := make(chan int)

    wg.Add(1)
    go func() {
        defer wg.Done()
        for number := range numberChan {
            squareChan <- number * number
        }
        close(squareChan)
    }()

    wg.Add(1)
    go func() {
        defer wg.Done()
        for square := range squareChan {
            productChan <- square * 2
        }
        close(productChan)
    }()

    wg.Add(1)
    go func() {
        defer wg.Done()
        for product := range productChan {
            fmt.Println("Результат:", product)
        }
    }()

    scanner := bufio.NewScanner(os.Stdin)
    fmt.Println("Введите число или 'стоп' для завершения:")
    for scanner.Scan() {
        text := scanner.Text()
        if strings.ToLower(text) == "стоп" {
            break
        }

        if number, err := strconv.Atoi(text); err == nil {
            numberChan <- number
        } else {
            fmt.Println("Ошибка: введите корректное число")
        }
    }

    close(numberChan)
    wg.Wait()
}

func t23_2() {
	exit := make(chan os.Signal)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
	stop := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		num := 1
		for {
			select {
			case <-stop:
				fmt.Println("Выхожу из программы...")
				return
			default:
				square := num * num
				fmt.Printf("Квадрат числа %d: %d\n", num, square)
				num++
				time.Sleep(time.Second) 
			}
		}
	}()
	<-exit
	close(stop)
	wg.Wait()
}

func main() {
	t23_2()
}