package bashColor

import "fmt"

type Color struct {
	Code Coder
}

func NewColor() Colorer {
	var color Colorer
	color = &Color{
		Code: NewCode(),
	}
	return color
}

func (c *Color) Format(color string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(color, fmt.Sprint(args...))
	}
	return sprint
}

func (c *Color) Print(f func(...interface{}) string, s string) {
	fmt.Println(f(s))
}

func (c *Color) Info(message ...interface{}) string {
	return c.Teal(message...)
}
func (c *Color) Success(message ...interface{}) string {
	return c.Green(message...)
}
func (c *Color) Warning(message ...interface{}) string {
	return c.Yellow(message...)
}
func (c *Color) Fatal(message ...interface{}) string {
	return c.Red(message...)
}

func (c *Color) Black(message ...interface{}) string {
	return c.Format(c.Code.GetColor(Black) + "%s" + c.Code.GetColor(Default))(message...)
}
func (c *Color) Red(message ...interface{}) string {
	return c.Format(c.Code.GetColor(Red) + "%s" + c.Code.GetColor(Default))(message...)
}
func (c *Color) Green(message ...interface{}) string {
	return c.Format(c.Code.GetColor(Green) + "%s" + c.Code.GetColor(Default))(message...)
}
func (c *Color) Yellow(message ...interface{}) string {
	return c.Format(c.Code.GetColor(Yellow) + "%s" + c.Code.GetColor(Default))(message...)
}
func (c *Color) Purple(message ...interface{}) string {
	return c.Format(c.Code.GetColor(Purple) + "%s" + c.Code.GetColor(Default))(message...)
}
func (c *Color) Magenta(message ...interface{}) string {
	return c.Format(c.Code.GetColor(Magenta) + "%s" + c.Code.GetColor(Default))(message...)
}
func (c *Color) Teal(message ...interface{}) string {
	return c.Format(c.Code.GetColor(Teal) + "%s" + c.Code.GetColor(Default))(message...)
}
func (c *Color) White(message ...interface{}) string {
	return c.Format(c.Code.GetColor(White) + "%s" + c.Code.GetColor(Default))(message...)
}