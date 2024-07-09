# go_numcalc

**Read this in other languages: [简体中文](./README.zh-CN.md)**

The go_numcalc package is a package developed in Go language. Its main function is to perform basic numerical processing such as data conversion, data grouping, and data smoothing for numerical type data.
All functions support only two data types: Int32 and Float32.

> [!TIP]
> This package is written in Go language and contains part of C++, that is, the cgo part

## Installation
Use the `go get` to install go_numcalc.
```shell
go get github.com/dingyuqi/go_numcalc
```

## Dependencies
There are two main external libraries used in the NumCalc package:

1. Go language: [Gonum](https://www.gonum.org/)
2. C++ language: [Armadillo](https://arma.sourceforge.net/docs.html#top)

Armadillo does not need to be installed and is called in Cgo as a static library.

## Usage
The following is a simple example of the `LogInt32()` method in conversion. The usage of other functions is consistent with this Demo.
1. Call the `NewCalculator()` method in conversion to initialize a numerical conversion object.
2. Call the `LogInt32()` method. This method will return the logarithm calculation result of the corresponding subscript of the slice `data`.


```go
package main

import (
	"github.com/dingyuqi/go_numcalc/src/conversion"
	"log"
)

func main() {
	data := []int32{1, 2, 3, 4, 5}
	c, err := conversion.NewCalculator()
	result, err := c.LogInt32(data)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("result is: ", result)
}
```
> [!TIP]
> This test code is in `example/example_test.go` and can be run directly.

## Project structure

1. **src**: src code
    - binning: data binning functions
    - conversion: data conversion functions
    - outlier: data outlier functions
    - smoothness: data smoothness functions
2. **pkg**: Used to place static compilation files and related header files of the C++ language Armadillo library.
3. **example**: Used to place Cgo calling samples and Go language calling samples.
4. **test**: Unit test data and test cases.

> [!TIP]
> If you want to test Cgo, you can directly execute the following command in the example folder:
>   ```shell
>    go build ./example.exe
>   ```
> `example_test.go` contains a pure Go language implementation of the same function (logarithmic calculation), used to compare the calculation speed of Cgo.

## License
[MIT](https://choosealicense.com/licenses/mit/) License