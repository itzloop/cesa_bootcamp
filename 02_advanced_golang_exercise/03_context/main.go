package main

import (
	"context"
	"errors"
	"fmt"
	"math"
	"sync"
	"time"
)

func HandleRequest(ctx context.Context, work int) {
	startTime := time.Now()
	defer func() {
		fmt.Printf("HandleRequest took %v to run\n", time.Since(startTime))
	}()

	wg := &sync.WaitGroup{}
	wg.Add(3)
	go Work1(work, wg)
	go Work2(work, wg)
	go Work3(work, wg)
	wg.Wait()
}

func Work1(work int, wg *sync.WaitGroup) error {
	defer wg.Done()
	userId := 0 // TODO
	if userId == 4 {
		return errors.New("we can't do work1 for this user!!!")
	}
	select {
	case <-time.After(time.Duration(work) * time.Second): // simulation of a time-consuming work
		fmt.Printf("the work1 for user %v is done\n", userId)
		return nil
	}
}

func Work2(work int, wg *sync.WaitGroup) {
	defer wg.Done()
	userId := 0 // TODO
	select {
	case <-time.After(time.Duration(math.Pow(2, float64(work))) * time.Second): // simulation of a time-consuming work
		fmt.Printf("the work2 for user %v is done\n", userId)
	}
}

func Work3(work int, wg *sync.WaitGroup) {
	defer wg.Done()
	userId := 0 // TODO
	select {
	case <-time.After(time.Duration(2*work) * time.Second): // simulation of a time-consuming work
		fmt.Printf("the work3 ror user %v is done\n", userId)
	}
}

func main() {
	ctx := context.WithValue(context.TODO(), "user_id", 1)
	HandleRequest(ctx, 2)

	ctx2 := context.WithValue(context.TODO(), "user_id", 2)
	HandleRequest(ctx2, 4)

	ctx3 := context.WithValue(context.TODO(), "user_id", 4)
	HandleRequest(ctx3, 4)
}
