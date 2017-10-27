package main

import (
	"log"
	"net/http"
	"time"
)

const (
	numPollers = 2
	pollInterval = 60 * time.Second
	statusInterval = 10 * time.Second
	errTimeout = 10 * time.Second

)

var urls = []string{
	"https://www.jfz.com",
	"https://passport.jinfuzi.com",
	"https://ty.jfz.com",
	"https://m.jfz.com",
	"https://broker.jfz.com",
	"http://yun.jinfuzi.com/",
}

//State represents the last-known state of a url
//State 表示一个url最后的已知状态

type State struct {
	url string
	status string
}

//StateMonitor 维护了一个映射，它存储了URL被轮询的状态，并每隔updateInterval 纳秒
//打印出当前的状态，它向资源状态的接收者返回一个chan State

func StateMonitor(updateInterval time.Duration) chan<- State {
	updates := make(chan State)
	urlStatus := make(map[string]string)
	ticker := time.NewTicker(updateInterval)

	go func() {
		for {
			select {
			case <- ticker.C:
				logState(urlStatus)

			case s := <- updates:
				urlStatus[s.url] = s.status
			}
		}
	}()

	return updates
}

// logState 打印出一个状态映射
func logState(s map[string]string) {
	log.Println("Current state:")
	for k, v := range s {
		log.Printf(" %s %s", k, v)
	}
}

//Resource 表示一个被此程序轮询的http url
type Resource struct {
	url string
	errCount int
}

//Poll为url执行一个http head请求，并返回http的状态字符串或一个错误字符串
func (r *Resource) Poll() string {
	resp, err := http.Head(r.url)
	if err != nil {
		log.Println("Error", r.url, err)
		r.errCount++
		return err.Error()
	}

	r.errCount = 0
	return resp.Status
}

func (r *Resource) Sleep(done chan<- *Resource) {
	time.Sleep(pollInterval + errTimeout * time.Duration(r.errCount))
	done <- r
}

func Poller(in <-chan *Resource, out chan<- *Resource, status chan<- State) {
	for r:= range in {
		s := r.Poll()
		status <- State{r.url, s}
		out <- r
	}
}

func main() {
	pending, complete := make(chan *Resource), make(chan *Resource)

	//启动StateMonitor
	status := StateMonitor(statusInterval)

	for i := 0; i < numPollers; i++ {
		go Poller(pending, complete, status)
	}

	go func() {
		for _, url := range urls {
			pending <- &Resource{url: url}
		}
	}()

	for r := range complete {
		go r.Sleep(pending)
	}



}

