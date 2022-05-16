package api

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

var cells = Cells{}

type (
	Cells map[string]Cell
	Cell  struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
		ID   string `json:"id"`
		Sex  string `json:"sex"`
	}
)

func (c *Cell) Parse(f *excelize.File) {
	sheetName := "Sheet1"
	c.Name, _ = f.GetCellValue(sheetName, "B1")
	s, _ := f.GetCellValue(sheetName, "B2")
	c.Age, _ = strconv.Atoi(s)
	c.ID, _ = f.GetCellValue(sheetName, "B3")
	c.Sex, _ = f.GetCellValue(sheetName, "B4")
}

func init() {
	go func() {
		filepath.Walk("./data",
			func(path string, info os.FileInfo, err error) error {
				var (
					cell     = new(Cell)
					filename = filepath.Base(path)
					ext      = ".xlsx"
				)
				if filepath.Ext(path) == ext {
					f, _ := excelize.OpenFile(path)
					cell.Parse(f)
					id := filename[:len(filename)-len(ext)]
					fmt.Printf("%v", cell)
					cells[id] = *cell
				}
				return nil
			})
	}()
}

func FetchMap(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, cells)
}

func FetchList(c *gin.Context) {
	var (
		i     int
		array = make([]string, len(cells))
	)
	for k, e := range cells {
		array[i] = fmt.Sprintf("%s %v", k, e)
		i++
	}
	c.IndentedJSON(http.StatusOK, array)
}
