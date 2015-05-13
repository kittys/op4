package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		t := new(Equipment)
		t.Name = fmt.Sprintf("Tester %v", i)
		t.Status = "running"
		t.Duration = time.Now().Sub(time.Date(2015, 5, 1, 0, 0, 0, 0, time.Local))
		fmt.Printf("%s\t%s\t%s\t%s\t%s\n", t.Name, t.Status, Round(t.Duration, time.Second), "47303", t.MTBF())
		time.Sleep(200 * time.Millisecond)
	}
}
