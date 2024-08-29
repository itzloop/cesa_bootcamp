package fanin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFanIn(t *testing.T) {
	var (
		channels         []chan int
		recvOnlyChannels []<-chan int
		channelCount     = 100
		n                = 500000
	)

	for i := 0; i < channelCount; i++ {
		ch := make(chan int)
		channels = append(channels, ch)
		recvOnlyChannels = append(recvOnlyChannels, ch)
	}

	ch3 := runFanIn(recvOnlyChannels...)

	// send some data on ch1 and ch2
	go func() {
		for i := 0; i < n; i++ {
			channels[i%channelCount] <- i + 1
		}

		for _, ch := range channels {
			close(ch)
		}
	}()

	var (
		expectedSum = ((n) * (n + 1)) / 2
		actualSum   = 0
	)
	for result := range ch3 {
		actualSum += result
	}

    assert.EqualValues(t, expectedSum, actualSum)
}
