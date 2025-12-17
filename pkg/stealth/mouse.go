package stealth

import (
	"math"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

// Point represents a 2D coordinate
type Point struct {
	X float64
	Y float64
}

// MoveMouse moves the mouse to a specific coordinate using Bezier curves
func (e *Engine) MoveMouse(page *rod.Page, toX, toY float64) error {
	return e.humanMove(page, toX, toY)
}

// BezierCurve generates points along a cubic Bezier curve
func BezierCurve(start, end, control1, control2 Point, steps int) []Point {
	path := make([]Point, steps)
	for i := 0; i < steps; i++ {
		t := float64(i) / float64(steps-1)

		u := 1 - t
		tt := t * t
		uu := u * u
		uuu := uu * u
		ttt := tt * t

		p := Point{
			X: uuu*start.X + 3*uu*t*control1.X + 3*u*tt*control2.X + ttt*end.X,
			Y: uuu*start.Y + 3*uu*t*control1.Y + 3*u*tt*control2.Y + ttt*end.Y,
		}
		path[i] = p
	}
	return path
}

// MoveToElement moves the mouse to an element using a human-like curve
func (e *Engine) MoveToElement(page *rod.Page, el *rod.Element) error {
	res, err := el.Eval(`() => {
        const rect = this.getBoundingClientRect();
        return {x: rect.x, y: rect.y, width: rect.width, height: rect.height};
    }`)
	if err != nil {
		return err
	}

	val := res.Value
	x := val.Get("x").Num()
	y := val.Get("y").Num()
	width := val.Get("width").Num()
	height := val.Get("height").Num()

	targetX := x + width*0.5 + (e.rnd.Float64()-0.5)*(width*0.8)
	targetY := y + height*0.5 + (e.rnd.Float64()-0.5)*(height*0.8)

	return e.humanMove(page, targetX, targetY)
}

func (e *Engine) humanMove(page *rod.Page, toX, toY float64) error {
	fromX := e.lastX
	fromY := e.lastY

	dist := math.Hypot(toX-fromX, toY-fromY)
	steps := int(dist / 10)
	if steps < 10 {
		steps = 10
	}

	spread := dist * 0.2
	c1 := Point{
		X: fromX + (toX-fromX)/3 + (e.rnd.Float64()-0.5)*spread,
		Y: fromY + (toY-fromY)/3 + (e.rnd.Float64()-0.5)*spread,
	}
	c2 := Point{
		X: fromX + 2*(toX-fromX)/3 + (e.rnd.Float64()-0.5)*spread,
		Y: fromY + 2*(toY-fromY)/3 + (e.rnd.Float64()-0.5)*spread,
	}

	path := BezierCurve(Point{fromX, fromY}, Point{toX, toY}, c1, c2, steps)

	for _, p := range path {
		page.Mouse.MoveLinear(proto.Point{X: p.X, Y: p.Y}, 1)

		if e.rnd.Float64() < 0.1 {
			time.Sleep(time.Millisecond * time.Duration(e.rnd.Intn(10)))
		}
	}

	page.Mouse.MoveLinear(proto.Point{X: toX, Y: toY}, 1)
	e.lastX = toX
	e.lastY = toY

	return nil
}
