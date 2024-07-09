# Binning
**Read this in other languages: [简体中文](./README.zh-CN.md)**

This package mainly encapsulates interfaces related to numerical classification.

## Structure

```go
package binning

type Calculator struct {
	width         int32 // Equal width grouping parameter: the total number of groups
	freq          int32 // Equal frequency grouping parameter: the number of values ​​contained in each group
	kMeansMaxIter int32 // Clustering grouping parameter: the maximum number of iterations of the KMeans algorithm
	kMeansK       int32 // Clustering grouping parameter: the k value of the KMeans algorithm
}

```

You can call ```func NewCalculator(optionalParam ...int32) Calculator``` to get the object.
If the length of the dynamic parameter is equal to 4, it will be initialized in the order of width, freq, kMeansMaxIter,
kMeansK. If the length is not 4, it will be automatically initialized to the default parameter.
For details about the default parameters, see the file ```src/static.go```

## Interface

### 1.ClusterWithKMeansInt32([]int32) ([]int32, error)

Use the K-Means algorithm to cluster and group data. The slice index returned by the function corresponds to the input
slice, and the slice element is the group number of the value of the index.
> [!IMPORTANT]
> 1. The input data length cannot be 0.
> 2. The k value cannot be greater than the number of data points.

### 2. ClusterWithKMeansFloat32([]float32) ([]int32, error)

Use the K-Means algorithm to cluster and group data. The slice index returned by the function corresponds to the input
slice, and the slice element is the group number of the value of the index.
> [!IMPORTANT]
> 1. The input data length cannot be 0.
> 2. The k value cannot be greater than the number of data points.

### 3. EqualWidthBinningInt32([]int32) ([]int32, error)

Group the data with equal width, the number of groups is Calculator.width. The function returns the interval index of
each data point.

> [!IMPORTANT]
> The length of the input data cannot be 0

### 4. EqualWidthBinningFloat32([]float32) ([]int32, error)

Group the data with equal width, the number of groups is Calculator.width. The function returns the interval index of
each data point.

> [!IMPORTANT]
> The length of the input data cannot be 0

### 5. EqualFrequencyBinningInt32([]int32) ([]int32, error)

Group the data with equal frequency, the number of data in each group is Calculator.freq. The function returns the
interval index of each data point.

### 6. EqualFrequencyBinningFloat32([]float32) ([]int32, error)

Group the data with equal frequency, The number of data points in each group is Calculator.freq. The function returns
the interval index of each data point.