# 数值分类

**其他语言版本: [English](./README.md)**

本包主要封装数值分类相关接口.

## 结构体

```go
package binning

type Calculator struct {
	width         int32 // 等宽分组参数: 总共分组的组别个数
	freq          int32 // 等频分组参数: 每组中包含的数值个数
	kMeansMaxIter int32 // 聚类分组参数: KMeans算法最大迭代次数
	kMeansK       int32 // 聚类分组参数: KMeans算法k值
}

```

可以调用 ```func NewCalculator(optionalParam ...int32) Calculator``` 获取对象.
动态参数如果长度等于4则会以 width, freq, kMeansMaxIter, kMeansK 的顺序进行初始化, 如果长度不为4则自动初始化为默认参数.
默认参数信息详见文件 ```src/static.go```

## 接口

### 1.ClusterWithKMeansInt32([]int32) ([]int32, error)

使用K-Means算法对数据进行聚类分组. 函数返回的切片下标与输入切片对应, 切片元素为该索引的值的分组组号.
> [!IMPORTANT]
>  1. 输入数据长度不可为0.
>  2. k值不可大于数据点的数量.

### 2. ClusterWithKMeansFloat32([]float32) ([]int32, error)

使用K-Means算法对数据进行聚类分组. 函数返回的切片下标与输入切片对应, 切片元素为该索引的值的分组组号.
> [!IMPORTANT]
>   1. 输入数据长度不可为0.
>   2. k值不可大于数据点的数量.

### 3. EqualWidthBinningInt32([]int32) ([]int32, error)

对数据进行等宽分组, 分组数量为Calculator.width. 函数返回每个数据点的区间索引.
> [!IMPORTANT]
> 输入的数据长度不得为0

### 4. EqualWidthBinningFloat32([]float32) ([]int32, error)

对数据进行等宽分组, 分组数量为Calculator.width. 函数返回每个数据点的区间索引.

> [!IMPORTANT]
> 输入的数据长度不得为0

### 5. EqualFrequencyBinningInt32([]int32) ([]int32, error)

对数据进行等频分组, 每组数据个数为Calculator.freq. 函数返回每个数据点的区间索引.

### 6. EqualFrequencyBinningFloat32([]float32) ([]int32, error)

对数据进行等频分组, 每组数据个数为Calculator.freq. 函数返回每个数据点的区间索引.

   

