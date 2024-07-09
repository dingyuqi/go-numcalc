package outlier

import (
	"errors"
	"github.com/dingyuqi/go_numcalc/src"
	"github.com/montanaflynn/stats"
	"gonum.org/v1/gonum/stat"
	"math"
	"sort"
)

type Calculator struct {
	threshold float32 // 判断离群值的阈值
}

// NewCalculator 初始化离群值判断阈值, 参数必须为正数.
func NewCalculator(threshold float32) (Calculator, error) {
	if threshold <= 0 {
		return Calculator{}, errors.New("参数必须为正数")
	}
	return Calculator{threshold: threshold}, nil
}

// DetectOutliersByStandardDeviationInt32
/*
根据标准差判断离群值.
计算数据点于平均值的标准差之间的差异, 超过阈值时视为离群值.
切片下标对应为true则表示该值为离群值, 为false则为正常值.
注意: data切片长度不可小于2.
*/
func (c *Calculator) DetectOutliersByStandardDeviationInt32(data []int32) ([]bool, error) {
	if len(data) < 2 {
		return nil, errors.New("数据量太少, 数量少于2时无法计算离群值")
	}
	float64Data := src.Int32ToFloat64(data)
	avg := float32(stat.Mean(float64Data, nil))
	std := float32(stat.StdDev(float64Data, nil))
	outliers := make([]bool, len(data))
	for i, v := range data {
		diff := float32(math.Abs(float64(float32(v) - avg)))
		outliers[i] = diff > c.threshold*std
	}
	return outliers, nil
}

// DetectOutliersByStandardDeviationFloat32
/*
根据标准差判断离群值.
计算数据点于平均值的标准差之间的差异, 超过阈值时视为离群值.
切片下标对应为true则表示该值为离群值, 为false则为正常值.
注意: data切片长度不可小于2.
*/
func (c *Calculator) DetectOutliersByStandardDeviationFloat32(data []float32) ([]bool, error) {
	if len(data) < 2 {
		return nil, errors.New("数据量太少, 数量少于2时无法计算离群值")
	}
	float64Data := src.Float32ToFloat64(data)
	avg := float32(stat.Mean(float64Data, nil))
	std := float32(stat.StdDev(float64Data, nil))
	outliers := make([]bool, len(data))
	for i, v := range data {
		diff := float32(math.Abs(float64(v - avg)))
		outliers[i] = diff > c.threshold*std
	}
	return outliers, nil
}

// DetectOutliersByQuartilesInt32
/*
根据四分位数判断离群值.
计算数据的四分位数和异常值范围，将超出上下边界的值视为离群值.
切片下标对应为true则表示该值为离群值, 为false则为正常值.
注意: data切片长度不可小于2.
*/
func (c *Calculator) DetectOutliersByQuartilesInt32(data []int32) ([]bool, error) {
	if len(data) < 2 {
		return nil, errors.New("数据量太少, 数量少于2时无法计算离群值")
	}
	floatData := make([]float64, len(data))
	for i, v := range data {
		floatData[i] = float64(v)
	}
	return DetectOutliersByQuartiles(floatData)
}

// DetectOutliersByQuartilesFloat32
/*
根据四分位数判断离群值.
计算数据的四分位数和异常值范围，将超出上下边界的值视为离群值.
切片下标对应为true则表示该值为离群值, 为false则为正常值.
注意: data切片长度不可小于2.
*/
func (c *Calculator) DetectOutliersByQuartilesFloat32(data []float32) ([]bool, error) {
	if len(data) < 2 {
		return nil, errors.New("数据量太少, 数量少于2时无法计算离群值")
	}
	floatData := make([]float64, len(data))
	for i, v := range data {
		floatData[i] = float64(v)
	}
	return DetectOutliersByQuartiles(floatData)
}

func DetectOutliersByQuartiles(floatData []float64) ([]bool, error) {
	sort.Float64s(floatData)

	q1, err := stats.Percentile(floatData, 25)
	if err != nil {
		return nil, err
	}
	q3, err := stats.Percentile(floatData, 75)
	if err != nil {
		return nil, err
	}

	iqr := q3 - q1
	lowerBound := q1 - 1.5*iqr
	upperBound := q3 + 1.5*iqr

	outliers := make([]bool, len(floatData))
	for i, v := range floatData {
		outliers[i] = v < lowerBound || v > upperBound
	}

	return outliers, nil
}
