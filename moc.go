package main

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "moc.h"
import "C"
import (
	"runtime"
	"strings"
	"unsafe"

	"github.com/therecipe/qt"
	std_core "github.com/therecipe/qt/core"
)

func cGoUnpackString(s C.struct_Moc_PackedString) string {
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_Moc_PackedString) []byte {
	if int(s.len) == -1 {
		gs := C.GoString(s.data)
		return *(*[]byte)(unsafe.Pointer(&gs))
	}
	return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len))
}
func unpackStringList(s string) []string {
	if len(s) == 0 {
		return make([]string, 0)
	}
	return strings.Split(s, "¡¦!")
}

type CtxObject_ITF interface {
	std_core.QObject_ITF
	CtxObject_PTR() *CtxObject
}

func (ptr *CtxObject) CtxObject_PTR() *CtxObject {
	return ptr
}

func (ptr *CtxObject) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *CtxObject) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromCtxObject(ptr CtxObject_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.CtxObject_PTR().Pointer()
	}
	return nil
}

func NewCtxObjectFromPointer(ptr unsafe.Pointer) (n *CtxObject) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(CtxObject)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *CtxObject:
			n = deduced

		case *std_core.QObject:
			n = &CtxObject{QObject: *deduced}

		default:
			n = new(CtxObject)
			n.SetPointer(ptr)
		}
	}
	return
}
func (this *CtxObject) Init() { this.init() }

//export callbackCtxObject40f1b6_Constructor
func callbackCtxObject40f1b6_Constructor(ptr unsafe.Pointer) {
	this := NewCtxObjectFromPointer(ptr)
	qt.Register(ptr, this)
	this.ConnectClicked(this.clicked)
	this.init()
}

//export callbackCtxObject40f1b6_Clicked
func callbackCtxObject40f1b6_Clicked(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "clicked"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *CtxObject) ConnectClicked(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "clicked") {
			C.CtxObject40f1b6_ConnectClicked(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "clicked"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "clicked", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clicked", unsafe.Pointer(&f))
		}
	}
}

func (ptr *CtxObject) DisconnectClicked() {
	if ptr.Pointer() != nil {
		C.CtxObject40f1b6_DisconnectClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "clicked")
	}
}

func (ptr *CtxObject) Clicked() {
	if ptr.Pointer() != nil {
		C.CtxObject40f1b6_Clicked(ptr.Pointer())
	}
}

//export callbackCtxObject40f1b6_SomeString
func callbackCtxObject40f1b6_SomeString(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "someString"); signal != nil {
		tempVal := (*(*func() string)(signal))()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewCtxObjectFromPointer(ptr).SomeStringDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *CtxObject) ConnectSomeString(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "someString"); signal != nil {
			f := func() string {
				(*(*func() string)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "someString", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "someString", unsafe.Pointer(&f))
		}
	}
}

func (ptr *CtxObject) DisconnectSomeString() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "someString")
	}
}

func (ptr *CtxObject) SomeString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.CtxObject40f1b6_SomeString(ptr.Pointer()))
	}
	return ""
}

func (ptr *CtxObject) SomeStringDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.CtxObject40f1b6_SomeStringDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackCtxObject40f1b6_SetSomeString
func callbackCtxObject40f1b6_SetSomeString(ptr unsafe.Pointer, someString C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setSomeString"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(someString))
	} else {
		NewCtxObjectFromPointer(ptr).SetSomeStringDefault(cGoUnpackString(someString))
	}
}

func (ptr *CtxObject) ConnectSetSomeString(f func(someString string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setSomeString"); signal != nil {
			f := func(someString string) {
				(*(*func(string))(signal))(someString)
				f(someString)
			}
			qt.ConnectSignal(ptr.Pointer(), "setSomeString", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setSomeString", unsafe.Pointer(&f))
		}
	}
}

func (ptr *CtxObject) DisconnectSetSomeString() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setSomeString")
	}
}

func (ptr *CtxObject) SetSomeString(someString string) {
	if ptr.Pointer() != nil {
		var someStringC *C.char
		if someString != "" {
			someStringC = C.CString(someString)
			defer C.free(unsafe.Pointer(someStringC))
		}
		C.CtxObject40f1b6_SetSomeString(ptr.Pointer(), C.struct_Moc_PackedString{data: someStringC, len: C.longlong(len(someString))})
	}
}

func (ptr *CtxObject) SetSomeStringDefault(someString string) {
	if ptr.Pointer() != nil {
		var someStringC *C.char
		if someString != "" {
			someStringC = C.CString(someString)
			defer C.free(unsafe.Pointer(someStringC))
		}
		C.CtxObject40f1b6_SetSomeStringDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: someStringC, len: C.longlong(len(someString))})
	}
}

//export callbackCtxObject40f1b6_SomeStringChanged
func callbackCtxObject40f1b6_SomeStringChanged(ptr unsafe.Pointer, someString C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "someStringChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(someString))
	}

}

func (ptr *CtxObject) ConnectSomeStringChanged(f func(someString string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "someStringChanged") {
			C.CtxObject40f1b6_ConnectSomeStringChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "someStringChanged"); signal != nil {
			f := func(someString string) {
				(*(*func(string))(signal))(someString)
				f(someString)
			}
			qt.ConnectSignal(ptr.Pointer(), "someStringChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "someStringChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *CtxObject) DisconnectSomeStringChanged() {
	if ptr.Pointer() != nil {
		C.CtxObject40f1b6_DisconnectSomeStringChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "someStringChanged")
	}
}

func (ptr *CtxObject) SomeStringChanged(someString string) {
	if ptr.Pointer() != nil {
		var someStringC *C.char
		if someString != "" {
			someStringC = C.CString(someString)
			defer C.free(unsafe.Pointer(someStringC))
		}
		C.CtxObject40f1b6_SomeStringChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: someStringC, len: C.longlong(len(someString))})
	}
}

func CtxObject_QRegisterMetaType() int {
	return int(int32(C.CtxObject40f1b6_CtxObject40f1b6_QRegisterMetaType()))
}

func (ptr *CtxObject) QRegisterMetaType() int {
	return int(int32(C.CtxObject40f1b6_CtxObject40f1b6_QRegisterMetaType()))
}

func CtxObject_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.CtxObject40f1b6_CtxObject40f1b6_QRegisterMetaType2(typeNameC)))
}

func (ptr *CtxObject) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.CtxObject40f1b6_CtxObject40f1b6_QRegisterMetaType2(typeNameC)))
}

func CtxObject_QmlRegisterType() int {
	return int(int32(C.CtxObject40f1b6_CtxObject40f1b6_QmlRegisterType()))
}

func (ptr *CtxObject) QmlRegisterType() int {
	return int(int32(C.CtxObject40f1b6_CtxObject40f1b6_QmlRegisterType()))
}

func CtxObject_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.CtxObject40f1b6_CtxObject40f1b6_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *CtxObject) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.CtxObject40f1b6_CtxObject40f1b6_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *CtxObject) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.CtxObject40f1b6___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *CtxObject) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.CtxObject40f1b6___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *CtxObject) __children_newList() unsafe.Pointer {
	return C.CtxObject40f1b6___children_newList(ptr.Pointer())
}

func (ptr *CtxObject) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.CtxObject40f1b6___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *CtxObject) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.CtxObject40f1b6___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *CtxObject) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.CtxObject40f1b6___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *CtxObject) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.CtxObject40f1b6___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *CtxObject) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.CtxObject40f1b6___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *CtxObject) __findChildren_newList() unsafe.Pointer {
	return C.CtxObject40f1b6___findChildren_newList(ptr.Pointer())
}

func (ptr *CtxObject) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.CtxObject40f1b6___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *CtxObject) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.CtxObject40f1b6___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *CtxObject) __findChildren_newList3() unsafe.Pointer {
	return C.CtxObject40f1b6___findChildren_newList3(ptr.Pointer())
}

func (ptr *CtxObject) __qFindChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.CtxObject40f1b6___qFindChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *CtxObject) __qFindChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.CtxObject40f1b6___qFindChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *CtxObject) __qFindChildren_newList2() unsafe.Pointer {
	return C.CtxObject40f1b6___qFindChildren_newList2(ptr.Pointer())
}

func NewCtxObject(parent std_core.QObject_ITF) *CtxObject {
	tmpValue := NewCtxObjectFromPointer(C.CtxObject40f1b6_NewCtxObject(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackCtxObject40f1b6_DestroyCtxObject
func callbackCtxObject40f1b6_DestroyCtxObject(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~CtxObject"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewCtxObjectFromPointer(ptr).DestroyCtxObjectDefault()
	}
}

func (ptr *CtxObject) ConnectDestroyCtxObject(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~CtxObject"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~CtxObject", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~CtxObject", unsafe.Pointer(&f))
		}
	}
}

func (ptr *CtxObject) DisconnectDestroyCtxObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~CtxObject")
	}
}

func (ptr *CtxObject) DestroyCtxObject() {
	if ptr.Pointer() != nil {
		C.CtxObject40f1b6_DestroyCtxObject(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *CtxObject) DestroyCtxObjectDefault() {
	if ptr.Pointer() != nil {
		C.CtxObject40f1b6_DestroyCtxObjectDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackCtxObject40f1b6_ChildEvent
func callbackCtxObject40f1b6_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*std_core.QChildEvent))(signal))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewCtxObjectFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *CtxObject) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CtxObject40f1b6_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackCtxObject40f1b6_ConnectNotify
func callbackCtxObject40f1b6_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewCtxObjectFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *CtxObject) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.CtxObject40f1b6_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackCtxObject40f1b6_CustomEvent
func callbackCtxObject40f1b6_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewCtxObjectFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *CtxObject) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CtxObject40f1b6_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackCtxObject40f1b6_DeleteLater
func callbackCtxObject40f1b6_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewCtxObjectFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *CtxObject) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.CtxObject40f1b6_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackCtxObject40f1b6_Destroyed
func callbackCtxObject40f1b6_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(obj))
	}
	qt.Unregister(ptr)

}

//export callbackCtxObject40f1b6_DisconnectNotify
func callbackCtxObject40f1b6_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewCtxObjectFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *CtxObject) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.CtxObject40f1b6_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackCtxObject40f1b6_Event
func callbackCtxObject40f1b6_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCtxObjectFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *CtxObject) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.CtxObject40f1b6_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackCtxObject40f1b6_EventFilter
func callbackCtxObject40f1b6_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCtxObjectFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *CtxObject) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.CtxObject40f1b6_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackCtxObject40f1b6_ObjectNameChanged
func callbackCtxObject40f1b6_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackCtxObject40f1b6_TimerEvent
func callbackCtxObject40f1b6_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*std_core.QTimerEvent))(signal))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewCtxObjectFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *CtxObject) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CtxObject40f1b6_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}
