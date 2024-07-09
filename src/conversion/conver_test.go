package conversion

import (
	"github.com/dingyuqi/go_numcalc/test"
	"testing"
)

func TestNewCalculator(t *testing.T) {
	t.Log(NewCalculator(10))
	t.Log(NewCalculator())
}

func TestCalculator_StandardizationInt32(t *testing.T) {
	c, _ := NewCalculator()
	t.Log(c.StandardizationInt32(test.DataPositiveInt32()))
	t.Log(c.StandardizationInt32(test.DataNegativeInt32()))
	t.Log(c.StandardizationInt32(test.DataSameInt32()))
	t.Log(c.StandardizationInt32(test.DataSameInt32()))
}

func TestCalculator_StandardizationFloat32(t *testing.T) {
	c, _ := NewCalculator()
	t.Log(c.StandardizationFloat32(test.DataPositiveFloat32()))
	t.Log(c.StandardizationFloat32(test.DataNegativeFloat32()))
	t.Log(c.StandardizationFloat32(test.DataSpecialFloat32()))
	t.Log(c.StandardizationFloat32(test.DataSameFloat32()))
}

func TestCalculator_ZScoreInt32(t *testing.T) {
	c, _ := NewCalculator()
	t.Log(c.ZScoreInt32(test.DataPositiveInt32()))
	t.Log(c.ZScoreInt32(test.DataNegativeInt32()))
	t.Log(c.ZScoreInt32(test.DataSpecialInt32()))
	t.Log(c.ZScoreInt32(test.DataSameInt32()))

}

func TestCalculator_ZScoreFloat32(t *testing.T) {
	c, _ := NewCalculator()
	t.Log(c.ZScoreFloat32(test.DataPositiveFloat32()))
	t.Log(c.ZScoreFloat32(test.DataNegativeFloat32()))
	t.Log(c.ZScoreFloat32(test.DataSpecialFloat32()))
	t.Log(c.ZScoreFloat32(test.DataSameFloat32()))
}

func TestCalculator_LogInt32(t *testing.T) {
	c, _ := NewCalculator()
	t.Log(c.LogInt32(test.DataPositiveInt32()))
	t.Log(c.LogInt32(test.DataNegativeInt32()))
	t.Log(c.LogInt32(test.DataSpecialInt32()))
}

func TestCalculator_LogFloat32(t *testing.T) {
	c, _ := NewCalculator()
	t.Log(c.LogFloat32(test.DataPositiveFloat32()))
	t.Log(c.LogFloat32(test.DataNegativeFloat32()))
	t.Log(c.LogFloat32(test.DataSpecialFloat32()))
}

func TestCalculator_SquareInt32(t *testing.T) {
	c, _ := NewCalculator()
	t.Log(c.SquareInt32(test.DataPositiveInt32()))
	t.Log(c.SquareInt32(test.DataNegativeInt32()))
	t.Log(c.SquareInt32(test.DataSpecialInt32()))
}

func TestCalculator_SquareFloat32(t *testing.T) {
	c, _ := NewCalculator()
	t.Log(c.SquareFloat32(test.DataPositiveFloat32()))
	t.Log(c.SquareFloat32(test.DataNegativeFloat32()))
	t.Log(c.SquareFloat32(test.DataSpecialFloat32()))
}
