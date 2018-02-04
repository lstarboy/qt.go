package qtgui

// /usr/include/qt/QtGui/qaccessible.h
// #include <qaccessible.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 4
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

type QAccessibleEvent struct {
	*qtrt.CObject
}

func (this *QAccessibleEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QAccessibleEvent) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQAccessibleEventFromPointer(cthis unsafe.Pointer) *QAccessibleEvent {
	return &QAccessibleEvent{&qtrt.CObject{cthis}}
}
func (*QAccessibleEvent) NewFromPointer(cthis unsafe.Pointer) *QAccessibleEvent {
	return NewQAccessibleEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qaccessible.h:668
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QAccessibleEvent(QObject *, QAccessible::Event)
func NewQAccessibleEvent(obj *qtcore.QObject /*777 QObject **/, typ int) *QAccessibleEvent {
	var convArg0 = obj.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAccessibleEventC2EP7QObjectN11QAccessible5EventE", qtrt.FFI_TYPE_POINTER, convArg0, typ)
	gopp.ErrPrint(err, rv)
	gothis := NewQAccessibleEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAccessibleEvent)
	return gothis
}

// /usr/include/qt/QtGui/qaccessible.h:684
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QAccessibleEvent(QAccessibleInterface *, QAccessible::Event)
func NewQAccessibleEvent_1(iface *QAccessibleInterface /*777 QAccessibleInterface **/, typ int) *QAccessibleEvent {
	var convArg0 = iface.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAccessibleEventC2EP20QAccessibleInterfaceN11QAccessible5EventE", qtrt.FFI_TYPE_POINTER, convArg0, typ)
	gopp.ErrPrint(err, rv)
	gothis := NewQAccessibleEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAccessibleEvent)
	return gothis
}

// /usr/include/qt/QtGui/qaccessible.h:699
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAccessibleEvent()
func DeleteQAccessibleEvent(this *QAccessibleEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAccessibleEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 32)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qaccessible.h:701
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QAccessible::Event type()
func (this *QAccessibleEvent) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QAccessibleEvent4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qaccessible.h:702
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QObject * object()
func (this *QAccessibleEvent) Object() *qtcore.QObject /*777 QObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QAccessibleEvent6objectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qaccessible.h:703
// index:0
// Public Visibility=Default Availability=Available
// [4] QAccessible::Id uniqueId()
func (this *QAccessibleEvent) UniqueId() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QAccessibleEvent8uniqueIdEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qaccessible.h:705
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setChild(int)
func (this *QAccessibleEvent) SetChild(chld int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAccessibleEvent8setChildEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), chld)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:706
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int child()
func (this *QAccessibleEvent) Child() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QAccessibleEvent5childEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qaccessible.h:708
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QAccessibleInterface * accessibleInterface()
func (this *QAccessibleEvent) AccessibleInterface() *QAccessibleInterface /*777 QAccessibleInterface **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QAccessibleEvent19accessibleInterfaceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQAccessibleInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

//  body block end
