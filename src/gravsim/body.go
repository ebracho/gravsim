package gravsim

import (
    "math"
    "fmt"
    "image"
)

type Body struct {
    P, V, A Vec2
    M, Radius float64
}

func (b Body) String() string {
    return fmt.Sprintf("{P: %s, V: %s, A: %s, M: %f}", b.P, b.V, b.A, b.M)
}

func (b *Body) Step(t float64) {
    b.V.addi(b.A.scale(t))
    b.P.addi(b.V.scale(t))
}

func (b *Body) Render(img *image.Paletted, pi uint8) {
    r := img.Rect
    // Limit render area to bounding box of body
    xmin := int(math.Max(float64(r.Min.X), b.P.X - b.Radius))
    ymin := int(math.Max(float64(r.Min.Y), b.P.Y - b.Radius))
    xmax := int(math.Min(float64(r.Max.X), b.P.X + b.Radius))
    ymax := int(math.Min(float64(r.Max.Y), b.P.Y + b.Radius))
    radius2 := b.Radius * b.Radius
    for i := xmin; i <= xmax; i++ {
        for j := ymin; j <= ymax; j++ {
            if b.P.dist2(Vec2{X: float64(i), Y: float64(j)}) < radius2 {
                img.SetColorIndex(i, j, pi)
            }
        }
    }
}
