package fileopera

import (
	"autotest/basefunc"
	"errors"

	"github.com/xuri/excelize"
)

type ExcelInfo struct {
	Obj         *excelize.File
	ListExample map[string]TestExample
}

type TestExample struct {
	Id          string
	Module      string
	CaseName    string
	Method      string
	Host        string
	ReqData     string
	AckRealData string
	ExecTime    string
}

// 表格纵列
const (
	Id           = "A"
	Module       = "B"
	CaseName     = "C"
	Method       = "D"
	Host         = "E"
	ReqData      = "F"
	AckAtendData = "G"
	AckRealData  = "H"
	ExecTime     = "I"
)

// module
const (
	Mod_Login = "登录"
)

func (excel *ExcelInfo) OpenTestExcel(excelName string) error {
	var err error
	excel.Obj, err = excelize.OpenFile(excelName)
	if err != nil {
		basefunc.Gbs_log.Printf("open excel:%s fail,err:%s", excelName, err.Error())
	}
	excel.ListExample = make(map[string]TestExample)
	return err
}

func (excel *ExcelInfo) ReadCellValue(sheetName string) error {
	if excel == nil {
		basefunc.Gbs_log.Printf("ReadCellValue fail,obj is nil")
		return errors.New("ReadCellValue fail,obj is nil")
	}
	rows, err := excel.Obj.GetRows(sheetName)
	if err != nil {
		basefunc.Gbs_log.Printf("getrows fail,err:%s", err.Error())
		return err
	}
	rowCnt := len(rows)

	titleList := rows[0]
	//colCnt := len(titleList)

	for i := 1; i < rowCnt; i++ {
		var data TestExample
		for j, title := range titleList {
			switch title {
			case "测试id":
				data.Id = rows[i][j]
				break
			case "模块":
				data.Module = rows[i][j]
				break
			case "用例名称":
				data.CaseName = rows[i][j]
				break
			case "协议":
				data.Method = rows[i][j]
				break
			case "请求地址":
				data.Host = rows[i][j]
				break
			case "请求数据":
				data.ReqData = rows[i][j]
				break
			}
		}
		excel.ListExample[data.Id] = data
	}

	return nil
	// for i := 1; i < rowCnt; i++ {
	// 	for j := 0; j < colCnt; j++ {
	// 		row[i][j]
	// 	}
	// }

}
