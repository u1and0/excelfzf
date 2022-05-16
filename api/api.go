package api

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/xuri/excelize/v2"
)

var cells = Cells{}

type (
	Cells []Cell
	Cell  struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
		ID   string `json:"id"`
		Sex  string `json:"sex"`
	}
)

func (c *Cell) Parse(f *excelize.File) {
	sheetName := "Sheet1"
	c.Name, _ = f.GetCellValue(sheetName, "A1")
	s, _ := f.GetCellValue(sheetName, "A2")
	c.Age, _ = strconv.Atoi(s)
	c.ID = f.GetCellValue(sheetName, "A3")
	c.Sex = f.GetCellValue(sheetName, "A4")
}

func init() {
	go func() {
		filepath.Walk("./data",
			func(path string, info os.FileInfo, err error) error {
				var (
					cell     = new(Cell)
					filename = filepath.Base(path)
				)
				if filepath.Ext(path) == ".xlsx" {
					f, _ := excelize.OpenFile(path)
					cell.Parse(f)
					id := filename[:len(filename)-5]
					cell[id] = *cell
				}
				return nil
			})
	}()
}
