//  header block begin
// /usr/include/qt/QtCore/qabstractanimation.h
// #include <qabstractanimation.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 29
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"

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
}

//  ext block end

//  body block begin
type QAnimationDriver struct {
	*QObject
}

func (this *QAnimationDriver) GetCthis() unsafe.Pointer {
	return this.QObject.GetCthis()
}

// /usr/include/qt/QtCore/qabstractanimation.h:135
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QAnimationDriver) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAnimationDriver10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:139
// index:0
// void QAnimationDriver(class QObject *)
func NewQAnimationDriver(parent unsafe.Pointer) *QAnimationDriver {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAnimationDriverC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQAnimationDriverFromPointer(cthis)
	return gothis
}
func NewQAnimationDriverFromPointer(cthis unsafe.Pointer) *QAnimationDriver {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QAnimationDriver{bcthis0}
}

// /usr/include/qt/QtCore/qabstractanimation.h:165
// index:1
// void QAnimationDriver(class QAnimationDriverPrivate &, class QObject *)
func NewQAnimationDriver_1(dd unsafe.Pointer, parent unsafe.Pointer) *QAnimationDriver {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAnimationDriverC2ER23QAnimationDriverPrivateP7QObject", ffiqt.FFI_TYPE_VOID, cthis, dd, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQAnimationDriverFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qabstractanimation.h:140
// index:0
// virtual
// void ~QAnimationDriver()
func DeleteQAnimationDriver(*QAnimationDriver) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAnimationDriverD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:142
// index:0
// virtual
// void advance()
func (this *QAnimationDriver) Advance() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAnimationDriver7advanceEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:144
// index:0
// void install()
func (this *QAnimationDriver) Install() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAnimationDriver7installEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:145
// index:0
// void uninstall()
func (this *QAnimationDriver) Uninstall() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAnimationDriver9uninstallEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:147
// index:0
// bool isRunning()
func (this *QAnimationDriver) IsRunning() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAnimationDriver9isRunningEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:149
// index:0
// virtual
// qint64 elapsed()
func (this *QAnimationDriver) Elapsed() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAnimationDriver7elapsedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:152
// index:0
// void setStartTime(qint64)
func (this *QAnimationDriver) SetStartTime(startTime int64) {
	// 0: (, startTime qint64), (&startTime)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAnimationDriver12setStartTimeEx", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &startTime)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:153
// index:0
// qint64 startTime()
func (this *QAnimationDriver) StartTime() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAnimationDriver9startTimeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:156
// index:0
// void started()
func (this *QAnimationDriver) Started() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAnimationDriver7startedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:157
// index:0
// void stopped()
func (this *QAnimationDriver) Stopped() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAnimationDriver7stoppedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:161
// index:0
// void advanceAnimation(qint64)
func (this *QAnimationDriver) AdvanceAnimation(timeStep int64) {
	// 0: (, timeStep qint64), (&timeStep)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAnimationDriver16advanceAnimationEx", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &timeStep)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:162
// index:0
// virtual
// void start()
func (this *QAnimationDriver) Start() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAnimationDriver5startEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:163
// index:0
// virtual
// void stop()
func (this *QAnimationDriver) Stop() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAnimationDriver4stopEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
