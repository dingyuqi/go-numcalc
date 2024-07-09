package binning

import (
	"errors"
	"github.com/dingyuqi/go_numcalc/src"
)

type Calculator struct {
	width         int32 // 等宽分组参数: 总共分组的组别个数
	freq          int32 // 等频分组参数: 每组中包含的数值个数
	kMeansMaxIter int32 // 聚类分组参数: KMeans算法最大迭代次数
	kMeansK       int32 //聚类分组参数: KMeans算法k值
}

// NewCalculator 初始化数据分组基础参数. 参数不为4个的返回设置默认参数的对象.
// 初始化参数必须为正数.
func NewCalculator(optionalParam ...int32) (Calculator, error) {
	if len(optionalParam) == 4 {
		for _, i := range optionalParam {
			if i <= 0 {
				return Calculator{}, errors.New("初始化参数必须为正数")
			}
		}
		return Calculator{
			width:         optionalParam[0],
			freq:          optionalParam[1],
			kMeansMaxIter: optionalParam[2],
			kMeansK:       optionalParam[3],
		}, nil
	} else {
		return Calculator{
			width:         src.BinningWidth,
			freq:          src.BinningFreq,
			kMeansMaxIter: src.KMeansMaxIter,
			kMeansK:       src.KMeansK,
		}, nil
	}
}
