package qtcore

// /usr/include/qt/QtCore/qjsondocument.h
// #include <qjsondocument.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 27
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

type QJsonParseError struct {
	*qtrt.CObject
}

func (this *QJsonParseError) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QJsonParseError) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQJsonParseErrorFromPointer(cthis unsafe.Pointer) *QJsonParseError {
	return &QJsonParseError{&qtrt.CObject{cthis}}
}
func (*QJsonParseError) NewFromPointer(cthis unsafe.Pointer) *QJsonParseError {
	return NewQJsonParseErrorFromPointer(cthis)
}

// /usr/include/qt/QtCore/qjsondocument.h:73
// index:0
// Public Visibility=Default Availability=Available
// [8] QString errorString()
func (this *QJsonParseError) ErrorString() *QString /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QJsonParseError11errorStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQString)
	return rv2
}

func DeleteQJsonParseError(this *QJsonParseError) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QJsonParseErrorD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

type QJsonParseError__ParseError = int

const QJsonParseError__NoError QJsonParseError__ParseError = 0
const QJsonParseError__UnterminatedObject QJsonParseError__ParseError = 1
const QJsonParseError__MissingNameSeparator QJsonParseError__ParseError = 2
const QJsonParseError__UnterminatedArray QJsonParseError__ParseError = 3
const QJsonParseError__MissingValueSeparator QJsonParseError__ParseError = 4
const QJsonParseError__IllegalValue QJsonParseError__ParseError = 5
const QJsonParseError__TerminationByNumber QJsonParseError__ParseError = 6
const QJsonParseError__IllegalNumber QJsonParseError__ParseError = 7
const QJsonParseError__IllegalEscapeSequence QJsonParseError__ParseError = 8
const QJsonParseError__IllegalUTF8String QJsonParseError__ParseError = 9
const QJsonParseError__UnterminatedString QJsonParseError__ParseError = 10
const QJsonParseError__MissingObject QJsonParseError__ParseError = 11
const QJsonParseError__DeepNesting QJsonParseError__ParseError = 12
const QJsonParseError__DocumentTooLarge QJsonParseError__ParseError = 13
const QJsonParseError__GarbageAtEnd QJsonParseError__ParseError = 14

//  body block end
