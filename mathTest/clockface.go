package mathTest

import (
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

// SecondHand is the unit vector of the second hand of an analogue clock at time `t`
// represented as a Point.
func SecondHand(t time.Time) Point {
	p := secondHandPoint(t)
	p = Point{p.X * 90, p.Y * 90}   // scale: SVGの時計の長さに合わせる
	p = Point{p.X, -p.Y}            // flip: 12時をスタート地点にするため、X軸上で反転させる
	p = Point{p.X + 150, p.Y + 150} // translate: 正しい位置に移動
	return p
}

func secondsInRadians(t time.Time) float64 {
	return (math.Pi / (30 / (float64(t.Second()))))
}

func secondHandPoint(t time.Time) Point {
	angle := secondsInRadians(t)
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}
