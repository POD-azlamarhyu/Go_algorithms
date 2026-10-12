package src

func BinarySearch[T Data](data *[]T, target T) *T {
	var left_val = 0
	var right_val = len(*data) - 1
	d := *data
	for right_val >= left_val {
		middle := left_val + (right_val - left_val) / 2
		if d[middle] == target {
			return &d[middle]
		}else if d[middle] > target {
			right_val = middle - 1
		}else if d[middle] < target {
			left_val = middle + 1
		}
	}
	return nil
}