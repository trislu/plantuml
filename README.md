# plantuml
PlantUML encoder&amp;decoder in go

# import

```golang
import (
    "github.com/trislu/plantuml"
)
```

# `Encode()` example

```golang
const text = `@startuml
Eve -> Bob : hello
@enduml`
var encoded string = plantuml.Encode(text)
var url string = "http://www.plantuml.com/plantuml/img/" + encoded
resp, err := http.Get(url)
```
The retrieved diagram:

![alt tag](http://www.plantuml.com/plantuml/img/SYWkIImgAStDuN8jIrNGjLDmoazIi5B8ICt9oUToICrBAStD0GG00F__)

# `Decode()` example

```golang
const cipher = "SYWkIImgAStDuN8jIrNGjLDmoazIi5B8ICt9oUToICrBAStD0GG00F__"
plain, err := plantuml.Decode(cipher)
log.Println(plain)
/*
@startuml
Eve -> Bob : hello
@enduml
*/
```

# test

```bash
go test -v
```