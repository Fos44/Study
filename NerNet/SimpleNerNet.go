package main

import (
	"fmt"
	"github.com/fxsjy/gonn/gonn"
)

func CreateNN()  {
	nn := gonn.DefaultNetwork(3, 16, 4, false)		//4 выходных нейрона, 16 скрытых нейронов

	input := [][] float64{
		{0.5, 1, 1}, {0.9, 1, 2}, {0.8, 0, 1},
		{0.3, 1, 1}, {0.6, 1, 2}, {0.4, 0, 1},
		{0.9, 1, 7}, {0.6, 1, 4}, {0.1, 0, 1},
		{0.6, 1, 0}, {1, 0, 0}}

	target := [][]float64{
		{1, 0, 0, 0}, {1, 0, 0, 0}, {1, 0, 0, 0},
		{0, 1, 0, 0}, {0, 1, 0, 0}, {0, 1, 0, 0},
		{0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 1, 0},
		{0, 0, 0, 1}, {0, 0, 0, 1}}

	nn.Train(input, target, 100000)

	gonn.DumpNN("FirstNN", nn)
	
}

func GetResult(output []float64) string  {
	var max float64 = -99999
	pos := -1

	//позиция нейрона с самым большим весом
	for i, value := range output {
		if value > max {
			max = value
			pos = i
		}
	}

	switch pos {
	case 0: return "Атаковать противника"
	case 1: return "Красться"
	case 2: return "Спасаться бегством"
	case 3: return "Пассивно защищаться"

		}
	return ""
}

func main() {

	CreateNN()
	nn := gonn.LoadNN("FirstNN")		//Загрузка НС из созданного файла

	var HealthPoint  = 0.7		//Здоровье
	var Weapon float64 = 1		//1 - есть оружие, 0 - без оружия
	var EnemyCount float64 = 1 //Колличество врагов

	out := nn.Forward([]float64{HealthPoint, Weapon, EnemyCount})

	fmt.Println(GetResult(out))

}
