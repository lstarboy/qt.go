//  header block begin
// /usr/include/qt/QtGui/qaccessibleplugin.h
// #include <qaccessibleplugin.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 10
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

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
}

//  ext block end

//  body block begin
type QAccessiblePlugin struct {
	*qtcore.QObject
}

func (this *QAccessiblePlugin) GetCthis() unsafe.Pointer {
	return this.QObject.GetCthis()
}

// /usr/include/qt/QtGui/qaccessibleplugin.h:61
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QAccessiblePlugin) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAccessiblePlugin10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessibleplugin.h:63
// index:0
// void QAccessiblePlugin(class QObject *)
func NewQAccessiblePlugin(parent unsafe.Pointer) *QAccessiblePlugin {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAccessiblePluginC1EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQAccessiblePluginFromPointer(cthis)
	return gothis
}
func NewQAccessiblePluginFromPointer(cthis unsafe.Pointer) *QAccessiblePlugin {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QAccessiblePlugin{bcthis0}
}

// /usr/include/qt/QtGui/qaccessibleplugin.h:64
// index:0
// virtual
// void ~QAccessiblePlugin()
func DeleteQAccessiblePlugin(*QAccessiblePlugin) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAccessiblePluginD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessibleplugin.h:66
// index:0
// pure virtual
// QAccessibleInterface * create(const class QString &, class QObject *)
func (this *QAccessiblePlugin) Create(key unsafe.Pointer, object unsafe.Pointer) {
	// 0: (, key const QString &, object QObject *), (key, object)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAccessiblePlugin6createERK7QStringP7QObject", ffiqt.FFI_TYPE_VOID, this.GetCthis(), key, object)
	gopp.ErrPrint(err, rv)
}

//  body block end
