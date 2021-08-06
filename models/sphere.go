package models

import (
	"fmt"
	"math"
)

type Sphere struct {
	position Point3D
	radius   float64
}

func NewSphere(position *Point3D, radius float64) *Sphere {
	return &Sphere{
		position: *position,
		radius:   radius,
	}
}

func (sphere *Sphere) Position() Point3D {
	return sphere.position
}

func (sphere *Sphere) Radius() float64 {
	return sphere.radius
}

func (sphere *Sphere) ToString() string {
	return fmt.Sprintf("Sphere(pos=%s, r=%.3f)", sphere.position.ToString(), sphere.radius)
}

func (sphere *Sphere) IsPointInside(point *Point3D) bool {
	return point.Distance(&sphere.position) < sphere.radius
}

func (sphere *Sphere) IsPointInsideNoSqrt(point *Point3D) bool {
	return point.DistancePowerOf2(&sphere.position) < math.Pow(sphere.radius, 2)
}

func (sphere *Sphere) HasIntersect(target *Sphere) bool {
	return sphere.position.Distance(&target.position) < (sphere.radius + target.radius)
}

func (sphere *Sphere) HasIntersectWithAABB3D(target *AABB3D) bool {
	p := Point3D{
		x: math.Max(target.MinX(), math.Min(sphere.position.x, target.MaxX())),
		y: math.Max(target.MinY(), math.Min(sphere.position.y, target.MaxY())),
		z: math.Max(target.MinZ(), math.Min(sphere.position.z, target.MaxZ())),
	}

	return p.Distance(&sphere.position) < sphere.radius
}

func (sphere *Sphere) HasIntersectWithAABB3DNoSqrt(target *AABB3D) bool {
	p := Point3D{
		x: math.Max(target.MinX(), math.Min(sphere.position.x, target.MaxX())),
		y: math.Max(target.MinY(), math.Min(sphere.position.y, target.MaxY())),
		z: math.Max(target.MinZ(), math.Min(sphere.position.z, target.MaxZ())),
	}

	return p.DistancePowerOf2(&sphere.position) < math.Pow(sphere.radius, 2)
}
