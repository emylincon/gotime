# gotime
convert string to go time.time

# Example
* Example 1
```golang

time, err := gotime.Parse("2016 07 25")
if err != nil {
    fmt.Println(err) // handle error
}
fmt.Println(time) //2016-07-25 00:00:00 +0000 UTC
```

* Example 2
```golang

time, err := gotime.Parse("01 January 1970 00:00:00 GMT")
if err != nil {
    fmt.Println(err) // handle error
}
fmt.Println(time) //1970-01-01 00:00:00 +0000 UTC
```
