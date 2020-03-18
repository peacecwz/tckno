# TC Identity Utils
Simple TC Identity number generator and validator (offline)

## Getting Started

Install package on your project 

```sh
$ go get -u github.com/peacecwz/tc-identity-utils
```

### Generate TC Identity Number

Generate TC Identity number. It's so simple

```go
package main

import (
    "fmt"
    "github.com/peacecwz/tckno"
)

func main(){
    tckNo := tckno.Generate()
    fmt.Printf(tckNo)
}
```

### Validate TC Identity Number

```go
package main

import (
    "fmt"
    "github.com/peacecwz/tckno"
)

func main() {
    tckNo := "29896722612"
    result, _ := tckno.Validate(tckNo)
    if result {
        fmt.Printf("%s identity number valid", tckNo)
    }
}
```

## License

This project is licensed under the MIT License

