# dbhwin32
Additional win32 wrappers for GO based on https://github.com/winlabs/gowin32


## Installation
```shell
go get -v github.com/DBHeise/dbhwin32
```

## Win API's implemented
- [MsiEnumProducts](https://msdn.microsoft.com/en-us/library/windows/desktop/aa370101(v=vs.85).aspx)
- [MsiEnumComponents](https://msdn.microsoft.com/en-us/library/windows/desktop/aa370097(v=vs.85).aspx)
- [MsiEnumPatches](https://msdn.microsoft.com/en-us/library/windows/desktop/aa370094(v=vs.85).aspx)
- [MsiEnumClients](https://msdn.microsoft.com/en-us/library/windows/desktop/aa370099(v=vs.85).aspx)

## Sample Code
```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/dbheise/dbhwin32"
)

func main() {
	products := dbhwin32.GetInstalledProducts()
	j, _ := json.Marshal(products)
	fmt.Printf("%s\n", j)
}
```