package main

import (
	"bytes"
	"excelgo/util"
	"flag"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

var infile *string = flag.String("i", "infile", "要转换的excel文件名")
var targetSheet *string = flag.String("s", "Sheet1", "要转换的Sheet，默认值为Sheet1")
var outfile *string = flag.String("o", "output.json", "目标json文件名")
var headfirst *bool = flag.Bool("h", true, "是否把excel文件的第一行设置为json目标文件的属性名")

func main() {
	flag.Parse()

	// 必须加参数
	if len(flag.Args()) != 0 {
		flag.Usage()
		return
	}

	// 必须有输入文件的参数
	if *infile == "infile" {
		// 请输入要转换的excel文件名称: xxx.xlsx
		fmt.Println("请输入要转换的excel文件名称: xxx.xlsx")
		return
	}

	f, err := excelize.OpenFile(*infile)
	if err != nil {
		fmt.Println("请确保您的文件格式正确,错误信息如下所示")
		fmt.Println(err)
		return
	}

	rows, err := f.GetRows(*targetSheet)
	// 处理标题:测试通过
	fields := []string{}
	// 如果第一行是标题，那么设置fields为第一行的内容，否则，默认设置fields1,fields2...为标题
	if *headfirst {
		head := rows[0]
		for i := 0; i < len(head); i++ {
			fields = append(fields, head[i])
		}
	} else {
		for i := 0; i < len(rows[0]); i++ {
			fields = append(fields, fmt.Sprintf("Field%d", i+1))
		}
	}

	jsonData := bytes.NewBuffer([]byte{})
	for i, row := range rows {
		if i == 0 && *headfirst {
			continue
		}
		jsonObj, err := util.String2json(fields, row)
		if err != nil {
			panic(err)
		}
		jsonData.Write(jsonObj)
		jsonData.WriteString(",\n")
	}

	fmt.Println(jsonData.String())
}

