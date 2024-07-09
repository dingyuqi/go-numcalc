package binning

import (
	"errors"
	"github.com/dingyuqi/go_numcalc/src"
	"math"
	"math/rand"
)

// ClusterWithKMeansInt32
/*
使用K-Means算法对数据进行聚类分组.
注意: 1. 输入数据长度不可为0. 2. k值不可大于数据点的数量.
*/
func (c *Calculator) ClusterWithKMeansInt32(data []int32) ([]int32, error) {
	n := len(data)
	if n == 0 {
		return nil, errors.New("输入数据长度为0")
	} else if int(c.kMeansK) > len(data) {
		return nil, errors.New("初始化k值大于数据点的数量")
	}
	return kMeans(src.Int32ToFloat64(data), c.kMeansK, c.kMeansMaxIter)
}

// ClusterWithKMeansFloat32
/*
使用K-Means算法对数据进行聚类分组.
注意: 1. 输入数据长度不可为0. 2. k值不可数据点的数量.
*/
func (c *Calculator) ClusterWithKMeansFloat32(data []float32) ([]int32, error) {
	n := len(data)
	if n == 0 {
		return nil, errors.New("输入数据长度为0")
	} else if int(c.kMeansK) > len(data) {
		return nil, errors.New("初始化k值大于数据点的数量")
	}
	return kMeans(src.Float32ToFloat64(data), c.kMeansK, c.kMeansMaxIter)
}

func kMeans(data []float64, k, maxIter int32) ([]int32, error) {
	n := len(data)
	// Initialize centroids randomly from data points
	centroids := make([]float64, k)
	perm := rand.Perm(n)
	for i := 0; i < int(k); i++ {
		centroids[i] = data[perm[i]]
	}

	labels := make([]int32, n)

	for iter := 0; iter < int(maxIter); iter++ {
		// Assign each data point to the closest centroid
		for i, v := range data {
			labels[i] = closestCentroid(v, centroids)
		}

		// Update centroids
		newCentroids := make([]float64, k)
		counts := make([]int, k)
		for i, label := range labels {
			newCentroids[label] += data[i]
			counts[label]++
		}
		for i := 0; i < int(k); i++ {
			if counts[i] > 0 {
				newCentroids[i] /= float64(counts[i])
			}
		}
		centroids = newCentroids
	}
	return labels, nil
}

func closestCentroid(point float64, centroids []float64) int32 {
	minIdx := 0
	minDist := math.Abs(point - centroids[0])
	for i := 1; i < len(centroids); i++ {
		dist := math.Abs(point - centroids[i])
		if dist < minDist {
			minDist = dist
			minIdx = i
		}
	}
	return int32(minIdx)
}
