package binning

import (
	"errors"
	"github.com/dingyuqi/go_numcalc/src"
	"gonum.org/v1/gonum/floats"
)

// EqualWidthBinningInt32
/*
对数据进行等宽分组, 分组数量为Calculator.width
返回每个数据点的区间索引.
注意: 输入的数据长度不得为0
*/
func (c *Calculator) EqualWidthBinningInt32(data []int32) ([]int32, error) {
	if len(data) == 0 {
		return nil, errors.New("输入数据长度为0")
	}

	// 查找数据的最小值和最大值
	float64Data := src.Int32ToFloat64(data)
	minVal, maxVal := int32(floats.Min(float64Data)), int32(floats.Max(float64Data))

	// 计算每个区间的宽度
	binWidth := float32(maxVal-minVal) / float32(c.width)

	// 分配每个数据点到相应的区间
	binningResult := make([]int32, len(data))
	for i, v := range data {
		if v == maxVal {
			binningResult[i] = c.width - 1 // 将最大值分配到最后一个区间
		} else {
			binningResult[i] = int32(float32(v-minVal) / binWidth)
		}
	}
	return binningResult, nil
}

// EqualWidthBinningFloat32
/*
对数据进行等宽分组, 分组数量为Calculator.width
返回每个数据点的区间索引.
注意: 输入的数据长度不得为0
*/
func (c *Calculator) EqualWidthBinningFloat32(data []float32) ([]int32, error) {
	if len(data) == 0 {
		return nil, errors.New("输入数据长度为0")
	}

	// 查找数据的最小值和最大值
	float64Data := src.Float32ToFloat64(data)
	minVal, maxVal := float32(floats.Min(float64Data)), float32(floats.Max(float64Data))
	// 计算每个区间的宽度
	binWidth := (maxVal - minVal) / float32(c.width)

	// 分配每个数据点到相应的区间
	binningResult := make([]int32, len(data))
	for i, v := range data {
		if v == maxVal {
			binningResult[i] = c.width - 1 // 将最大值分配到最后一个区间
		} else {
			binningResult[i] = int32((v - minVal) / binWidth)
		}
	}

	return binningResult, nil
}
