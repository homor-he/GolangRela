// WARNING! All changes made in this file will be lost!
package uitogo

import (
	"github.com/therecipe/qt/widgets"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/core"
)

type UIPerClient struct {
	VerticalLayout2 *widgets.QVBoxLayout
	VerticalLayout *widgets.QVBoxLayout
	TopTabWidget *widgets.QTabWidget
	Login *widgets.QWidget
	VerticalLayout3 *widgets.QVBoxLayout
	LoginSub *widgets.QTabWidget
	AccountLogin *widgets.QWidget
	AccountLabel *widgets.QLabel
	PwdLabel *widgets.QLabel
	AccountInput *widgets.QLineEdit
	PwdInput *widgets.QLineEdit
	Check *widgets.QPushButton
	Clear *widgets.QPushButton
	Logout *widgets.QPushButton
	St1 *widgets.QWidget
	T2 *widgets.QWidget
	T3 *widgets.QWidget
	T4 *widgets.QWidget
	T5 *widgets.QWidget
	T6 *widgets.QWidget
	T7 *widgets.QWidget
	T8 *widgets.QWidget
	T9 *widgets.QWidget
	T10 *widgets.QWidget
	Log *widgets.QPlainTextEdit
}

func (this *UIPerClient) SetupUI(PerClient *widgets.QFrame) {
	PerClient.SetObjectName("PerClient")
	PerClient.SetGeometry(core.NewQRect4(0, 0, 720, 520))
	PerClient.SetMinimumSize(core.NewQSize2(720, 520))
	this.VerticalLayout2 = widgets.NewQVBoxLayout2(PerClient)
	this.VerticalLayout2.SetObjectName("verticalLayout_2")
	this.VerticalLayout2.SetContentsMargins(0, 0, 0, 0)
	this.VerticalLayout2.SetSpacing(0)
	this.VerticalLayout = widgets.NewQVBoxLayout2(PerClient)
	this.VerticalLayout.SetObjectName("verticalLayout")
	this.VerticalLayout.SetContentsMargins(0, 0, 0, 0)
	this.VerticalLayout.SetSpacing(1)
	this.TopTabWidget = widgets.NewQTabWidget(PerClient)
	this.TopTabWidget.SetObjectName("TopTabWidget")
	this.TopTabWidget.SetTabsClosable(false)
	this.TopTabWidget.SetTabBarAutoHide(false)
	this.Login = widgets.NewQWidget(this.TopTabWidget, core.Qt__Widget)
	this.Login.SetObjectName("Login")
	this.VerticalLayout3 = widgets.NewQVBoxLayout2(this.Login)
	this.VerticalLayout3.SetObjectName("verticalLayout_3")
	this.VerticalLayout3.SetContentsMargins(0, 0, 0, 0)
	this.VerticalLayout3.SetSpacing(0)
	this.LoginSub = widgets.NewQTabWidget(this.Login)
	this.LoginSub.SetObjectName("LoginSub")
	this.LoginSub.SetMinimumSize(core.NewQSize2(711, 361))
	this.AccountLogin = widgets.NewQWidget(this.LoginSub, core.Qt__Widget)
	this.AccountLogin.SetObjectName("AccountLogin")
	this.AccountLabel = widgets.NewQLabel(this.AccountLogin, core.Qt__Widget)
	this.AccountLabel.SetObjectName("AccountLabel")
	this.AccountLabel.SetGeometry(core.NewQRect4(150, 80, 91, 41))
	var font *gui.QFont
	font = gui.NewQFont()
	font.SetPointSize(15)
	this.AccountLabel.SetFont(font)
	this.PwdLabel = widgets.NewQLabel(this.AccountLogin, core.Qt__Widget)
	this.PwdLabel.SetObjectName("PwdLabel")
	this.PwdLabel.SetGeometry(core.NewQRect4(150, 150, 91, 41))
	font = gui.NewQFont()
	font.SetPointSize(15)
	this.PwdLabel.SetFont(font)
	this.PwdLabel.SetLayoutDirection(core.Qt__RightToLeft)
	this.PwdLabel.SetAlignment(core.Qt__AlignRight | core.Qt__AlignTrailing | core.Qt__AlignVCenter)
	this.AccountInput = widgets.NewQLineEdit(this.AccountLogin)
	this.AccountInput.SetObjectName("AccountInput")
	this.AccountInput.SetGeometry(core.NewQRect4(250, 80, 250, 41))
	this.PwdInput = widgets.NewQLineEdit(this.AccountLogin)
	this.PwdInput.SetObjectName("PwdInput")
	this.PwdInput.SetGeometry(core.NewQRect4(250, 150, 250, 41))
	this.Check = widgets.NewQPushButton(this.AccountLogin)
	this.Check.SetObjectName("Check")
	this.Check.SetGeometry(core.NewQRect4(180, 240, 93, 28))
	this.Clear = widgets.NewQPushButton(this.AccountLogin)
	this.Clear.SetObjectName("Clear")
	this.Clear.SetGeometry(core.NewQRect4(310, 240, 93, 28))
	this.Logout = widgets.NewQPushButton(this.AccountLogin)
	this.Logout.SetObjectName("Logout")
	this.Logout.SetGeometry(core.NewQRect4(440, 240, 93, 28))
	this.LoginSub.AddTab(this.AccountLogin, "")
	this.St1 = widgets.NewQWidget(this.LoginSub, core.Qt__Widget)
	this.St1.SetObjectName("St1")
	this.LoginSub.AddTab(this.St1, "")
	this.VerticalLayout3.AddWidget(this.LoginSub, 0, 0)
	this.TopTabWidget.AddTab(this.Login, "")
	this.T2 = widgets.NewQWidget(this.TopTabWidget, core.Qt__Widget)
	this.T2.SetObjectName("T2")
	this.TopTabWidget.AddTab(this.T2, "")
	this.T3 = widgets.NewQWidget(this.TopTabWidget, core.Qt__Widget)
	this.T3.SetObjectName("T3")
	this.TopTabWidget.AddTab(this.T3, "")
	this.T4 = widgets.NewQWidget(this.TopTabWidget, core.Qt__Widget)
	this.T4.SetObjectName("T4")
	this.TopTabWidget.AddTab(this.T4, "")
	this.T5 = widgets.NewQWidget(this.TopTabWidget, core.Qt__Widget)
	this.T5.SetObjectName("T5")
	this.TopTabWidget.AddTab(this.T5, "")
	this.T6 = widgets.NewQWidget(this.TopTabWidget, core.Qt__Widget)
	this.T6.SetObjectName("T6")
	this.TopTabWidget.AddTab(this.T6, "")
	this.T7 = widgets.NewQWidget(this.TopTabWidget, core.Qt__Widget)
	this.T7.SetObjectName("T7")
	this.TopTabWidget.AddTab(this.T7, "")
	this.T8 = widgets.NewQWidget(this.TopTabWidget, core.Qt__Widget)
	this.T8.SetObjectName("T8")
	this.TopTabWidget.AddTab(this.T8, "")
	this.T9 = widgets.NewQWidget(this.TopTabWidget, core.Qt__Widget)
	this.T9.SetObjectName("T9")
	this.TopTabWidget.AddTab(this.T9, "")
	this.T10 = widgets.NewQWidget(this.TopTabWidget, core.Qt__Widget)
	this.T10.SetObjectName("T10")
	this.TopTabWidget.AddTab(this.T10, "")
	this.VerticalLayout.AddWidget(this.TopTabWidget, 0, 0)
	this.Log = widgets.NewQPlainTextEdit(PerClient)
	this.Log.SetObjectName("Log")
	this.VerticalLayout.AddWidget(this.Log, 0, 0)
	this.VerticalLayout.SetStretch(0, 3)
	this.VerticalLayout.SetStretch(1, 1)
	this.VerticalLayout2.AddLayout(this.VerticalLayout, 0)


    this.RetranslateUi(PerClient)
	this.TopTabWidget.SetCurrentIndex(0)
	this.LoginSub.SetCurrentIndex(0)
}

func (this *UIPerClient) RetranslateUi(PerClient *widgets.QFrame) {
    _translate := core.QCoreApplication_Translate
	PerClient.SetWindowTitle(_translate("PerClient", "Frame", "", -1))
	this.Login.SetToolTip(_translate("PerClient", "<html><head/><body><p>123</p></body></html>", "", -1))
	this.AccountLabel.SetText(_translate("PerClient", "用户名:", "", -1))
	this.PwdLabel.SetText(_translate("PerClient", "密码:", "", -1))
	this.Check.SetText(_translate("PerClient", "确 定", "", -1))
	this.Clear.SetText(_translate("PerClient", "清除文本", "", -1))
	this.Logout.SetText(_translate("PerClient", "登 出", "", -1))
	this.LoginSub.SetTabText(this.LoginSub.IndexOf(this.AccountLogin), _translate("PerClient", "账密登录", "", -1))
	this.LoginSub.SetTabText(this.LoginSub.IndexOf(this.St1), _translate("PerClient", "其他", "", -1))
	this.TopTabWidget.SetTabText(this.TopTabWidget.IndexOf(this.Login), _translate("PerClient", "登录", "", -1))
	this.TopTabWidget.SetTabText(this.TopTabWidget.IndexOf(this.T2), _translate("PerClient", "Page", "", -1))
	this.TopTabWidget.SetTabText(this.TopTabWidget.IndexOf(this.T3), _translate("PerClient", "Page", "", -1))
	this.TopTabWidget.SetTabText(this.TopTabWidget.IndexOf(this.T4), _translate("PerClient", "Page", "", -1))
	this.TopTabWidget.SetTabText(this.TopTabWidget.IndexOf(this.T5), _translate("PerClient", "Page", "", -1))
	this.TopTabWidget.SetTabText(this.TopTabWidget.IndexOf(this.T6), _translate("PerClient", "Page", "", -1))
	this.TopTabWidget.SetTabText(this.TopTabWidget.IndexOf(this.T7), _translate("PerClient", "Page", "", -1))
	this.TopTabWidget.SetTabText(this.TopTabWidget.IndexOf(this.T8), _translate("PerClient", "Page", "", -1))
	this.TopTabWidget.SetTabText(this.TopTabWidget.IndexOf(this.T9), _translate("PerClient", "Page", "", -1))
	this.TopTabWidget.SetTabText(this.TopTabWidget.IndexOf(this.T10), _translate("PerClient", "Page", "", -1))
}
