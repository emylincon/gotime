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
fmt.Println(time) //1970-01-01 00:00:00 +0000 GMT
```

* Group Examples
```golang
times := []string{"2011-10-10T14:48:00 GMT", "2011-10-10T14:48:00"}
for _, v := range times {
		fmt.Println(gotime.Parse(v))
	}

```
`OUTPUT`
```
2011-10-10 14:48:00 +0000 GMT, nil
2011-10-10 14:48:00 +0000 UTC, nil

```
# Official Documentation
<a style="text-decoration:none" href="https://godoc.org/github.com/emylincon/gotime" target="_blank">
    <img src="https://godoc.org/github.com/emylincon/gotime?status.svg" alt="GoDoc" />
</a>
