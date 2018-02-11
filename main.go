package main

import (
    "os"

    gs "gravsim"
)

func main() {
    bodies := []gs.Body{
        gs.Body{
            P: gs.Vec2{150.0,150.0},
            V: gs.Vec2{10.0,0.0},
            A: gs.Vec2{0.0,0.0},
            M: 50000.0,
            Radius: 15.0,
        },
        gs.Body{
            P: gs.Vec2{350.0,150.0},
            V: gs.Vec2{0.0,10.0},
            A: gs.Vec2{0.0,0.0},
            M: 50000.0,
            Radius: 15.0,
        },
        gs.Body{
            P: gs.Vec2{150.0,350.0},
            V: gs.Vec2{0.0,-10.0},
            A: gs.Vec2{0.0,0.0},
            M: 50000.0,
            Radius: 15.0,
        },
        gs.Body{
            P: gs.Vec2{350.0,350.0},
            V: gs.Vec2{-10.0,0.0},
            A: gs.Vec2{0.0,0.0},
            M: 50000.0,
            Radius: 15.0,
        },
    }
    system := &gs.System{
        Bodies: bodies,
        G: 1.0,
    }
    simulation := &gs.Simulation{
        System: system,
        X: 0,
        Y: 0,
        Dx: 500,
        Dy: 500,
        StepInterval: 0.04,
        StepsPerFrame: 1,
        TotalSteps: 2000,
    }
    f, err := os.Create("render.gif")
    if err != nil {
        panic(err)
    }
    simulation.Render(f)
}

