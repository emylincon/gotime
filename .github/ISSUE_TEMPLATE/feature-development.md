---
name: Feature Development
about: Feature opened by developer
title: "[Feature]: "
labels: feature
assignees: emylincon

---

**Describe feature to implement**
A clear and concise description of what you want to happen.

**Example**
Add example code if applicable.
```golang
import (
"github.com/emylincon/gotime"
)

func main(){
    time, err := gotime.Parse("2016 07 25")
    if err != nil {
        fmt.Println(err) // handle error
    }
    fmt.Println(time) // 2016-07-25 00:00:00 +0000 UTC
}
```
