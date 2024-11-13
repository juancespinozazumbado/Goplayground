# Working with files in go

example

```go
package main

import (
	"files/files"
	"fmt"
)

func main() {

	res := files.OpenFile("path/to/your-file.extension")
	fmt.Printf("response: %+v", res)
}

```
