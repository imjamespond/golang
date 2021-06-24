package testprogressbar

import (
	"time"

	pb "github.com/cheggaaa/pb/v3"
)

func TestPb() {
	count := 10000
	bar := pb.StartNew(count)
	for i := 0; i < count; i++ {
		bar.Increment()
		time.Sleep(time.Millisecond)
	}
	bar.Finish()
}
