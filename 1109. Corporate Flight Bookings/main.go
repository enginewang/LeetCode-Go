package main

func corpFlightBookings(bookings [][]int, n int) []int {
	result := make([]int, n)
	diffArray := make([]int, n)
	for _, booking := range bookings{
		first := booking[0]
		last := booking[1]
		seats := booking[2]
		diffArray[first-1] += seats
		if last < n{
			diffArray[last] -= seats
		}
	}
	cur := 0
	for i, diff := range diffArray{
		cur += diff
		result[i] = cur
	}
	return result
}

func main() {
	corpFlightBookings([][]int{ {1,2,10}, {2,3,20}, {2,5,25} }, 5)
}
