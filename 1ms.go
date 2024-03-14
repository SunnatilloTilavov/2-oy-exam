package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

func rand1(ch chan<- int, a, b, count int) {
    defer close(ch)
    for i := 0; i < count; i++ {
        num := rand.Intn(b-a+1) + a
        ch <- num
        time.Sleep(2 * time.Second)
    }
}

func Squares(ch1 <-chan int, ch2 chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()
    for num := range ch1 {
        square := num * num
        ch2 <- square
    }
    close(ch2)
}

func main() {
    ch1 := make(chan int)
    ch2 := make(chan int)

    var wg sync.WaitGroup
    wg.Add(2)
    rand.Seed(time.Now().UnixNano())
    go rand1(ch1, 10, 30, rand.Intn(6)+5)
    go Squares(ch1, ch2, &wg)
    go func() {
        defer wg.Done()
        for numb := range ch2 {
            fmt.Println("Kv:", numb)
            time.Sleep(2 * time.Second)
        }
    }()

    wg.Wait()
}
