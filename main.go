package main

import (
	"autotest/tcpconn"
	"autotest/uitogo"
	"os"

	"github.com/therecipe/qt/widgets"
)

func main() {
	// 界面
	widgets.NewQApplication(len(os.Args), os.Args)
	var mainWindow_ui uitogo.MY_UIMainWindow
	mainWindow_ui.SetupUI(nil)
	mainWindow_ui.Show()
	widgets.QApplication_Exec()
	for _, client := range tcpconn.Gbs_clientList {
		client.NetConn.Close()
	}

	//非界面，读取excel
	// basefunc.Gbs_log.Init()

	// var excelInfo fileopera.ExcelInfo
	// if excelInfo.OpenTestExcel("./自动化测试.xlsx") != nil {
	// 	return
	// }
	// err := excelInfo.ReadCellValue("内测tcp登录")
	// if err != nil {
	// 	return
	// }

	// connectCnt := len(excelInfo.ListExample)
	// if connectCnt > tcpconn.MaxConnect {
	// 	connectCnt = tcpconn.MaxConnect
	// }
	// clientSlice := make([]tcpconn.Client, connectCnt)
	// clientIndex := 0
	// for _, val := range excelInfo.ListExample {
	// 	if tcpconn.BuildConnect(&clientSlice[clientIndex], val.Method, val.Host) != nil {
	// 		return
	// 	}
	// 	clientIndex++
	// }

	// for {
	// 	clientIndex = 0
	// 	//var socketClient tcpconn.Client
	// 	for _, val := range excelInfo.ListExample {
	// 		switch val.Module {
	// 		case fileopera.Mod_Login:
	// 			{
	// 				tcpconn.SendGeneralLoginPkg(&clientSlice[clientIndex], val.ReqData)
	// 				clientIndex = (clientIndex + 1) % len(excelInfo.ListExample)
	// 				time.Sleep(time.Duration(100) * time.Millisecond)
	// 				break
	// 			}
	// 		}
	// 	}
	// }

	// time.Sleep(time.Duration(100000) * time.Second)
	// //socketClient.NetConn.Close()
	// for _, client := range clientSlice {
	// 	client.NetConn.Close()
	// }

}
