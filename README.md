# TC Identity Utils
Simple TC Identity number generator and validator (offline / online)

## Getting Started

Install package on your machine 

```
$ go get -u github.com/peacecwz/tc-identity-utils
```

### Generate TC Identity Number

Generate TC Identity number. It's so simple

```
package main

import (
	"fmt"
	"github.com/peacecwz/tc-identity-utils"
)

func main(){
	fmt.Printf(Generate())
}
```

### Validate TC Identity Number


```
package main

import (
	"fmt"
	"github.com/peacecwz/tc-identity-utils"
)

func main(){
  result := Validate("29896722612")
  if result {
    fmt.Print("Is Valid")
  }
}
```

### Verify TC Identity Number

It's coming soon

## License

This project is licensed under the MIT License

