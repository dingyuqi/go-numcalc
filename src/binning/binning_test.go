package binning

import (
	"github.com/dingyuqi/go_numcalc/test"
	"testing"
)

// todo: 考虑切片长度问题
func TestNewCalculator(t *testing.T) {
	t.Log(NewCalculator())
	t.Log(NewCalculator(10))
	t.Log(NewCalculator(10, 20, 30))
}

func TestCalculator_ClusterWithKMeansInt32(t *testing.T) {
	c, _ := NewCalculator()
	t.Log(c.ClusterWithKMeansInt32(test.DataPositiveInt32()))
	t.Log(c.ClusterWithKMeansInt32(test.DataSameInt32()))
	t.Log(c.ClusterWithKMeansInt32(test.DataSpecialInt32()))
}

func TestCalculator_ClusterWithKMeansFloat32(t *testing.T) {
	c, _ := NewCalculator()
	t.Log(c.ClusterWithKMeansFloat32(test.DataPositiveFloat32()))
	t.Log(c.ClusterWithKMeansFloat32(test.DataSameFloat32()))
	t.Log(c.ClusterWithKMeansFloat32(test.DataSpecialFloat32()))
}
func TestCalculator_EqualFrequencyBinningInt32(t *testing.T) {
	c, _ := NewCalculator()
	t.Log(c.EqualFrequencyBinningInt32(test.DataPositiveInt32()))
	t.Log(c.EqualFrequencyBinningInt32(test.DataSameInt32()))
	t.Log(c.EqualFrequencyBinningInt32(test.DataSpecialInt32()))
}
func TestCalculator_EqualFrequencyBinningFloat32(t *testing.T) {
	c, _ := NewCalculator()
	t.Log(c.EqualFrequencyBinningFloat32(test.DataPositiveFloat32()))
	t.Log(c.EqualFrequencyBinningFloat32(test.DataSameFloat32()))
	t.Log(c.EqualFrequencyBinningFloat32(test.DataSpecialFloat32()))
}
func TestCalculator_EqualWidthBinningInt32(t *testing.T) {
	c, _ := NewCalculator()
	t.Log(c.EqualWidthBinningInt32(test.DataPositiveInt32()))
	t.Log(c.EqualWidthBinningInt32(test.DataSameInt32()))
	t.Log(c.EqualWidthBinningInt32(test.DataSpecialInt32()))
}
func TestCalculator_EqualWidthBinningFloat32(t *testing.T) {
	c, _ := NewCalculator()
	t.Log(c.EqualWidthBinningFloat32(test.DataPositiveFloat32()))
	t.Log(c.EqualWidthBinningFloat32(test.DataSameFloat32()))
	t.Log(c.EqualWidthBinningFloat32(test.DataSpecialFloat32()))
}
