package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
    totalBirds := 0
    lenSlice := len(birdsPerDay)

    for i := 0; i < lenSlice; i++ {
        totalBirds += birdsPerDay[i]
    }
    return totalBirds
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
    totalBirds := 0
    lenSlice := len(birdsPerDay)

    for i := 0; i < lenSlice; i++ {
        if i >= (week - 1) * 7 && i <= week * 7 {
	        totalBirds += birdsPerDay[i]
        }
    }
    return totalBirds
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
    lenSlice := len(birdsPerDay)

    for i := 0; i < lenSlice; i++ {
        if i % 2 == 0 {
            birdsPerDay[i] += 1
        }
    }
    return birdsPerDay
}
