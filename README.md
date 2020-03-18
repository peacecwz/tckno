# TC Identity Utils
Simple TC Identity number generator and validator (offline / online)

## Getting Started

Install package on your machine 

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
    tckNo := Generate()
	fmt.Printf()
}
```

### Validate TC Identity Number


```go
package main

import (
	"fmt"
	"github.com/peacecwz/tc-identity-utils"
)

func main(){
  result := Validate("29896722612")
  if result {
    fmt.Print("It's Valid")
  }
}
```

### Verify TC Identity Number

Coming soon

## License

This project is licensed under the MIT License

