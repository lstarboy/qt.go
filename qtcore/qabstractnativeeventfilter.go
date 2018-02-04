package qtcore

// /usr/include/qt/QtCore/qabstractnativeeventfilter.h
// #include <qabstractnativeeventfilter.h>
// #include <QtCore>

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
}

//  ext block end

//  body block begin

type QAbstractNativeEventFilter struct {
	*qtrt.CObject
}

func (this *QAbstractNativeEventFilter) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QAbstractNativeEventFilter) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQAbstractNativeEventFilterFromPointer(cthis unsafe.Pointer) *QAbstractNativeEventFilter {
	return &QAbstractNativeEventFilter{&qtrt.CObject{cthis}}
}
func (*QAbstractNativeEventFilter) NewFromPointer(cthis unsafe.Pointer) *QAbstractNativeEventFilter {
	return NewQAbstractNativeEventFilterFromPointer(cthis)
}

// /usr/include/qt/QtCore/qabstractnativeeventfilter.h:52
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractNativeEventFilter()
func NewQAbstractNativeEventFilter() *QAbstractNativeEventFilter {
	rv, err := qtrt.InvokeQtFunc6("_ZN26QAbstractNativeEventFilterC1Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQAbstractNativeEventFilterFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAbstractNativeEventFilter)
	return gothis
}

// /usr/include/qt/QtCore/qabstractnativeeventfilter.h:53
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAbstractNativeEventFilter()
func DeleteQAbstractNativeEventFilter(this *QAbstractNativeEventFilter) {
	rv, err := qtrt.InvokeQtFunc6("_ZN26QAbstractNativeEventFilterD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qabstractnativeeventfilter.h:55
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool nativeEventFilter(const QByteArray &, void *, long *)
func (this *QAbstractNativeEventFilter) NativeEventFilter(eventType *QByteArray, message unsafe.Pointer /*666*/, result unsafe.Pointer /*666*/) bool {
	var convArg0 = eventType.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN26QAbstractNativeEventFilter17nativeEventFilterERK10QByteArrayPvPl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, message, &result)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

//  body block end
