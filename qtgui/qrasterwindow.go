package qtgui

// /usr/include/qt/QtGui/qrasterwindow.h
// #include <qrasterwindow.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
}

//  ext block end

//  body block begin
// int metric(enum QPaintDevice::PaintDeviceMetric)
func (this *QRasterWindow) InheritMetric(f func(metric int) int) {
	qtrt.SetAllInheritCallback(this, "metric", f)
}

// QPaintDevice * redirected(class QPoint *)
func (this *QRasterWindow) InheritRedirected(f func(arg0 *qtcore.QPoint /*777 QPoint **/) unsafe.Pointer /*666*/) {
	qtrt.SetAllInheritCallback(this, "redirected", f)
}

type QRasterWindow struct {
	*QPaintDeviceWindow
}

func (this *QRasterWindow) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QPaintDeviceWindow.GetCthis()
	}
}
func (this *QRasterWindow) SetCthis(cthis unsafe.Pointer) {
	this.QPaintDeviceWindow = NewQPaintDeviceWindowFromPointer(cthis)
}
func NewQRasterWindowFromPointer(cthis unsafe.Pointer) *QRasterWindow {
	bcthis0 := NewQPaintDeviceWindowFromPointer(cthis)
	return &QRasterWindow{bcthis0}
}
func (*QRasterWindow) NewFromPointer(cthis unsafe.Pointer) *QRasterWindow {
	return NewQRasterWindowFromPointer(cthis)
}

// /usr/include/qt/QtGui/qrasterwindow.h:52
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QRasterWindow) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QRasterWindow10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qrasterwindow.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QRasterWindow(QWindow *)
func NewQRasterWindow(parent *QWindow /*777 QWindow **/) *QRasterWindow {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QRasterWindowC2EP7QWindow", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQRasterWindowFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qrasterwindow.h:57
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QRasterWindow()
func DeleteQRasterWindow(this *QRasterWindow) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QRasterWindowD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 64)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qrasterwindow.h:60
// index:0
// Protected virtual Visibility=Default Availability=Available
// [4] int metric(enum QPaintDevice::PaintDeviceMetric)
func (this *QRasterWindow) Metric(metric int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QRasterWindow6metricEN12QPaintDevice17PaintDeviceMetricE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), metric)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qrasterwindow.h:61
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QPaintDevice * redirected(QPoint *)
func (this *QRasterWindow) Redirected(arg0 *qtcore.QPoint /*777 QPoint **/) *QPaintDevice /*777 QPaintDevice **/ {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QRasterWindow10redirectedEP6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQPaintDeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

//  body block end
