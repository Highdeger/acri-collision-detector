package main

import (
	"./tests"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	tests.PointsInSphere()
	tests.PointsInAABB3D()
}
