# 数值转换

本包主要封装数值转换相关接口.

## 结构体

```go
package conversion

type Calculator struct {
	logBase float32 //对数计算的底数, 默认为自然数e
}

```

可以调用 ```func NewCalculator(optionalParam ...float32) Calculator ``` 获取对象. 动态参数如果存在则设置为logBase, 若不存在则默认设置为自然数e.

## 接口

1. **StandardizationInt32([]int32) ([]float32, error)**
    ```text
    最大最小值标准化.
    将int32类型的一组数据处理过后数据均落在［0，1］区间内, 处理后输出不再是整数.
    计算公式为: normalized=(maxValue−minValue)/(value−minValue)
    注意: 无法对零范围(最大最小值相等)的数组进行规范化.
    ```

2. **StandardizationFloat32([]float32) ([]float32, error)**
     ```text
    最大最小值标准化.
    将float32类型的一组数据处理过后数据均落在［0，1］区间内.
    计算公式为: normalized=(maxValue−minValue)/(value−minValue)
    注意: 无法对零范围(最大最小值相等)的数组进行规范化.
    ```
3. **ZScoreInt32([]int32) ([]float32, error)**
      ```text
   Z-Score标准化.
   计算公式为: Z=(X−μ)/σ
   X 是原始数据, μ 是数据的平均值, σ 是数据的标准差
   注意: 无法对零范围(最大最小值相等)的数组进行规范化.
    ```
4. **ZScoreFloat32([]float32) ([]float32, error)**
    ```text
    Z-Score标准化.
   计算公式为: Z=(X−μ)/σ
   X 是原始数据, μ 是数据的平均值, σ 是数据的标准差
   注意: 无法对零范围(最大最小值相等)的数组进行规范化.
    ```
5. **LogInt32([]int32) ([]float32, error)**
      ```text
   计算对数.
   计算公式为: log(base, x), base默认为自然数e.
   注意: 需保证输入参数均为正数, 否则:
        Log(+Inf) = +Inf
        Log(0) = -Inf
        Log(x < 0) = NaN
        Log(NaN) = NaN
    ```
6. **LogFloat32([]float32) ([]float32, error)**
     ```text
   计算对数.
   计算公式为: log(base, x), base默认为自然数e.
   注意: 需保证输入参数均为正数, 否则:
        Log(+Inf) = +Inf
        Log(0) = -Inf
        Log(x < 0) = NaN
        Log(NaN) = NaN
    ```
7. **SquareInt32([]int32) ([]float32, error)**
      ```text
   计算平方根.
   计算公式为: √x
   注意: 需保证输入参数均为非负实数, 否则:
        Sqrt(+Inf) = +Inf
        Sqrt(±0) = ±0
        Sqrt(x < 0) = NaN
        Sqrt(NaN) = NaN
    ```
8. **SquareFloat32([]float32) ([]float32, error)**
      ```text
   计算平方根.
   计算公式为: √x
   注意: 需保证输入参数均为非负实数, 否则:
        Sqrt(+Inf) = +Inf
        Sqrt(±0) = ±0
        Sqrt(x < 0) = NaN
        Sqrt(NaN) = NaN
    ```