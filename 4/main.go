package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

/*
Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.
Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ
завершения работы всех воркеров.
*/

/*
Выбор: отмена через контекст, так как контекст это удобный инструмент передачи завершения работы дочерних функций из вызываемых
Так пример задачи идеально подходит под назначение этого инструмента
*/

var wg sync.WaitGroup

// раелизация воркера: пока канал открыт, читаем значения из канала
func worker(workerNum int, ch <-chan int) {
	defer wg.Done()
	for n := range ch {
		// вывод принятого значения из канала в формате номер воркера : значение
		fmt.Println(workerNum, " : ", n)
	}
}

// generator - функция генерации новых значений (последовательный целые положительные)
func generator(ctx context.Context, ch chan<- int) {
	defer close(ch)
	num := 0
	for {
		select {
		// принятие сигнала отмены контекста - завершение работы генератора
		case <-ctx.Done():
			fmt.Println("Done")
			return
		// вариант генерации нового значения и запись в канал
		case ch <- num:
			num++
			time.Sleep(time.Millisecond * 500) // для задержек, чтобы результат был наглядным
		}
	}
}

func main() {

	var countWorkers int
	_, err := fmt.Scan(&countWorkers)
	if err != nil {
		fmt.Println(err)
		return
	}

	// формирование контекста с отменой
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan int)

	go generator(ctx, ch)

	//запуск воркеров в количестве countWorkers
	wg.Add(countWorkers)
	for i := 0; i < countWorkers; i++ {
		go worker(i, ch)
	}

	// обработка сигнала завершения
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	fmt.Println("Stop signal")

	// отмена генерации, закрытие канала, остановка работы воркеров
	cancel()
	// ждем завершения всех воркеров
	wg.Wait()
	fmt.Println("Finish")
}
