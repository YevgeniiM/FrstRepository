package main

import (
	"fmt"
)

/*
Задача 1. Баллы ЕГЭ
Сумма проходных баллов в институт составляет 275 баллов.
Напишите программу, которая запрашивает у пользователя результаты ЕГЭ по трём экзаменам и проверяет, поступил ли он в институт или нет.
*/

func main() {
	passingScore := 500
	var subjectScore1 int
	fmt.Println("Введите Ваш бал по первому предмету:")
	fmt.Scan(&subjectScore1)
	var subjectScore2 int
	fmt.Println("Введите Ваш бал по второму предмету:")
	fmt.Scan(&subjectScore2)

	var subjectScore3 int
	fmt.Println("Введите Ваш бал по третьему предмету:")
	fmt.Scan(&subjectScore3)
	totalScores := subjectScore1 + subjectScore2 + subjectScore3
	if totalScores >= passingScore {
		fmt.Println("Поздравляем! Вы поступили! Вы набрали", totalScores, "баллов из", passingScore, "необходимых.")
	} else {

		fmt.Println("Для поступления в этом году Вам не хватило", passingScore-totalScores, "баллов до", passingScore, "необходимых. Ждем Вас в следующем году.")
	}
}
