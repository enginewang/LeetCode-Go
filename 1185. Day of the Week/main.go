package main

func dayOfTheWeek(day int, month int, year int) string {
	dayName := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	start := 5
	monthDay := []int{31,28,31,30,31,30,31,31,30,31,30,31}
	addYear := (year - 1971)*1 + int((year - 1971 + 2)/4)
	start = (start + addYear)%7
	if (year%4 == 0 && year%100 !=0) || year%400 == 0{
		monthDay[1] = 29
	}
	days := 0
	for i := 0; i < month-1; i++{
		days += monthDay[i]
	}
	start = (start + days + day-1)%7
	return dayName[start]
}

func main() {
	dayOfTheWeek(1,3,2100)
}
