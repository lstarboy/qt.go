//  header block begin
// /usr/include/qt/QtGui/qpolygon.h
// #include <qpolygon.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 72
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
type QPolygon struct {
	*qtrt.CObject
}

func (this *QPolygon) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtGui/qpolygon.h:59
// index:0
// inline
// void QPolygon()
func NewQPolygon() *QPolygon {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPolygonC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQPolygonFromPointer(cthis)
	return gothis
}
func NewQPolygonFromPointer(cthis unsafe.Pointer) *QPolygon {
	return &QPolygon{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qpolygon.h:61
// index:1
// inline
// void QPolygon(int)
func NewQPolygon_1(size int) *QPolygon {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPolygonC2Ei", ffiqt.FFI_TYPE_VOID, cthis, &size)
	gopp.ErrPrint(err, rv)
	gothis := NewQPolygonFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qpolygon.h:66
// index:2
// void QPolygon(const class QRect &, _Bool)
func NewQPolygon_2(r unsafe.Pointer, closed bool) *QPolygon {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPolygonC2ERK5QRectb", ffiqt.FFI_TYPE_VOID, cthis, r, &closed)
	gopp.ErrPrint(err, rv)
	gothis := NewQPolygonFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qpolygon.h:67
// index:3
// void QPolygon(int, const int *)
func NewQPolygon_3(nPoints int, points unsafe.Pointer) *QPolygon {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPolygonC2EiPKi", ffiqt.FFI_TYPE_VOID, cthis, &nPoints, points)
	gopp.ErrPrint(err, rv)
	gothis := NewQPolygonFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qpolygon.h:60
// index:0
// inline
// void ~QPolygon()
func DeleteQPolygon(*QPolygon) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPolygonD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:74
// index:0
// inline
// void swap(class QPolygon &)
func (this *QPolygon) Swap(other unsafe.Pointer) {
	// 0: (, other QPolygon &), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPolygon4swapERS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:78
// index:0
// void translate(int, int)
func (this *QPolygon) Translate(dx int, dy int) {
	// 0: (, dx int, dy int), (&dx, &dy)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPolygon9translateEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &dx, &dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:79
// index:1
// void translate(const class QPoint &)
func (this *QPolygon) Translate_1(offset unsafe.Pointer) {
	// 1: (, offset const QPoint &), (offset)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPolygon9translateERK6QPoint", ffiqt.FFI_TYPE_VOID, this.GetCthis(), offset)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:81
// index:0
// QPolygon translated(int, int)
func (this *QPolygon) Translated(dx int, dy int) {
	// 0: (, dx int, dy int), (&dx, &dy)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPolygon10translatedEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &dx, &dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:82
// index:1
// inline
// QPolygon translated(const class QPoint &)
func (this *QPolygon) Translated_1(offset unsafe.Pointer) {
	// 1: (, offset const QPoint &), (offset)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPolygon10translatedERK6QPoint", ffiqt.FFI_TYPE_VOID, this.GetCthis(), offset)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:84
// index:0
// QRect boundingRect()
func (this *QPolygon) BoundingRect() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPolygon12boundingRectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:86
// index:0
// void point(int, int *, int *)
func (this *QPolygon) Point(i int, x unsafe.Pointer, y unsafe.Pointer) {
	// 0: (, i int, x int *, y int *), (&i, x, y)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPolygon5pointEiPiS0_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &i, x, y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:87
// index:1
// QPoint point(int)
func (this *QPolygon) Point_1(i int) {
	// 1: (, i int), (&i)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPolygon5pointEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &i)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:88
// index:0
// void setPoint(int, int, int)
func (this *QPolygon) SetPoint(index int, x int, y int) {
	// 0: (, index int, x int, y int), (&index, &x, &y)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPolygon8setPointEiii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index, &x, &y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:89
// index:1
// void setPoint(int, const class QPoint &)
func (this *QPolygon) SetPoint_1(index int, p unsafe.Pointer) {
	// 1: (, index int, p const QPoint &), (&index, p)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPolygon8setPointEiRK6QPoint", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index, p)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:90
// index:0
// void setPoints(int, const int *)
func (this *QPolygon) SetPoints(nPoints int, points unsafe.Pointer) {
	// 0: (, nPoints int, points const int *), (&nPoints, points)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPolygon9setPointsEiPKi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &nPoints, points)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:92
// index:0
// void putPoints(int, int, const int *)
func (this *QPolygon) PutPoints(index int, nPoints int, points unsafe.Pointer) {
	// 0: (, index int, nPoints int, points const int *), (&index, &nPoints, points)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPolygon9putPointsEiiPKi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index, &nPoints, points)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:94
// index:1
// void putPoints(int, int, const class QPolygon &, int)
func (this *QPolygon) PutPoints_1(index int, nPoints int, from unsafe.Pointer, fromIndex int) {
	// 1: (, index int, nPoints int, from const QPolygon &, fromIndex int), (&index, &nPoints, from, &fromIndex)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPolygon9putPointsEiiRKS_i", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index, &nPoints, from, &fromIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:96
// index:0
// bool containsPoint(const class QPoint &, Qt::FillRule)
func (this *QPolygon) ContainsPoint(pt unsafe.Pointer, fillRule int) {
	// 0: (, pt const QPoint &, fillRule Qt::FillRule), (pt, &fillRule)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPolygon13containsPointERK6QPointN2Qt8FillRuleE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pt, &fillRule)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:98
// index:0
// QPolygon united(const class QPolygon &)
func (this *QPolygon) United(r unsafe.Pointer) {
	// 0: (, r const QPolygon &), (r)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPolygon6unitedERKS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), r)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:99
// index:0
// QPolygon intersected(const class QPolygon &)
func (this *QPolygon) Intersected(r unsafe.Pointer) {
	// 0: (, r const QPolygon &), (r)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPolygon11intersectedERKS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), r)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:100
// index:0
// QPolygon subtracted(const class QPolygon &)
func (this *QPolygon) Subtracted(r unsafe.Pointer) {
	// 0: (, r const QPolygon &), (r)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPolygon10subtractedERKS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), r)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:102
// index:0
// bool intersects(const class QPolygon &)
func (this *QPolygon) Intersects(r unsafe.Pointer) {
	// 0: (, r const QPolygon &), (r)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPolygon10intersectsERKS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), r)
	gopp.ErrPrint(err, rv)
}

//  body block end
