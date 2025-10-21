package main

import (
	"fmt"

	"github.com/raye17/subDemo/calc"
	"github.com/raye17/subDemo/greeting"
	"github.com/raye17/subDemo/times"
	"github.com/raye17/subDemo/utils"
)

func main() {
	s := greeting.Hello("sxy")
	fmt.Println(s)
	greeting.Hello2("sxy")
	fmt.Println(calc.Add(1, 2))
	fmt.Println(times.NowTime())
	fmt.Println(utils.NewUuid())
	fmt.Println(utils.NewUuidShort())
}
