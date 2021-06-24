package testprogressbar

import (
	"testing"
	"time"

	pb "github.com/cheggaaa/pb/v3"
)

func Test1(t *testing.T) {
	count := 1000
	// create and start new bar
	// bar := pb.StartNew(count)

	// start bar from 'default' template
	// bar := pb.Default.Start(count)

	// start bar from 'simple' template
	// bar := pb.Simple.Start(count)

	// start bar from 'full' template
	bar := pb.Full.Start(count)

	for i := 0; i < count; i++ {
		bar.Increment()
		time.Sleep(time.Second)
	}
	bar.Finish()
}
