// K-means clustering algorithm adapted from the JavaScript implementation by
// Heather Arthur: https://github.com/harthur/clusterfck

package kmeans

import (
	"math"
	"math/rand"
	"sync"
)

// Calculate distance between 2 4-dimensional coordinates
func distance4d(coords1 []float64, coords2 []float64) float64 {
	l1 := math.Pow(coords2[0]-coords1[0], 2.0)
	l2 := math.Pow(coords2[1]-coords1[1], 2.0)
	l3 := math.Pow(coords2[2]-coords1[2], 2.0)
	l4 := math.Pow(coords2[3]-coords1[3], 2.0)
	return math.Sqrt(l1 + l2 + l3 + l4)
}

// Select k points randomly from a slice of points to use as centroids
func randomCentroids(points [][]float64, k int) [][]float64 {
	var centroids [][]float64
	centroids = make([][]float64, k, k)

	for i := 0; i < k; i++ {
		centroids[i] = points[rand.Intn(len(points))]
	}

	return centroids
}

// Identify the closest centroid to the point
func classify(point []float64, centroids [][]float64) int {
	var (
		min   float64
		index int
	)

	min = math.MaxFloat64
	index = 0

	for i := 0; i < len(centroids); i++ {
		dist := distance4d(point, centroids[i])
		if dist < min {
			min = dist
			index = i
		}
	}

	return index
}

// Perform k-means clustering on a set of 4-d coordinates
func Cluster(points [][]float64, k int) ([][]float64, [][][]float64) {
	var (
		centroidAssignment []int
		centroids          [][]float64
		clusters           [][][]float64
		movement           bool
		numPoints          int
		g, h, i, j, l      int
		wg                 sync.WaitGroup
	)

	// Create initial state
	numPoints = len(points)
	centroids = randomCentroids(points, k)
	centroidAssignment = make([]int, numPoints, numPoints)
	clusters = make([][][]float64, k, k)
	movement = true

	for movement {
		// update point-to-centroid assignments
		wg.Add(numPoints)
		for i = 0; i < numPoints; i++ {
			go func(i int) {
				centroidAssignment[i] = classify(points[i], centroids)
			}(i)
		}
		wg.Done()

		// update location of each centroid
		movement = false
		for j = 0; j < k; j++ {
			// assign each point to the correct cluster[k]
			assigned := make([][]float64, 0)
			for g = 0; g < numPoints; g++ {
				if centroidAssignment[g] == j {
					assigned = append(assigned, points[g])
				}
			}

			lenAssigned := len(assigned)
			if lenAssigned == 0 {
				continue
			}

			curCentroid := centroids[j]
			centroidLen := len(curCentroid)
			newCentroid := make([]float64, centroidLen, centroidLen)

			// update each centroid pos to be average pos for each point coord
			for h = 0; h < centroidLen; h++ {
				var sum float64
				sum = 0
				for l = 0; l < lenAssigned; l++ {
					sum += assigned[l][h]
				}

				newCentroid[h] = sum / float64(lenAssigned)

				if newCentroid[h] != curCentroid[h] {
					movement = true
				}
			}

			centroids[j] = newCentroid
			clusters[j] = assigned
		}
	}

	return centroids, clusters
}
