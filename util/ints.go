package util

// Sort3 returns the given three integers sorted by value (smallest first).
func Sort3(x, y, z int) (a, b, c int) {
	if x <= y { // x <= y
		if x <= z { // x <= y, x <= z
			if y <= z { // x <= y, x <= z, y <= z
				return x, y, z
			} else { // x <= y, x <= z, y > z
				return x, z, y
			}
		} else { // x <= y, x > z
			return z, x, y
		}
	} else { // x > y
		if y <= z { // x > y, y <= z
			if x <= z { // x > y, y <= z, x <= z
				return y, x, z
			} else { // x > y, y <= z, x > z
				return y, z, x
			}
		} else { // x > y, y > z
			return z, y, x
		}
	}
}
