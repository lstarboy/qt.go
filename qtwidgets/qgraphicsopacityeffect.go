package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicseffect.h
// #include <qgraphicseffect.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 20
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"

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
// void draw(class QPainter *)
func (this *QGraphicsOpacityEffect) InheritDraw(f func(painter *qtgui.QPainter /*777 QPainter **/)) {
	qtrt.SetAllInheritCallback(this, "draw", f)
}

type QGraphicsOpacityEffect struct {
	*QGraphicsEffect
}

func (this *QGraphicsOpacityEffect) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QGraphicsEffect.GetCthis()
	}
}
func (this *QGraphicsOpacityEffect) SetCthis(cthis unsafe.Pointer) {
	this.QGraphicsEffect = NewQGraphicsEffectFromPointer(cthis)
}
func NewQGraphicsOpacityEffectFromPointer(cthis unsafe.Pointer) *QGraphicsOpacityEffect {
	bcthis0 := NewQGraphicsEffectFromPointer(cthis)
	return &QGraphicsOpacityEffect{bcthis0}
}
func (*QGraphicsOpacityEffect) NewFromPointer(cthis unsafe.Pointer) *QGraphicsOpacityEffect {
	return NewQGraphicsOpacityEffectFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:254
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QGraphicsOpacityEffect) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QGraphicsOpacityEffect10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:258
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsOpacityEffect(QObject *)
func NewQGraphicsOpacityEffect(parent *qtcore.QObject /*777 QObject **/) *QGraphicsOpacityEffect {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN22QGraphicsOpacityEffectC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsOpacityEffectFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:259
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGraphicsOpacityEffect()
func DeleteQGraphicsOpacityEffect(this *QGraphicsOpacityEffect) {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QGraphicsOpacityEffectD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:261
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal opacity()
func (this *QGraphicsOpacityEffect) Opacity() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QGraphicsOpacityEffect7opacityEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:262
// index:0
// Public Visibility=Default Availability=Available
// [8] QBrush opacityMask()
func (this *QGraphicsOpacityEffect) OpacityMask() *qtgui.QBrush /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QGraphicsOpacityEffect11opacityMaskEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:265
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOpacity(qreal)
func (this *QGraphicsOpacityEffect) SetOpacity(opacity float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QGraphicsOpacityEffect10setOpacityEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), opacity)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:266
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOpacityMask(const QBrush &)
func (this *QGraphicsOpacityEffect) SetOpacityMask(mask *qtgui.QBrush) {
	var convArg0 = mask.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN22QGraphicsOpacityEffect14setOpacityMaskERK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:269
// index:0
// Public Visibility=Default Availability=Available
// [-2] void opacityChanged(qreal)
func (this *QGraphicsOpacityEffect) OpacityChanged(opacity float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QGraphicsOpacityEffect14opacityChangedEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), opacity)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:270
// index:0
// Public Visibility=Default Availability=Available
// [-2] void opacityMaskChanged(const QBrush &)
func (this *QGraphicsOpacityEffect) OpacityMaskChanged(mask *qtgui.QBrush) {
	var convArg0 = mask.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN22QGraphicsOpacityEffect18opacityMaskChangedERK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:273
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void draw(QPainter *)
func (this *QGraphicsOpacityEffect) Draw(painter *qtgui.QPainter /*777 QPainter **/) {
	var convArg0 = painter.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN22QGraphicsOpacityEffect4drawEP8QPainter", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
