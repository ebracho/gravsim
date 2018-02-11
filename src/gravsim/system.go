package gravsim

import (
    "fmt"
    "image"
    "math"
)

type System struct {
    Bodies []Body
    G float64
}

func (s System) String() string {
    return fmt.Sprintf("{G: %f, Bodies: %s}", s.G, s.Bodies)
}

func (s *System) updateAccelerations() {
    for i := range s.Bodies {
        s.Bodies[i].A = Vec2{X: 0.0, Y: 0.0}
    }
    for i := 0; i < len(s.Bodies); i++ {
        for j := i+1; j < len(s.Bodies); j++ {
            bi := s.Bodies[i]
            bj := s.Bodies[j]
            d2 := math.Max(bi.P.dist2(bj.P), math.Pow(bi.Radius + bj.Radius, 2))
            s.Bodies[i].A.addi(bj.P.subtract(bi.P).norm().scale(s.G * bj.M / d2))
            s.Bodies[j].A.addi(bi.P.subtract(bj.P).norm().scale(s.G * bi.M / d2))
        }
    }
}

func (s *System) Step(t float64) {
    s.updateAccelerations()
    for i := range s.Bodies {
        s.Bodies[i].Step(t)
    }
}

func (s *System) Render(img *image.Paletted, pi uint8) {
    for i := range s.Bodies {
        s.Bodies[i].Render(img, pi)
    }
}
