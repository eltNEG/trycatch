# Trycatch

`trycatch` functionality in golang. Inspired by [lo](https://github.com/samber/lo)

### Install

```bash
go get github.com/eltneg/trycatch
```

### Usage

```go
package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/eltneg/trycatch"
)

type Err struct {
	Message string
	Code    string
}

func main() {
	go func() {
		for {
			fmt.Println("hello")
			time.Sleep(10 * time.Second)
		}
	}()

	go func() {
		r := Do1()
		if r.HasError() {
			fmt.Printf("Error in Do1. message: [%s]; code: [%s], value: [%s]\n", r.Err, r.Err.Code, r.Val)
			return
		}
		fmt.Println("Do is successful: ", r.Val)
	}()

	for {

	}

}

func OpenWithError(v int) (int, error) {
	if v == 1 {
		return 1, nil
	}
	return 0, errors.New("open error")
}

func handleOpenError(r *trycatch.Result[string, Err], v int, err error) {
	er := &Err{
		Message: err.Error(),
		Code:    "OPEN_HANDLE_ERROR",
	}
	r.Err = er
	return
}

func Do1() (result *trycatch.Result[string, Err]) {
	result = trycatch.Init(result)
	defer trycatch.Catch(result)

	v := trycatch.Try(result, handleOpenError)(OpenWithError(1))
	fmt.Println("No error occured with value: ", v)
	result.Val = "value1" // This value is preseved
	v = trycatch.Try(result, handleOpenError)(OpenWithError(2))
	fmt.Print("This should not be printed")
	panic("Do1 function is down: " + fmt.Sprint(v))
}

```

### Spec

specification

### Contributing

- Holla on twitter @eltneg
- Fork the project
- Fix open issues or request new features
- Raise pull request against the master branch

### References

- https://deedlefake.com/experimenting-with-error-handling-via-generics-in-go.html
- https://www.reddit.com/r/golang/comments/hefpj8/experimenting_with_error_handling_via_generics/
- https://dev.to/dean/go-2-draft-error-handling-3loo
- https://melvinbiamont.medium.com/how-generics-can-change-error-handling-in-go-34f47347925a
- https://github.com/samber/lo

### License

Copyright © 2022 Hakeem Orewole.

This project is MIT licensed.
