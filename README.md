# gotime
convert string to go time.time

# Example
```golang

time, err := gotime.Parse("2016 07 25")
if err != nil {
    fmt.Println(err) // handle error
}
fmt.Println(time) //2016-07-25 00:00:00 +0000 UTC
```
