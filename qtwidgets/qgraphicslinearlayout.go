package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h
// #include <qgraphicslinearlayout.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 26
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

type QGraphicsLinearLayout struct {
	*QGraphicsLayout
}

func (this *QGraphicsLinearLayout) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QGraphicsLayout.GetCthis()
	}
}
func (this *QGraphicsLinearLayout) SetCthis(cthis unsafe.Pointer) {
	this.QGraphicsLayout = NewQGraphicsLayoutFromPointer(cthis)
}
func NewQGraphicsLinearLayoutFromPointer(cthis unsafe.Pointer) *QGraphicsLinearLayout {
	bcthis0 := NewQGraphicsLayoutFromPointer(cthis)
	return &QGraphicsLinearLayout{bcthis0}
}
func (*QGraphicsLinearLayout) NewFromPointer(cthis unsafe.Pointer) *QGraphicsLinearLayout {
	return NewQGraphicsLinearLayoutFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsLinearLayout(QGraphicsLayoutItem *)
func NewQGraphicsLinearLayout(parent *QGraphicsLayoutItem /*777 QGraphicsLayoutItem **/) *QGraphicsLinearLayout {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGraphicsLinearLayoutC2EP19QGraphicsLayoutItem", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsLinearLayoutFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsLinearLayout)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:57
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsLinearLayout(Qt::Orientation, QGraphicsLayoutItem *)
func NewQGraphicsLinearLayout_1(orientation int, parent *QGraphicsLayoutItem /*777 QGraphicsLayoutItem **/) *QGraphicsLinearLayout {
	var convArg1 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGraphicsLinearLayoutC2EN2Qt11OrientationEP19QGraphicsLayoutItem", qtrt.FFI_TYPE_POINTER, orientation, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsLinearLayoutFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsLinearLayout)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:58
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGraphicsLinearLayout()
func DeleteQGraphicsLinearLayout(this *QGraphicsLinearLayout) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGraphicsLinearLayoutD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:60
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOrientation(Qt::Orientation)
func (this *QGraphicsLinearLayout) SetOrientation(orientation int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGraphicsLinearLayout14setOrientationEN2Qt11OrientationE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), orientation)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:61
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::Orientation orientation()
func (this *QGraphicsLinearLayout) Orientation() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QGraphicsLinearLayout11orientationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:63
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void addItem(QGraphicsLayoutItem *)
func (this *QGraphicsLinearLayout) AddItem(item *QGraphicsLayoutItem /*777 QGraphicsLayoutItem **/) {
	var convArg0 = item.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGraphicsLinearLayout7addItemEP19QGraphicsLayoutItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:64
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void addStretch(int)
func (this *QGraphicsLinearLayout) AddStretch(stretch int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGraphicsLinearLayout10addStretchEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), stretch)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertItem(int, QGraphicsLayoutItem *)
func (this *QGraphicsLinearLayout) InsertItem(index int, item *QGraphicsLayoutItem /*777 QGraphicsLayoutItem **/) {
	var convArg1 = item.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGraphicsLinearLayout10insertItemEiP19QGraphicsLayoutItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:67
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertStretch(int, int)
func (this *QGraphicsLinearLayout) InsertStretch(index int, stretch int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGraphicsLinearLayout13insertStretchEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, stretch)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeItem(QGraphicsLayoutItem *)
func (this *QGraphicsLinearLayout) RemoveItem(item *QGraphicsLayoutItem /*777 QGraphicsLayoutItem **/) {
	var convArg0 = item.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGraphicsLinearLayout10removeItemEP19QGraphicsLayoutItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:70
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void removeAt(int)
func (this *QGraphicsLinearLayout) RemoveAt(index int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGraphicsLinearLayout8removeAtEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSpacing(qreal)
func (this *QGraphicsLinearLayout) SetSpacing(spacing float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGraphicsLinearLayout10setSpacingEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), spacing)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:73
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal spacing()
func (this *QGraphicsLinearLayout) Spacing() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QGraphicsLinearLayout7spacingEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setItemSpacing(int, qreal)
func (this *QGraphicsLinearLayout) SetItemSpacing(index int, spacing float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGraphicsLinearLayout14setItemSpacingEid", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, spacing)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:75
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal itemSpacing(int)
func (this *QGraphicsLinearLayout) ItemSpacing(index int) float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QGraphicsLinearLayout11itemSpacingEi", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStretchFactor(QGraphicsLayoutItem *, int)
func (this *QGraphicsLinearLayout) SetStretchFactor(item *QGraphicsLayoutItem /*777 QGraphicsLayoutItem **/, stretch int) {
	var convArg0 = item.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGraphicsLinearLayout16setStretchFactorEP19QGraphicsLayoutItemi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, stretch)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:78
// index:0
// Public Visibility=Default Availability=Available
// [4] int stretchFactor(QGraphicsLayoutItem *)
func (this *QGraphicsLinearLayout) StretchFactor(item *QGraphicsLayoutItem /*777 QGraphicsLayoutItem **/) int {
	var convArg0 = item.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QGraphicsLinearLayout13stretchFactorEP19QGraphicsLayoutItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAlignment(QGraphicsLayoutItem *, Qt::Alignment)
func (this *QGraphicsLinearLayout) SetAlignment(item *QGraphicsLayoutItem /*777 QGraphicsLayoutItem **/, alignment int) {
	var convArg0 = item.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGraphicsLinearLayout12setAlignmentEP19QGraphicsLayoutItem6QFlagsIN2Qt13AlignmentFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, alignment)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:81
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::Alignment alignment(QGraphicsLayoutItem *)
func (this *QGraphicsLinearLayout) Alignment(item *QGraphicsLayoutItem /*777 QGraphicsLayoutItem **/) int {
	var convArg0 = item.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QGraphicsLinearLayout9alignmentEP19QGraphicsLayoutItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:83
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setGeometry(const QRectF &)
func (this *QGraphicsLinearLayout) SetGeometry(rect *qtcore.QRectF) {
	var convArg0 = rect.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGraphicsLinearLayout11setGeometryERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:85
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int count()
func (this *QGraphicsLinearLayout) Count() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QGraphicsLinearLayout5countEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:86
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QGraphicsLayoutItem * itemAt(int)
func (this *QGraphicsLinearLayout) ItemAt(index int) *QGraphicsLayoutItem /*777 QGraphicsLayoutItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QGraphicsLinearLayout6itemAtEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQGraphicsLayoutItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:88
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void invalidate()
func (this *QGraphicsLinearLayout) Invalidate() {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGraphicsLinearLayout10invalidateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:89
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QSizeF sizeHint(Qt::SizeHint, const QSizeF &)
func (this *QGraphicsLinearLayout) SizeHint(which int, constraint *qtcore.QSizeF) *qtcore.QSizeF /*123*/ {
	var convArg1 = constraint.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QGraphicsLinearLayout8sizeHintEN2Qt8SizeHintERK6QSizeF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), which, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:95
// index:0
// Public Visibility=Default Availability=Available
// [-2] void dump(int)
func (this *QGraphicsLinearLayout) Dump(indent int) {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QGraphicsLinearLayout4dumpEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), indent)
	gopp.ErrPrint(err, rv)
}

//  body block end
