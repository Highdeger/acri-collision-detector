package models

import (
	"fmt"
	"math"
)

type Point3D struct {
	x float64
	y float64
	z float64
}

func NewPoint(x, y, z float64) *Point3D {
	return &Point3D{
		x: x,
		y: y,
		z: z,
	}
}

func (p *Point3D) X() float64 {
	return p.x
}

func (p *Point3D) Y() float64 {
	return p.y
}

func (p *Point3D) Z() float64 {
	return p.z
}

func (p *Point3D) ToString() string {
	return fmt.Sprintf("Point3D(%.3f, %.3f, %.3f)", p.x, p.y, p.z)
}

func (p *Point3D) Distance(to *Point3D) float64 {
	return math.Sqrt(math.Pow(p.x-to.x, 2) + math.Pow(p.y-to.y, 2) + math.Pow(p.z-to.z, 2))
}

func (p *Point3D) DistancePowerOf2(to *Point3D) float64 {
	return math.Pow(p.x-to.x, 2) + math.Pow(p.y-to.y, 2) + math.Pow(p.z-to.z, 2)
}
