# Outlier
**Read this in other languages: [简体中文](./README.zh-CN.md)**
This package mainly encapsulates outlier judgment related interfaces.

## Structure

```go
package outlier

type Calculator struct {
	threshold float32 // Threshold for judging outliers
}
```
You can call ```NewCalculator(threshold float32) Calculator``` to get the object.

## Interface

### 1. DetectOutliersByStandardDeviationInt32([]int32) ([]bool, error)
Judging outliers based on standard deviation.

Calculate the difference between the standard deviation of the data point and the mean value. If it exceeds the threshold, it is considered an outlier.

If the corresponding slice subscript is true, it means that the value is an outlier, and if it is false, it is a normal value.

> [!IMPORTANT]
> The length of the input slice cannot be less than 2.

### 2. DetectOutliersByStandardDeviationFloat32([]float32) ([]bool, error)
Determine outliers based on standard deviation.

Calculate the difference between the standard deviation of the data point and the mean value, and consider it an outlier when it exceeds the threshold.
If the slice subscript corresponds to true, it means that the value is an outlier, and if it is false, it is a normal value.

> [!IMPORTANT]
> The length of the input slice cannot be less than 2.

### 3. DetectOutliersByQuartilesInt32([]int32) ([]bool, error)

Determine outliers based on quartiles.

Calculate the quartiles and outlier range of the data, and consider values ​​beyond the upper and lower boundaries as outliers.
If the slice subscript corresponds to true, it means that the value is an outlier, and if it is false, it is a normal value.

> [!IMPORTANT]
> The length of the input slice cannot be less than 2.

### 4. DetectOutliersByQuartilesFloat32([]float32) ([]bool, error)
Judging outliers based on quartiles.

Calculate the quartiles and outlier range of the data, and treat values ​​beyond the upper and lower boundaries as outliers.
If the slice subscript is true, it means that the value is an outlier, and if it is false, it is a normal value.
> [!IMPORTANT]
> The length of the input slice cannot be less than 2.