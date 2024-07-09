# 数值转换
**其他语言版本: [English](./README.md)**
本包主要封装离群值判断相关接口.

## 结构体

```go
package outlier

type Calculator struct {
	threshold float32 // 判断离群值的阈值
}
```
可以调用 ```NewCalculator(threshold float32) Calculator``` 获取对象.

## 接口

### 1. DetectOutliersByStandardDeviationInt32([]int32) ([]bool, error)
根据标准差判断离群值.

计算数据点于平均值的标准差之间的差异, 超过阈值时视为离群值.
切片下标对应为true则表示该值为离群值, 为false则为正常值.

> [!IMPORTANT]
> 入参切片长度不可小于2.

### 2. DetectOutliersByStandardDeviationFloat32([]float32) ([]bool, error)
根据标准差判断离群值.

计算数据点于平均值的标准差之间的差异, 超过阈值时视为离群值.
切片下标对应为true则表示该值为离群值, 为false则为正常值.

> [!IMPORTANT]
> 入参切片长度不可小于2.

### 3. DetectOutliersByQuartilesInt32([]int32) ([]bool, error)

根据四分位数判断离群值.

计算数据的四分位数和异常值范围，将超出上下边界的值视为离群值.
切片下标对应为true则表示该值为离群值, 为false则为正常值.
> [!IMPORTANT]
> 入参切片长度不可小于2.


### 4. DetectOutliersByQuartilesFloat32([]float32) ([]bool, error)
根据四分位数判断离群值.

计算数据的四分位数和异常值范围，将超出上下边界的值视为离群值.
切片下标对应为true则表示该值为离群值, 为false则为正常值.
> [!IMPORTANT]
> 入参切片长度不可小于2.
