//  header block begin
// /usr/include/qt/QtGui/qdesktopservices.h
// #include <qdesktopservices.h>
// #include <QtGui>
package qtgui

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
type QDesktopServices struct {
	*qtrt.CObject
}

func (this *QDesktopServices) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtGui/qdesktopservices.h:59
// index:0
// static
// bool openUrl(const class QUrl &)
func (this *QDesktopServices) OpenUrl(url unsafe.Pointer) {
	// 0: (url const QUrl &), (url)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDesktopServices7openUrlERK4QUrl", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QDesktopServices_OpenUrl(url unsafe.Pointer) {
	// 0: (url const QUrl &), (url)
	var nilthis *QDesktopServices
	nilthis.OpenUrl(url)
}

// /usr/include/qt/QtGui/qdesktopservices.h:60
// index:0
// static
// void setUrlHandler(const class QString &, class QObject *, const char *)
func (this *QDesktopServices) SetUrlHandler(scheme unsafe.Pointer, receiver unsafe.Pointer, method unsafe.Pointer) {
	// 0: (scheme const QString &, receiver QObject *, method const char *), (scheme, receiver, method)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDesktopServices13setUrlHandlerERK7QStringP7QObjectPKc", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QDesktopServices_SetUrlHandler(scheme unsafe.Pointer, receiver unsafe.Pointer, method unsafe.Pointer) {
	// 0: (scheme const QString &, receiver QObject *, method const char *), (scheme, receiver, method)
	var nilthis *QDesktopServices
	nilthis.SetUrlHandler(scheme, receiver, method)
}

// /usr/include/qt/QtGui/qdesktopservices.h:61
// index:0
// static
// void unsetUrlHandler(const class QString &)
func (this *QDesktopServices) UnsetUrlHandler(scheme unsafe.Pointer) {
	// 0: (scheme const QString &), (scheme)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QDesktopServices15unsetUrlHandlerERK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QDesktopServices_UnsetUrlHandler(scheme unsafe.Pointer) {
	// 0: (scheme const QString &), (scheme)
	var nilthis *QDesktopServices
	nilthis.UnsetUrlHandler(scheme)
}

//  body block end
