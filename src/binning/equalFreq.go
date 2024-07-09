package binning

import (
	"errors"
	"github.com/dingyuqi/go_numcalc/src"
	"math"
	"sort"
)

// EqualFrequencyBinningInt32
/*
对数据进行等频分组, 每组数据个数为Calculator.freq
返回每个数据点的区间索引.
*/
func (c *Calculator) EqualFrequencyBinningInt32(data []int32) ([]int32, error) {
	if len(data) == 0 {
		return nil, errors.New("输入数据长度为0")
	}

	// 复制并对数据进行排序
	sortedData := make([]int32, len(data))
	copy(sortedData, data)
	sort.Slice(sortedData, func(i, j int) bool {
		return sortedData[i] < sortedData[j]
	})
	binSize := math.Ceil(float64(int32(len(data)) / c.freq))
	binningResult := make([]int32, len(data))
	for i, v := range data {
		pos := sort.SearchInts(src.Int32ToInt(sortedData), int(v))
		binningResult[i] = int32(math.Floor(float64(pos) / binSize))
	}
	return binningResult, nil
}

// EqualFrequencyBinningFloat32
/*
对数据进行等频分组, 每组数据个数为Calculator.freq
返回每个数据点的区间索引.
*/
func (c *Calculator) EqualFrequencyBinningFloat32(data []float32) ([]int32, error) {
	if len(data) == 0 {
		return nil, errors.New("输入数据长度为0")
	}
	// 复制并对数据进行排序
	sortedData := make([]float32, len(data))
	copy(sortedData, data)
	sort.Slice(sortedData, func(i, j int) bool {
		return sortedData[i] < sortedData[j]
	})
	binSize := math.Ceil(float64(int32(len(data)) / c.freq))
	binningResult := make([]int32, len(data))
	for i, v := range data {
		pos := sort.SearchFloat64s(src.Float32ToFloat64(sortedData), float64(v))
		binningResult[i] = int32(math.Floor(float64(pos) / binSize))
	}
	return binningResult, nil
}
