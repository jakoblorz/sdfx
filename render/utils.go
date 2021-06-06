//-----------------------------------------------------------------------------

//-----------------------------------------------------------------------------

package render

import "github.com/jakoblorz/sdfx/sdf"

//-----------------------------------------------------------------------------

const tolerance = 1e-9
const epsilon = 1e-12

//-----------------------------------------------------------------------------

// nextCombination generates the next k-length combination of 0 to n-1. (returns false when done).
func nextCombination(n int, a []int) bool {
	k := len(a)
	m := 0
	i := 0
	for {
		i++
		if i > k {
			return false
		}
		if a[k-i] < n-i {
			m = a[k-i]
			for j := i; j >= 1; j-- {
				m++
				a[k-j] = m
			}
			return true
		}
	}
}

// mapCombinations applies a function f to each k-length combination from 0 to n-1.
func mapCombinations(n, k int, f func([]int)) {
	if k >= 0 && n >= k {
		a := make([]int, k)
		for i := range a {
			a[i] = i
		}
		for {
			f(a)
			if nextCombination(n, a) == false {
				break
			}
		}
	}
}

//-----------------------------------------------------------------------------

func RandomInUnitSphere(rnd Rnd) sdf.V3 {
	for {
		p := sdf.V3{2.0*rnd.Float64() - 1.0, 2.0*rnd.Float64() - 1.0, 2.0*rnd.Float64() - 1.0}
		if p.Dot(p) < 1.0 {
			return p
		}
	}
}

func RandomInUnitDisk(rnd Rnd) sdf.V3 {
	for {
		p := sdf.V3{2.0*rnd.Float64() - 1.0, 2.0*rnd.Float64() - 1.0, 0}
		if p.Dot(p) < 1.0 {
			return p
		}
	}
}
