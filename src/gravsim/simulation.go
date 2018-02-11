package gravsim

import (
    "image"
    "image/color"
    "image/gif"
    "io"
    "math"
)

type Simulation struct {
    System *System        // System to simulate
    X, Y, Dx, Dy int      // Position and dimensions of simulation camera
    StepInterval float64  // Time between simulation steps
    StepsPerFrame int     // Number of steps to run per frame render
    TotalSteps int        // Number of steps to run for simulation
}

func (s Simulation) renderFrame() *image.Paletted {
    rect := image.Rect(s.X, s.Y, s.Dx, s.Dy)
    p := color.Palette{color.White, color.Black}
    img := image.NewPaletted(rect, p)
    s.System.Render(img, 1)
    return img
}

// Render gif animation of simulation
func (s Simulation) Render(out io.Writer) {
    nframes := s.TotalSteps / s.StepsPerFrame
    delay := int(math.Ceil(s.StepInterval * float64(s.StepsPerFrame) * 100.0))
    anim := gif.GIF{
        LoopCount: nframes,
        //Image: make([]*image.Paletted, nframes),
        //Delay: make([]int, nframes),
    }
    for t := 0; t < s.TotalSteps; t++ {
        s.System.Step(s.StepInterval)
        if t % s.StepsPerFrame == 0 {
            anim.Image = append(anim.Image, s.renderFrame())
            anim.Delay = append(anim.Delay, delay)
        }
    }
    gif.EncodeAll(out, &anim)
}
