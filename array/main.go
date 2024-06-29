package array

func ForEach[T any](arr []T, callback func(T, int, []T)) {
	for i := 0; i < len(arr); i++ {
		callback(arr[i], i, arr)
	}
}

func Map[T any](arr []T, callback func(T, int, []T) T) []T {
	var _arr []T
	for i := 0; i < len(arr); i++ {
		item := callback(arr[i], i, arr)
		_arr = append(_arr, item)
	}

	return _arr
}

func Filter[T any](arr []T, callback func(T, int, []T) bool) []T {
	var _arr []T
	for i := 0; i < len(arr); i++ {
		item := arr[i]
		isTrue := callback(arr[i], i, arr)

		if isTrue {
			_arr = append(_arr, item)
		}

	}

	return _arr
}

func Reduce[T any](arr []T, callback func([]T, T, int, []T) []T, wrapper []T) []T {
	for i := 0; i < len(arr); i++ {
		wrapper = callback(wrapper, arr[i], i, arr)
	}

	return wrapper
}

func ReduceRight[T any](arr []T, callback func([]T, T, int, []T) []T, wrapper []T) []T {
	for i := len(arr) - 1; i >= 0; i-- {
		wrapper = callback(wrapper, arr[i], i, arr)
	}

	return wrapper
}

func Every[T any](arr []T, callback func(T, int, []T) bool) bool {
	for i := 0; i < len(arr); i++ {
		isTrue := callback(arr[i], i, arr)

		if !isTrue {
			return false
		}
	}

	return true
}

func Some[T any](arr []T, callback func(T, int, []T) bool) bool {
	for i := 0; i < len(arr); i++ {
		isTrue := callback(arr[i], i, arr)

		if isTrue {
			return true
		}
	}

	return false
}

func Push[T any](arr *[]T, args ...T) int {
	*arr = append(*arr, args...)
	return len(*arr)
}

func Unshift[T any](arr *[]T, item T) int {
	*arr = append([]T{item}, *arr...)
	return len(*arr)
}

func Pop[T any](arr *[]T) T {
	lastIndex := len(*arr) - 1
	lastItem := (*arr)[lastIndex]

	*arr = (*arr)[:lastIndex]

	return lastItem
}

func Shift[T any](arr *[]T) T {
	firstItem := (*arr)[0]
	*arr = (*arr)[1:]
	return firstItem
}

func Reverse[T any](arr *[]T) []T {
	for i, j := 0, len(*arr)-1; i < j; i, j = i+1, j-1 {
		(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
	}

	return *arr
}

/*
*
{1, 2, 3, 4, 5} = s => { 1, 2, 300, 5 } { 4, 5 }

	array.Splice(s, 2, 2, 300)
*/
func Splice[T any](arr *[]T, tIdx, dCount int, items ...T) []T {
	var deletedItems []T

	for i := 0; i < len(*arr); i++ {
		if i >= tIdx && i <= tIdx+(dCount-1) {
			deletedItems = append(deletedItems, (*arr)[i])
		}
	}

	arrTmp := append((*arr)[:tIdx], items...)
	*arr = append(arrTmp, (*arr)[tIdx+dCount:]...)

	return deletedItems
}

func Slice[T any](arr []T, st, end int) []T {
	var newArr []T

	newArr = arr[st:end]
	return newArr
}

func Find[T any](arr []T, callback func(item T) bool) *T {
	for i := 0; i < len(arr); i++ {
		isTrue := callback(arr[i])

		if isTrue {
			return &arr[i]
		}
	}

	return nil
}

func FindIndex[T any](arr []T, callback func(item T) bool) int {
	for i := 0; i < len(arr); i++ {
		isTrue := callback(arr[i])

		if isTrue {
			return i
		}
	}

	return -1
}
