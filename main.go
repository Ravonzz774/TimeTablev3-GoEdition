package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/antchfx/htmlquery"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/html"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(cors.Default())

	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static/")
	// Конечная точка для получения расписания
	r.GET("/", func(c *gin.Context) {
		// Создание экземпляра расписания
		schedule := updateSchedule()
		// Отправка HTML-страницы на основе шаблона
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Date1": schedule.DateList[0].Monday,
			"Date2": schedule.DateList[0].Tuesday,
			"Date3": schedule.DateList[0].Wednesday,
			"Date4": schedule.DateList[0].Thursday,
			"Date5": schedule.DateList[0].Friday,
			"Date6": schedule.DateList[0].Saturday,

			"Monday":    schedule.Monday,
			"Tuesday":   schedule.Tuesday,
			"Wednesday": schedule.Wednesday,
			"Thursday":  schedule.Thursday,
			"Friday":    schedule.Friday,
			"Saturday":  schedule.Saturday,

		})
	})

	r.GET("/next", func(c *gin.Context) {
		// Создание экземпляра расписания
		schedule := updateNextWeekSchedule()
		// Отправка HTML-страницы на основе шаблона
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Date1": schedule.DateList[0].Monday,
			"Date2": schedule.DateList[0].Tuesday,
			"Date3": schedule.DateList[0].Wednesday,
			"Date4": schedule.DateList[0].Thursday,
			"Date5": schedule.DateList[0].Friday,
			"Date6": schedule.DateList[0].Saturday,

			"Monday":    schedule.Monday,
			"Tuesday":   schedule.Tuesday,
			"Wednesday": schedule.Wednesday,
			"Thursday":  schedule.Thursday,
			"Friday":    schedule.Friday,
			"Saturday":  schedule.Saturday,

		})
	})

	// Запускаем веб-сервер
	go func() {
		if err := r.Run(":80"); err != nil {
			log.Fatal(err)
		}
	}()

	// Устанавливаем таймер для обновления каждые 5 минут
	ticker := time.NewTicker(5 * time.Minute)

	// Запускаем бесконечный цикл, который будет обновлять расписание по истечении каждого интервала времени
	for range ticker.C {
		updateSchedule()
		updateNextWeekSchedule()
	}
}

func updateSchedule() Schedule {

	currentWeekNumber := getCurrentWeekNumber()

	// Конвертируем номер следующей недели в формат, используемый в URL
	url := fmt.Sprintf("https://is.radiotech.su/blocks/manage_groups/website/view.php?gr=151&week=%d&dep=1",
		currentWeekNumber)

	doc, err := htmlquery.LoadURL(url)
	if err != nil {
		log.Fatal(err)
	}

	schedule := Schedule{}

	//Понедельник
	for i := 1; i <= 6; i++ {
		lesson := Lesson{
			Name:       removeOtherName(searchTag(doc, fmt.Sprintf("/html/body/div[1]/div[3]/div[1]/table/tbody/tr/td[1]/div[2]/div[%d]/div[2]/div[1]/span", i))),
			Instructor: removeOtherName(searchTag(doc, fmt.Sprintf("/html/body/div[1]/div[3]/div[1]/table/tbody/tr/td[1]/div[2]/div[%d]/div[2]/div[2]/div[1]/span", i))),
			StartTime:  removeOtherName(searchTag(doc, fmt.Sprintf("/html/body/div[1]/div[3]/div[1]/table/tbody/tr/td[1]/div[2]/div[%d]/div[1]/div[2]", i))),
			EndTime:    removeOtherName(searchTag(doc, fmt.Sprintf("/html/body/div[1]/div[3]/div[1]/table/tbody/tr/td[1]/div[2]/div[%d]/div[1]/div[3]", i))),
			Classroom:  removeOtherName(searchTag(doc, fmt.Sprintf("/html/body/div[1]/div[3]/div[1]/table/tbody/tr/td[1]/div[2]/div[%d]/div[2]/div[2]/div[2]/span", i))),
			Number:     fmt.Sprintf("%d", i),
		}

		schedule.Monday = append(schedule.Monday, lesson)

	}

	//Вторник
	for i := 1; i <= 6; i++ {
		lesson := Lesson{
			Name:       removeOtherName(searchTag(doc, fmt.Sprintf("/html/body/div[1]/div[3]/div[1]/table/tbody/tr/td[2]/div[2]/div[%d]/div[2]/div[1]/span", i))),
			Instructor: removeOtherName(searchTag(doc, fmt.Sprintf("/html/body/div[1]/div[3]/div[1]/table/tbody/tr/td[2]/div[2]/div[%d]/div[2]/div[2]/div[1]/span", i))),
			StartTime:  removeOtherName(searchTag(doc, fmt.Sprintf("/html/body/div[1]/div[3]/div[1]/table/tbody/tr/td[2]/div[2]/div[%d]/div[1]/div[2]", i))),
			EndTime:    removeOtherName(searchTag(doc, fmt.Sprintf("/html/body/div[1]/div[3]/div[1]/table/tbody/tr/td[2]/div[2]/div[%d]/div[1]/div[3]", i))),
			Classroom:  removeOtherName(searchTag(doc, fmt.Sprintf("/html/body/div[1]/div[3]/div[1]/table/tbody/tr/td[2]/div[2]/div[%d]/div[2]/div[2]/div[2]/span", i))),
			Number:     fmt.Sprintf("%d", i),
		}

		schedule.Tuesday = append(schedule.Tuesday, lesson)
	}

	//Среда
	for i := 1; i <= 6; i++ {
		lesson := Lesson{
			Name:       removeOtherName(searchTag(doc, fmt.Sprintf("/html/body/div[1]/div[3]/div[1]/table/tbody/tr/td[3]/div[2]/div[%d]/div[2]/div[1]/span", i))),
			Instructor: removeOtherName(searchTag(doc, fmt.Sprintf("/html/body/div[1]/div[3]/div[1]/table/tbody/tr/td[3]/div[2]/div[%d]/div[2]/div[2]/div[1]/span", i))),
			StartTime:  removeOtherName(searchTag(doc, fmt.Sprintf("/html/body/div[1]/div[3]/div[1]/table/tbody/tr/td[3]/div[2]/div[%d]/div[1]/div[2]", i))),
			EndTime:    removeOtherName(searchTag(doc, fmt.Sprintf("/html/body/div[1]/div[3]/div[1]/table/tbody/tr/td[3]/div[2]/div[%d]/div[1]/div[3]", i))),
			Classroom:  removeOtherName(searchTag(doc, fmt.Sprintf("/html/body/div[1]/div[3]/div[1]/table/tbody/tr/td[3]/div[2]/div[%d]/div[2]/div[2]/div[2]/span", i))),
			Number:     fmt.Sprintf("%d", i),
		}

		schedule.Wednesday = append(schedule.Wednesday, lesson)

	}

	//Четверг
	for i := 1; i <= 6; i++ {
		lesson := Lesson{
			Name:       removeOtherName(searchTag(doc, fmt.Sprintf("/html/body/div[1]/div[3]/div[2]/table/tbody/tr/td[1]/div[2]/div[%d]/div[2]/div[1]/span", i))),
			Instructor: removeOtherName(searchTag(doc, fmt.Sprintf("/html/body/div[1]/div[3]/div[2]/table/tbody/tr/td[1]/div[2]/div[%d]/div[2]/div[2]/div[1]/span", i))),
			StartTime:  removeOtherName(searchTag(doc, fmt.Sprintf("/html/body/div[1]/div[3]/div[2]/table/tbody/tr/td[1]/div[2]/div[%d]/div[1]/div[2]", i))),
			EndTime:    removeOtherName(searchTag(doc, fmt.Sprintf("/html/body/div[1]/div[3]/div[2]/table/tbody/tr/td[1]/div[2]/div[%d]/div[1]/div[3]", i))),
			Classroom:  removeOtherName(searchTag(doc, fmt.Sprintf("/html/body/div[1]/div[3]/div[2]/table/tbody/tr/td[1]/div[2]/div[%d]/div[2]/div[2]/div[2]/span", i))),
			Number:     fmt.Sprintf("%d", i),
		}

		schedule.Thursday = append(schedule.Thursday, lesson)

	}

	//Пятница
	for i := 1; i <= 6; i++ {
		lesson := Lesson{
			Name:       removeOtherName(searchTag(doc, fmt.Sprintf("/html/body/div[1]/div[3]/div[2]/table/tbody/tr/td[2]/div[2]/div[%d]/div[2]/div[1]/span", i))),
			Instructor: removeOtherName(searchTag(doc, fmt.Sprintf("/html/body/div[1]/div[3]/div[2]/table/tbody/tr/td[2]/div[2]/div[%d]/div[2]/div[2]/div[1]/span", i))),
			StartTime:  removeOtherName(searchTag(doc, fmt.Sprintf("/html/body/div[1]/div[3]/div[2]/table/tbody/tr/td[2]/div[2]/div[%d]/div[1]/div[2]", i))),
			EndTime:    removeOtherName(searchTag(doc, fmt.Sprintf("/html/body/div[1]/div[3]/div[2]/table/tbody/tr/td[2]/div[2]/div[%d]/div[1]/div[3]", i))),
			Classroom:  removeOtherName(searchTag(doc, fmt.Sprintf("/html/body/div[1]/div[3]/div[2]/table/tbody/tr/td[2]/div[2]/div[%d]/div[2]/div[2]/div[2]/span", i))),
			Number:     fmt.Sprintf("%d", i),
		}

		schedule.Friday = append(schedule.Friday, lesson)
	}

	//Суббота
	for i := 1; i <= 6; i++ {
		lesson := Lesson{
			Name:       removeOtherName(searchTag(doc, fmt.Sprintf("/html/body/div[1]/div[3]/div[2]/table/tbody/tr/td[3]/div[2]/div[%d]/div[2]/div[1]/span", i))),
			Instructor: removeOtherName(searchTag(doc, fmt.Sprintf("/html/body/div[1]/div[3]/div[2]/table/tbody/tr/td[3]/div[2]/div[%d]/div[2]/div[2]/div[1]/span", i))),
			StartTime:  removeOtherName(searchTag(doc, fmt.Sprintf("/html/body/div[1]/div[3]/div[2]/table/tbody/tr/td[3]/div[2]/div[%d]/div[1]/div[2]", i))),
			EndTime:    removeOtherName(searchTag(doc, fmt.Sprintf("/html/body/div[1]/div[3]/div[2]/table/tbody/tr/td[3]/div[2]/div[%d]/div[1]/div[3]", i))),
			Classroom:  removeOtherName(searchTag(doc, fmt.Sprintf("/html/body/div[1]/div[3]/div[2]/table/tbody/tr/td[3]/div[2]/div[%d]/div[2]/div[2]/div[2]/span", i))),
			Number:     fmt.Sprintf("%d", i),
		}

		schedule.Saturday = append(schedule.Saturday, lesson)
	}

	datelist := DateList{
		Monday:    searchTag(doc, "/html/body/div[1]/div[3]/div[1]/table/tbody/tr/td[1]/div[1]/span"),
		Tuesday:   searchTag(doc, "/html/body/div[1]/div[3]/div[1]/table/tbody/tr/td[2]/div[1]/span"),
		Wednesday: searchTag(doc, "/html/body/div[1]/div[3]/div[1]/table/tbody/tr/td[3]/div[1]/span"),
		Thursday:  searchTag(doc, "/html/body/div[1]/div[3]/div[2]/table/tbody/tr/td[1]/div[1]/span"),
		Friday:    searchTag(doc, "/html/body/div[1]/div[3]/div[2]/table/tbody/tr/td[2]/div[1]/span"),
		Saturday:  searchTag(doc, "/html/body/div[1]/div[3]/div[2]/table/tbody/tr/td[3]/div[1]/span"),
	}

	schedule.DateList = append(schedule.DateList, datelist)

	return schedule
}

func updateNextWeekSchedule() Schedule {
	currentWeekNumber := getCurrentWeekNumber()
	nextWeekNumber := currentWeekNumber + 1

	// Конвертируем номер следующей недели в формат, используемый в URL
	nextWeekURL := fmt.Sprintf("https://is.radiotech.su/blocks/manage_groups/website/view.php?gr=151&week=%d&dep=1",
		nextWeekNumber)

	// Загружаем HTML следующей недели
	doc, err := htmlquery.LoadURL(nextWeekURL)
	if err != nil {
		log.Fatal(err)
	}

	schedule := Schedule{}

	//Понедельник
	for i := 1; i <= 6; i++ {
		lesson := Lesson{
			Name:       removeOtherName(searchTag(doc, fmt.Sprintf("/html/body/div[1]/div[3]/div[1]/table/tbody/tr/td[1]/div[2]/div[%d]/div[2]/div[1]/span", i))),
			Instructor: removeOtherName(searchTag(doc, fmt.Sprintf("/html/body/div[1]/div[3]/div[1]/table/tbody/tr/td[1]/div[2]/div[%d]/div[2]/div[2]/div[1]/span", i))),
			StartTime:  removeOtherName(searchTag(doc, fmt.Sprintf("/html/body/div[1]/div[3]/div[1]/table/tbody/tr/td[1]/div[2]/div[%d]/div[1]/div[2]", i))),
			EndTime:    removeOtherName(searchTag(doc, fmt.Sprintf("/html/body/div[1]/div[3]/div[1]/table/tbody/tr/td[1]/div[2]/div[%d]/div[1]/div[3]", i))),
			Classroom:  removeOtherName(searchTag(doc, fmt.Sprintf("/html/body/div[1]/div[3]/div[1]/table/tbody/tr/td[1]/div[2]/div[%d]/div[2]/div[2]/div[2]/span", i))),
			Number:     fmt.Sprintf("%d", i),
		}

		schedule.Monday = append(schedule.Monday, lesson)
	}

	//Вторник
	for i := 1; i <= 6; i++ {
		lesson := Lesson{
			Name:       removeOtherName(searchTag(doc, fmt.Sprintf("/html/body/div[1]/div[3]/div[1]/table/tbody/tr/td[2]/div[2]/div[%d]/div[2]/div[1]/span", i))),
			Instructor: removeOtherName(searchTag(doc, fmt.Sprintf("/html/body/div[1]/div[3]/div[1]/table/tbody/tr/td[2]/div[2]/div[%d]/div[2]/div[2]/div[1]/span", i))),
			StartTime:  removeOtherName(searchTag(doc, fmt.Sprintf("/html/body/div[1]/div[3]/div[1]/table/tbody/tr/td[2]/div[2]/div[%d]/div[1]/div[2]", i))),
			EndTime:    removeOtherName(searchTag(doc, fmt.Sprintf("/html/body/div[1]/div[3]/div[1]/table/tbody/tr/td[2]/div[2]/div[%d]/div[1]/div[3]", i))),
			Classroom:  removeOtherName(searchTag(doc, fmt.Sprintf("/html/body/div[1]/div[3]/div[1]/table/tbody/tr/td[2]/div[2]/div[%d]/div[2]/div[2]/div[2]/span", i))),
			Number:     fmt.Sprintf("%d", i),
		}

		schedule.Tuesday = append(schedule.Tuesday, lesson)
	}

	//Среда
	for i := 1; i <= 6; i++ {
		lesson := Lesson{
			Name:       removeOtherName(searchTag(doc, fmt.Sprintf("/html/body/div[1]/div[3]/div[1]/table/tbody/tr/td[3]/div[2]/div[%d]/div[2]/div[1]/span", i))),
			Instructor: removeOtherName(searchTag(doc, fmt.Sprintf("/html/body/div[1]/div[3]/div[1]/table/tbody/tr/td[3]/div[2]/div[%d]/div[2]/div[2]/div[1]/span", i))),
			StartTime:  removeOtherName(searchTag(doc, fmt.Sprintf("/html/body/div[1]/div[3]/div[1]/table/tbody/tr/td[3]/div[2]/div[%d]/div[1]/div[2]", i))),
			EndTime:    removeOtherName(searchTag(doc, fmt.Sprintf("/html/body/div[1]/div[3]/div[1]/table/tbody/tr/td[3]/div[2]/div[%d]/div[1]/div[3]", i))),
			Classroom:  removeOtherName(searchTag(doc, fmt.Sprintf("/html/body/div[1]/div[3]/div[1]/table/tbody/tr/td[3]/div[2]/div[%d]/div[2]/div[2]/div[2]/span", i))),
			Number:     fmt.Sprintf("%d", i),
		}

		schedule.Wednesday = append(schedule.Wednesday, lesson)
	}

	//Четверг
	for i := 1; i <= 6; i++ {
		lesson := Lesson{
			Name:       removeOtherName(searchTag(doc, fmt.Sprintf("/html/body/div[1]/div[3]/div[2]/table/tbody/tr/td[1]/div[2]/div[%d]/div[2]/div[1]/span", i))),
			Instructor: removeOtherName(searchTag(doc, fmt.Sprintf("/html/body/div[1]/div[3]/div[2]/table/tbody/tr/td[1]/div[2]/div[%d]/div[2]/div[2]/div[1]/span", i))),
			StartTime:  removeOtherName(searchTag(doc, fmt.Sprintf("/html/body/div[1]/div[3]/div[2]/table/tbody/tr/td[1]/div[2]/div[%d]/div[1]/div[2]", i))),
			EndTime:    removeOtherName(searchTag(doc, fmt.Sprintf("/html/body/div[1]/div[3]/div[2]/table/tbody/tr/td[1]/div[2]/div[%d]/div[1]/div[3]", i))),
			Classroom:  removeOtherName(searchTag(doc, fmt.Sprintf("/html/body/div[1]/div[3]/div[2]/table/tbody/tr/td[1]/div[2]/div[%d]/div[2]/div[2]/div[2]/span", i))),
			Number:     fmt.Sprintf("%d", i),
		}

		schedule.Thursday = append(schedule.Thursday, lesson)

	}

	//Пятница
	for i := 1; i <= 6; i++ {
		lesson := Lesson{
			Name:       removeOtherName(searchTag(doc, fmt.Sprintf("/html/body/div[1]/div[3]/div[2]/table/tbody/tr/td[2]/div[2]/div[%d]/div[2]/div[1]/span", i))),
			Instructor: removeOtherName(searchTag(doc, fmt.Sprintf("/html/body/div[1]/div[3]/div[2]/table/tbody/tr/td[2]/div[2]/div[%d]/div[2]/div[2]/div[1]/span", i))),
			StartTime:  removeOtherName(searchTag(doc, fmt.Sprintf("/html/body/div[1]/div[3]/div[2]/table/tbody/tr/td[2]/div[2]/div[%d]/div[1]/div[2]", i))),
			EndTime:    removeOtherName(searchTag(doc, fmt.Sprintf("/html/body/div[1]/div[3]/div[2]/table/tbody/tr/td[2]/div[2]/div[%d]/div[1]/div[3]", i))),
			Classroom:  removeOtherName(searchTag(doc, fmt.Sprintf("/html/body/div[1]/div[3]/div[2]/table/tbody/tr/td[2]/div[2]/div[%d]/div[2]/div[2]/div[2]/span", i))),
			Number:     fmt.Sprintf("%d", i),
		}

		schedule.Friday = append(schedule.Friday, lesson)
	}

	//Суббота
	for i := 1; i <= 6; i++ {
		lesson := Lesson{
			Name:       removeOtherName(searchTag(doc, fmt.Sprintf("/html/body/div[1]/div[3]/div[2]/table/tbody/tr/td[3]/div[2]/div[%d]/div[2]/div[1]/span", i))),
			Instructor: removeOtherName(searchTag(doc, fmt.Sprintf("/html/body/div[1]/div[3]/div[2]/table/tbody/tr/td[3]/div[2]/div[%d]/div[2]/div[2]/div[1]/span", i))),
			StartTime:  removeOtherName(searchTag(doc, fmt.Sprintf("/html/body/div[1]/div[3]/div[2]/table/tbody/tr/td[3]/div[2]/div[%d]/div[1]/div[2]", i))),
			EndTime:    removeOtherName(searchTag(doc, fmt.Sprintf("/html/body/div[1]/div[3]/div[2]/table/tbody/tr/td[3]/div[2]/div[%d]/div[1]/div[3]", i))),
			Classroom:  removeOtherName(searchTag(doc, fmt.Sprintf("/html/body/div[1]/div[3]/div[2]/table/tbody/tr/td[3]/div[2]/div[%d]/div[2]/div[2]/div[2]/span", i))),
			Number:     fmt.Sprintf("%d", i),
		}

		schedule.Saturday = append(schedule.Saturday, lesson)
	}

	datelist := DateList{
		Monday:    searchTag(doc, "/html/body/div[1]/div[3]/div[1]/table/tbody/tr/td[1]/div[1]/span"),
		Tuesday:   searchTag(doc, "/html/body/div[1]/div[3]/div[1]/table/tbody/tr/td[2]/div[1]/span"),
		Wednesday: searchTag(doc, "/html/body/div[1]/div[3]/div[1]/table/tbody/tr/td[3]/div[1]/span"),
		Thursday:  searchTag(doc, "/html/body/div[1]/div[3]/div[2]/table/tbody/tr/td[1]/div[1]/span"),
		Friday:    searchTag(doc, "/html/body/div[1]/div[3]/div[2]/table/tbody/tr/td[2]/div[1]/span"),
		Saturday:  searchTag(doc, "/html/body/div[1]/div[3]/div[2]/table/tbody/tr/td[3]/div[1]/span"),
	}

	schedule.DateList = append(schedule.DateList, datelist)

	return schedule
}

func searchTag(doc *html.Node, xpath string) string {
	node, err := htmlquery.Query(doc, xpath)

	if err != nil {
		log.Fatal(err)
	}

	if node != nil {
		return htmlquery.InnerText(node)
	}

	return ""
}

type Schedule struct {
	Monday    []Lesson
	Tuesday   []Lesson
	Wednesday []Lesson
	Thursday  []Lesson
	Friday    []Lesson
	Saturday  []Lesson
	DateList  []DateList
}

type Lesson struct {
	Name       string
	Instructor string
	StartTime  string
	EndTime    string
	Classroom  string
	Number     string
}

type DateList struct {
	Monday    string
	Tuesday   string
	Wednesday string
	Thursday  string
	Friday    string
	Saturday  string
}

func removeOtherName(str string) string {
	result := strings.ReplaceAll(str, "(09.02.07)", "")
	result = strings.ReplaceAll(result, "ОГСЭ.05 ", "")
	return result
}

func getCurrentWeekNumber() int {
	currentYear, currentWeek := time.Now().ISOWeek()

	// Начало сентября
	septemberStartDate := time.Date(currentYear, time.September, 1, 0, 0, 0, 0, time.UTC)

	// Разность в неделях от начала сентября до текущей даты
	_, septemberStartWeek := septemberStartDate.ISOWeek()

	weekDifference := currentWeek - septemberStartWeek + 1

	return weekDifference
}

// /html/body/div[1]/div[3]/div[1]/table/tbody/tr/td[1]/div[2]/div[2]/div[2]/div[1]/span

// /html/body/div[1]/div[3]/div[1]/table/tbody/tr/td[1]/div[2]/div[2]/div[2]/div[1]/span
