package main

import (
	"log"
	"net/http"
	"strings"
	"github.com/PuerkitoBio/goquery"
	"fmt"
	"strconv"
	"sort"
)

type Food struct {
	Station string
	Name string
	Cals int
}

func main() {
	// Make HTTP GET request
	response, err := http.Get("https://carleton.cafebonappetit.com/cafe/sayles-hill-cafe/2021-07-09")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	doc, _ := goquery.NewDocumentFromReader(response.Body)
	sel := doc.Find(".c-tab__content-inner")

	var menu []Food
	var cur_station, cur_name string

	choice := sel.Eq(8)
	choice.Find(`.site-panel__daypart-station-title,
							.site-panel__daypart-item-title,
							.site-panel__daypart-item-calories`).
		Each(func(_ int, e *goquery.Selection) {
			str := strings.ReplaceAll(strings.ReplaceAll(e.
				Text(), "\n", ""), "\t", "")
      if e.HasClass("site-panel__daypart-item-calories") {
        cal, err := strconv.Atoi(str[:len(str)-5])
				if err != nil {
					cal = -1
				}
				menu = append(menu, Food{cur_station, cur_name, cal})
      } else if e.HasClass("site-panel__daypart-station-title") {
				cur_station = str
			} else if e.HasClass("site-panel__daypart-item-title") {
				cur_name = str
			}
		})

	fmt.Print("Enter sorting criteria: type gtl/ltg/'filter' to enter calorie ranges. ")
	var input string
	fmt.Scanf("%s", &input)
	
	if input == "gtl" {
		sort.Slice(menu, func(i, j int) bool {
			return menu[i].Cals > menu[j].Cals
		})
		for i := 0; i < len(menu); i++ {
			fmt.Printf("%s \033[31m%d\033[0m\n", menu[i].Name, menu[i].Cals)
		}
	} else if input == "ltg" {
		sort.Slice(menu, func(i, j int) bool {
			return menu[i].Cals < menu[j].Cals
		})
		for i := 0; i < len(menu); i++ {
			fmt.Printf("%s \033[31m%d\033[0m\n", menu[i].Name, menu[i].Cals)
		}
    
	} else if input == "filter" {
		fmt.Print("Enter filter range: ")
		var ramin, ramax int
		fmt.Scanf("%d-%d", &ramin, &ramax)

    sort.Slice(menu, func(i, j int) bool {
			return menu[i].Cals > menu[j].Cals
		})
		for i := 0; i < len(menu); i++ {
			if menu[i].Cals >= ramin && menu[i].Cals <= ramax{
        fmt.Printf("%s \033[31m%d\033[0m\n", menu[i].Name, menu[i].Cals)
      }
		}
	} else {
    for i := 0; i < len(menu); i++ {
      fmt.Printf("%s \033[31m%d\033[0m\n", menu[i].Name, menu[i].Cals)
    }
  }

}