package rectangle

type Rectangle struct {
	Width  float64
	Height float64
	Area   float64
}

// NewRectangle is a constructor for Rectangle
func NewRectangle(width float64, height float64) *Rectangle {
	return &Rectangle{
		Width:  width,
		Height: height,
	}
}

func (rect *Rectangle) CalcArea() {
	rect.Area = rect.Width * rect.Height
}
