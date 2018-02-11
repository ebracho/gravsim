package gravsim

import (
    "math"
    "fmt"
)

type Vec2 struct {
    X, Y float64
}

func (v Vec2) String() string {
    return fmt.Sprintf("{X: %f, Y: %f}", v.X, v.Y)
}

func (v Vec2) add(vo Vec2) Vec2 {
    return Vec2{X: v.X+vo.X, Y: v.Y+vo.Y}
}

func (v *Vec2) addi(vo Vec2) {
    v.X += vo.X
    v.Y += vo.Y
}

func (v Vec2) subtract(vo Vec2) Vec2 {
    return Vec2{X: v.X-vo.X, Y: v.Y-vo.Y}
}

func (v *Vec2) subtracti(vo Vec2) {
    v.X -= vo.X
    v.Y -= vo.Y
}

func (v Vec2) scale(s float64) Vec2 {
    return Vec2{X: s*v.X, Y: s*v.Y}
}

func (v Vec2) scalei(s float64) {
    v.X *= s
    v.Y *= s
}

func (v Vec2) length2() float64 {
    return math.Pow(v.X, 2) + math.Pow(v.Y, 2)
}

func (v Vec2) length() float64 {
    return math.Sqrt(v.length2())
}

func (v Vec2) dist2(vo Vec2) float64 {
    return v.subtract(vo).length2()
}

func (v Vec2) dist(vo Vec2) float64 {
    return v.subtract(vo).length()
}

func (v Vec2) norm() Vec2 {
    return v.scale(1.0/v.length())
}
