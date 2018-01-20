//  header block begin
// /usr/include/qt/QtWidgets/qformlayout.h
// #include <qformlayout.h>
// #include <QtWidgets>
package qtwidgets

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
import "qtgui"

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
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin
type QFormLayout struct {
	*QLayout
}

func (this *QFormLayout) GetCthis() unsafe.Pointer {
	return this.QLayout.GetCthis()
}

// /usr/include/qt/QtWidgets/qformlayout.h:55
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QFormLayout) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFormLayout10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:91
// index:0
// void QFormLayout(class QWidget *)
func NewQFormLayout(parent unsafe.Pointer) *QFormLayout {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayoutC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQFormLayoutFromPointer(cthis)
	return gothis
}
func NewQFormLayoutFromPointer(cthis unsafe.Pointer) *QFormLayout {
	bcthis0 := NewQLayoutFromPointer(cthis)
	return &QFormLayout{bcthis0}
}

// /usr/include/qt/QtWidgets/qformlayout.h:92
// index:0
// virtual
// void ~QFormLayout()
func DeleteQFormLayout(*QFormLayout) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayoutD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:94
// index:0
// void setFieldGrowthPolicy(enum QFormLayout::FieldGrowthPolicy)
func (this *QFormLayout) SetFieldGrowthPolicy(policy int) {
	// 0: (, policy QFormLayout::FieldGrowthPolicy), (&policy)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayout20setFieldGrowthPolicyENS_17FieldGrowthPolicyE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &policy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:95
// index:0
// QFormLayout::FieldGrowthPolicy fieldGrowthPolicy()
func (this *QFormLayout) FieldGrowthPolicy() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFormLayout17fieldGrowthPolicyEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:96
// index:0
// void setRowWrapPolicy(enum QFormLayout::RowWrapPolicy)
func (this *QFormLayout) SetRowWrapPolicy(policy int) {
	// 0: (, policy QFormLayout::RowWrapPolicy), (&policy)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayout16setRowWrapPolicyENS_13RowWrapPolicyE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &policy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:97
// index:0
// QFormLayout::RowWrapPolicy rowWrapPolicy()
func (this *QFormLayout) RowWrapPolicy() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFormLayout13rowWrapPolicyEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:98
// index:0
// void setLabelAlignment(Qt::Alignment)
func (this *QFormLayout) SetLabelAlignment(alignment int) {
	// 0: (, QFlags<Qt::AlignmentFlag> alignment), (&alignment)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayout17setLabelAlignmentE6QFlagsIN2Qt13AlignmentFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &alignment)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:99
// index:0
// Qt::Alignment labelAlignment()
func (this *QFormLayout) LabelAlignment() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFormLayout14labelAlignmentEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:100
// index:0
// void setFormAlignment(Qt::Alignment)
func (this *QFormLayout) SetFormAlignment(alignment int) {
	// 0: (, QFlags<Qt::AlignmentFlag> alignment), (&alignment)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayout16setFormAlignmentE6QFlagsIN2Qt13AlignmentFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &alignment)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:101
// index:0
// Qt::Alignment formAlignment()
func (this *QFormLayout) FormAlignment() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFormLayout13formAlignmentEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:103
// index:0
// void setHorizontalSpacing(int)
func (this *QFormLayout) SetHorizontalSpacing(spacing int) {
	// 0: (, spacing int), (&spacing)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayout20setHorizontalSpacingEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &spacing)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:104
// index:0
// int horizontalSpacing()
func (this *QFormLayout) HorizontalSpacing() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFormLayout17horizontalSpacingEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:105
// index:0
// void setVerticalSpacing(int)
func (this *QFormLayout) SetVerticalSpacing(spacing int) {
	// 0: (, spacing int), (&spacing)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayout18setVerticalSpacingEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &spacing)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:106
// index:0
// int verticalSpacing()
func (this *QFormLayout) VerticalSpacing() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFormLayout15verticalSpacingEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:108
// index:0
// int spacing()
func (this *QFormLayout) Spacing() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFormLayout7spacingEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:109
// index:0
// void setSpacing(int)
func (this *QFormLayout) SetSpacing(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayout10setSpacingEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:111
// index:0
// void addRow(class QWidget *, class QWidget *)
func (this *QFormLayout) AddRow(label unsafe.Pointer, field unsafe.Pointer) {
	// 0: (, label QWidget *, field QWidget *), (label, field)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayout6addRowEP7QWidgetS1_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), label, field)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:112
// index:1
// void addRow(class QWidget *, class QLayout *)
func (this *QFormLayout) AddRow_1(label unsafe.Pointer, field unsafe.Pointer) {
	// 1: (, label QWidget *, field QLayout *), (label, field)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayout6addRowEP7QWidgetP7QLayout", ffiqt.FFI_TYPE_VOID, this.GetCthis(), label, field)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:113
// index:2
// void addRow(const class QString &, class QWidget *)
func (this *QFormLayout) AddRow_2(labelText unsafe.Pointer, field unsafe.Pointer) {
	// 2: (, labelText const QString &, field QWidget *), (labelText, field)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayout6addRowERK7QStringP7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), labelText, field)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:114
// index:3
// void addRow(const class QString &, class QLayout *)
func (this *QFormLayout) AddRow_3(labelText unsafe.Pointer, field unsafe.Pointer) {
	// 3: (, labelText const QString &, field QLayout *), (labelText, field)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayout6addRowERK7QStringP7QLayout", ffiqt.FFI_TYPE_VOID, this.GetCthis(), labelText, field)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:115
// index:4
// void addRow(class QWidget *)
func (this *QFormLayout) AddRow_4(widget unsafe.Pointer) {
	// 4: (, widget QWidget *), (widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayout6addRowEP7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:116
// index:5
// void addRow(class QLayout *)
func (this *QFormLayout) AddRow_5(layout unsafe.Pointer) {
	// 5: (, layout QLayout *), (layout)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayout6addRowEP7QLayout", ffiqt.FFI_TYPE_VOID, this.GetCthis(), layout)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:118
// index:0
// void insertRow(int, class QWidget *, class QWidget *)
func (this *QFormLayout) InsertRow(row int, label unsafe.Pointer, field unsafe.Pointer) {
	// 0: (, row int, label QWidget *, field QWidget *), (&row, label, field)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayout9insertRowEiP7QWidgetS1_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, label, field)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:119
// index:1
// void insertRow(int, class QWidget *, class QLayout *)
func (this *QFormLayout) InsertRow_1(row int, label unsafe.Pointer, field unsafe.Pointer) {
	// 1: (, row int, label QWidget *, field QLayout *), (&row, label, field)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayout9insertRowEiP7QWidgetP7QLayout", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, label, field)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:120
// index:2
// void insertRow(int, const class QString &, class QWidget *)
func (this *QFormLayout) InsertRow_2(row int, labelText unsafe.Pointer, field unsafe.Pointer) {
	// 2: (, row int, labelText const QString &, field QWidget *), (&row, labelText, field)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayout9insertRowEiRK7QStringP7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, labelText, field)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:121
// index:3
// void insertRow(int, const class QString &, class QLayout *)
func (this *QFormLayout) InsertRow_3(row int, labelText unsafe.Pointer, field unsafe.Pointer) {
	// 3: (, row int, labelText const QString &, field QLayout *), (&row, labelText, field)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayout9insertRowEiRK7QStringP7QLayout", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, labelText, field)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:122
// index:4
// void insertRow(int, class QWidget *)
func (this *QFormLayout) InsertRow_4(row int, widget unsafe.Pointer) {
	// 4: (, row int, widget QWidget *), (&row, widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayout9insertRowEiP7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:123
// index:5
// void insertRow(int, class QLayout *)
func (this *QFormLayout) InsertRow_5(row int, layout unsafe.Pointer) {
	// 5: (, row int, layout QLayout *), (&row, layout)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayout9insertRowEiP7QLayout", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, layout)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:125
// index:0
// void removeRow(int)
func (this *QFormLayout) RemoveRow(row int) {
	// 0: (, row int), (&row)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayout9removeRowEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:126
// index:1
// void removeRow(class QWidget *)
func (this *QFormLayout) RemoveRow_1(widget unsafe.Pointer) {
	// 1: (, widget QWidget *), (widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayout9removeRowEP7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:127
// index:2
// void removeRow(class QLayout *)
func (this *QFormLayout) RemoveRow_2(layout unsafe.Pointer) {
	// 2: (, layout QLayout *), (layout)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayout9removeRowEP7QLayout", ffiqt.FFI_TYPE_VOID, this.GetCthis(), layout)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:129
// index:0
// QFormLayout::TakeRowResult takeRow(int)
func (this *QFormLayout) TakeRow(row int) {
	// 0: (, row int), (&row)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayout7takeRowEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:130
// index:1
// QFormLayout::TakeRowResult takeRow(class QWidget *)
func (this *QFormLayout) TakeRow_1(widget unsafe.Pointer) {
	// 1: (, widget QWidget *), (widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayout7takeRowEP7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:131
// index:2
// QFormLayout::TakeRowResult takeRow(class QLayout *)
func (this *QFormLayout) TakeRow_2(layout unsafe.Pointer) {
	// 2: (, layout QLayout *), (layout)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayout7takeRowEP7QLayout", ffiqt.FFI_TYPE_VOID, this.GetCthis(), layout)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:133
// index:0
// void setItem(int, enum QFormLayout::ItemRole, class QLayoutItem *)
func (this *QFormLayout) SetItem(row int, role int, item unsafe.Pointer) {
	// 0: (, row int, role QFormLayout::ItemRole, item QLayoutItem *), (&row, &role, item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayout7setItemEiNS_8ItemRoleEP11QLayoutItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, &role, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:134
// index:0
// void setWidget(int, enum QFormLayout::ItemRole, class QWidget *)
func (this *QFormLayout) SetWidget(row int, role int, widget unsafe.Pointer) {
	// 0: (, row int, role QFormLayout::ItemRole, widget QWidget *), (&row, &role, widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayout9setWidgetEiNS_8ItemRoleEP7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, &role, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:135
// index:0
// void setLayout(int, enum QFormLayout::ItemRole, class QLayout *)
func (this *QFormLayout) SetLayout(row int, role int, layout unsafe.Pointer) {
	// 0: (, row int, role QFormLayout::ItemRole, layout QLayout *), (&row, &role, layout)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayout9setLayoutEiNS_8ItemRoleEP7QLayout", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, &role, layout)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:137
// index:0
// QLayoutItem * itemAt(int, enum QFormLayout::ItemRole)
func (this *QFormLayout) ItemAt(row int, role int) {
	// 0: (, row int, role QFormLayout::ItemRole), (&row, &role)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFormLayout6itemAtEiNS_8ItemRoleE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:146
// index:1
// virtual
// QLayoutItem * itemAt(int)
func (this *QFormLayout) ItemAt_1(index int) {
	// 1: (, index int), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFormLayout6itemAtEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:138
// index:0
// void getItemPosition(int, int *, enum QFormLayout::ItemRole *)
func (this *QFormLayout) GetItemPosition(index int, rowPtr unsafe.Pointer, rolePtr unsafe.Pointer) {
	// 0: (, index int, rowPtr int *, rolePtr QFormLayout::ItemRole *), (&index, rowPtr, rolePtr)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFormLayout15getItemPositionEiPiPNS_8ItemRoleE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index, rowPtr, rolePtr)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:139
// index:0
// void getWidgetPosition(class QWidget *, int *, enum QFormLayout::ItemRole *)
func (this *QFormLayout) GetWidgetPosition(widget unsafe.Pointer, rowPtr unsafe.Pointer, rolePtr unsafe.Pointer) {
	// 0: (, widget QWidget *, rowPtr int *, rolePtr QFormLayout::ItemRole *), (widget, rowPtr, rolePtr)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFormLayout17getWidgetPositionEP7QWidgetPiPNS_8ItemRoleE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), widget, rowPtr, rolePtr)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:140
// index:0
// void getLayoutPosition(class QLayout *, int *, enum QFormLayout::ItemRole *)
func (this *QFormLayout) GetLayoutPosition(layout unsafe.Pointer, rowPtr unsafe.Pointer, rolePtr unsafe.Pointer) {
	// 0: (, layout QLayout *, rowPtr int *, rolePtr QFormLayout::ItemRole *), (layout, rowPtr, rolePtr)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFormLayout17getLayoutPositionEP7QLayoutPiPNS_8ItemRoleE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), layout, rowPtr, rolePtr)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:141
// index:0
// QWidget * labelForField(class QWidget *)
func (this *QFormLayout) LabelForField(field unsafe.Pointer) {
	// 0: (, field QWidget *), (field)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFormLayout13labelForFieldEP7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), field)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:142
// index:1
// QWidget * labelForField(class QLayout *)
func (this *QFormLayout) LabelForField_1(field unsafe.Pointer) {
	// 1: (, field QLayout *), (field)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFormLayout13labelForFieldEP7QLayout", ffiqt.FFI_TYPE_VOID, this.GetCthis(), field)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:145
// index:0
// virtual
// void addItem(class QLayoutItem *)
func (this *QFormLayout) AddItem(item unsafe.Pointer) {
	// 0: (, item QLayoutItem *), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayout7addItemEP11QLayoutItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:147
// index:0
// virtual
// QLayoutItem * takeAt(int)
func (this *QFormLayout) TakeAt(index int) {
	// 0: (, index int), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayout6takeAtEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:149
// index:0
// virtual
// void setGeometry(const class QRect &)
func (this *QFormLayout) SetGeometry(rect unsafe.Pointer) {
	// 0: (, rect const QRect &), (rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayout11setGeometryERK5QRect", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:150
// index:0
// virtual
// QSize minimumSize()
func (this *QFormLayout) MinimumSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFormLayout11minimumSizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:151
// index:0
// virtual
// QSize sizeHint()
func (this *QFormLayout) SizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFormLayout8sizeHintEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:152
// index:0
// virtual
// void invalidate()
func (this *QFormLayout) Invalidate() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFormLayout10invalidateEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:154
// index:0
// virtual
// bool hasHeightForWidth()
func (this *QFormLayout) HasHeightForWidth() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFormLayout17hasHeightForWidthEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:155
// index:0
// virtual
// int heightForWidth(int)
func (this *QFormLayout) HeightForWidth(width int) {
	// 0: (, width int), (&width)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFormLayout14heightForWidthEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:156
// index:0
// virtual
// Qt::Orientations expandingDirections()
func (this *QFormLayout) ExpandingDirections() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFormLayout19expandingDirectionsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:157
// index:0
// virtual
// int count()
func (this *QFormLayout) Count() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFormLayout5countEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qformlayout.h:159
// index:0
// int rowCount()
func (this *QFormLayout) RowCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFormLayout8rowCountEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
