package models

import "fmt"

type AABB3D struct {
	position Point3D
	xLen     float64
	yLen     float64
	zLen     float64
}

func NewAABB3D(position *Point3D, xLen, yLen, zLen float64) *AABB3D {
	return &AABB3D{
		position: *position,
		xLen:     xLen,
		yLen:     yLen,
		zLen:     zLen,
	}
}

func (box *AABB3D) ToString() string {
	return fmt.Sprintf("AABB3D(pos=%s, xLen=%.3f, yLen=%.3f, zLen=%.3f)", box.position.ToString(), box.xLen, box.yLen, box.zLen)
}

func (box *AABB3D) MinX() float64 {
	return box.position.x - (box.xLen / 2)
}

func (box *AABB3D) MaxX() float64 {
	return box.position.x + (box.xLen / 2)
}

func (box *AABB3D) MinY() float64 {
	return box.position.y - (box.yLen / 2)
}

func (box *AABB3D) MaxY() float64 {
	return box.position.y + (box.yLen / 2)
}

func (box *AABB3D) MinZ() float64 {
	return box.position.z - (box.zLen / 2)
}

func (box *AABB3D) MaxZ() float64 {
	return box.position.z + (box.zLen / 2)
}

func (box *AABB3D) IsPointInside(point *Point3D) bool {
	return (point.x >= box.MinX() && point.x <= box.MaxX()) &&
		(point.y >= box.MinY() && point.y <= box.MaxY()) &&
		(point.z >= box.MinZ() && point.z <= box.MaxZ())
}

func (box *AABB3D) HasIntersect(target *AABB3D) bool {
	return (box.MinX() <= target.MaxX() && box.MaxX() >= target.MinX()) &&
		(box.MinY() <= target.MaxY() && box.MaxY() >= target.MinY()) &&
		(box.MinZ() <= target.MaxZ() && box.MaxZ() >= target.MinZ())
}

func (box *AABB3D) HasIntersectWithSphere(target *Sphere) bool {
	return target.HasIntersectWithAABB3D(box)
}

func (box *AABB3D) HasIntersectWithSphereNoSqrt(target *Sphere) bool {
	return target.HasIntersectWithAABB3DNoSqrt(box)
}
