# useragents
[![Go Reference](https://pkg.go.dev/badge/github.com/gebleksengek/useragents.svg)](https://pkg.go.dev/github.com/gebleksengek/useragents)
 
Generate random user-agent



## Installation
```
go get -u github.com/gebleksengek/useragents
```

## Usage
### Generate a random user-agent
```go
package main

import (
    "fmt"
    "github.com/gebleksengek/useragents"
)

func main() {
   ua := useragents.Random() 
   fmt.Println(ua)
}
```

### Generate a random user-agent with latest version
```go
package main

import (
    "fmt"
    "github.com/gebleksengek/useragents"
)

func main() {
   ua := useragents.RandomLatest() 
   fmt.Println(ua)
}
```

### Generate a random firefox user-agent
```go
package main

import (
    "fmt"
    "github.com/gebleksengek/useragents"
)

func main() {
   ua := useragents.Firefox() 
   fmt.Println(ua)
}
```