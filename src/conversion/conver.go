package conversion

import (
	"errors"
	"fmt"
	"github.com/dingyuqi/go_numcalc/src"
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/stat"
	"math"
)

type Calculator struct {
	logBase float32 //对数计算的底数, 默认为自然数e
}

// NewCalculator optionalParam仅接受一个参数, 用于初始化对数计算中的底数
func NewCalculator(optionalParam ...float32) (Calculator, error) {
	if len(optionalParam) == 0 {
		return Calculator{
			logBase: math.E,
		}, nil
	} else {
		base := optionalParam[0]
		if base <= 0 || base == 1 {
			return Calculator{}, errors.New("计算对数的底数必须为正数且不能为1")
		}
		return Calculator{
			logBase: optionalParam[0],
		}, nil
	}
}

//StandardizationInt32
/*
最大最小值标准化.
将int32类型的一组数据处理过后数据均落在［0，1］区间内, 处理后输出不再是整数.
计算公式为: normalized=(maxValue−minValue)/(value−minValue)
注意: 无法对零范围(最大最小值相等)的数组进行规范化.
*/
func (c *Calculator) StandardizationInt32(data []int32) ([]float32, error) {
	float64Data := src.Int32ToFloat64(data)
	minVal, maxVal := int32(floats.Min(float64Data)), int32(floats.Max(float64Data))
	if maxVal == minVal {
		return nil, errors.New("无法对零范围(最大最小值相等)的数组进行规范化")
	}
	var normalized []float32
	for _, v := range data {
		normalized = append(normalized, float32(v-minVal)/float32(maxVal-minVal))
	}

	return normalized, nil
}

//StandardizationFloat32
/*
最大最小值标准化.
将float32类型的一组数据处理过后数据均落在［0，1］区间内.
计算公式为: normalized=(maxValue−minValue)/(value−minValue)
注意: 无法对零范围(最大最小值相等)的数组进行规范化.
*/
func (c *Calculator) StandardizationFloat32(data []float32) ([]float32, error) {
	float64Data := src.Float32ToFloat64(data)
	minVal, maxVal := float32(floats.Min(float64Data)), float32(floats.Max(float64Data))
	if maxVal == minVal {
		return nil, errors.New("无法对零范围(最大最小值相等)的数组进行规范化")
	}
	var normalized []float32
	for _, v := range data {
		normalized = append(normalized, (v-minVal)/(maxVal-minVal))
	}

	return normalized, nil
}

// ZScoreInt32
/*
Z-Score标准化.
计算公式为: Z=(X−μ)/σ
X 是原始数据, μ 是数据的平均值, σ 是数据的标准差
注意: 无法对零范围(最大最小值相等)的数组进行规范化.
*/
func (c *Calculator) ZScoreInt32(data []int32) ([]float32, error) {
	float64Data := src.Int32ToFloat64(data)
	m := float32(stat.Mean(float64Data, nil))
	s := float32(stat.StdDev(float64Data, nil))

	if s == 0 {
		return nil, fmt.Errorf("无法对零范围(最大最小值相等)的数组进行Z-Score规范化")
	}

	var normalized []float32
	for _, v := range data {
		normalized = append(normalized, (float32(v)-m)/s)
	}

	return normalized, nil
}

// ZScoreFloat32
/*
Z-Score标准化.
计算公式为: Z=(X−μ)/σ
X 是原始数据, μ 是数据的平均值, σ 是数据的标准差
注意: 无法对零范围(最大最小值相等)的数组进行规范化.
*/
func (c *Calculator) ZScoreFloat32(data []float32) ([]float32, error) {
	float64Data := src.Float32ToFloat64(data)
	m := float32(stat.Mean(float64Data, nil))
	s := float32(stat.StdDev(float64Data, nil))

	if s == 0 {
		return nil, fmt.Errorf("无法对零范围(最大最小值相等)的数组进行Z-Score规范化")
	}

	var normalized []float32
	for _, v := range data {
		normalized = append(normalized, (v-m)/s)
	}

	return normalized, nil
}

// LogInt32
/*
计算对数.
计算公式为: log(base, x), base默认为自然数e.
注意: 需保证输入参数均为正数, 否则:
Log(+Inf) = +Inf
Log(0) = -Inf
Log(x < 0) = NaN
Log(NaN) = NaN
*/
func (c *Calculator) LogInt32(data []int32) ([]float32, error) {
	var log []float32
	for _, v := range data {
		log = append(log, float32(math.Log(float64(v))/math.Log(float64(c.logBase))))
	}
	return log, nil
}

// LogFloat32
/*
计算对数.
计算公式为: log(base, x), base默认为自然数e.
注意: 需保证输入参数均为正数, 否则:
Log(+Inf) = +Inf
Log(0) = -Inf
Log(x < 0) = NaN
Log(NaN) = NaN
*/
func (c *Calculator) LogFloat32(data []float32) ([]float32, error) {
	var log []float32
	for _, v := range data {
		log = append(log, float32(math.Log(float64(v))/math.Log(float64(c.logBase))))
	}
	return log, nil
}

// SquareInt32
/*
计算平方根.
计算公式为: √x
注意: 需保证输入参数均为非负实数, 否则:
Sqrt(+Inf) = +Inf
Sqrt(±0) = ±0
Sqrt(x < 0) = NaN
Sqrt(NaN) = NaN
*/
func (c *Calculator) SquareInt32(data []int32) ([]float32, error) {
	var sqrt []float32
	for _, v := range data {
		sqrt = append(sqrt, float32(math.Sqrt(float64(v))))
	}
	return sqrt, nil
}

// SquareFloat32
/*
计算平方根.
计算公式为: √x
注意: 需保证输入参数均为非负实数, 否则:
Sqrt(+Inf) = +Inf
Sqrt(±0) = ±0
Sqrt(x < 0) = NaN
Sqrt(NaN) = NaN
*/
func (c *Calculator) SquareFloat32(data []float32) ([]float32, error) {
	var sqrt []float32
	for _, v := range data {
		sqrt = append(sqrt, float32(math.Sqrt(float64(v))))
	}
	return sqrt, nil
}
