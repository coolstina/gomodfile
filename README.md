# gomodfile

Detect directory is Go project


## Installation

```shell
go get -u github.com/coolstina/gomodfile
```

## Example

```go
package main

import (
	"fmt"
	"log"

	"github.com/coolstina/gomodfile"
)

func main() {
	goProject, err := gomodfile.IsGoProject()
	if err != nil {
		log.Fatalf("can't get current directory is Go project")
	}

	if !goProject {
		log.Printf("get current directory is not Go project")
		return
	}

	fmt.Printf("get current directory is Go project")
}
```

```shell
get current directory is Go project
```