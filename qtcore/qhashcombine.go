package qtcore

// /usr/include/qt/QtCore/qhashfunctions.h
// #include <qhashfunctions.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 83
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

type QHashCombine struct {
	*qtrt.CObject
}

func (this *QHashCombine) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QHashCombine) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQHashCombineFromPointer(cthis unsafe.Pointer) *QHashCombine {
	return &QHashCombine{&qtrt.CObject{cthis}}
}
func (*QHashCombine) NewFromPointer(cthis unsafe.Pointer) *QHashCombine {
	return NewQHashCombineFromPointer(cthis)
}

func DeleteQHashCombine(this *QHashCombine) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QHashCombineD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end
