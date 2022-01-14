package main

import (
	"encoding/json"
	"fmt"
	"github.com/xuri/excelize/v2"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/download", download)
	http.ListenAndServe(":8080", nil)
}

// 上传文件
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	filename := header.Filename
	println(filename)
	//导入文件到数据库
	rows, head := importExcel(io.Reader(file))

	bytes, err := json.Marshal(rows)
	println(string(bytes))
	marshal, err := json.Marshal(head)
	println(string(marshal))
}

// 下载
func download(w http.ResponseWriter, r *http.Request) {
	//示例 获取数据二维数组
	f, err := excelize.OpenFile("excel/Book1.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	head := rows[0]
	rowValues := rows[1:]
	//示例 获取数据二维数组结束

	ff := exportExcel(head, rowValues)
	w.Header().Add("Content-Type", "application/octet-stream")
	w.Header().Add("Content-Disposition", "attachment; filename=\""+"1.xlsx"+"\"")
	if err != nil {
		fmt.Println("Read File Err:", err.Error())
	} else {
		//服务器传输文件
		ff.Write(w)
	}
}

// 导出excel
func exportExcel(head []string, rowValues [][]string) *excelize.File {
	//新建文件
	ff := excelize.NewFile()
	defer func() {
		// Close the spreadsheet.
		if err := ff.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	for i := range head {
		//写入表头
		switch i {
		case 0:
			ff.SetCellValue("Sheet1", fmt.Sprintf("A%d", 1), head[i])
			break
		case 1:
			ff.SetCellValue("Sheet1", fmt.Sprintf("B%d", 1), head[i])
			break
		}
		//写入数据
		for j := range rowValues {
			//println(rowValues[j][i])
			switch i {
			case 0:
				ff.SetCellValue("Sheet1", fmt.Sprintf("A%d", j+2), rowValues[j][i])
				break
			case 1:
				ff.SetCellValue("Sheet1", fmt.Sprintf("B%d", j+2), rowValues[j][i])
				break
			}
		}
	}
	//本地写入文件 可以不要
	err := ff.SaveAs("test.xlsx")
	if err != nil {
		fmt.Println("Read File Err:", err.Error())
	}
	return ff
}

// 读取excel
func importExcel(reader io.Reader) ([][]string, []string) {

	f, err := excelize.OpenReader(reader)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// Get value from cell by given worksheet name and axis.
	cell, err := f.GetCellValue("Sheet1", "B2")
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	fmt.Println(cell)
	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	//表头
	head := rows[0]
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
	// rows[1:] 数据  head 表头
	return rows[1:], head
}
