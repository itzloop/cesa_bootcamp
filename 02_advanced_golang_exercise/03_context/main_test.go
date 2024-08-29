package main

import (
	"context"
	"testing"
	"time"
)

func TestNormalExecution(t *testing.T) {
	ctx := context.WithValue(context.Background(), "user_id", 1)
	startTime := time.Now()
	HandleRequest(ctx, 2)
	execTime := time.Since(startTime)

	expectedDuration := 4 * time.Second
	tolerance := 100 * time.Millisecond

	if execTime < expectedDuration-tolerance || execTime > expectedDuration+tolerance {
		t.Errorf("Execution time was %v, expected around %v", execTime, expectedDuration)
	}
}

func TestTimeout(t *testing.T) {
	ctx := context.WithValue(context.Background(), "user_id", 2)
	startTime := time.Now()
	HandleRequest(ctx, 4)
	execTime := time.Since(startTime)

	expectedDuration := 10 * time.Second
	tolerance := 100 * time.Millisecond

	if execTime < expectedDuration-tolerance || execTime > expectedDuration+tolerance {
		t.Errorf("Execution time was %v, expected around %v", execTime, expectedDuration)
	}
}

func TestCancel(t *testing.T) {
	ctx := context.WithValue(context.Background(), "user_id", 4)
	startTime := time.Now()
	HandleRequest(ctx, 4)
	execTime := time.Since(startTime)

	expectedDuration := 0 * time.Second
	tolerance := 100 * time.Millisecond

	if execTime < expectedDuration-tolerance || execTime > expectedDuration+tolerance {
		t.Errorf("Execution time was %v, expected around %v", execTime, expectedDuration)
	}
}
