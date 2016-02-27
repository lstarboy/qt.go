package qtcore
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
// src-file: /QtCore/qtemporarydir.h
// dst-file: /src/core/qtemporarydir.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
import "fmt"
import "reflect"
import "unsafe"
import "runtime"
import "qtrt"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QTemporaryDir::~QTemporaryDir();
extern void C_ZN13QTemporaryDirD2Ev(void* qthis); // 4
  // proto:  void QTemporaryDir::setAutoRemove(bool b);
extern void C_ZN13QTemporaryDir13setAutoRemoveEb(void* qthis, bool arg0); // 4
  // proto:  void QTemporaryDir::QTemporaryDir();
extern void* C_ZN13QTemporaryDirC2Ev(); // 3
  // proto:  void QTemporaryDir::QTemporaryDir(const QString & templateName);
extern void* C_ZN13QTemporaryDirC2ERK7QString(void* arg0); // 3
  // proto:  bool QTemporaryDir::isValid();
extern bool C_ZNK13QTemporaryDir7isValidEv(void* qthis); // 4
  // proto:  bool QTemporaryDir::remove();
extern bool C_ZN13QTemporaryDir6removeEv(void* qthis); // 4
  // proto:  bool QTemporaryDir::autoRemove();
extern bool C_ZNK13QTemporaryDir10autoRemoveEv(void* qthis); // 4
  // proto:  QString QTemporaryDir::path();
extern void* C_ZNK13QTemporaryDir4pathEv(void* qthis); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
  if false {reflect.TypeOf(runtime.Version)}
}

// class sizeof(QTemporaryDir)=1
type QTemporaryDir struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// ~QTemporaryDir()
func (this *QTemporaryDir) Free(args ...interface{}) () {
  // ~QTemporaryDir()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTemporaryDirD0Ev
    // invoke: void ~QTemporaryDir()
    if this != nil && this.Qclsinst != nil {
      C.C_ZN13QTemporaryDirD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QTemporaryDir", "~QTemporaryDir", args)
  }

  return
}

// setAutoRemove(_Bool)
func (this *QTemporaryDir) SetAutoRemove(args ...interface{}) () {
  // setAutoRemove(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTemporaryDir13setAutoRemoveEb
    // invoke: void setAutoRemove(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN13QTemporaryDir13setAutoRemoveEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTemporaryDir", "setAutoRemove", args)
  }

  return
}

// QTemporaryDir()
func GcfreeQTemporaryDir(this *QTemporaryDir) {
  qtrt.UniverseFree(this)
}
func NewQTemporaryDir(args ...interface{}) *QTemporaryDir {
  // QTemporaryDir()
  // QTemporaryDir(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTemporaryDirC1Ev
    // invoke: void QTemporaryDir()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN13QTemporaryDirC2Ev()
    this := &QTemporaryDir{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQTemporaryDir)
    return this
  case 1:
    // invoke: _ZN13QTemporaryDirC1ERK7QString
    // invoke: void QTemporaryDir(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN13QTemporaryDirC2ERK7QString(arg0)
    this := &QTemporaryDir{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQTemporaryDir)
    return this
  default:
    qtrt.ErrorResolve("QTemporaryDir", "QTemporaryDir", args)
  }

  return nil // QTemporaryDir{Qclsinst:qthis}
}

// isValid()
func (this *QTemporaryDir) IsValid(args ...interface{}) (ret interface{}) {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTemporaryDir7isValidEv
    // invoke: bool isValid()
    var ret0 = C.C_ZNK13QTemporaryDir7isValidEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTemporaryDir", "isValid", args)
  }

  return
}

// remove()
func (this *QTemporaryDir) Remove(args ...interface{}) (ret interface{}) {
  // remove()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QTemporaryDir6removeEv
    // invoke: bool remove()
    var ret0 = C.C_ZN13QTemporaryDir6removeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTemporaryDir", "remove", args)
  }

  return
}

// autoRemove()
func (this *QTemporaryDir) AutoRemove(args ...interface{}) (ret interface{}) {
  // autoRemove()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTemporaryDir10autoRemoveEv
    // invoke: bool autoRemove()
    var ret0 = C.C_ZNK13QTemporaryDir10autoRemoveEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTemporaryDir", "autoRemove", args)
  }

  return
}

// path()
func (this *QTemporaryDir) Path(args ...interface{}) (ret interface{}) {
  // path()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QTemporaryDir4pathEv
    // invoke: QString path()
    var ret0 = C.C_ZNK13QTemporaryDir4pathEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTemporaryDir", "path", args)
  }

  return
}

// <= body block end

