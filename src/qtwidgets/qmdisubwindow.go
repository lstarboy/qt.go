//  header block begin
// /usr/include/qt/QtWidgets/qmdisubwindow.h
// #include <qmdisubwindow.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 46
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"
import "qtgui"

func init() {
	if false {
		reflect.TypeOf(123)
	}
	if false {
		reflect.TypeOf(unsafe.Sizeof(0))
	}
	if false {
		fmt.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin
type QMdiSubWindow struct {
	*QWidget
}

func (this *QMdiSubWindow) GetCthis() unsafe.Pointer {
	return this.QWidget.GetCthis()
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:57
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QMdiSubWindow) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMdiSubWindow10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:69
// index:0
// void QMdiSubWindow(class QWidget *, Qt::WindowFlags)
func NewQMdiSubWindow(parent unsafe.Pointer, flags int) *QMdiSubWindow {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindowC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", ffiqt.FFI_TYPE_VOID, cthis, parent, &flags)
	gopp.ErrPrint(err, rv)
	gothis := NewQMdiSubWindowFromPointer(cthis)
	return gothis
}
func NewQMdiSubWindowFromPointer(cthis unsafe.Pointer) *QMdiSubWindow {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QMdiSubWindow{bcthis0}
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:70
// index:0
// virtual
// void ~QMdiSubWindow()
func DeleteQMdiSubWindow(*QMdiSubWindow) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindowD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:72
// index:0
// virtual
// QSize sizeHint()
func (this *QMdiSubWindow) SizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMdiSubWindow8sizeHintEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:73
// index:0
// virtual
// QSize minimumSizeHint()
func (this *QMdiSubWindow) MinimumSizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMdiSubWindow15minimumSizeHintEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:75
// index:0
// void setWidget(class QWidget *)
func (this *QMdiSubWindow) SetWidget(widget unsafe.Pointer) {
	// 0: (, widget QWidget *), (widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow9setWidgetEP7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:76
// index:0
// QWidget * widget()
func (this *QMdiSubWindow) Widget() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMdiSubWindow6widgetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:78
// index:0
// QWidget * maximizedButtonsWidget()
func (this *QMdiSubWindow) MaximizedButtonsWidget() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMdiSubWindow22maximizedButtonsWidgetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:79
// index:0
// QWidget * maximizedSystemMenuIconWidget()
func (this *QMdiSubWindow) MaximizedSystemMenuIconWidget() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMdiSubWindow29maximizedSystemMenuIconWidgetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:81
// index:0
// bool isShaded()
func (this *QMdiSubWindow) IsShaded() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMdiSubWindow8isShadedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:83
// index:0
// void setOption(enum QMdiSubWindow::SubWindowOption, _Bool)
func (this *QMdiSubWindow) SetOption(option int, on bool) {
	// 0: (, option QMdiSubWindow::SubWindowOption, on bool), (&option, &on)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow9setOptionENS_15SubWindowOptionEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &option, &on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:84
// index:0
// bool testOption(enum QMdiSubWindow::SubWindowOption)
func (this *QMdiSubWindow) TestOption(arg0 int) {
	// 0: (, QMdiSubWindow::SubWindowOption arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMdiSubWindow10testOptionENS_15SubWindowOptionE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:86
// index:0
// void setKeyboardSingleStep(int)
func (this *QMdiSubWindow) SetKeyboardSingleStep(step int) {
	// 0: (, step int), (&step)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow21setKeyboardSingleStepEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &step)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:87
// index:0
// int keyboardSingleStep()
func (this *QMdiSubWindow) KeyboardSingleStep() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMdiSubWindow18keyboardSingleStepEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:89
// index:0
// void setKeyboardPageStep(int)
func (this *QMdiSubWindow) SetKeyboardPageStep(step int) {
	// 0: (, step int), (&step)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow19setKeyboardPageStepEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &step)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:90
// index:0
// int keyboardPageStep()
func (this *QMdiSubWindow) KeyboardPageStep() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMdiSubWindow16keyboardPageStepEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:93
// index:0
// void setSystemMenu(class QMenu *)
func (this *QMdiSubWindow) SetSystemMenu(systemMenu unsafe.Pointer) {
	// 0: (, systemMenu QMenu *), (systemMenu)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow13setSystemMenuEP5QMenu", ffiqt.FFI_TYPE_VOID, this.GetCthis(), systemMenu)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:94
// index:0
// QMenu * systemMenu()
func (this *QMdiSubWindow) SystemMenu() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMdiSubWindow10systemMenuEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:97
// index:0
// QMdiArea * mdiArea()
func (this *QMdiSubWindow) MdiArea() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMdiSubWindow7mdiAreaEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:100
// index:0
// void windowStateChanged(Qt::WindowStates, Qt::WindowStates)
func (this *QMdiSubWindow) WindowStateChanged(oldState int, newState int) {
	// 0: (, QFlags<Qt::WindowState> oldState, QFlags<Qt::WindowState> newState), (&oldState, &newState)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow18windowStateChangedE6QFlagsIN2Qt11WindowStateEES3_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &oldState, &newState)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:101
// index:0
// void aboutToActivate()
func (this *QMdiSubWindow) AboutToActivate() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow15aboutToActivateEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:105
// index:0
// void showSystemMenu()
func (this *QMdiSubWindow) ShowSystemMenu() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow14showSystemMenuEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:107
// index:0
// void showShaded()
func (this *QMdiSubWindow) ShowShaded() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow10showShadedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:110
// index:0
// virtual
// bool eventFilter(class QObject *, class QEvent *)
func (this *QMdiSubWindow) EventFilter(object unsafe.Pointer, event unsafe.Pointer) {
	// 0: (, object QObject *, event QEvent *), (object, event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow11eventFilterEP7QObjectP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), object, event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:111
// index:0
// virtual
// bool event(class QEvent *)
func (this *QMdiSubWindow) Event(event unsafe.Pointer) {
	// 0: (, event QEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow5eventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:112
// index:0
// virtual
// void showEvent(class QShowEvent *)
func (this *QMdiSubWindow) ShowEvent(showEvent unsafe.Pointer) {
	// 0: (, showEvent QShowEvent *), (showEvent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow9showEventEP10QShowEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), showEvent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:113
// index:0
// virtual
// void hideEvent(class QHideEvent *)
func (this *QMdiSubWindow) HideEvent(hideEvent unsafe.Pointer) {
	// 0: (, hideEvent QHideEvent *), (hideEvent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow9hideEventEP10QHideEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), hideEvent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:114
// index:0
// virtual
// void changeEvent(class QEvent *)
func (this *QMdiSubWindow) ChangeEvent(changeEvent unsafe.Pointer) {
	// 0: (, changeEvent QEvent *), (changeEvent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow11changeEventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), changeEvent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:115
// index:0
// virtual
// void closeEvent(class QCloseEvent *)
func (this *QMdiSubWindow) CloseEvent(closeEvent unsafe.Pointer) {
	// 0: (, closeEvent QCloseEvent *), (closeEvent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow10closeEventEP11QCloseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), closeEvent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:116
// index:0
// virtual
// void leaveEvent(class QEvent *)
func (this *QMdiSubWindow) LeaveEvent(leaveEvent unsafe.Pointer) {
	// 0: (, leaveEvent QEvent *), (leaveEvent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow10leaveEventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), leaveEvent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:117
// index:0
// virtual
// void resizeEvent(class QResizeEvent *)
func (this *QMdiSubWindow) ResizeEvent(resizeEvent unsafe.Pointer) {
	// 0: (, resizeEvent QResizeEvent *), (resizeEvent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow11resizeEventEP12QResizeEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), resizeEvent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:118
// index:0
// virtual
// void timerEvent(class QTimerEvent *)
func (this *QMdiSubWindow) TimerEvent(timerEvent unsafe.Pointer) {
	// 0: (, timerEvent QTimerEvent *), (timerEvent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow10timerEventEP11QTimerEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), timerEvent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:119
// index:0
// virtual
// void moveEvent(class QMoveEvent *)
func (this *QMdiSubWindow) MoveEvent(moveEvent unsafe.Pointer) {
	// 0: (, moveEvent QMoveEvent *), (moveEvent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow9moveEventEP10QMoveEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), moveEvent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:120
// index:0
// virtual
// void paintEvent(class QPaintEvent *)
func (this *QMdiSubWindow) PaintEvent(paintEvent unsafe.Pointer) {
	// 0: (, paintEvent QPaintEvent *), (paintEvent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), paintEvent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:121
// index:0
// virtual
// void mousePressEvent(class QMouseEvent *)
func (this *QMdiSubWindow) MousePressEvent(mouseEvent unsafe.Pointer) {
	// 0: (, mouseEvent QMouseEvent *), (mouseEvent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow15mousePressEventEP11QMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), mouseEvent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:122
// index:0
// virtual
// void mouseDoubleClickEvent(class QMouseEvent *)
func (this *QMdiSubWindow) MouseDoubleClickEvent(mouseEvent unsafe.Pointer) {
	// 0: (, mouseEvent QMouseEvent *), (mouseEvent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow21mouseDoubleClickEventEP11QMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), mouseEvent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:123
// index:0
// virtual
// void mouseReleaseEvent(class QMouseEvent *)
func (this *QMdiSubWindow) MouseReleaseEvent(mouseEvent unsafe.Pointer) {
	// 0: (, mouseEvent QMouseEvent *), (mouseEvent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow17mouseReleaseEventEP11QMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), mouseEvent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:124
// index:0
// virtual
// void mouseMoveEvent(class QMouseEvent *)
func (this *QMdiSubWindow) MouseMoveEvent(mouseEvent unsafe.Pointer) {
	// 0: (, mouseEvent QMouseEvent *), (mouseEvent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow14mouseMoveEventEP11QMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), mouseEvent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:125
// index:0
// virtual
// void keyPressEvent(class QKeyEvent *)
func (this *QMdiSubWindow) KeyPressEvent(keyEvent unsafe.Pointer) {
	// 0: (, keyEvent QKeyEvent *), (keyEvent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow13keyPressEventEP9QKeyEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), keyEvent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:127
// index:0
// virtual
// void contextMenuEvent(class QContextMenuEvent *)
func (this *QMdiSubWindow) ContextMenuEvent(contextMenuEvent unsafe.Pointer) {
	// 0: (, contextMenuEvent QContextMenuEvent *), (contextMenuEvent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow16contextMenuEventEP17QContextMenuEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), contextMenuEvent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:129
// index:0
// virtual
// void focusInEvent(class QFocusEvent *)
func (this *QMdiSubWindow) FocusInEvent(focusInEvent unsafe.Pointer) {
	// 0: (, focusInEvent QFocusEvent *), (focusInEvent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow12focusInEventEP11QFocusEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), focusInEvent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:130
// index:0
// virtual
// void focusOutEvent(class QFocusEvent *)
func (this *QMdiSubWindow) FocusOutEvent(focusOutEvent unsafe.Pointer) {
	// 0: (, focusOutEvent QFocusEvent *), (focusOutEvent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow13focusOutEventEP11QFocusEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), focusOutEvent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:131
// index:0
// virtual
// void childEvent(class QChildEvent *)
func (this *QMdiSubWindow) ChildEvent(childEvent unsafe.Pointer) {
	// 0: (, childEvent QChildEvent *), (childEvent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow10childEventEP11QChildEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), childEvent)
	gopp.ErrPrint(err, rv)
}

//  body block end
