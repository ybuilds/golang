package utils

func Swap(x *int, y *int) {
	var temp int = *y;
	*y = *x;
	*x = temp;
}