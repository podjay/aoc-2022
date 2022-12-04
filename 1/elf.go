package main

func NewElf() Elf {
	return Elf{
		calories: make([]int64, 0),
	}
}

type Calories []int64

type Elf struct {
	calories Calories
}

func (e *Elf) AddCalorie(calorie int64) {
	e.calories = append(e.calories, calorie)
}

func (e *Elf) GetCalorieCount() int64 {
	total := int64(0)
	for _, calorie := range e.calories {
		total += calorie
	}
	return total
}
