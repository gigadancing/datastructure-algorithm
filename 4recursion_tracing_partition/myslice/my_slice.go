package myslice

type MySlice []int

func (ms *MySlice) Len() int {
	return len(*ms)
}

func (ms *MySlice) Swap(i, j int) {
	(*ms)[i], (*ms)[j] = (*ms)[j], (*ms)[i]
}

func (ms *MySlice) Less(i, j int) bool {
	return (*ms)[i] < (*ms)[j]
}
