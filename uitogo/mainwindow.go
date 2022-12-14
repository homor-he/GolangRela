package uitogo

import (
	"autotest/image"
	"autotest/tcpconn"
	"fmt"
	"strconv"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type MY_UIMainWindow struct {
	MainWindow    *widgets.QMainWindow    //主窗口
	DesktopWidget *widgets.QDesktopWidget //获取当前屏幕大小
	DesktopRect   *core.QRect

	Menubar        *widgets.QMenuBar //菜单栏
	FileMenu       *widgets.QMenu    //文件菜单
	BuildTcpAction *widgets.QAction  //文件下的退出菜单
	ExitAction     *widgets.QAction  //文件下的退出菜单

	Statusbar *widgets.QStatusBar //状态栏
	//WidgetList []*widgets.QWidget  //子窗口

	//主界面整体布局
	Centralwidget    *widgets.QWidget //主界面
	GridLayout       *widgets.QGridLayout
	VLayoutCW        *widgets.QVBoxLayout
	TopLine          *widgets.QFrame
	HLayoutInCW      *widgets.QHBoxLayout
	TreeWidget       *widgets.QTreeWidget
	VLayoutInHLayout *widgets.QVBoxLayout
	TopTabWidget     *widgets.QTabWidget
	Tab              *widgets.QWidget
	NextTabWidget    *widgets.QTabWidget
	Tab3             *widgets.QWidget
	Log              *widgets.QPlainTextEdit
	BottomLine       *widgets.QFrame

	BuildTcpWidget     MY_UIBuildTcpWidget
	BTW_TreeWidgetCont TreeWidget_Cont //树窗口目录

	UserWidget MY_UIPerClient
}

type TreeWidget_Cont struct {
	TopLevelItems map[string]int            //string为toplevel的名字，int为topItem索引
	SubItems      map[int]*ToplevelSubItems //int为topItem索引,ToplevelSubItems为对应的toplevel下的所有子item
}

type ItemInfo struct {
	ItemIndex  int
	Connect    *tcpconn.Client
	UserWidget *MY_UIPerClient
}

type ToplevelSubItems struct {
	SubItems map[int]*ItemInfo //int为subItem的索引
}

type FixName string

const (
	TreeWidget_tcp FixName = FixName("tcp连接")
)

func (this *MY_UIMainWindow) SetupUI(window *widgets.QMainWindow) {
	this.DesktopWidget = widgets.NewQDesktopWidget()
	this.DesktopRect = this.DesktopWidget.AvailableGeometry(0)

	this.MainWindow = widgets.NewQMainWindow(nil, 0)
	this.MainWindow.SetWindowTitle("测试工具1.0")
	this.MainWindow.SetGeometry(core.NewQRect4(0, 0, 1280, 760))
	this.MainWindow.StatusBar().ShowMessage("状态栏", -1)

	//菜单栏相关
	this.MenuBarRela()

	//主界面整体布局
	this.Centralwidget = widgets.NewQWidget(this.MainWindow, core.Qt__Widget)
	this.Centralwidget.SetObjectName("Centralwidget")
	this.Centralwidget.SetGeometry(core.NewQRect4(0, 26, 1280, 734))
	this.Centralwidget.SetEnabled(true)
	this.Centralwidget.SetMouseTracking(true)
	this.GridLayout = widgets.NewQGridLayout(this.Centralwidget)
	this.GridLayout.SetObjectName("gridLayout")
	this.GridLayout.SetContentsMargins(0, 0, 0, 0)
	this.GridLayout.SetSpacing(0)
	this.VLayoutCW = widgets.NewQVBoxLayout2(this.Centralwidget)
	this.VLayoutCW.SetObjectName("VLayout_CW")
	this.VLayoutCW.SetContentsMargins(0, 0, 0, 0)
	this.VLayoutCW.SetSpacing(0)
	this.TopLine = widgets.NewQFrame(this.Centralwidget, core.Qt__Widget)
	this.TopLine.SetObjectName("TopLine")
	this.TopLine.SetFrameShadow(widgets.QFrame__Sunken)
	this.TopLine.SetFrameShape(widgets.QFrame__HLine)
	this.VLayoutCW.AddWidget(this.TopLine, 0, 0)
	this.HLayoutInCW = widgets.NewQHBoxLayout2(this.Centralwidget)
	this.HLayoutInCW.SetObjectName("HLayoutInCW")
	this.HLayoutInCW.SetContentsMargins(0, 0, 0, 0)
	this.HLayoutInCW.SetSpacing(5)
	this.HLayoutInCW.SetSizeConstraint(widgets.QLayout__SetNoConstraint)

	this.TreeWidget = widgets.NewQTreeWidget(this.Centralwidget)
	this.TreeWidget.SetObjectName("TreeWidget")
	//var treeItem1 *widgets.QTreeWidgetItem
	//treeItem1 = widgets.NewQTreeWidgetItem3(this.TreeWidget, 0)
	//this.TreeWidget.SetHeaderItem(treeItem1)
	treeItem := widgets.NewQTreeWidgetItem3(this.TreeWidget, 0)
	this.TreeWidget.AddTopLevelItem(treeItem)
	this.TreeWidget.TopLevelItem(this.TreeWidget.IndexOfTopLevelItem(treeItem)).SetText(0, string(TreeWidget_tcp))
	this.BTW_TreeWidgetCont.TopLevelItems = make(map[string]int)
	this.BTW_TreeWidgetCont.SubItems = make(map[int]*ToplevelSubItems)
	this.BTW_TreeWidgetCont.TopLevelItems[string(TreeWidget_tcp)] = 0
	//添加TreeWidget右键点击事件
	this.TreeWidget.ConnectContextMenuEvent(func(event *gui.QContextMenuEvent) {
		this.MouseRightClickOnTreeWidget()
	})
	this.HLayoutInCW.AddWidget(this.TreeWidget, 0, 0)

	this.VLayoutInHLayout = widgets.NewQVBoxLayout2(this.Centralwidget)
	this.VLayoutInHLayout.SetObjectName("VLayoutInHLayout")
	this.VLayoutInHLayout.SetContentsMargins(0, 0, 0, 0)
	this.VLayoutInHLayout.SetSpacing(2)
	this.TopTabWidget = widgets.NewQTabWidget(this.Centralwidget)
	this.TopTabWidget.SetObjectName("TopTabWidget")
	this.TopTabWidget.SetMinimumSize(core.NewQSize2(836, 276))
	this.Tab = widgets.NewQWidget(this.TopTabWidget, core.Qt__Widget)
	this.Tab.SetObjectName("Tab")
	this.TopTabWidget.AddTab(this.Tab, "")
	this.VLayoutInHLayout.AddWidget(this.TopTabWidget, 0, 0)
	this.NextTabWidget = widgets.NewQTabWidget(this.Centralwidget)
	this.NextTabWidget.SetObjectName("NextTabWidget")
	this.NextTabWidget.SetMinimumSize(core.NewQSize2(836, 277))
	this.Tab3 = widgets.NewQWidget(this.NextTabWidget, core.Qt__Widget)
	this.Tab3.SetObjectName("Tab3")
	this.NextTabWidget.AddTab(this.Tab3, "")
	this.VLayoutInHLayout.AddWidget(this.NextTabWidget, 0, 0)
	this.Log = widgets.NewQPlainTextEdit(this.Centralwidget)
	this.Log.SetObjectName("Log")
	this.VLayoutInHLayout.AddWidget(this.Log, 0, 0)
	this.VLayoutInHLayout.SetStretch(0, 4)
	this.VLayoutInHLayout.SetStretch(1, 4)
	this.HLayoutInCW.AddLayout(this.VLayoutInHLayout, 0)
	this.HLayoutInCW.SetStretch(0, 1)
	this.HLayoutInCW.SetStretch(1, 4)
	this.VLayoutCW.AddLayout(this.HLayoutInCW, 0)
	this.BottomLine = widgets.NewQFrame(this.Centralwidget, core.Qt__Widget)
	this.BottomLine.SetObjectName("BottomLine")
	this.BottomLine.SetFrameShadow(widgets.QFrame__Sunken)
	this.BottomLine.SetFrameShape(widgets.QFrame__HLine)
	this.VLayoutCW.AddWidget(this.BottomLine, 0, 0)
	this.GridLayout.AddLayout2(this.VLayoutCW, 0, 0, 1, 1, 0)
	this.MainWindow.SetCentralWidget(this.Centralwidget)

	this.RetranslateUi(this.MainWindow)
	this.TopTabWidget.SetCurrentIndex(0)
	this.NextTabWidget.SetCurrentIndex(0)
}

func (this *MY_UIMainWindow) MenuBarRela() {
	this.Menubar = this.MainWindow.MenuBar()
	//添加菜单
	this.FileMenu = this.Menubar.AddMenu2("&文件")
	//添加按钮
	this.BuildTcpAction = this.FileMenu.AddAction("&建立tcp连接")
	this.BuildTcpAction.ConnectTriggered(func(checked bool) {
		this.OpenBuildTcpWidget()
	})

	this.ExitAction = this.FileMenu.AddAction("&退出")
	this.ExitAction.SetStatusTip("退出程序")
	this.ExitAction.ConnectTriggered(func(checked bool) {
		this.MainWindow.Close()
	})
}

func (this *MY_UIMainWindow) Show() {
	winWidth := this.MainWindow.Width()
	winHeight := this.MainWindow.Height()
	this.MainWindow.SetGeometry(core.NewQRect4(this.DesktopRect.Width()/2-winWidth/2, this.DesktopRect.Height()/2-winHeight/2, winWidth, winHeight))
	this.MainWindow.Show()
}

func (this *MY_UIMainWindow) OpenBuildTcpWidget() {
	this.BuildTcpWidget.Frame = widgets.NewQFrame(nil, core.Qt__Widget)
	this.BuildTcpWidget.Frame.SetWindowFlags(core.Qt__Dialog)
	//主窗口关闭时，子窗口会自动退出
	this.BuildTcpWidget.Frame.SetAttribute(core.Qt__WA_QuitOnClose, false)
	this.BuildTcpWidget.SetupUI(this.BuildTcpWidget.Frame)
	this.BuildTcpWidget.SetParent(this)
	//居中屏幕显示
	frameWidth := this.BuildTcpWidget.Frame.Width()
	frameHeight := this.BuildTcpWidget.Frame.Height()
	this.BuildTcpWidget.Frame.SetGeometry(core.NewQRect4(this.DesktopRect.Width()/2-frameWidth/2, this.DesktopRect.Height()/2-frameHeight/2, frameWidth, frameHeight))
	this.BuildTcpWidget.Show()
}

func (this *MY_UIMainWindow) RetranslateUi(MainWindow *widgets.QMainWindow) {
	_translate := core.QCoreApplication_Translate
	MainWindow.SetWindowTitle(_translate("MainWindow", "测试工具1.0", "", -1))
	var sortingEnabled bool
	sortingEnabled = this.TreeWidget.IsSortingEnabled()
	this.TreeWidget.HeaderItem().SetText(0, _translate("MainWindow", "连接菜单", "", -1))
	this.TreeWidget.HeaderItem().SetText(1, _translate("MainWindow", "索引", "", -1))
	this.TreeWidget.SetSortingEnabled(sortingEnabled)
	this.TopTabWidget.SetTabText(this.TopTabWidget.IndexOf(this.Tab), _translate("MainWindow", "Tab 1", "", -1))
	this.NextTabWidget.SetTabText(this.NextTabWidget.IndexOf(this.Tab3), _translate("MainWindow", "Tab 1", "", -1))
}

func (this *MY_UIMainWindow) AddLog(format string, v ...any) {
	log := fmt.Sprintf(format, v...)
	this.Log.AppendPlainText(log)
}

func (this *MY_UIMainWindow) MouseRightClickOnTreeWidget() {
	popMenu := widgets.NewQMenu(this.TreeWidget)

	cursor := this.MainWindow.Cursor()
	//chooseItem := this.TreeWidget.ItemAt(cursor.Pos())
	currItem := this.TreeWidget.CurrentItem()
	index := this.TreeWidget.IndexOfTopLevelItem(currItem)

	actions := make([]*widgets.QAction, 0)
	if (currItem != nil) && (index < 0) {
		//获取父item
		parentIndex := this.TreeWidget.IndexOfTopLevelItem(currItem.Parent())
		childIndex := currItem.Parent().IndexOfChild(currItem)
		itemInfo := this.BTW_TreeWidgetCont.SubItems[parentIndex].SubItems[childIndex]
		actionOpen := popMenu.AddAction("&打开子界面")
		actionOpen.ConnectTriggered(func(checked bool) {
			name := "tcp:" + itemInfo.Connect.Addr + "__" + strconv.Itoa(childIndex)
			this.PopMenuOpenUserClient(itemInfo, name)
		})

		actionReconnect := popMenu.AddAction("&重连")
		actionReconnect.ConnectTriggered(func(checked bool) {
			go this.PopMenuReconnect(itemInfo, currItem)
		})

		actionDel := popMenu.AddAction("&删除")
		actionDel.ConnectTriggered(func(checked bool) {
			currItem.SetHidden(true)
			if itemInfo.UserWidget != nil {
				itemInfo.UserWidget.Frame.Close()

			}
			if itemInfo.Connect.NetConn != nil {
				itemInfo.Connect.NetConn.Close()
			}
			delete(this.BTW_TreeWidgetCont.SubItems[parentIndex].SubItems, childIndex)
		})

		actions = append(actions, actionOpen)
		actions = append(actions, actionReconnect)
		actions = append(actions, actionDel)

		var height int
		for _, action := range actions {
			height += popMenu.ActionGeometry(action).Height() + 2
		}

		popMenu.SetGeometry2(cursor.Pos().Rx(), cursor.Pos().Ry(), popMenu.ActionGeometry(actionOpen).Width(), height)
		popMenu.Show()
	}
}

func (this *MY_UIMainWindow) PopMenuOpenUserClient(info *ItemInfo, name string) {
	info.UserWidget = new(MY_UIPerClient)
	info.UserWidget.Frame = widgets.NewQFrame(nil, core.Qt__Widget)
	info.UserWidget.Frame.SetAttribute(core.Qt__WA_QuitOnClose, false)
	info.UserWidget.SetupUI(info.UserWidget.Frame)
	info.UserWidget.Frame.SetWindowTitle(name)
	info.UserWidget.Show()

	info.UserWidget.SetLink(info)
}

func (this *MY_UIMainWindow) PopMenuReconnect(itemInfo *ItemInfo, currItem *widgets.QTreeWidgetItem) {
	oneByte := make([]byte, 1)
	var err error
	if itemInfo.Connect.NetConn != nil {
		_, err = itemInfo.Connect.NetConn.Read(oneByte)
		if err != nil {
			err = itemInfo.Connect.Connect()
			if err == nil {
				currItem.SetIcon(0, gui.NewQIcon5(image.YesCheck))
				this.AddLog("与地址[%s]重新建立tcp连接成功", itemInfo.Connect.Addr)
			}
		}
	} else {
		err = itemInfo.Connect.Connect()
		if err == nil {
			currItem.SetIcon(0, gui.NewQIcon5(image.YesCheck))
			this.AddLog("与地址[%s]重新建立tcp连接成功", itemInfo.Connect.Addr)
		}
	}
}
