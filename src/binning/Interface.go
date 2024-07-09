package binning

// ClusterWithKMeans K-Means聚类分组接口
type ClusterWithKMeans interface {
	ClusterWithKMeansInt32([]int32) ([]int32, error)
	ClusterWithKMeansFloat32([]float32) ([]int32, error)
}

// EqualWidthBinning 等宽分组接口
type EqualWidthBinning interface {
	EqualWidthBinningInt32([]int32) ([]int32, error)
	EqualWidthBinningFloat32([]float32) ([]int32, error)
}

// EqualFrequencyBinning 等频分组接口
type EqualFrequencyBinning interface {
	EqualFrequencyBinningInt32([]int32) ([]int32, error)
	EqualFrequencyBinningFloat32([]float32) ([]int32, error)
}
