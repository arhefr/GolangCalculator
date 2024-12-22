package pkg

// Убирает элемент из слайса по индексу
func RemoveElement(slice []float64, s int) []float64 {
	return append(slice[:s], slice[s+1:]...)
}
