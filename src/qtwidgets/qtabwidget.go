package qtwidgets
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
// src-file: /QtWidgets/qtabwidget.h
// dst-file: /src/widgets/qtabwidget.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
import "fmt"
import "reflect"
import "unsafe"
import "runtime"
import "qtrt"
import "qtcore"
import "qtgui"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  QString QTabWidget::tabToolTip(int index);
extern void* C_ZNK10QTabWidget10tabToolTipEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QTabWidget::tabBarAutoHide();
extern bool C_ZNK10QTabWidget14tabBarAutoHideEv(void* qthis); // 4
  // proto:  bool QTabWidget::isTabEnabled(int index);
extern bool C_ZNK10QTabWidget12isTabEnabledEi(void* qthis, int32_t arg0); // 4
  // proto:  QString QTabWidget::tabWhatsThis(int index);
extern void* C_ZNK10QTabWidget12tabWhatsThisEi(void* qthis, int32_t arg0); // 4
  // proto:  QIcon QTabWidget::tabIcon(int index);
extern void* C_ZNK10QTabWidget7tabIconEi(void* qthis, int32_t arg0); // 4
  // proto:  QSize QTabWidget::minimumSizeHint();
extern void* C_ZNK10QTabWidget15minimumSizeHintEv(void* qthis); // 4
  // proto:  void QTabWidget::setCurrentIndex(int index);
extern void C_ZN10QTabWidget15setCurrentIndexEi(void* qthis, int32_t arg0); // 4
  // proto:  QWidget * QTabWidget::currentWidget();
extern void* C_ZNK10QTabWidget13currentWidgetEv(void* qthis); // 4
  // proto:  void QTabWidget::setMovable(bool movable);
extern void C_ZN10QTabWidget10setMovableEb(void* qthis, bool arg0); // 4
  // proto:  bool QTabWidget::tabsClosable();
extern bool C_ZNK10QTabWidget12tabsClosableEv(void* qthis); // 4
  // proto:  void QTabWidget::setTabIcon(int index, const QIcon & icon);
extern void C_ZN10QTabWidget10setTabIconEiRK5QIcon(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QTabWidget::setCurrentWidget(QWidget * widget);
extern void C_ZN10QTabWidget16setCurrentWidgetEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  void QTabWidget::~QTabWidget();
extern void C_ZN10QTabWidgetD2Ev(void* qthis); // 4
  // proto:  int QTabWidget::insertTab(int index, QWidget * widget, const QString & );
extern int32_t C_ZN10QTabWidget9insertTabEiP7QWidgetRK7QString(void* qthis, int32_t arg0, void* arg1, void* arg2); // 4
  // proto:  int QTabWidget::insertTab(int index, QWidget * widget, const QIcon & icon, const QString & label);
extern int32_t C_ZN10QTabWidget9insertTabEiP7QWidgetRK5QIconRK7QString(void* qthis, int32_t arg0, void* arg1, void* arg2, void* arg3); // 4
  // proto:  QTabWidget::TabShape QTabWidget::tabShape();
extern void C_ZNK10QTabWidget8tabShapeEv(void* qthis); // 4
  // proto:  int QTabWidget::indexOf(QWidget * widget);
extern int32_t C_ZNK10QTabWidget7indexOfEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  QTabBar * QTabWidget::tabBar();
extern void* C_ZNK10QTabWidget6tabBarEv(void* qthis); // 4
  // proto:  void QTabWidget::setTabToolTip(int index, const QString & tip);
extern void C_ZN10QTabWidget13setTabToolTipEiRK7QString(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QTabWidget::QTabWidget(QWidget * parent);
extern void* C_ZN10QTabWidgetC2EP7QWidget(void* arg0); // 3
  // proto:  void QTabWidget::setTabBarAutoHide(bool enabled);
extern void C_ZN10QTabWidget17setTabBarAutoHideEb(void* qthis, bool arg0); // 4
  // proto:  bool QTabWidget::documentMode();
extern bool C_ZNK10QTabWidget12documentModeEv(void* qthis); // 4
  // proto:  QWidget * QTabWidget::widget(int index);
extern void* C_ZNK10QTabWidget6widgetEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTabWidget::removeTab(int index);
extern void C_ZN10QTabWidget9removeTabEi(void* qthis, int32_t arg0); // 4
  // proto:  QSize QTabWidget::iconSize();
extern void* C_ZNK10QTabWidget8iconSizeEv(void* qthis); // 4
  // proto:  void QTabWidget::setTabText(int index, const QString & );
extern void C_ZN10QTabWidget10setTabTextEiRK7QString(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  QString QTabWidget::tabText(int index);
extern void* C_ZNK10QTabWidget7tabTextEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTabWidget::setTabsClosable(bool closeable);
extern void C_ZN10QTabWidget15setTabsClosableEb(void* qthis, bool arg0); // 4
  // proto:  QSize QTabWidget::sizeHint();
extern void* C_ZNK10QTabWidget8sizeHintEv(void* qthis); // 4
  // proto:  QTabWidget::TabPosition QTabWidget::tabPosition();
extern void C_ZNK10QTabWidget11tabPositionEv(void* qthis); // 4
  // proto:  bool QTabWidget::usesScrollButtons();
extern bool C_ZNK10QTabWidget17usesScrollButtonsEv(void* qthis); // 4
  // proto:  void QTabWidget::setDocumentMode(bool set);
extern void C_ZN10QTabWidget15setDocumentModeEb(void* qthis, bool arg0); // 4
  // proto:  int QTabWidget::count();
extern int32_t C_ZNK10QTabWidget5countEv(void* qthis); // 4
  // proto:  void QTabWidget::setUsesScrollButtons(bool useButtons);
extern void C_ZN10QTabWidget20setUsesScrollButtonsEb(void* qthis, bool arg0); // 4
  // proto:  const QMetaObject * QTabWidget::metaObject();
extern void C_ZNK10QTabWidget10metaObjectEv(void* qthis); // 4
  // proto:  Qt::TextElideMode QTabWidget::elideMode();
extern void C_ZNK10QTabWidget9elideModeEv(void* qthis); // 4
  // proto:  bool QTabWidget::hasHeightForWidth();
extern bool C_ZNK10QTabWidget17hasHeightForWidthEv(void* qthis); // 4
  // proto:  void QTabWidget::setTabWhatsThis(int index, const QString & text);
extern void C_ZN10QTabWidget15setTabWhatsThisEiRK7QString(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QTabWidget::clear();
extern void C_ZN10QTabWidget5clearEv(void* qthis); // 4
  // proto:  void QTabWidget::setIconSize(const QSize & size);
extern void C_ZN10QTabWidget11setIconSizeERK5QSize(void* qthis, void* arg0); // 4
  // proto:  bool QTabWidget::isMovable();
extern bool C_ZNK10QTabWidget9isMovableEv(void* qthis); // 4
  // proto:  int QTabWidget::heightForWidth(int width);
extern int32_t C_ZNK10QTabWidget14heightForWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTabWidget::setTabEnabled(int index, bool );
extern void C_ZN10QTabWidget13setTabEnabledEib(void* qthis, int32_t arg0, bool arg1); // 4
  // proto:  int QTabWidget::currentIndex();
extern int32_t C_ZNK10QTabWidget12currentIndexEv(void* qthis); // 4
  // proto:  int QTabWidget::addTab(QWidget * widget, const QIcon & icon, const QString & label);
extern int32_t C_ZN10QTabWidget6addTabEP7QWidgetRK5QIconRK7QString(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  int QTabWidget::addTab(QWidget * widget, const QString & );
extern int32_t C_ZN10QTabWidget6addTabEP7QWidgetRK7QString(void* qthis, void* arg0, void* arg1); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {qtcore.KeepMe()}
  if false {qtgui.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
  if false {reflect.TypeOf(runtime.Version)}
}

// class sizeof(QTabWidget)=1
type QTabWidget struct {
  /*qbase*/ QWidget;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _tabCloseRequested QTabWidget_tabCloseRequested_signal;
//  _tabBarDoubleClicked QTabWidget_tabBarDoubleClicked_signal;
//  _tabBarClicked QTabWidget_tabBarClicked_signal;
//  _currentChanged QTabWidget_currentChanged_signal;
}

// tabToolTip(int)
func (this *QTabWidget) TabToolTip(args ...interface{}) (ret interface{}) {
  // tabToolTip(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTabWidget10tabToolTipEi
    // invoke: QString tabToolTip(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QTabWidget10tabToolTipEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTabWidget", "tabToolTip", args)
  }

  return
}

// tabBarAutoHide()
func (this *QTabWidget) TabBarAutoHide(args ...interface{}) (ret interface{}) {
  // tabBarAutoHide()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTabWidget14tabBarAutoHideEv
    // invoke: bool tabBarAutoHide()
    var ret0 = C.C_ZNK10QTabWidget14tabBarAutoHideEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTabWidget", "tabBarAutoHide", args)
  }

  return
}

// isTabEnabled(int)
func (this *QTabWidget) IsTabEnabled(args ...interface{}) (ret interface{}) {
  // isTabEnabled(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTabWidget12isTabEnabledEi
    // invoke: bool isTabEnabled(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QTabWidget12isTabEnabledEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTabWidget", "isTabEnabled", args)
  }

  return
}

// tabWhatsThis(int)
func (this *QTabWidget) TabWhatsThis(args ...interface{}) (ret interface{}) {
  // tabWhatsThis(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTabWidget12tabWhatsThisEi
    // invoke: QString tabWhatsThis(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QTabWidget12tabWhatsThisEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTabWidget", "tabWhatsThis", args)
  }

  return
}

// tabIcon(int)
func (this *QTabWidget) TabIcon(args ...interface{}) (ret interface{}) {
  // tabIcon(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTabWidget7tabIconEi
    // invoke: QIcon tabIcon(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QTabWidget7tabIconEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QIcon{}) // "QIcon"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTabWidget", "tabIcon", args)
  }

  return
}

// minimumSizeHint()
func (this *QTabWidget) MinimumSizeHint(args ...interface{}) (ret interface{}) {
  // minimumSizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTabWidget15minimumSizeHintEv
    // invoke: QSize minimumSizeHint()
    var ret0 = C.C_ZNK10QTabWidget15minimumSizeHintEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTabWidget", "minimumSizeHint", args)
  }

  return
}

// setCurrentIndex(int)
func (this *QTabWidget) SetCurrentIndex(args ...interface{}) () {
  // setCurrentIndex(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTabWidget15setCurrentIndexEi
    // invoke: void setCurrentIndex(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN10QTabWidget15setCurrentIndexEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTabWidget", "setCurrentIndex", args)
  }

  return
}

// currentWidget()
func (this *QTabWidget) CurrentWidget(args ...interface{}) (ret interface{}) {
  // currentWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTabWidget13currentWidgetEv
    // invoke: QWidget * currentWidget()
    var ret0 = C.C_ZNK10QTabWidget13currentWidgetEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTabWidget", "currentWidget", args)
  }

  return
}

// setMovable(_Bool)
func (this *QTabWidget) SetMovable(args ...interface{}) () {
  // setMovable(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTabWidget10setMovableEb
    // invoke: void setMovable(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN10QTabWidget10setMovableEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTabWidget", "setMovable", args)
  }

  return
}

// tabsClosable()
func (this *QTabWidget) TabsClosable(args ...interface{}) (ret interface{}) {
  // tabsClosable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTabWidget12tabsClosableEv
    // invoke: bool tabsClosable()
    var ret0 = C.C_ZNK10QTabWidget12tabsClosableEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTabWidget", "tabsClosable", args)
  }

  return
}

// setTabIcon(int, const class QIcon &)
func (this *QTabWidget) SetTabIcon(args ...interface{}) () {
  // setTabIcon(int, const class QIcon &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(qtgui.QIcon{}) // "const QIcon &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTabWidget10setTabIconEiRK5QIcon
    // invoke: void setTabIcon(int, const class QIcon &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtgui.QIcon).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN10QTabWidget10setTabIconEiRK5QIcon(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTabWidget", "setTabIcon", args)
  }

  return
}

// setCurrentWidget(class QWidget *)
func (this *QTabWidget) SetCurrentWidget(args ...interface{}) () {
  // setCurrentWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTabWidget16setCurrentWidgetEP7QWidget
    // invoke: void setCurrentWidget(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN10QTabWidget16setCurrentWidgetEP7QWidget(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTabWidget", "setCurrentWidget", args)
  }

  return
}

// ~QTabWidget()
func (this *QTabWidget) Free(args ...interface{}) () {
  // ~QTabWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTabWidgetD0Ev
    // invoke: void ~QTabWidget()
    if this != nil && this.Qclsinst != nil {
      C.C_ZN10QTabWidgetD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QTabWidget", "~QTabWidget", args)
  }

  return
}

// insertTab(int, class QWidget *, const class QString &)
func (this *QTabWidget) InsertTab(args ...interface{}) (ret interface{}) {
  // insertTab(int, class QWidget *, const class QString &)
  // insertTab(int, class QWidget *, const class QIcon &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[0][2] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1][2] = reflect.TypeOf(qtgui.QIcon{}) // "const QIcon &"
  vtys[1][3] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTabWidget9insertTabEiP7QWidgetRK7QString
    // invoke: int insertTab(int, class QWidget *, const class QString &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QWidget).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZN10QTabWidget9insertTabEiP7QWidgetRK7QString(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN10QTabWidget9insertTabEiP7QWidgetRK5QIconRK7QString
    // invoke: int insertTab(int, class QWidget *, const class QIcon &, const class QString &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QWidget).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtgui.QIcon).Qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg3)}
    var ret0 = C.C_ZN10QTabWidget9insertTabEiP7QWidgetRK5QIconRK7QString(this.Qclsinst, arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTabWidget", "insertTab", args)
  }

  return
}

// tabShape()
func (this *QTabWidget) TabShape(args ...interface{}) () {
  // tabShape()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTabWidget8tabShapeEv
    // invoke: QTabWidget::TabShape tabShape()
    C.C_ZNK10QTabWidget8tabShapeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTabWidget", "tabShape", args)
  }

  return
}

// indexOf(class QWidget *)
func (this *QTabWidget) IndexOf(args ...interface{}) (ret interface{}) {
  // indexOf(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTabWidget7indexOfEP7QWidget
    // invoke: int indexOf(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QTabWidget7indexOfEP7QWidget(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTabWidget", "indexOf", args)
  }

  return
}

// tabBar()
func (this *QTabWidget) TabBar(args ...interface{}) (ret interface{}) {
  // tabBar()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTabWidget6tabBarEv
    // invoke: QTabBar * tabBar()
    var ret0 = C.C_ZNK10QTabWidget6tabBarEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTabBar{}) // "QTabBar *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTabWidget", "tabBar", args)
  }

  return
}

// setTabToolTip(int, const class QString &)
func (this *QTabWidget) SetTabToolTip(args ...interface{}) () {
  // setTabToolTip(int, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTabWidget13setTabToolTipEiRK7QString
    // invoke: void setTabToolTip(int, const class QString &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN10QTabWidget13setTabToolTipEiRK7QString(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTabWidget", "setTabToolTip", args)
  }

  return
}

// QTabWidget(class QWidget *)
func GcfreeQTabWidget(this *QTabWidget) {
  qtrt.UniverseFree(this)
}
func NewQTabWidget(args ...interface{}) *QTabWidget {
  // QTabWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTabWidgetC1EP7QWidget
    // invoke: void QTabWidget(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN10QTabWidgetC2EP7QWidget(arg0)
    this := &QTabWidget{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQTabWidget)
    return this
  default:
    qtrt.ErrorResolve("QTabWidget", "QTabWidget", args)
  }

  return nil // QTabWidget{Qclsinst:qthis}
}

// setTabBarAutoHide(_Bool)
func (this *QTabWidget) SetTabBarAutoHide(args ...interface{}) () {
  // setTabBarAutoHide(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTabWidget17setTabBarAutoHideEb
    // invoke: void setTabBarAutoHide(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN10QTabWidget17setTabBarAutoHideEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTabWidget", "setTabBarAutoHide", args)
  }

  return
}

// documentMode()
func (this *QTabWidget) DocumentMode(args ...interface{}) (ret interface{}) {
  // documentMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTabWidget12documentModeEv
    // invoke: bool documentMode()
    var ret0 = C.C_ZNK10QTabWidget12documentModeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTabWidget", "documentMode", args)
  }

  return
}

// widget(int)
func (this *QTabWidget) Widget(args ...interface{}) (ret interface{}) {
  // widget(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTabWidget6widgetEi
    // invoke: QWidget * widget(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QTabWidget6widgetEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTabWidget", "widget", args)
  }

  return
}

// removeTab(int)
func (this *QTabWidget) RemoveTab(args ...interface{}) () {
  // removeTab(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTabWidget9removeTabEi
    // invoke: void removeTab(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN10QTabWidget9removeTabEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTabWidget", "removeTab", args)
  }

  return
}

// iconSize()
func (this *QTabWidget) IconSize(args ...interface{}) (ret interface{}) {
  // iconSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTabWidget8iconSizeEv
    // invoke: QSize iconSize()
    var ret0 = C.C_ZNK10QTabWidget8iconSizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTabWidget", "iconSize", args)
  }

  return
}

// setTabText(int, const class QString &)
func (this *QTabWidget) SetTabText(args ...interface{}) () {
  // setTabText(int, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTabWidget10setTabTextEiRK7QString
    // invoke: void setTabText(int, const class QString &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN10QTabWidget10setTabTextEiRK7QString(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTabWidget", "setTabText", args)
  }

  return
}

// tabText(int)
func (this *QTabWidget) TabText(args ...interface{}) (ret interface{}) {
  // tabText(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTabWidget7tabTextEi
    // invoke: QString tabText(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QTabWidget7tabTextEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTabWidget", "tabText", args)
  }

  return
}

// setTabsClosable(_Bool)
func (this *QTabWidget) SetTabsClosable(args ...interface{}) () {
  // setTabsClosable(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTabWidget15setTabsClosableEb
    // invoke: void setTabsClosable(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN10QTabWidget15setTabsClosableEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTabWidget", "setTabsClosable", args)
  }

  return
}

// sizeHint()
func (this *QTabWidget) SizeHint(args ...interface{}) (ret interface{}) {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTabWidget8sizeHintEv
    // invoke: QSize sizeHint()
    var ret0 = C.C_ZNK10QTabWidget8sizeHintEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTabWidget", "sizeHint", args)
  }

  return
}

// tabPosition()
func (this *QTabWidget) TabPosition(args ...interface{}) () {
  // tabPosition()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTabWidget11tabPositionEv
    // invoke: QTabWidget::TabPosition tabPosition()
    C.C_ZNK10QTabWidget11tabPositionEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTabWidget", "tabPosition", args)
  }

  return
}

// usesScrollButtons()
func (this *QTabWidget) UsesScrollButtons(args ...interface{}) (ret interface{}) {
  // usesScrollButtons()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTabWidget17usesScrollButtonsEv
    // invoke: bool usesScrollButtons()
    var ret0 = C.C_ZNK10QTabWidget17usesScrollButtonsEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTabWidget", "usesScrollButtons", args)
  }

  return
}

// setDocumentMode(_Bool)
func (this *QTabWidget) SetDocumentMode(args ...interface{}) () {
  // setDocumentMode(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTabWidget15setDocumentModeEb
    // invoke: void setDocumentMode(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN10QTabWidget15setDocumentModeEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTabWidget", "setDocumentMode", args)
  }

  return
}

// count()
func (this *QTabWidget) Count(args ...interface{}) (ret interface{}) {
  // count()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTabWidget5countEv
    // invoke: int count()
    var ret0 = C.C_ZNK10QTabWidget5countEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTabWidget", "count", args)
  }

  return
}

// setUsesScrollButtons(_Bool)
func (this *QTabWidget) SetUsesScrollButtons(args ...interface{}) () {
  // setUsesScrollButtons(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTabWidget20setUsesScrollButtonsEb
    // invoke: void setUsesScrollButtons(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN10QTabWidget20setUsesScrollButtonsEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTabWidget", "setUsesScrollButtons", args)
  }

  return
}

// metaObject()
func (this *QTabWidget) MetaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTabWidget10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK10QTabWidget10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTabWidget", "metaObject", args)
  }

  return
}

// elideMode()
func (this *QTabWidget) ElideMode(args ...interface{}) () {
  // elideMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTabWidget9elideModeEv
    // invoke: Qt::TextElideMode elideMode()
    C.C_ZNK10QTabWidget9elideModeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTabWidget", "elideMode", args)
  }

  return
}

// hasHeightForWidth()
func (this *QTabWidget) HasHeightForWidth(args ...interface{}) (ret interface{}) {
  // hasHeightForWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTabWidget17hasHeightForWidthEv
    // invoke: bool hasHeightForWidth()
    var ret0 = C.C_ZNK10QTabWidget17hasHeightForWidthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTabWidget", "hasHeightForWidth", args)
  }

  return
}

// setTabWhatsThis(int, const class QString &)
func (this *QTabWidget) SetTabWhatsThis(args ...interface{}) () {
  // setTabWhatsThis(int, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTabWidget15setTabWhatsThisEiRK7QString
    // invoke: void setTabWhatsThis(int, const class QString &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN10QTabWidget15setTabWhatsThisEiRK7QString(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTabWidget", "setTabWhatsThis", args)
  }

  return
}

// clear()
func (this *QTabWidget) Clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTabWidget5clearEv
    // invoke: void clear()
    C.C_ZN10QTabWidget5clearEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTabWidget", "clear", args)
  }

  return
}

// setIconSize(const class QSize &)
func (this *QTabWidget) SetIconSize(args ...interface{}) () {
  // setIconSize(const class QSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QSize{}) // "const QSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTabWidget11setIconSizeERK5QSize
    // invoke: void setIconSize(const class QSize &)
    var arg0 = args[0].(*qtcore.QSize).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN10QTabWidget11setIconSizeERK5QSize(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTabWidget", "setIconSize", args)
  }

  return
}

// isMovable()
func (this *QTabWidget) IsMovable(args ...interface{}) (ret interface{}) {
  // isMovable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTabWidget9isMovableEv
    // invoke: bool isMovable()
    var ret0 = C.C_ZNK10QTabWidget9isMovableEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTabWidget", "isMovable", args)
  }

  return
}

// heightForWidth(int)
func (this *QTabWidget) HeightForWidth(args ...interface{}) (ret interface{}) {
  // heightForWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTabWidget14heightForWidthEi
    // invoke: int heightForWidth(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QTabWidget14heightForWidthEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTabWidget", "heightForWidth", args)
  }

  return
}

// setTabEnabled(int, _Bool)
func (this *QTabWidget) SetTabEnabled(args ...interface{}) () {
  // setTabEnabled(int, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTabWidget13setTabEnabledEib
    // invoke: void setTabEnabled(int, _Bool)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C.C_ZN10QTabWidget13setTabEnabledEib(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTabWidget", "setTabEnabled", args)
  }

  return
}

// currentIndex()
func (this *QTabWidget) CurrentIndex(args ...interface{}) (ret interface{}) {
  // currentIndex()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTabWidget12currentIndexEv
    // invoke: int currentIndex()
    var ret0 = C.C_ZNK10QTabWidget12currentIndexEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTabWidget", "currentIndex", args)
  }

  return
}

// addTab(class QWidget *, const class QIcon &, const class QString &)
func (this *QTabWidget) AddTab(args ...interface{}) (ret interface{}) {
  // addTab(class QWidget *, const class QIcon &, const class QString &)
  // addTab(class QWidget *, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[0][1] = reflect.TypeOf(qtgui.QIcon{}) // "const QIcon &"
  vtys[0][2] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1][1] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTabWidget6addTabEP7QWidgetRK5QIconRK7QString
    // invoke: int addTab(class QWidget *, const class QIcon &, const class QString &)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtgui.QIcon).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZN10QTabWidget6addTabEP7QWidgetRK5QIconRK7QString(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN10QTabWidget6addTabEP7QWidgetRK7QString
    // invoke: int addTab(class QWidget *, const class QString &)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN10QTabWidget6addTabEP7QWidgetRK7QString(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTabWidget", "addTab", args)
  }

  return
}

// <= body block end

