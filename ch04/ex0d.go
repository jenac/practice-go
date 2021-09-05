package main

import (
	"fmt"
	"time"
)


func TimeUsage() {
	now:=time.Now()
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Date())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())
}

func TimeOperate() {
	start := time.Now()
	time.Sleep(1 * time.Second)
	fmt.Println(time.Now().Sub(start))
	fmt.Println(time.Now().Format("2006-01-02 15:04:05")) // has to be 2006-01-02 15:04:05
	fmt.Println(start.Round(time.Second))
	fmt.Println(start.Truncate(time.Second))
	startingTime:="1991-12-25 19:00:00"
	birthday, _:= time.ParseInLocation("2006-01-02 15:04:05", startingTime, time.Local)
	fmt.Println(birthday.String())
	
}

func TimeAdd() {
	now:=time.Now()
	oneDayBefore:=now.AddDate(0,0,-1)
	fmt.Println(now.String(), oneDayBefore.String())
}

func Timer() {
	ticker:= time.NewTicker(time.Second)
	defer ticker.Stop()
	done:=make(chan bool)
	go func ()  {
		time.Sleep(10*time.Second)
		done <- true
	}()
	for {
		select {
		case <- done:
			fmt.Println("Done!")
			return
		case t:= <- ticker.C:
			fmt.Println("Current time:", t)
		}
	}

}

func main() {
	TimeUsage()
	fmt.Println("----------")
	TimeOperate()
	fmt.Println("----------")
	TimeAdd()
	fmt.Println("----------")
	Timer()
}