<h1>
<p align="center">
  <span style="font-style:italic;font-family:Papyrus">gotime</span> ๐โโ๏ธโ
</p>
</h1>

<p align="center">
Convert string to go <code>time.time</code>
</p>


![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/emylincon/gotime?sort=semver&style=for-the-badge)
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/emylincon/gotime/Go?style=for-the-badge)
![GitHub](https://img.shields.io/github/license/emylincon/gotime?style=for-the-badge)
![Go Report Card](https://goreportcard.com/badge/github.com/emylincon/gotime?style=for-the-badge)
[![pre-commit](https://img.shields.io/badge/pre--commit-enabled-brightgreen?logo=pre-commit&logoColor=white&style=for-the-badge)](https://github.com/pre-commit/pre-commit)


# Examples ๐
๐ &nbsp; Example 1๏ธโฃ
```golang

time, err := gotime.Parse("2016 07 25")
if err != nil {
    fmt.Println(err) // handle error
}
fmt.Println(time) //2016-07-25 00:00:00 +0000 UTC
```

๐ &nbsp; Example 2๏ธโฃ
```golang

time, err := gotime.Parse("01 January 1970 00:00:00 GMT")
if err != nil {
    fmt.Println(err) // handle error
}
fmt.Println(time) //1970-01-01 00:00:00 +0000 GMT
```

๐ &nbsp; Group Examples ๐๏ธ
```golang
times := []string{"August 7, 2014", "2011-10-10T14:48:00 GMT", "2011-10-10T14:48:00", "2011-10-10T14:48:00.000+09:00", "January 1, 09 00:00:00.00 GMT"}
for _, v := range times {
		fmt.Println(gotime.Parse(v))
	}

```
๐ด &nbsp; `OUTPUT` ๐งพ
```
2014-08-07 00:00:00 +0000 UTC, nil
2011-10-10 14:48:00 +0000 GMT, nil
2011-10-10 14:48:00 +0000 UTC, nil
2011-10-10 14:48:00 +9000 UTC, nil
2009-01-01 00:00:00 +0000 GMT, nil

```
# Official Documentation ๐
<a style="text-decoration:none" href="https://godoc.org/github.com/emylincon/gotime" target="_blank">
    <img src="https://godoc.org/github.com/emylincon/gotime?status.svg" alt="GoDoc" />
</a>
