// WARNING! All changes made in this file will be lost!
package uitogo

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type UIMainwindowMainWindow struct {
	Centralwidget *widgets.QWidget
	GridLayout *widgets.QGridLayout
	VLayoutCW *widgets.QVBoxLayout
	TopLine *widgets.QFrame
	HLayoutInCW *widgets.QHBoxLayout
	TreeWidget *widgets.QTreeWidget
	VLayoutInHLayout *widgets.QVBoxLayout
	TopTabWidget *widgets.QTabWidget
	Tab *widgets.QWidget
	NextTabWidget *widgets.QTabWidget
	Tab3 *widgets.QWidget
	Log *widgets.QPlainTextEdit
	BottomLine *widgets.QFrame
}

func (this *UIMainwindowMainWindow) SetupUI(MainWindow *widgets.QMainWindow) {
	MainWindow.SetObjectName("MainWindow")
	MainWindow.SetGeometry(core.NewQRect4(0, 0, 1057, 736))
	MainWindow.SetMouseTracking(true)
	MainWindow.SetAcceptDrops(true)
	MainWindow.SetDocumentMode(false)
	MainWindow.SetDockNestingEnabled(false)
	this.Centralwidget = widgets.NewQWidget(MainWindow, core.Qt__Widget)
	this.Centralwidget.SetObjectName("Centralwidget")
	this.Centralwidget.SetEnabled(true)
	this.Centralwidget.SetMouseTracking(false)
	this.Centralwidget.SetAcceptDrops(false)
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
	var treeItem1 *widgets.QTreeWidgetItem
	treeItem1 = widgets.NewQTreeWidgetItem3(this.TreeWidget, 0)
	this.TreeWidget.SetHeaderItem(treeItem1)
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
	this.HLayoutInCW.SetStretch(1, 4)
	this.VLayoutCW.AddLayout(this.HLayoutInCW, 0)
	this.BottomLine = widgets.NewQFrame(this.Centralwidget, core.Qt__Widget)
	this.BottomLine.SetObjectName("BottomLine")
	this.BottomLine.SetFrameShadow(widgets.QFrame__Sunken)
	this.BottomLine.SetFrameShape(widgets.QFrame__HLine)
	this.VLayoutCW.AddWidget(this.BottomLine, 0, 0)
	this.GridLayout.AddLayout2(this.VLayoutCW, 0, 0, 1, 1, 0)
	MainWindow.SetCentralWidget(this.Centralwidget)


    this.RetranslateUi(MainWindow)
	this.TopTabWidget.SetCurrentIndex(0)
	this.NextTabWidget.SetCurrentIndex(0)
}

func (this *UIMainwindowMainWindow) RetranslateUi(MainWindow *widgets.QMainWindow) {
    _translate := core.QCoreApplication_Translate
	MainWindow.SetWindowTitle(_translate("MainWindow", "测试工具1.0", "", -1))
	var sortingEnabled bool
	sortingEnabled = this.TreeWidget.IsSortingEnabled()
	this.TreeWidget.HeaderItem().SetText(0, _translate("MainWindow", "连接菜单", "", -1))
	this.TreeWidget.HeaderItem().SetText(1, _translate("MainWindow", "索引", "", -1))
	this.TreeWidget.SetSortingEnabled(sortingEnabled)
	this.TopTabWidget.SetTabText(this.TopTabWidget.IndexOf(this.Tab), _translate("MainWindow", "Tab 1", "", -1))
	this.NextTabWidget.SetTabText(this.NextTabWidget.IndexOf(this.Tab3), _translate("MainWindow", "Tab 1", "", -1))
	this.Log.SetPlainText(_translate("MainWindow", "中文\n中文", "", -1))
}
