package golang_united_school_homework

import "fmt"

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes)+1 > b.shapesCapacity {
		return fmt.Errorf("error")
	}
	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i > len(b.shapes)-1 {
		return nil, fmt.Errorf("error")
	} else {
		return b.shapes[i], nil
	}

}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	sh, err := b.GetByIndex(i)

	if err == nil {
		b.shapes = remove(b.shapes, i)
	}

	return sh, err
}

func remove(s []Shape, i int) []Shape {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	sh, err := b.GetByIndex(i)

	if err == nil {
		b.shapes = replaceByIndex(b.shapes, i, shape)
	}

	return sh, err
}

func replaceByIndex(s []Shape, i int, shape Shape) []Shape {
	s[i] = shape
	return s
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var res float64

	for _, k := range b.shapes {
		res += k.CalcPerimeter()
	}

	return res

}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var res float64

	for _, k := range b.shapes {
		res += k.CalcArea()
	}

	return res
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	var cnt int

	lst := make([]Shape, 0)

	for _, v := range b.shapes {
		if v.GetType() != "Circle" {
			lst = append(lst, v)
		} else {
			cnt++
		}
	}

	b.shapes = lst

	if cnt > 0 {
		return nil
	} else {
		return fmt.Errorf("error")
	}
}
