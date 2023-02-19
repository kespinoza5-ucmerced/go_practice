package math

// Point declare a struct with only members
type Point struct {
	x, y float64
}

// NewPoint constructor, creates a new point
func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

// DistanceToOrigin return distance from x & y vals belonging to a point
func (p Point) DistanceToOrigin() float64 {
	return p.x*p.x + p.y*p.y
}

// Double without needing to return, can modify values of a point by reference
func (p *Point) Double() {
	p.x *= 2
	p.y *= 2
}
