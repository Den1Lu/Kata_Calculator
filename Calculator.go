package main
import (
"fmt"
"bufio"
"os"
"strings"
"strconv"
)
func main() {
scan := bufio.NewScanner(os.Stdin)
_ = scan.Scan()
a := scan.Text() // Приняли строку
b := strings.Split(a, " ") // Разбили строку на строчный массив
c := len(b) // Создали переменную, инициализировали ее длинной строчного массива
var rez int // Создали три переменные, бля чисел и для результата вычисления
var rimrez, rimrez1, rimrez2 string // Добавляем три переменные для вывода результата в римских цифрах
k1, _ := strconv.Atoi(string(b[0])) // Преобразуем элементы строчного массива в число
k2, _ := strconv.Atoi(string(b[2]))
if c > 3 || c < 3 {
panic("ошибка ввода") // Если строчный массив состоит бельше чем из трех значений то выводим panic и завершаем программу
}
if (k1 == 0 && k2 != 0) || (k1 != 0 && k2 == 0) {
panic("ошибка ввода")
}
if k1 == 0 && k2 == 0 {
rimnum := map[string]int { // Создали карту и проинициализировали ее римскими цифрами
"I": 1,
"II": 2,
"III": 3,
"IV": 4,
"V": 5,
"VI": 6,
"VII": 7,
"VIII": 8,
"IX": 9,
"X": 10,
"XX": 20,
"XXX": 30,
"XL": 40,
"L": 50,
"LX": 60,
"LXX": 70,
"LXXX": 80,
"LC": 90,
"C": 100,
}
for key, value := range rimnum { // Через цикл прогнали карту и сравнили все ее ключи с первым значением нашего страчного массива
if string(b[0]) == key {
k1 = value // Если какой то ключ совпадает то переменной присваиваем значение этого ключа
}
if string(b[2]) == key { // Делаем все то же самое только с третьим значением нашего строчного массива
k2 = value
}
}
if k1 == 0 || k2 == 0 || k1 > 10 || k2 > 10 { // Проверка на то что в строке все цифры римские и на величину значени
panic("ошибка ввода")
}
switch string(b[1]) { // Делаем цикл через оператор switch (если второе значение массива равно одному из 4 знаков, присваиваем переменной rez итоговое значение вычисления)
case "-":
rez = k1 - k2
case "+":
rez = k1 + k2
case "*":
rez = k1 * k2
case "/":
rez = k1 / k2
}
if rez < 1 { // Проверка результата вычисления на арабских цифрах
panic("ошибка ввода")
}
if rez < 11 && rez > 99 {
for key, value := range rimnum {
if rez == value {
rimrez = key
fmt.Print(rimrez)
}
}
} else {
rez1 := rez % 10
rez = rez / 10
rez2 := rez * 10
for key, value := range rimnum {
if rez1 == value {
rimrez1 = key
}
}
for key, value := range rimnum {
if rez2 == value {
rimrez2 = key
}
}
rimrez3 := rimrez2 + rimrez1
fmt.Print(rimrez3)
}
} else {
if k1 < 1 || k2 < 1 || k1 > 10 || k2 > 10 { // Проверка на величину значения
panic("ошибка ввода")
}
switch string(b[1]) { // Делаем цикл через оператор switch (если второе значение массива равно одному из 4 знаков, присваиваем переменной rez итоговое значение вычисления)
case "-":
rez = k1 - k2
case "+":
rez = k1 + k2
case "*":
rez = k1 * k2
case "/":
rez = k1 / k2
}
fmt.Print(rez)
}
}
    

    
 
    

