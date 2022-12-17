<h1>
<p align="center">
  <span style="font-style:italic;font-family:Papyrus">gotime</span> 🏃‍♂️⌛
</p>
</h1>

<p align="center">
Convert string to go <code>time.time</code>
</p>


![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/emylincon/gotime?sort=semver&style=for-the-badge)
![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/emylincon/gotime/go.yml?branch=main&style=for-the-badge)
![GitHub](https://img.shields.io/github/license/emylincon/gotime?style=for-the-badge)
![Go Report Card](https://goreportcard.com/badge/github.com/emylincon/gotime?style=for-the-badge)
[![pre-commit](https://img.shields.io/badge/pre--commit-enabled-brightgreen?logo=pre-commit&logoColor=white&style=for-the-badge)](https://github.com/pre-commit/pre-commit)


# Examples 👇
👉 &nbsp; Example 1️⃣
```golang

time, err := gotime.Parse("2016 07 25")
if err != nil {
    fmt.Println(err) // handle error
}
fmt.Println(time) //2016-07-25 00:00:00 +0000 UTC
```

👉 &nbsp; Example 2️⃣
```golang

time, err := gotime.Parse("01 January 1970 00:00:00 GMT")
if err != nil {
    fmt.Println(err) // handle error
}
fmt.Println(time) //1970-01-01 00:00:00 +0000 GMT
```

👉 &nbsp; Group Examples 🗂️
```golang
times := []string{"August 7, 2014", "2011-10-10T14:48:00 GMT", "2011-10-10T14:48:00", "2011-10-10T14:48:00.000+09:00", "January 1, 09 00:00:00.00 GMT"}
for _, v := range times {
		fmt.Println(gotime.Parse(v))
	}

```
🔴 &nbsp; `OUTPUT` 🧾
```
2014-08-07 00:00:00 +0000 UTC, nil
2011-10-10 14:48:00 +0000 GMT, nil
2011-10-10 14:48:00 +0000 UTC, nil
2011-10-10 14:48:00 +9000 UTC, nil
2009-01-01 00:00:00 +0000 GMT, nil

```
# Official Documentation 📜
<a style="text-decoration:none" href="https://godoc.org/github.com/emylincon/gotime" target="_blank">
    <img src="https://godoc.org/github.com/emylincon/gotime?status.svg" alt="GoDoc" />
</a>
