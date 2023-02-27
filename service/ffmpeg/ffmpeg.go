package ffmpeg

import (
	"errors"
	"fmt"
	"http-server/utils"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
)

type Helper struct {
	once  sync.Once
	wait  sync.WaitGroup
	mutex sync.Mutex
}
type Service struct {
	queue chan *Job
}
type Job struct {
	Type int
	Data interface{}
}
type JobDl struct {
	M3U8 string `binding:"required"`
	File string `binding:"required"`
}
type JobConv struct {
	File string `binding:"required"`
}

var (
	helper  Helper
	service *Service
	log     *string
	cmd     *exec.Cmd
)

func setLog(str *string) {
	defer helper.mutex.Unlock()
	helper.mutex.Lock()
	log = str
}
func setCmd(c *exec.Cmd) {
	defer helper.mutex.Unlock()
	helper.mutex.Lock()
	if nil == c && cmd != nil {
		cmd.Process.Kill()
	}
	cmd = c
}

func (self *Service) Start() {
	go func() {
		for {
			utils.ProtectRun(func() {
				job, ok := <-self.queue
				if ok {
					if 0 == job.Type {
						time.Sleep(time.Second)
						fmt.Println("done", *job)
					} else if 1 == job.Type {
						dl := job.Data.(*JobDl)
						// https://stackoverflow.com/questions/37091316/how-to-get-the-realtime-output-for-a-shell-command-in-golang
						// Looks like ffmpeg sends all diagnostic messages (the "console output") to stderr instead of stdout. Below code works for me.
						cmd, exec := utils.ExecCommandDir(nil, "bash", os.Getenv("HOME")+"/dl-m3u8.sh", dl.M3U8, dl.File)

						setCmd(cmd)
						exec(func(str string) {
							if strings.Index(str, "out_time=") == 0 {
								fmt.Println("scan", str)
								setLog(&str)
							}
						})
						setCmd(nil)
						fmt.Println("done", *job)
					} else if 2 == job.Type {
						dl := job.Data.(*JobConv)
						cmd, exec := utils.ExecCommandDir(nil, "bash", os.Getenv("HOME")+"/convert.sh", dl.File)
						setCmd(cmd)
						exec(func(str string) {
							if strings.Index(str, "out_time=") == 0 {
								fmt.Println("scan", str)
								setLog(&str)
							}
						})
						setCmd(nil)
						fmt.Println("done", *job)
					}
				}
			})
		}
	}()
}

func (self *Service) Kill() {
	setCmd(nil)
	utils.ExecCommand("bash", "-c", "ps -ef|grep \"ffmpeg\"|awk '{print $2}'|xargs kill -9")
	setLog(nil)
}

func (self *Service) Add(job *Job) error {
	if len(self.queue) > 8 {
		return errors.New("queue is too large")
	}
	self.queue <- job
	return nil
}

func (self *Service) Log() (string, error) {
	defer helper.mutex.Unlock()
	helper.mutex.Lock()
	if log == nil || cmd == nil {
		return "", errors.New("log is null")
	}
	return *log, nil
}

func GetInstance() *Service {
	/*
		Because no call to Do returns until the one call to f returns, if f causes Do to be called, it will deadlock.
		If f panics, Do considers it to have returned; future calls of Do return without calling f. 如果panics最好确保返回，将来Do不会再call f
	*/
	helper.once.Do(func() {
		service = new(Service)
		service.queue = make(chan *Job, 32)
		fmt.Println("GetInstance")
	})

	return service
}
