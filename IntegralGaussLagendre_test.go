// IntegralGaussLagendre_test
/*
------------------------------------------------------
作者   : Black Ghost
日期   : 2018-12-12
版本   : 0.0.0
------------------------------------------------------
    不超过8次的Gauss-Lagendre求积分公式
理论：
    对于积分
    b
    |f(x)dx
    a

    使用n+1次Lagendre多项式的零点作为高斯点，可获得代数
    精度为2n+1的高斯型求积公式
    其中区间[a, b]需要预先转换为区间[-1, 1]

    参考 李信真, 车刚明, 欧阳洁, 等. 计算方法. 西北工业大学
       出版社, 2000, pp 162-164.
------------------------------------------------------
输入   :
    fun     被积分函数
    a, b    积分区间
    n       求积分公式次数
输出   :
    sol     解
    err     解出标志：false-未解出或达到步数上限；
                     true-全部解出
------------------------------------------------------
*/

package goNum_test

import (
	"math"
	"testing"

	"github.com/chfenger/goNum"
)

func IntegralGaussLagendre(fun func(float64) float64, a, b float64, n int) (float64, bool) {
	/*
		不超过8次的Gauss-Lagendre求积分公式
		输入   :
		    fun     被积分函数
		    a, b    积分区间
		    n       求积分公式次数
		输出   :
		    sol     解
		    err     解出标志：false-未解出或达到步数上限；
		                     true-全部解出
	*/
	//判断n范围
	if (n < 1) || (n > 8) {
		panic("Error in goNum.IntegralGaussLagendre: n is a not correct input")
	}

	xi := [][]float64{
		{0.0},
		{-0.5773502692, 0.5773502692},
		{-0.7745966692, 0.0, 0.7745966692},
		{-0.8611363116, -0.3399810436, 0.3399810436, 0.8611363116},
		{-0.9061798459, -0.5384693101, 0.0, 0.5384693101, 0.9061798459},
		{-0.9324695142, -0.6612093865, -0.2386191861, 0.2386191861, 0.6612093865, 0.9324695142},
		{-0.9491079123, -0.7415311856, -0.4058451514, 0.0, 0.4058451514, 0.7415311856, 0.9491079123},
		{-0.9602898566, -0.7966664774, -0.5255324099, -0.1834346425, 0.1834346425, 0.5255324099, 0.7966664774, 0.9602898566},
	}
	Ai := [][]float64{
		{2.0},
		{1.0, 1.0},
		{0.555555555555556, 0.888888888888889, 0.555555555555556},
		{0.3478548451, 0.6521451549, 0.6521451549, 0.3478548451},
		{0.2369268851, 0.4786286705, 0.568888889, 0.4786286705, 0.2369268851},
		{0.1713244924, 0.3607615730, 0.4679139346, 0.4679139346, 0.3607615730, 0.1713244924},
		{0.1294849662, 0.2797053915, 0.3818300505, 0.4179591837, 0.3818300505, 0.2797053915, 0.1294849662},
		{0.1012285363, 0.2223810345, 0.3137066459, 0.3626837834, 0.3626837834, 0.3137066459, 0.2223810345, 0.1012285363},
	}

	//区间转换
	c := (b - a) / 2.0
	d := (a + b) / 2.0

	switch n {
	case 1:
		sol := 0.0
		for i := 0; i < len(xi[0]); i++ {
			sol += Ai[0][i] * fun(d+c*xi[0][i])
		}
		return c * sol, true
	case 2:
		sol := 0.0
		for i := 0; i < len(xi[1]); i++ {
			sol += Ai[1][i] * fun(d+c*xi[1][i])
		}
		return c * sol, true
	case 3:
		sol := 0.0
		for i := 0; i < len(xi[2]); i++ {
			sol += Ai[2][i] * fun(d+c*xi[2][i])
		}
		return c * sol, true
	case 4:
		sol := 0.0
		for i := 0; i < len(xi[3]); i++ {
			sol += Ai[3][i] * fun(d+c*xi[3][i])
		}
		return c * sol, true
	case 5:
		sol := 0.0
		for i := 0; i < len(xi[4]); i++ {
			sol += Ai[4][i] * fun(d+c*xi[4][i])
		}
		return c * sol, true
	case 6:
		sol := 0.0
		for i := 0; i < len(xi[5]); i++ {
			sol += Ai[5][i] * fun(d+c*xi[5][i])
		}
		return c * sol, true
	case 7:
		sol := 0.0
		for i := 0; i < len(xi[6]); i++ {
			sol += Ai[6][i] * fun(d+c*xi[6][i])
		}
		return c * sol, true
	case 8:
		sol := 0.0
		for i := 0; i < len(xi[7]); i++ {
			sol += Ai[7][i] * fun(d+c*xi[7][i])
		}
		return c * sol, true
	default:
		return 0.0, false
	}
}

func fun38(x float64) float64 {
	return x * x * math.Cos(x)
}

func BenchmarkIntegralGaussLagendre(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goNum.IntegralGaussLagendre(fun38, 0.0, math.Pi/2.0, 4) //0.467402...
	}
}