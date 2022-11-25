package Triangle

func Triangle(a, b, c int) string {
	if a <= 0 || b <= 0 || c <= 0 || a+b <= c || b+c <= a || a+c <= b {
		return "Not a triangle"
	}
	switch {
	case a == b && a == c && b == c:
		return "equilateral"
	case a == b || b == c || c == a:
		return "isosceles"
	default:
		return "scalene"

	}
}
