package searches

func binarySearch(arr []int, target int, start int, end int) int {
	if start > end {
		return -1
	}
	mid := (end + start) / 2

	if arr[mid] > target {
		return binarySearch(arr, target, start, mid)
	}
	if arr[mid] < target {
		return binarySearch(arr, target, mid+1, end)
	}
	return mid
}
