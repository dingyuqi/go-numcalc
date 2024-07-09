package main

// #cgo CXXFLAGS: -I../pkg/include -std=c++11
// #cgo LDFLAGS: -L../pkg/ -larmadillo
// #include "logTransform.hpp"
import "C"
import (
	"fmt"
	"time"
	"unsafe"
)

func main() {
	data := make([]float64, 0)
	for i := -10; i < 1000; i++ {
		data = append(data, float64(i))
	}
	fmt.Println(data)

	// 将 Go 切片转换为 C 数组
	dataPtr := (*C.double)(unsafe.Pointer(&data[0]))

	// 创建用于接收结果的 C 数组
	result := make([]float64, len(data))
	resultPtr := (*C.double)(unsafe.Pointer(&result[0]))
	st := time.Now()
	// 调用 C 的包装函数
	C.logTransform(dataPtr, resultPtr, C.int(len(data)))
	dur := time.Since(st)
	fmt.Println(result)
	// 打印结果
	fmt.Println("耗时: ", dur.Seconds(), "秒")
}
