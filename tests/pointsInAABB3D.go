package tests

import (
	"../models"
	"fmt"
	"math/rand"
	"time"
)

func PointsInAABB3D() {
	pointCount := 200
	boxSize := float64(80)
	aabbXLen := float64(79)
	aabbYLen := float64(79)
	aabbZLen := float64(79)

	points := make([]*models.Point3D, pointCount)
	for i := 0; i < pointCount; i++ {
		points[i] = models.NewPoint(rand.Float64()*boxSize, rand.Float64()*boxSize, rand.Float64()*boxSize)
	}

	aabb3d := models.NewAABB3D(models.NewPoint(boxSize/2, boxSize/2, boxSize/2), aabbXLen, aabbYLen, aabbZLen)

	fmt.Printf("Start to check %d points in %s\n", pointCount, aabb3d.ToString())

	totalCorrect := 0
	totalInside := 0
	totalOutside := 0
	var averageTime int64 = 0

	for i := 0; i < pointCount; i++ {
		t := time.Now()
		isInside := aabb3d.IsPointInside(points[i])
		t1 := time.Since(t)

		averageTime = ((averageTime * int64(totalCorrect)) + t1.Nanoseconds()) / int64(totalCorrect+1)
		totalCorrect++

		if isInside {
			totalInside++
		} else {
			totalOutside++
		}
	}

	fmt.Printf("\tcorrect answers: %d (inside: %d, outside: %d)\n", totalCorrect, totalInside, totalOutside)
	fmt.Printf("\taverage time: %dns\n", averageTime)
}
