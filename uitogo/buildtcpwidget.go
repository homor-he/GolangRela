package uitogo

import (
	"autotest/image"
	"autotest/tcpconn"
	"strconv"
	"time"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type MY_UIBuildTcpWidget struct {
	Frame      *widgets.QFrame
	TextEdit   *widgets.QTextEdit
	PushButton *widgets.QPushButton
	Label      *widgets.QLabel

	Parent *MY_UIMainWindow
}

func (this *MY_UIBuildTcpWidget) SetupUI(Frame *widgets.QFrame) {
	Frame.SetObjectName("Frame")
	Frame.SetGeometry(core.NewQRect4(0, 0, 447, 319))
	var sizePolicy *widgets.QSizePolicy
	sizePolicy = widgets.NewQSizePolicy2(widgets.QSizePolicy__Fixed, widgets.QSizePolicy__Fixed, widgets.QSizePolicy__DefaultType)
	sizePolicy.SetHorizontalStretch(0)
	sizePolicy.SetVerticalStretch(0)
	sizePolicy.SetHeightForWidth(Frame.SizePolicy().HasHeightForWidth())
	Frame.SetSizePolicy(sizePolicy)
	this.TextEdit = widgets.NewQTextEdit(Frame)
	this.TextEdit.SetObjectName("TextEdit")
	this.TextEdit.SetGeometry(core.NewQRect4(110, 100, 271, 31))
	this.TextEdit.ConnectTextChanged(func() {
		this.TextEdit.SetFontPointSize(12)
	})

	this.PushButton = widgets.NewQPushButton(Frame)
	this.PushButton.SetObjectName("PushButton")
	this.PushButton.SetGeometry(core.NewQRect4(150, 200, 150, 40))
	//添加按钮事件
	this.PushButton.ConnectClicked(func(checked bool) {
		this.ButtonCheck()
	})

	this.Label = widgets.NewQLabel(Frame, core.Qt__Widget)
	this.Label.SetObjectName("Label")
	this.Label.SetGeometry(core.NewQRect4(40, 100, 61, 31))
	var font *gui.QFont
	font = gui.NewQFont()
	font.SetPointSize(15)
	font.SetWeight(50)
	this.Label.SetFont(font)
	this.Label.SetLayoutDirection(core.Qt__LeftToRight)

	this.RetranslateUi(Frame)

}

func (this *MY_UIBuildTcpWidget) RetranslateUi(Frame *widgets.QFrame) {
	_translate := core.QCoreApplication_Translate
	Frame.SetWindowTitle(_translate("Frame", "建立tcp连接", "", -1))
	this.TextEdit.SetHtml(_translate("Frame", "<!DOCTYPE HTML PUBLIC \"-//W3C//DTD HTML 4.0//EN\" \"http://www.w3.org/TR/REC-html40/strict.dtd\">\n<html><head><meta name=\"qrichtext\" content=\"1\" /><style type=\"text/css\">\np, li { white-space: pre-wrap; }\n</style></head><body style=\" font-family:'SimSun'; font-size:9pt; font-weight:400; font-style:normal;\">\n<p style=\" margin-top:0px; margin-bottom:0px; margin-left:0px; margin-right:0px; -qt-block-indent:0; text-indent:0px;\"><span style=\" font-size:12pt;\">192.168.2.46:30001</span></p></body></html>", "", -1))
	this.PushButton.SetText(_translate("Frame", "确  定", "", -1))
	this.Label.SetText(_translate("Frame", "Host", "", -1))
}

func (this *MY_UIBuildTcpWidget) Show() {
	this.Frame.Show()
}

func (this *MY_UIBuildTcpWidget) SetParent(parent *MY_UIMainWindow) {
	this.Parent = parent
}

func (this *MY_UIBuildTcpWidget) ButtonCheck() {
	//todo 更新主窗口左侧界面  关闭当前窗口 创建新窗口 显示日志
	var client = new(tcpconn.Client)
	client.Addr = this.TextEdit.ToPlainText()
	client.Network = "tcp"
	err := client.Connect()
	if err != nil {
		this.Parent.AddLog("与地址[%s]建立tcp连接失败,请检查网络等原因", client.Addr)
		return
	}
	err = tcpconn.SendOriginePkg(client)
	if err == nil {
		this.Parent.AddLog("与地址[%s]建立tcp连接成功", client.Addr)
		this.Frame.Close()

		toplevelIndex := this.Parent.BTW_TreeWidgetCont.TopLevelItems[string(TreeWidget_tcp)]
		topItem := this.Parent.TreeWidget.TopLevelItem(toplevelIndex)
		topItem.SetText(0, string(TreeWidget_tcp))

		childIndex := topItem.ChildCount()
		child := widgets.NewQTreeWidgetItem(0)
		child.SetText(0, client.Addr)
		child.SetText(1, strconv.Itoa(childIndex))
		child.SetIcon(0, gui.NewQIcon5(image.YesCheck))
		topItem.AddChild(child)

		//保存item-tcp连接-窗口等对应关系
		itemsMap := this.Parent.BTW_TreeWidgetCont.SubItems
		_, ok := itemsMap[toplevelIndex]
		if !ok {
			var topLevelSubItems = new(ToplevelSubItems)
			topLevelSubItems.SubItems = make(map[int]*ItemInfo)

			var itemInfo = new(ItemInfo)
			itemInfo.ItemIndex = childIndex
			itemInfo.Connect = client

			topLevelSubItems.SubItems[childIndex] = itemInfo
			itemsMap[toplevelIndex] = topLevelSubItems
		} else {
			var itemInfo = new(ItemInfo)
			itemInfo.ItemIndex = childIndex
			itemInfo.Connect = client

			itemsMap[toplevelIndex].SubItems[childIndex] = itemInfo
		}

		go this.CheckTcpDisconnect(client, topItem, child)

	} else {
		this.Parent.AddLog("与地址[%s]建立tcp连接失败,请检查网络等原因", client.Addr)
	}

}

func (this *MY_UIBuildTcpWidget) CheckTcpDisconnect(client *tcpconn.Client, topLevelItem *widgets.QTreeWidgetItem, child *widgets.QTreeWidgetItem) {
	for {
		time.Sleep(time.Duration(5) * time.Second)
		err := tcpconn.SendCltHeartPkg(client)
		if err != nil {
			if client.NetConn != nil {
				client.NetConn.Close()
			}
			//topLevelItem.RemoveChild(item)
			child.SetIcon(0, gui.NewQIcon5(image.Cross))
			this.Parent.AddLog("与地址[%s]断开了连接", client.Addr)
			break
		}
	}
}
