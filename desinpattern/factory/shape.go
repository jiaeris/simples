package factory

type Shape interface {
	Draw()
}

func GetShape() Shape {

	return nil
}

type Circle struct {
}

func (c *Circle) Draw() {

}

type Rectangle struct {

}

func (c *Rectangle) Draw() {

}

type Square struct {
}

func (c *Square) Draw() {

}
