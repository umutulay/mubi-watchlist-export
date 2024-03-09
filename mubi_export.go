package main

import (
	"fmt"

	"github.com/jdennes/mubi"
	"github.com/xuri/excelize/v2"
)

func main() {
	api := mubi.NewMubiAPI()
	userId := int64(1) // Change this to the userId you want
	perPage := 25      // It always gets 25 contents, so it does not matter which number put in here
	f := excelize.NewFile()
	sheetName := "Sheet1"

	// Define column headers
	f.SetCellValue(sheetName, "A1", "Title")
	f.SetCellValue(sheetName, "B1", "Year")

	row := 2 // Start from the second row to account for headers

	/* every page has 25 films, so you can divide number of films in the watchlist by 25
	to find how many pages you have in watchlist*/
	for page := 1; page <= 1; page++ {
		fmt.Println("Page:", page)
		watchlistItems := api.GetWatchlist(userId, page, perPage)
		for _, item := range watchlistItems {
			// Write to Excel file
			f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), item.Film.Title)
			f.SetCellValue(sheetName, fmt.Sprintf("B%d", row), item.Film.Year)
			row++
		}
	}

	// Save the Excel file
	if err := f.SaveAs("Watchlist.xlsx"); err != nil {
		fmt.Println(err)
	}
}
