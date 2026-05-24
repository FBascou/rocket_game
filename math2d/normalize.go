package math2d

func Normalize(v Vector2) Vector2 {
	mag := Magnitude(v)

	if mag == 0 {
		return Vector2{}
	}

	return Vector2{
		X: v.X / mag,
		Y: v.Y / mag,
	}
}

func Direction(from, to Vector2) Vector2 {
	return Normalize(Sub(to, from))
}
