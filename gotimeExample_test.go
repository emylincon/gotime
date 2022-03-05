package gotime_test

import (
	"fmt"

	"github.com/emylincon/gotime"
)

func Example() {
	time, err := gotime.Parse("2016 07 25")
	if err != nil {
		fmt.Println(err) // handle error
	}
	fmt.Println(time)
	time, err = gotime.Parse("25 07 2016")
	if err != nil {
		fmt.Println(err) // handle error
	}
	fmt.Println(time)

	//Output:
	//2016-07-25 00:00:00 +0000 UTC
	//2016-07-25 00:00:00 +0000 UTC

}
