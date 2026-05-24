package math2d

func CircleCollision(
	posA Vector2,
	radiusA float64,
	posB Vector2,
	radiusB float64,
) bool {
	r := radiusA + radiusB

	return DistanceSquared(posA, posB) <= r*r
}

func PointInRect(
	point Vector2,
	x, y, w, h float64,
) bool {
	return point.X >= x &&
		point.X <= x+w &&
		point.Y >= y &&
		point.Y <= y+h
}
