package vector

type Vector struct {
	X float64
	Y float64
}

func Add(a, b Vector) Vector {
	return Vector{a.X + b.X, a.Y + b.Y}
}

func Substract(a, b Vector) Vector {
	return Vector{a.X - b.X, a.Y - b.Y}
}

func Scale(a Vector, multiplier float64) Vector {
	return Vector{a.X * multiplier, a.Y * multiplier}
}
