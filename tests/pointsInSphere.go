package tests

import (
	"../models"
	"fmt"
	"math/rand"
	"time"
)

func PointsInSphere() {
	pointCount := 10000
	boxSize := float64(80)
	sphereRadius := float64(25)

	points := make([]*models.Point3D, pointCount)
	for i := 0; i < pointCount; i++ {
		points[i] = models.NewPoint(rand.Float64()*boxSize, rand.Float64()*boxSize, rand.Float64()*boxSize)
	}

	sphere := models.NewSphere(models.NewPoint(boxSize/2, boxSize/2, boxSize/2), sphereRadius)

	fmt.Printf("Start to check %d points in %s\n", pointCount, sphere.ToString())

	totalCorrect := 0
	totalInside := 0
	totalOutside := 0
	var averageTimeNormal int64 = 0
	var averageTimeNoSqrt int64 = 0
	fasterNormal := 0
	fasterNoSqrt := 0

	for i := 0; i < pointCount; i++ {
		t := time.Now()
		isInside := sphere.IsPointInside(points[i])
		t1 := time.Since(t)
		t = time.Now()
		isInsideNoSqrt := sphere.IsPointInsideNoSqrt(points[i])
		t2 := time.Since(t)

		if isInside == isInsideNoSqrt {
			if t1.Nanoseconds() < t2.Nanoseconds() {
				fasterNormal++
			} else {
				fasterNoSqrt++
			}
			averageTimeNormal = ((averageTimeNormal * int64(totalCorrect)) + t1.Nanoseconds()) / int64(totalCorrect+1)
			averageTimeNoSqrt = ((averageTimeNoSqrt * int64(totalCorrect)) + t2.Nanoseconds()) / int64(totalCorrect+1)
			totalCorrect++
			if isInside {
				totalInside++
			} else {
				totalOutside++
			}
		}
	}

	fmt.Printf("\tcorrect answers: %d (inside: %d, outside: %d)\n", totalCorrect, totalInside, totalOutside)
	fmt.Printf("\taverageTimeNormal: %dns, averageTimeNoSqrt: %dns\n", averageTimeNormal, averageTimeNoSqrt)
	fmt.Printf("\tin %d cases normal is faster and in %d cases nosqrt\n", fasterNormal, fasterNoSqrt)
}
