// WARNING! All changes made in this file will be lost!
package uitogo

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type UIFrame struct {
	TextEdit   *widgets.QTextEdit
	PushButton *widgets.QPushButton
	Label      *widgets.QLabel
}

func (this *UIFrame) SetupUI(Frame *widgets.QFrame) {
	Frame.SetObjectName("Frame")
	Frame.SetGeometry(core.NewQRect4(0, 0, 447, 319))
	var sizePolicy *widgets.QSizePolicy
	sizePolicy = widgets.NewQSizePolicy2(widgets.QSizePolicy__Fixed, widgets.QSizePolicy__Fixed, widgets.QSizePolicy__DefaultType)
	sizePolicy.SetHorizontalStretch(0)
	sizePolicy.SetVerticalStretch(0)
	sizePolicy.SetHeightForWidth(Frame.SizePolicy().HasHeightForWidth())
	Frame.SetSizePolicy(sizePolicy)
	var font *gui.QFont
	font = gui.NewQFont()
	font.SetPointSize(12)
	Frame.SetFont(font)
	this.TextEdit = widgets.NewQTextEdit(Frame)
	this.TextEdit.SetObjectName("TextEdit")
	this.TextEdit.SetGeometry(core.NewQRect4(110, 100, 271, 31))
	this.PushButton = widgets.NewQPushButton(Frame)
	this.PushButton.SetObjectName("PushButton")
	this.PushButton.SetGeometry(core.NewQRect4(150, 200, 150, 40))
	this.Label = widgets.NewQLabel(Frame, core.Qt__Widget)
	this.Label.SetObjectName("Label")
	this.Label.SetGeometry(core.NewQRect4(40, 100, 61, 31))
	font = gui.NewQFont()
	font.SetPointSize(15)
	font.SetWeight(50)
	this.Label.SetFont(font)
	this.Label.SetLayoutDirection(core.Qt__LeftToRight)

	this.RetranslateUi(Frame)

}

func (this *UIFrame) RetranslateUi(Frame *widgets.QFrame) {
	_translate := core.QCoreApplication_Translate
	Frame.SetWindowTitle(_translate("Frame", "建立tcp连接", "", -1))
	this.TextEdit.SetHtml(_translate("Frame", "<!DOCTYPE HTML PUBLIC \"-//W3C//DTD HTML 4.0//EN\" \"http://www.w3.org/TR/REC-html40/strict.dtd\">\n<html><head><meta name=\"qrichtext\" content=\"1\" /><style type=\"text/css\">\np, li { white-space: pre-wrap; }\n</style></head><body style=\" font-family:'SimSun'; font-size:12pt; font-weight:400; font-style:normal;\">\n<p style=\" margin-top:0px; margin-bottom:0px; margin-left:0px; margin-right:0px; -qt-block-indent:0; text-indent:0px;\">192.168.109.81:30001</p></body></html>", "", -1))
	this.PushButton.SetText(_translate("Frame", "确  定", "", -1))
	this.Label.SetText(_translate("Frame", "Host", "", -1))
}
