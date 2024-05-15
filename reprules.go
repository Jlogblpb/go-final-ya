package main

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

func NextDate(now time.Time, date string, repeat string) (string, error) {
	var rep []string
	//Проверяем есть ли правило повторения, если отсутствует возвращаем ошибку
	if len(repeat) > 0 { //Если привило больше больше 0, продолжаем работу если
		//меньше возвращаем ошибку
		rep = strings.Split(repeat, " ") //Разбираем правило на символы через пробел
	} else {
		err := errors.New("Правило повторения отсутствует")
		return "", err // больше возвращаем ошибку
	}
	timBase, err := time.Parse("20060102", date) //Получаем дату из строки
	if err != nil {                              //если строка с датой не корректная
		return "", err //возвращаем ошибку
	}
	//Проверяем режим работы год или день
	if rep[0] == "y" { //Если год то прибавляем к изначалной дате год проверяя
		// что полученная дата будет больше чем сегодняшняя
		for i := 0; timBase.Before(now); i++ { //ненравится переменная t хотя они одного типа
			timBase = timBase.AddDate(1, 0, 0)
		}
		result := timBase.Format("20060102")
		return result, nil
	}
	if rep[0] == "d" { // Если не год, то день, но проверяем вводимые джанные на приавильность
		days, err := strconv.Atoi(rep[1]) // преобразуем строку в число, количество дней
		if err != nil {
			return "", err // Если вместо дней данные не корректные возвращаем ошибку
		}
		if days > 400 { //максимальное количество дней для переноса собития 400 если дней
			err := errors.New("Перенос события более чем на 400 дней недопустим")
			return "", err // больше возвращаем ошибку
		}
		for i := 0; timBase.Before(now); i++ { //ненравится переменная t хотя они одного типа
			timBase = timBase.AddDate(0, 0, days)
		}
		result := timBase.Format("20060102")
		return result, nil
	}
	return "", err
}
