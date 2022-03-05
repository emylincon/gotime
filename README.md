# gotime
Convert string to go `time.time`


![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/emylincon/gotime?sort=semver&style=for-the-badge)
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/emylincon/gotime/Go?style=for-the-badge)
![GitHub](https://img.shields.io/github/license/emylincon/gotime?style=for-the-badge)
![Go Report Card](https://goreportcard.com/badge/github.com/emylincon/gotime?style=for-the-badge)
[![pre-commit](https://img.shields.io/badge/pre--commit-enabled-brightgreen?logo=pre-commit&logoColor=white&style=for-the-badge)](https://github.com/pre-commit/pre-commit)


# Examples
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
# Official Documentation
![godoc](https://godoc.org/github.com/emylincon/golist?status.svg)
