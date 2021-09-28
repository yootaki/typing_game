package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

var (
	question string
	score    int
)

/* CountDown */
func countdown() {
	for i := 3; i > 0; i-- {
		fmt.Printf("%d ", i)
		time.Sleep(time.Second)
	}
	fmt.Println("Let's Gooooo!")
}

/* Random string generater */
func init() {
	rand.Seed(time.Now().UnixNano())
}

var rs1Letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPWRSTUVWXYZ")

func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = rs1Letters[rand.Intn(len(rs1Letters))]
	}
	return string(b)
}

/* 5文字のランダム文字列生成 */
func get_q() {
	question = RandString(5)
	fmt.Printf("\n- %s -\n", question)
	fmt.Print("> ")
}

/* 標準入力からの文字列受け取り */
func input(r io.Reader) <-chan string {
	ch := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch <- s.Text()
		}
		close(ch)
	}()
	return ch
}

func main() {
	fmt.Println("-------------\n Typing Game \n-------------")
	countdown()
	bc := context.Background()
	t := 30 * time.Second
	ctx, cancel := context.WithTimeout(bc, t)
	defer cancel()

	get_q()

	ch := input(os.Stdin)
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("\n\n-------------\n %d Points!!! \n-------------\n", score)
			return
		case v := <-ch:
			if v == question {
				score++
				fmt.Println("Goodjob!!!")
			} else {
				fmt.Println("Baaad!!!")
			}
			get_q()
		}
	}
}
