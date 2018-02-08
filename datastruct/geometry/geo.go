package geometry

import "math"


type Point struct {
    X, Y float64
}

func Distance(p, q Point)float64 {
    return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func(self Point)Distance(q Point)float64 {
    return math.Hypot(self.X-q.X, self.Y-q.Y)
}


type Path []Point

func(self Path)Distance()float64 {
    sum := 0.0
    for i := range self {
        if i > 0 {
            sum += self[i-1].Distance(self[i])
        }
    }
    return sum
}        
