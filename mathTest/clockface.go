package mathTest

import (
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

const secondHandLength = 90
const clockCenterX = 150
const clockCenterY = 150

// SecondHand is the unit vector of the second hand of an analogue clock at time `t`
// represented as a Point.
func SecondHand(t time.Time) Point {
	p := secondHandPoint(t)
	p = Point{p.X * secondHandLength, p.Y * secondHandLength} // scale: SVGの時計の長さに合わせる
	p = Point{p.X, -p.Y}                                      // flip: 12時をスタート地点にするため、X軸上で反転させる
	p = Point{p.X + clockCenterX, p.Y + clockCenterY}         // translate: 座標の中心点を時計の中心に変換
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
