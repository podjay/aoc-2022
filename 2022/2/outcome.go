package main

type Outcome interface {
	Points() int
}

type Win struct {
}

func (w Win) Points() int {
	return 6
}

type Loss struct {
}

func (l Loss) Points() int {
	return 0
}

type Draw struct {
}

func (d Draw) Points() int {
	return 3
}
