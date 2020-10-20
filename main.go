package main

//func main() {
//	f := excelize.NewFile()
//	// A、B、C、D、E、F、G、H、I、J、K、L、M、N、O、P、Q、R、S、T、U、V、W、X、Y、Z
//	column :=[]string{"A","B","C","D","E","F","G","H","I","J","K","L","M","N","O","P","Q","R","S","T","U","V","W","X","Y","Z","AA","AB","AC","AD","AE","AF","AG","AH","AI","AJ","AK","AL","AM","AN","AO","AP","AQ","AR","AS","AT","AU","AV","AW","AX","AY","AZ"}
//	title :=[]string{"ID","title","num","name","age"}
//	// fields :=[]string{"ID","title","num","name","age"}
//
//	type data struct {
//		ID int
//		title string
//		num int
//		name string
//		age int
//	}
//	list := [3]data{}
//	list[0] = data{
//		ID : 1,
//		title : "阿斯顿发生好看多了",
//		num:  2,
//		name:  "212",
//		age:   2,
//	}
//	list[1] =data{
//		ID : 2,
//		title : "四点发开胡老师",
//		num:  3,
//		name:  "212",
//		age:   2,
//	}
//	list[2] = data{
//		ID : 3,
//		title : "奥德赛克静安寺",
//		num: 4,
//		name:  "212",
//		age:   2,
//	}
//
//
//	// Create a new sheet.
//	index := f.NewSheet("Sheet1")
//	// Set value of a cell.
//	for k,v :=range title {
//		f.SetCellValue("Sheet1", column[k]+"1", v)
//	}
//
//	for k,v :=range list {
//		i :=k+2
//		s:= strconv.Itoa(i)
//		val :=reflect.ValueOf(v)
//		num := val.NumField()
//		// fmt.Println(column[0]+s)
//		for i:=0;i<num ;i++  {
//			f.SetCellValue("Sheet1", column[i]+s,val.Field(i))
//		}
//	}
//
//	// Set active sheet of the workbook.
//	f.SetActiveSheet(index)
//	// Save xlsx file by the given path.
//	if err := f.SaveAs("Book1.xlsx"); err != nil {
//		fmt.Println(err)
//	}
//}
// func main() {
//	f, err := excelize.OpenFile("Book1.xlsx")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	// Get value from cell by given worksheet name and axis.
//	//cell := f.GetCellValue("Sheet1", "B2")
//	//
//	//fmt.Println(cell)
//	// Get all the rows in the Sheet1.
//	rows:= f.GetRows("Sheet1")
//	for k, row := range rows {
//		if k>0 {
//			for _, colCell := range row {
//
//				fmt.Println(colCell)
//			}
//		}
//
//		fmt.Println()
//	}
// }


// f.SetCellValue("Sheet1", "A1", "日期")
// f.SetCellValue("Sheet1", "A2", "")
// f.SetCellValue("Sheet1", "B1", "首套自住")
// f.SetCellValue("Sheet1", "B2", "虎鲸添加量")
// f.SetCellValue("Sheet1", "C2", "推销售端量")
// f.SetCellValue("Sheet1", "D1", "投资机会")
// f.SetCellValue("Sheet1", "D2", "虎鲸添加量")
// f.SetCellValue("Sheet1", "E2", "推销售端量")
// f.SetCellValue("Sheet1", "F2", "精准量")
// f.SetCellValue("Sheet1", "G2", "未回复量")
//
// f.MergeCell("Sheet1", "A1", "A2")
// f.MergeCell("Sheet1", "B1", "C1")
// f.MergeCell("Sheet1", "D1", "G1")
// style,err :=f.NewStyle(`{"alignment":{"horizontal":"center","vertical":"center"}}`)
// if err != nil {
// println(err.Error())
// }
// f.SetCellStyle("Sheet1", "A1", "A2", style)
// f.SetCellStyle("Sheet1", "B1", "C1", style)
// f.SetCellStyle("Sheet1", "D1", "G1", style)
