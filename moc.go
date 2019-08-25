package main

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "moc.h"
import "C"
import (
	"errors"
	"runtime"
	"strings"
	"unsafe"

	"github.com/therecipe/qt"
	std_core "github.com/therecipe/qt/core"
	std_gui "github.com/therecipe/qt/gui"
	std_quick "github.com/therecipe/qt/quick"
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

type ItemTemplate_ITF interface {
	std_quick.QQuickItem_ITF
	ItemTemplate_PTR() *ItemTemplate
}

func (ptr *ItemTemplate) ItemTemplate_PTR() *ItemTemplate {
	return ptr
}

func (ptr *ItemTemplate) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickItem_PTR().Pointer()
	}
	return nil
}

func (ptr *ItemTemplate) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QQuickItem_PTR().SetPointer(p)
	}
}

func PointerFromItemTemplate(ptr ItemTemplate_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.ItemTemplate_PTR().Pointer()
	}
	return nil
}

func NewItemTemplateFromPointer(ptr unsafe.Pointer) (n *ItemTemplate) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(ItemTemplate)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *ItemTemplate:
			n = deduced

		case *std_quick.QQuickItem:
			n = &ItemTemplate{QQuickItem: *deduced}

		default:
			n = new(ItemTemplate)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackItemTemplate40f1b6_Constructor
func callbackItemTemplate40f1b6_Constructor(ptr unsafe.Pointer) {
	this := NewItemTemplateFromPointer(ptr)
	qt.Register(ptr, this)
	this.ConnectSendBool(this.sendBool)
	this.ConnectSendInt(this.sendInt)
	this.ConnectSendFloat(this.sendFloat)
	this.ConnectSendDouble(this.sendDouble)
	this.ConnectSendString(this.sendString)
	this.ConnectSendError(this.sendError)
	this.ConnectSendVariantListMap(this.sendVariantListMap)
	this.ConnectSendItemTemplate(this.sendItemTemplate)
	this.ConnectSendItem(this.sendItem)
	this.ConnectSendObject(this.sendObject)
}

//export callbackItemTemplate40f1b6_SendBool
func callbackItemTemplate40f1b6_SendBool(ptr unsafe.Pointer, v0 C.char, v1 C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "sendBool"); signal != nil {
		(*(*func(bool, []bool))(signal))(int8(v0) != 0, func(l C.struct_Moc_PackedList) []bool {
			out := make([]bool, int(l.len))
			tmpList := NewItemTemplateFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__sendBool_v1_atList(i)
			}
			return out
		}(v1))
	}

}

func (ptr *ItemTemplate) ConnectSendBool(f func(v0 bool, v1 []bool)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "sendBool") {
			C.ItemTemplate40f1b6_ConnectSendBool(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "sendBool"); signal != nil {
			f := func(v0 bool, v1 []bool) {
				(*(*func(bool, []bool))(signal))(v0, v1)
				f(v0, v1)
			}
			qt.ConnectSignal(ptr.Pointer(), "sendBool", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sendBool", unsafe.Pointer(&f))
		}
	}
}

func (ptr *ItemTemplate) DisconnectSendBool() {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6_DisconnectSendBool(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "sendBool")
	}
}

func (ptr *ItemTemplate) SendBool(v0 bool, v1 []bool) {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6_SendBool(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(v0))), func() unsafe.Pointer {
			tmpList := NewItemTemplateFromPointer(NewItemTemplateFromPointer(nil).__sendBool_v1_newList())
			for _, v := range v1 {
				tmpList.__sendBool_v1_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackItemTemplate40f1b6_SendInt
func callbackItemTemplate40f1b6_SendInt(ptr unsafe.Pointer, v0 C.int, v1 C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "sendInt"); signal != nil {
		(*(*func(int, []int))(signal))(int(int32(v0)), func(l C.struct_Moc_PackedList) []int {
			out := make([]int, int(l.len))
			tmpList := NewItemTemplateFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__sendInt_v1_atList(i)
			}
			return out
		}(v1))
	}

}

func (ptr *ItemTemplate) ConnectSendInt(f func(v0 int, v1 []int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "sendInt") {
			C.ItemTemplate40f1b6_ConnectSendInt(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "sendInt"); signal != nil {
			f := func(v0 int, v1 []int) {
				(*(*func(int, []int))(signal))(v0, v1)
				f(v0, v1)
			}
			qt.ConnectSignal(ptr.Pointer(), "sendInt", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sendInt", unsafe.Pointer(&f))
		}
	}
}

func (ptr *ItemTemplate) DisconnectSendInt() {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6_DisconnectSendInt(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "sendInt")
	}
}

func (ptr *ItemTemplate) SendInt(v0 int, v1 []int) {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6_SendInt(ptr.Pointer(), C.int(int32(v0)), func() unsafe.Pointer {
			tmpList := NewItemTemplateFromPointer(NewItemTemplateFromPointer(nil).__sendInt_v1_newList())
			for _, v := range v1 {
				tmpList.__sendInt_v1_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackItemTemplate40f1b6_SendFloat
func callbackItemTemplate40f1b6_SendFloat(ptr unsafe.Pointer, v0 C.float) {
	if signal := qt.GetSignal(ptr, "sendFloat"); signal != nil {
		(*(*func(float32))(signal))(float32(v0))
	}

}

func (ptr *ItemTemplate) ConnectSendFloat(f func(v0 float32)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "sendFloat") {
			C.ItemTemplate40f1b6_ConnectSendFloat(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "sendFloat"); signal != nil {
			f := func(v0 float32) {
				(*(*func(float32))(signal))(v0)
				f(v0)
			}
			qt.ConnectSignal(ptr.Pointer(), "sendFloat", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sendFloat", unsafe.Pointer(&f))
		}
	}
}

func (ptr *ItemTemplate) DisconnectSendFloat() {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6_DisconnectSendFloat(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "sendFloat")
	}
}

func (ptr *ItemTemplate) SendFloat(v0 float32) {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6_SendFloat(ptr.Pointer(), C.float(v0))
	}
}

//export callbackItemTemplate40f1b6_SendDouble
func callbackItemTemplate40f1b6_SendDouble(ptr unsafe.Pointer, v0 C.double, v1 C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "sendDouble"); signal != nil {
		(*(*func(float64, []float64))(signal))(float64(v0), func(l C.struct_Moc_PackedList) []float64 {
			out := make([]float64, int(l.len))
			tmpList := NewItemTemplateFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__sendDouble_v1_atList(i)
			}
			return out
		}(v1))
	}

}

func (ptr *ItemTemplate) ConnectSendDouble(f func(v0 float64, v1 []float64)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "sendDouble") {
			C.ItemTemplate40f1b6_ConnectSendDouble(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "sendDouble"); signal != nil {
			f := func(v0 float64, v1 []float64) {
				(*(*func(float64, []float64))(signal))(v0, v1)
				f(v0, v1)
			}
			qt.ConnectSignal(ptr.Pointer(), "sendDouble", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sendDouble", unsafe.Pointer(&f))
		}
	}
}

func (ptr *ItemTemplate) DisconnectSendDouble() {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6_DisconnectSendDouble(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "sendDouble")
	}
}

func (ptr *ItemTemplate) SendDouble(v0 float64, v1 []float64) {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6_SendDouble(ptr.Pointer(), C.double(v0), func() unsafe.Pointer {
			tmpList := NewItemTemplateFromPointer(NewItemTemplateFromPointer(nil).__sendDouble_v1_newList())
			for _, v := range v1 {
				tmpList.__sendDouble_v1_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackItemTemplate40f1b6_SendString
func callbackItemTemplate40f1b6_SendString(ptr unsafe.Pointer, v0 C.struct_Moc_PackedString, v1 C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "sendString"); signal != nil {
		(*(*func(string, []string))(signal))(cGoUnpackString(v0), unpackStringList(cGoUnpackString(v1)))
	}

}

func (ptr *ItemTemplate) ConnectSendString(f func(v0 string, v1 []string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "sendString") {
			C.ItemTemplate40f1b6_ConnectSendString(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "sendString"); signal != nil {
			f := func(v0 string, v1 []string) {
				(*(*func(string, []string))(signal))(v0, v1)
				f(v0, v1)
			}
			qt.ConnectSignal(ptr.Pointer(), "sendString", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sendString", unsafe.Pointer(&f))
		}
	}
}

func (ptr *ItemTemplate) DisconnectSendString() {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6_DisconnectSendString(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "sendString")
	}
}

func (ptr *ItemTemplate) SendString(v0 string, v1 []string) {
	if ptr.Pointer() != nil {
		var v0C *C.char
		if v0 != "" {
			v0C = C.CString(v0)
			defer C.free(unsafe.Pointer(v0C))
		}
		v1C := C.CString(strings.Join(v1, "¡¦!"))
		defer C.free(unsafe.Pointer(v1C))
		C.ItemTemplate40f1b6_SendString(ptr.Pointer(), C.struct_Moc_PackedString{data: v0C, len: C.longlong(len(v0))}, C.struct_Moc_PackedString{data: v1C, len: C.longlong(len(strings.Join(v1, "¡¦!")))})
	}
}

//export callbackItemTemplate40f1b6_SendError
func callbackItemTemplate40f1b6_SendError(ptr unsafe.Pointer, v0 C.struct_Moc_PackedString, v1 C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "sendError"); signal != nil {
		(*(*func(error, []error))(signal))(errors.New(cGoUnpackString(v0)), func(l C.struct_Moc_PackedList) []error {
			out := make([]error, int(l.len))
			tmpList := NewItemTemplateFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__sendError_v1_atList(i)
			}
			return out
		}(v1))
	}

}

func (ptr *ItemTemplate) ConnectSendError(f func(v0 error, v1 []error)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "sendError") {
			C.ItemTemplate40f1b6_ConnectSendError(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "sendError"); signal != nil {
			f := func(v0 error, v1 []error) {
				(*(*func(error, []error))(signal))(v0, v1)
				f(v0, v1)
			}
			qt.ConnectSignal(ptr.Pointer(), "sendError", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sendError", unsafe.Pointer(&f))
		}
	}
}

func (ptr *ItemTemplate) DisconnectSendError() {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6_DisconnectSendError(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "sendError")
	}
}

func (ptr *ItemTemplate) SendError(v0 error, v1 []error) {
	if ptr.Pointer() != nil {
		v0C := C.CString(func() string {
			tmp := v0
			if tmp != nil {
				return tmp.Error()
			}
			return ""
		}())
		defer C.free(unsafe.Pointer(v0C))
		C.ItemTemplate40f1b6_SendError(ptr.Pointer(), C.struct_Moc_PackedString{data: v0C, len: C.longlong(-1)}, func() unsafe.Pointer {
			tmpList := NewItemTemplateFromPointer(NewItemTemplateFromPointer(nil).__sendError_v1_newList())
			for _, v := range v1 {
				tmpList.__sendError_v1_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackItemTemplate40f1b6_SendVariantListMap
func callbackItemTemplate40f1b6_SendVariantListMap(ptr unsafe.Pointer, v0 unsafe.Pointer, v1 C.struct_Moc_PackedList, v2 C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "sendVariantListMap"); signal != nil {
		(*(*func(*std_core.QVariant, []*std_core.QVariant, map[string]*std_core.QVariant))(signal))(std_core.NewQVariantFromPointer(v0), func(l C.struct_Moc_PackedList) []*std_core.QVariant {
			out := make([]*std_core.QVariant, int(l.len))
			tmpList := NewItemTemplateFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__sendVariantListMap_v1_atList(i)
			}
			return out
		}(v1), func(l C.struct_Moc_PackedList) map[string]*std_core.QVariant {
			out := make(map[string]*std_core.QVariant, int(l.len))
			tmpList := NewItemTemplateFromPointer(l.data)
			for i, v := range tmpList.__sendVariantListMap_v2_keyList() {
				out[v] = tmpList.__sendVariantListMap_v2_atList(v, i)
			}
			return out
		}(v2))
	}

}

func (ptr *ItemTemplate) ConnectSendVariantListMap(f func(v0 *std_core.QVariant, v1 []*std_core.QVariant, v2 map[string]*std_core.QVariant)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "sendVariantListMap") {
			C.ItemTemplate40f1b6_ConnectSendVariantListMap(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "sendVariantListMap"); signal != nil {
			f := func(v0 *std_core.QVariant, v1 []*std_core.QVariant, v2 map[string]*std_core.QVariant) {
				(*(*func(*std_core.QVariant, []*std_core.QVariant, map[string]*std_core.QVariant))(signal))(v0, v1, v2)
				f(v0, v1, v2)
			}
			qt.ConnectSignal(ptr.Pointer(), "sendVariantListMap", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sendVariantListMap", unsafe.Pointer(&f))
		}
	}
}

func (ptr *ItemTemplate) DisconnectSendVariantListMap() {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6_DisconnectSendVariantListMap(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "sendVariantListMap")
	}
}

func (ptr *ItemTemplate) SendVariantListMap(v0 std_core.QVariant_ITF, v1 []*std_core.QVariant, v2 map[string]*std_core.QVariant) {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6_SendVariantListMap(ptr.Pointer(), std_core.PointerFromQVariant(v0), func() unsafe.Pointer {
			tmpList := NewItemTemplateFromPointer(NewItemTemplateFromPointer(nil).__sendVariantListMap_v1_newList())
			for _, v := range v1 {
				tmpList.__sendVariantListMap_v1_setList(v)
			}
			return tmpList.Pointer()
		}(), func() unsafe.Pointer {
			tmpList := NewItemTemplateFromPointer(NewItemTemplateFromPointer(nil).__sendVariantListMap_v2_newList())
			for k, v := range v2 {
				tmpList.__sendVariantListMap_v2_setList(k, std_core.NewQVariant1(v))
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackItemTemplate40f1b6_SendItemTemplate
func callbackItemTemplate40f1b6_SendItemTemplate(ptr unsafe.Pointer, v0 unsafe.Pointer, v1 C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "sendItemTemplate"); signal != nil {
		(*(*func(*ItemTemplate, []*ItemTemplate))(signal))(NewItemTemplateFromPointer(v0), func(l C.struct_Moc_PackedList) []*ItemTemplate {
			out := make([]*ItemTemplate, int(l.len))
			tmpList := NewItemTemplateFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__sendItemTemplate_v1_atList(i)
			}
			return out
		}(v1))
	}

}

func (ptr *ItemTemplate) ConnectSendItemTemplate(f func(v0 *ItemTemplate, v1 []*ItemTemplate)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "sendItemTemplate") {
			C.ItemTemplate40f1b6_ConnectSendItemTemplate(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "sendItemTemplate"); signal != nil {
			f := func(v0 *ItemTemplate, v1 []*ItemTemplate) {
				(*(*func(*ItemTemplate, []*ItemTemplate))(signal))(v0, v1)
				f(v0, v1)
			}
			qt.ConnectSignal(ptr.Pointer(), "sendItemTemplate", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sendItemTemplate", unsafe.Pointer(&f))
		}
	}
}

func (ptr *ItemTemplate) DisconnectSendItemTemplate() {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6_DisconnectSendItemTemplate(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "sendItemTemplate")
	}
}

func (ptr *ItemTemplate) SendItemTemplate(v0 ItemTemplate_ITF, v1 []*ItemTemplate) {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6_SendItemTemplate(ptr.Pointer(), PointerFromItemTemplate(v0), func() unsafe.Pointer {
			tmpList := NewItemTemplateFromPointer(NewItemTemplateFromPointer(nil).__sendItemTemplate_v1_newList())
			for _, v := range v1 {
				tmpList.__sendItemTemplate_v1_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackItemTemplate40f1b6_SendItem
func callbackItemTemplate40f1b6_SendItem(ptr unsafe.Pointer, v0 unsafe.Pointer, v1 C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "sendItem"); signal != nil {
		(*(*func(*std_quick.QQuickItem, []*std_quick.QQuickItem))(signal))(std_quick.NewQQuickItemFromPointer(v0), func(l C.struct_Moc_PackedList) []*std_quick.QQuickItem {
			out := make([]*std_quick.QQuickItem, int(l.len))
			tmpList := NewItemTemplateFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__sendItem_v1_atList(i)
			}
			return out
		}(v1))
	}

}

func (ptr *ItemTemplate) ConnectSendItem(f func(v0 *std_quick.QQuickItem, v1 []*std_quick.QQuickItem)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "sendItem") {
			C.ItemTemplate40f1b6_ConnectSendItem(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "sendItem"); signal != nil {
			f := func(v0 *std_quick.QQuickItem, v1 []*std_quick.QQuickItem) {
				(*(*func(*std_quick.QQuickItem, []*std_quick.QQuickItem))(signal))(v0, v1)
				f(v0, v1)
			}
			qt.ConnectSignal(ptr.Pointer(), "sendItem", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sendItem", unsafe.Pointer(&f))
		}
	}
}

func (ptr *ItemTemplate) DisconnectSendItem() {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6_DisconnectSendItem(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "sendItem")
	}
}

func (ptr *ItemTemplate) SendItem(v0 std_quick.QQuickItem_ITF, v1 []*std_quick.QQuickItem) {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6_SendItem(ptr.Pointer(), std_quick.PointerFromQQuickItem(v0), func() unsafe.Pointer {
			tmpList := NewItemTemplateFromPointer(NewItemTemplateFromPointer(nil).__sendItem_v1_newList())
			for _, v := range v1 {
				tmpList.__sendItem_v1_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackItemTemplate40f1b6_SendObject
func callbackItemTemplate40f1b6_SendObject(ptr unsafe.Pointer, v0 unsafe.Pointer, v1 C.struct_Moc_PackedList) {
	if signal := qt.GetSignal(ptr, "sendObject"); signal != nil {
		(*(*func(*std_core.QObject, []*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(v0), func(l C.struct_Moc_PackedList) []*std_core.QObject {
			out := make([]*std_core.QObject, int(l.len))
			tmpList := NewItemTemplateFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__sendObject_v1_atList(i)
			}
			return out
		}(v1))
	}

}

func (ptr *ItemTemplate) ConnectSendObject(f func(v0 *std_core.QObject, v1 []*std_core.QObject)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "sendObject") {
			C.ItemTemplate40f1b6_ConnectSendObject(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "sendObject"); signal != nil {
			f := func(v0 *std_core.QObject, v1 []*std_core.QObject) {
				(*(*func(*std_core.QObject, []*std_core.QObject))(signal))(v0, v1)
				f(v0, v1)
			}
			qt.ConnectSignal(ptr.Pointer(), "sendObject", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sendObject", unsafe.Pointer(&f))
		}
	}
}

func (ptr *ItemTemplate) DisconnectSendObject() {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6_DisconnectSendObject(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "sendObject")
	}
}

func (ptr *ItemTemplate) SendObject(v0 std_core.QObject_ITF, v1 []*std_core.QObject) {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6_SendObject(ptr.Pointer(), std_core.PointerFromQObject(v0), func() unsafe.Pointer {
			tmpList := NewItemTemplateFromPointer(NewItemTemplateFromPointer(nil).__sendObject_v1_newList())
			for _, v := range v1 {
				tmpList.__sendObject_v1_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackItemTemplate40f1b6_SomeString
func callbackItemTemplate40f1b6_SomeString(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "someString"); signal != nil {
		tempVal := (*(*func() string)(signal))()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewItemTemplateFromPointer(ptr).SomeStringDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *ItemTemplate) ConnectSomeString(f func() string) {
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

func (ptr *ItemTemplate) DisconnectSomeString() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "someString")
	}
}

func (ptr *ItemTemplate) SomeString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ItemTemplate40f1b6_SomeString(ptr.Pointer()))
	}
	return ""
}

func (ptr *ItemTemplate) SomeStringDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ItemTemplate40f1b6_SomeStringDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackItemTemplate40f1b6_SetSomeString
func callbackItemTemplate40f1b6_SetSomeString(ptr unsafe.Pointer, someString C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setSomeString"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(someString))
	} else {
		NewItemTemplateFromPointer(ptr).SetSomeStringDefault(cGoUnpackString(someString))
	}
}

func (ptr *ItemTemplate) ConnectSetSomeString(f func(someString string)) {
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

func (ptr *ItemTemplate) DisconnectSetSomeString() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setSomeString")
	}
}

func (ptr *ItemTemplate) SetSomeString(someString string) {
	if ptr.Pointer() != nil {
		var someStringC *C.char
		if someString != "" {
			someStringC = C.CString(someString)
			defer C.free(unsafe.Pointer(someStringC))
		}
		C.ItemTemplate40f1b6_SetSomeString(ptr.Pointer(), C.struct_Moc_PackedString{data: someStringC, len: C.longlong(len(someString))})
	}
}

func (ptr *ItemTemplate) SetSomeStringDefault(someString string) {
	if ptr.Pointer() != nil {
		var someStringC *C.char
		if someString != "" {
			someStringC = C.CString(someString)
			defer C.free(unsafe.Pointer(someStringC))
		}
		C.ItemTemplate40f1b6_SetSomeStringDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: someStringC, len: C.longlong(len(someString))})
	}
}

//export callbackItemTemplate40f1b6_SomeStringChanged
func callbackItemTemplate40f1b6_SomeStringChanged(ptr unsafe.Pointer, someString C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "someStringChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(someString))
	}

}

func (ptr *ItemTemplate) ConnectSomeStringChanged(f func(someString string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "someStringChanged") {
			C.ItemTemplate40f1b6_ConnectSomeStringChanged(ptr.Pointer())
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

func (ptr *ItemTemplate) DisconnectSomeStringChanged() {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6_DisconnectSomeStringChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "someStringChanged")
	}
}

func (ptr *ItemTemplate) SomeStringChanged(someString string) {
	if ptr.Pointer() != nil {
		var someStringC *C.char
		if someString != "" {
			someStringC = C.CString(someString)
			defer C.free(unsafe.Pointer(someStringC))
		}
		C.ItemTemplate40f1b6_SomeStringChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: someStringC, len: C.longlong(len(someString))})
	}
}

func ItemTemplate_QRegisterMetaType() int {
	return int(int32(C.ItemTemplate40f1b6_ItemTemplate40f1b6_QRegisterMetaType()))
}

func (ptr *ItemTemplate) QRegisterMetaType() int {
	return int(int32(C.ItemTemplate40f1b6_ItemTemplate40f1b6_QRegisterMetaType()))
}

func ItemTemplate_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.ItemTemplate40f1b6_ItemTemplate40f1b6_QRegisterMetaType2(typeNameC)))
}

func (ptr *ItemTemplate) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.ItemTemplate40f1b6_ItemTemplate40f1b6_QRegisterMetaType2(typeNameC)))
}

func ItemTemplate_QmlRegisterType() int {
	return int(int32(C.ItemTemplate40f1b6_ItemTemplate40f1b6_QmlRegisterType()))
}

func (ptr *ItemTemplate) QmlRegisterType() int {
	return int(int32(C.ItemTemplate40f1b6_ItemTemplate40f1b6_QmlRegisterType()))
}

func ItemTemplate_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.ItemTemplate40f1b6_ItemTemplate40f1b6_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *ItemTemplate) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.ItemTemplate40f1b6_ItemTemplate40f1b6_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *ItemTemplate) __childItems_atList(i int) *std_quick.QQuickItem {
	if ptr.Pointer() != nil {
		tmpValue := std_quick.NewQQuickItemFromPointer(C.ItemTemplate40f1b6___childItems_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ItemTemplate) __childItems_setList(i std_quick.QQuickItem_ITF) {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6___childItems_setList(ptr.Pointer(), std_quick.PointerFromQQuickItem(i))
	}
}

func (ptr *ItemTemplate) __childItems_newList() unsafe.Pointer {
	return C.ItemTemplate40f1b6___childItems_newList(ptr.Pointer())
}

func (ptr *ItemTemplate) __grabTouchPoints_ids_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.ItemTemplate40f1b6___grabTouchPoints_ids_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *ItemTemplate) __grabTouchPoints_ids_setList(i int) {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6___grabTouchPoints_ids_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *ItemTemplate) __grabTouchPoints_ids_newList() unsafe.Pointer {
	return C.ItemTemplate40f1b6___grabTouchPoints_ids_newList(ptr.Pointer())
}

func (ptr *ItemTemplate) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.ItemTemplate40f1b6___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ItemTemplate) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ItemTemplate) __children_newList() unsafe.Pointer {
	return C.ItemTemplate40f1b6___children_newList(ptr.Pointer())
}

func (ptr *ItemTemplate) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.ItemTemplate40f1b6___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *ItemTemplate) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *ItemTemplate) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.ItemTemplate40f1b6___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *ItemTemplate) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.ItemTemplate40f1b6___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ItemTemplate) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ItemTemplate) __findChildren_newList() unsafe.Pointer {
	return C.ItemTemplate40f1b6___findChildren_newList(ptr.Pointer())
}

func (ptr *ItemTemplate) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.ItemTemplate40f1b6___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ItemTemplate) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ItemTemplate) __findChildren_newList3() unsafe.Pointer {
	return C.ItemTemplate40f1b6___findChildren_newList3(ptr.Pointer())
}

func (ptr *ItemTemplate) __qFindChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.ItemTemplate40f1b6___qFindChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ItemTemplate) __qFindChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6___qFindChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ItemTemplate) __qFindChildren_newList2() unsafe.Pointer {
	return C.ItemTemplate40f1b6___qFindChildren_newList2(ptr.Pointer())
}

func (ptr *ItemTemplate) __sendBool_v1_atList(i int) bool {
	if ptr.Pointer() != nil {
		return int8(C.ItemTemplate40f1b6___sendBool_v1_atList(ptr.Pointer(), C.int(int32(i)))) != 0
	}
	return false
}

func (ptr *ItemTemplate) __sendBool_v1_setList(i bool) {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6___sendBool_v1_setList(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(i))))
	}
}

func (ptr *ItemTemplate) __sendBool_v1_newList() unsafe.Pointer {
	return C.ItemTemplate40f1b6___sendBool_v1_newList(ptr.Pointer())
}

func (ptr *ItemTemplate) __sendInt_v1_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.ItemTemplate40f1b6___sendInt_v1_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *ItemTemplate) __sendInt_v1_setList(i int) {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6___sendInt_v1_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *ItemTemplate) __sendInt_v1_newList() unsafe.Pointer {
	return C.ItemTemplate40f1b6___sendInt_v1_newList(ptr.Pointer())
}

func (ptr *ItemTemplate) __sendDouble_v1_atList(i int) float64 {
	if ptr.Pointer() != nil {
		return float64(C.ItemTemplate40f1b6___sendDouble_v1_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return 0
}

func (ptr *ItemTemplate) __sendDouble_v1_setList(i float64) {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6___sendDouble_v1_setList(ptr.Pointer(), C.double(i))
	}
}

func (ptr *ItemTemplate) __sendDouble_v1_newList() unsafe.Pointer {
	return C.ItemTemplate40f1b6___sendDouble_v1_newList(ptr.Pointer())
}

func (ptr *ItemTemplate) __sendError_v1_atList(i int) error {
	if ptr.Pointer() != nil {
		return errors.New(cGoUnpackString(C.ItemTemplate40f1b6___sendError_v1_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return errors.New("")
}

func (ptr *ItemTemplate) __sendError_v1_setList(i error) {
	if ptr.Pointer() != nil {
		iC := C.CString(func() string {
			tmp := i
			if tmp != nil {
				return tmp.Error()
			}
			return ""
		}())
		defer C.free(unsafe.Pointer(iC))
		C.ItemTemplate40f1b6___sendError_v1_setList(ptr.Pointer(), C.struct_Moc_PackedString{data: iC, len: C.longlong(-1)})
	}
}

func (ptr *ItemTemplate) __sendError_v1_newList() unsafe.Pointer {
	return C.ItemTemplate40f1b6___sendError_v1_newList(ptr.Pointer())
}

func (ptr *ItemTemplate) __sendVariantListMap_v1_atList(i int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.ItemTemplate40f1b6___sendVariantListMap_v1_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *ItemTemplate) __sendVariantListMap_v1_setList(i std_core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6___sendVariantListMap_v1_setList(ptr.Pointer(), std_core.PointerFromQVariant(i))
	}
}

func (ptr *ItemTemplate) __sendVariantListMap_v1_newList() unsafe.Pointer {
	return C.ItemTemplate40f1b6___sendVariantListMap_v1_newList(ptr.Pointer())
}

func (ptr *ItemTemplate) __sendVariantListMap_v2_atList(v string, i int) *std_core.QVariant {
	if ptr.Pointer() != nil {
		var vC *C.char
		if v != "" {
			vC = C.CString(v)
			defer C.free(unsafe.Pointer(vC))
		}
		tmpValue := std_core.NewQVariantFromPointer(C.ItemTemplate40f1b6___sendVariantListMap_v2_atList(ptr.Pointer(), C.struct_Moc_PackedString{data: vC, len: C.longlong(len(v))}, C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *ItemTemplate) __sendVariantListMap_v2_setList(key string, i std_core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		C.ItemTemplate40f1b6___sendVariantListMap_v2_setList(ptr.Pointer(), C.struct_Moc_PackedString{data: keyC, len: C.longlong(len(key))}, std_core.PointerFromQVariant(i))
	}
}

func (ptr *ItemTemplate) __sendVariantListMap_v2_newList() unsafe.Pointer {
	return C.ItemTemplate40f1b6___sendVariantListMap_v2_newList(ptr.Pointer())
}

func (ptr *ItemTemplate) __sendVariantListMap_v2_keyList() []string {
	if ptr.Pointer() != nil {
		return func(l C.struct_Moc_PackedList) []string {
			out := make([]string, int(l.len))
			tmpList := NewItemTemplateFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____sendVariantListMap_v2_keyList_atList(i)
			}
			return out
		}(C.ItemTemplate40f1b6___sendVariantListMap_v2_keyList(ptr.Pointer()))
	}
	return make([]string, 0)
}

func (ptr *ItemTemplate) __sendItemTemplate_v1_atList(i int) *ItemTemplate {
	if ptr.Pointer() != nil {
		tmpValue := NewItemTemplateFromPointer(C.ItemTemplate40f1b6___sendItemTemplate_v1_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ItemTemplate) __sendItemTemplate_v1_setList(i ItemTemplate_ITF) {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6___sendItemTemplate_v1_setList(ptr.Pointer(), PointerFromItemTemplate(i))
	}
}

func (ptr *ItemTemplate) __sendItemTemplate_v1_newList() unsafe.Pointer {
	return C.ItemTemplate40f1b6___sendItemTemplate_v1_newList(ptr.Pointer())
}

func (ptr *ItemTemplate) __sendItem_v1_atList(i int) *std_quick.QQuickItem {
	if ptr.Pointer() != nil {
		tmpValue := std_quick.NewQQuickItemFromPointer(C.ItemTemplate40f1b6___sendItem_v1_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ItemTemplate) __sendItem_v1_setList(i std_quick.QQuickItem_ITF) {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6___sendItem_v1_setList(ptr.Pointer(), std_quick.PointerFromQQuickItem(i))
	}
}

func (ptr *ItemTemplate) __sendItem_v1_newList() unsafe.Pointer {
	return C.ItemTemplate40f1b6___sendItem_v1_newList(ptr.Pointer())
}

func (ptr *ItemTemplate) __sendObject_v1_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.ItemTemplate40f1b6___sendObject_v1_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *ItemTemplate) __sendObject_v1_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6___sendObject_v1_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *ItemTemplate) __sendObject_v1_newList() unsafe.Pointer {
	return C.ItemTemplate40f1b6___sendObject_v1_newList(ptr.Pointer())
}

func (ptr *ItemTemplate) ____sendVariantListMap_v2_keyList_atList(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.ItemTemplate40f1b6_____sendVariantListMap_v2_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *ItemTemplate) ____sendVariantListMap_v2_keyList_setList(i string) {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		C.ItemTemplate40f1b6_____sendVariantListMap_v2_keyList_setList(ptr.Pointer(), C.struct_Moc_PackedString{data: iC, len: C.longlong(len(i))})
	}
}

func (ptr *ItemTemplate) ____sendVariantListMap_v2_keyList_newList() unsafe.Pointer {
	return C.ItemTemplate40f1b6_____sendVariantListMap_v2_keyList_newList(ptr.Pointer())
}

func NewItemTemplate(parent std_quick.QQuickItem_ITF) *ItemTemplate {
	tmpValue := NewItemTemplateFromPointer(C.ItemTemplate40f1b6_NewItemTemplate(std_quick.PointerFromQQuickItem(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackItemTemplate40f1b6_DestroyItemTemplate
func callbackItemTemplate40f1b6_DestroyItemTemplate(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~ItemTemplate"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewItemTemplateFromPointer(ptr).DestroyItemTemplateDefault()
	}
}

func (ptr *ItemTemplate) ConnectDestroyItemTemplate(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~ItemTemplate"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~ItemTemplate", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~ItemTemplate", unsafe.Pointer(&f))
		}
	}
}

func (ptr *ItemTemplate) DisconnectDestroyItemTemplate() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~ItemTemplate")
	}
}

func (ptr *ItemTemplate) DestroyItemTemplate() {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6_DestroyItemTemplate(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *ItemTemplate) DestroyItemTemplateDefault() {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6_DestroyItemTemplateDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackItemTemplate40f1b6_ActiveFocusChanged
func callbackItemTemplate40f1b6_ActiveFocusChanged(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "activeFocusChanged"); signal != nil {
		(*(*func(bool))(signal))(int8(vbo) != 0)
	}

}

//export callbackItemTemplate40f1b6_ActiveFocusOnTabChanged
func callbackItemTemplate40f1b6_ActiveFocusOnTabChanged(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "activeFocusOnTabChanged"); signal != nil {
		(*(*func(bool))(signal))(int8(vbo) != 0)
	}

}

//export callbackItemTemplate40f1b6_AntialiasingChanged
func callbackItemTemplate40f1b6_AntialiasingChanged(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "antialiasingChanged"); signal != nil {
		(*(*func(bool))(signal))(int8(vbo) != 0)
	}

}

//export callbackItemTemplate40f1b6_BaselineOffsetChanged
func callbackItemTemplate40f1b6_BaselineOffsetChanged(ptr unsafe.Pointer, vqr C.double) {
	if signal := qt.GetSignal(ptr, "baselineOffsetChanged"); signal != nil {
		(*(*func(float64))(signal))(float64(vqr))
	}

}

//export callbackItemTemplate40f1b6_ChildMouseEventFilter
func callbackItemTemplate40f1b6_ChildMouseEventFilter(ptr unsafe.Pointer, item unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "childMouseEventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_quick.QQuickItem, *std_core.QEvent) bool)(signal))(std_quick.NewQQuickItemFromPointer(item), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewItemTemplateFromPointer(ptr).ChildMouseEventFilterDefault(std_quick.NewQQuickItemFromPointer(item), std_core.NewQEventFromPointer(event)))))
}

func (ptr *ItemTemplate) ChildMouseEventFilterDefault(item std_quick.QQuickItem_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.ItemTemplate40f1b6_ChildMouseEventFilterDefault(ptr.Pointer(), std_quick.PointerFromQQuickItem(item), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackItemTemplate40f1b6_ChildrenRectChanged
func callbackItemTemplate40f1b6_ChildrenRectChanged(ptr unsafe.Pointer, vqr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childrenRectChanged"); signal != nil {
		(*(*func(*std_core.QRectF))(signal))(std_core.NewQRectFFromPointer(vqr))
	}

}

//export callbackItemTemplate40f1b6_ClassBegin
func callbackItemTemplate40f1b6_ClassBegin(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "classBegin"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewItemTemplateFromPointer(ptr).ClassBeginDefault()
	}
}

func (ptr *ItemTemplate) ClassBeginDefault() {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6_ClassBeginDefault(ptr.Pointer())
	}
}

//export callbackItemTemplate40f1b6_ClipChanged
func callbackItemTemplate40f1b6_ClipChanged(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "clipChanged"); signal != nil {
		(*(*func(bool))(signal))(int8(vbo) != 0)
	}

}

//export callbackItemTemplate40f1b6_ComponentComplete
func callbackItemTemplate40f1b6_ComponentComplete(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "componentComplete"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewItemTemplateFromPointer(ptr).ComponentCompleteDefault()
	}
}

func (ptr *ItemTemplate) ComponentCompleteDefault() {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6_ComponentCompleteDefault(ptr.Pointer())
	}
}

//export callbackItemTemplate40f1b6_ContainmentMaskChanged
func callbackItemTemplate40f1b6_ContainmentMaskChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "containmentMaskChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackItemTemplate40f1b6_Contains
func callbackItemTemplate40f1b6_Contains(ptr unsafe.Pointer, point unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "contains"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QPointF) bool)(signal))(std_core.NewQPointFFromPointer(point)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewItemTemplateFromPointer(ptr).ContainsDefault(std_core.NewQPointFFromPointer(point)))))
}

func (ptr *ItemTemplate) ContainsDefault(point std_core.QPointF_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.ItemTemplate40f1b6_ContainsDefault(ptr.Pointer(), std_core.PointerFromQPointF(point))) != 0
	}
	return false
}

//export callbackItemTemplate40f1b6_DragEnterEvent
func callbackItemTemplate40f1b6_DragEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragEnterEvent"); signal != nil {
		(*(*func(*std_gui.QDragEnterEvent))(signal))(std_gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewItemTemplateFromPointer(ptr).DragEnterEventDefault(std_gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *ItemTemplate) DragEnterEventDefault(event std_gui.QDragEnterEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6_DragEnterEventDefault(ptr.Pointer(), std_gui.PointerFromQDragEnterEvent(event))
	}
}

//export callbackItemTemplate40f1b6_DragLeaveEvent
func callbackItemTemplate40f1b6_DragLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragLeaveEvent"); signal != nil {
		(*(*func(*std_gui.QDragLeaveEvent))(signal))(std_gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewItemTemplateFromPointer(ptr).DragLeaveEventDefault(std_gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *ItemTemplate) DragLeaveEventDefault(event std_gui.QDragLeaveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6_DragLeaveEventDefault(ptr.Pointer(), std_gui.PointerFromQDragLeaveEvent(event))
	}
}

//export callbackItemTemplate40f1b6_DragMoveEvent
func callbackItemTemplate40f1b6_DragMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dragMoveEvent"); signal != nil {
		(*(*func(*std_gui.QDragMoveEvent))(signal))(std_gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewItemTemplateFromPointer(ptr).DragMoveEventDefault(std_gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *ItemTemplate) DragMoveEventDefault(event std_gui.QDragMoveEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6_DragMoveEventDefault(ptr.Pointer(), std_gui.PointerFromQDragMoveEvent(event))
	}
}

//export callbackItemTemplate40f1b6_DropEvent
func callbackItemTemplate40f1b6_DropEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "dropEvent"); signal != nil {
		(*(*func(*std_gui.QDropEvent))(signal))(std_gui.NewQDropEventFromPointer(event))
	} else {
		NewItemTemplateFromPointer(ptr).DropEventDefault(std_gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *ItemTemplate) DropEventDefault(event std_gui.QDropEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6_DropEventDefault(ptr.Pointer(), std_gui.PointerFromQDropEvent(event))
	}
}

//export callbackItemTemplate40f1b6_EnabledChanged
func callbackItemTemplate40f1b6_EnabledChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "enabledChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackItemTemplate40f1b6_Event
func callbackItemTemplate40f1b6_Event(ptr unsafe.Pointer, ev unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(ev)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewItemTemplateFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(ev)))))
}

func (ptr *ItemTemplate) EventDefault(ev std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.ItemTemplate40f1b6_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(ev))) != 0
	}
	return false
}

//export callbackItemTemplate40f1b6_FocusChanged
func callbackItemTemplate40f1b6_FocusChanged(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "focusChanged"); signal != nil {
		(*(*func(bool))(signal))(int8(vbo) != 0)
	}

}

//export callbackItemTemplate40f1b6_FocusInEvent
func callbackItemTemplate40f1b6_FocusInEvent(ptr unsafe.Pointer, vqf unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusInEvent"); signal != nil {
		(*(*func(*std_gui.QFocusEvent))(signal))(std_gui.NewQFocusEventFromPointer(vqf))
	} else {
		NewItemTemplateFromPointer(ptr).FocusInEventDefault(std_gui.NewQFocusEventFromPointer(vqf))
	}
}

func (ptr *ItemTemplate) FocusInEventDefault(vqf std_gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6_FocusInEventDefault(ptr.Pointer(), std_gui.PointerFromQFocusEvent(vqf))
	}
}

//export callbackItemTemplate40f1b6_FocusOutEvent
func callbackItemTemplate40f1b6_FocusOutEvent(ptr unsafe.Pointer, vqf unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "focusOutEvent"); signal != nil {
		(*(*func(*std_gui.QFocusEvent))(signal))(std_gui.NewQFocusEventFromPointer(vqf))
	} else {
		NewItemTemplateFromPointer(ptr).FocusOutEventDefault(std_gui.NewQFocusEventFromPointer(vqf))
	}
}

func (ptr *ItemTemplate) FocusOutEventDefault(vqf std_gui.QFocusEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6_FocusOutEventDefault(ptr.Pointer(), std_gui.PointerFromQFocusEvent(vqf))
	}
}

//export callbackItemTemplate40f1b6_GeometryChanged
func callbackItemTemplate40f1b6_GeometryChanged(ptr unsafe.Pointer, newGeometry unsafe.Pointer, oldGeometry unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "geometryChanged"); signal != nil {
		(*(*func(*std_core.QRectF, *std_core.QRectF))(signal))(std_core.NewQRectFFromPointer(newGeometry), std_core.NewQRectFFromPointer(oldGeometry))
	} else {
		NewItemTemplateFromPointer(ptr).GeometryChangedDefault(std_core.NewQRectFFromPointer(newGeometry), std_core.NewQRectFFromPointer(oldGeometry))
	}
}

func (ptr *ItemTemplate) GeometryChangedDefault(newGeometry std_core.QRectF_ITF, oldGeometry std_core.QRectF_ITF) {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6_GeometryChangedDefault(ptr.Pointer(), std_core.PointerFromQRectF(newGeometry), std_core.PointerFromQRectF(oldGeometry))
	}
}

//export callbackItemTemplate40f1b6_HeightChanged
func callbackItemTemplate40f1b6_HeightChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "heightChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackItemTemplate40f1b6_HoverEnterEvent
func callbackItemTemplate40f1b6_HoverEnterEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hoverEnterEvent"); signal != nil {
		(*(*func(*std_gui.QHoverEvent))(signal))(std_gui.NewQHoverEventFromPointer(event))
	} else {
		NewItemTemplateFromPointer(ptr).HoverEnterEventDefault(std_gui.NewQHoverEventFromPointer(event))
	}
}

func (ptr *ItemTemplate) HoverEnterEventDefault(event std_gui.QHoverEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6_HoverEnterEventDefault(ptr.Pointer(), std_gui.PointerFromQHoverEvent(event))
	}
}

//export callbackItemTemplate40f1b6_HoverLeaveEvent
func callbackItemTemplate40f1b6_HoverLeaveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hoverLeaveEvent"); signal != nil {
		(*(*func(*std_gui.QHoverEvent))(signal))(std_gui.NewQHoverEventFromPointer(event))
	} else {
		NewItemTemplateFromPointer(ptr).HoverLeaveEventDefault(std_gui.NewQHoverEventFromPointer(event))
	}
}

func (ptr *ItemTemplate) HoverLeaveEventDefault(event std_gui.QHoverEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6_HoverLeaveEventDefault(ptr.Pointer(), std_gui.PointerFromQHoverEvent(event))
	}
}

//export callbackItemTemplate40f1b6_HoverMoveEvent
func callbackItemTemplate40f1b6_HoverMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "hoverMoveEvent"); signal != nil {
		(*(*func(*std_gui.QHoverEvent))(signal))(std_gui.NewQHoverEventFromPointer(event))
	} else {
		NewItemTemplateFromPointer(ptr).HoverMoveEventDefault(std_gui.NewQHoverEventFromPointer(event))
	}
}

func (ptr *ItemTemplate) HoverMoveEventDefault(event std_gui.QHoverEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6_HoverMoveEventDefault(ptr.Pointer(), std_gui.PointerFromQHoverEvent(event))
	}
}

//export callbackItemTemplate40f1b6_ImplicitHeightChanged
func callbackItemTemplate40f1b6_ImplicitHeightChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "implicitHeightChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackItemTemplate40f1b6_ImplicitWidthChanged
func callbackItemTemplate40f1b6_ImplicitWidthChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "implicitWidthChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackItemTemplate40f1b6_InputMethodEvent
func callbackItemTemplate40f1b6_InputMethodEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "inputMethodEvent"); signal != nil {
		(*(*func(*std_gui.QInputMethodEvent))(signal))(std_gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewItemTemplateFromPointer(ptr).InputMethodEventDefault(std_gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *ItemTemplate) InputMethodEventDefault(event std_gui.QInputMethodEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6_InputMethodEventDefault(ptr.Pointer(), std_gui.PointerFromQInputMethodEvent(event))
	}
}

//export callbackItemTemplate40f1b6_InputMethodQuery
func callbackItemTemplate40f1b6_InputMethodQuery(ptr unsafe.Pointer, query C.longlong) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "inputMethodQuery"); signal != nil {
		return std_core.PointerFromQVariant((*(*func(std_core.Qt__InputMethodQuery) *std_core.QVariant)(signal))(std_core.Qt__InputMethodQuery(query)))
	}

	return std_core.PointerFromQVariant(NewItemTemplateFromPointer(ptr).InputMethodQueryDefault(std_core.Qt__InputMethodQuery(query)))
}

func (ptr *ItemTemplate) InputMethodQueryDefault(query std_core.Qt__InputMethodQuery) *std_core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQVariantFromPointer(C.ItemTemplate40f1b6_InputMethodQueryDefault(ptr.Pointer(), C.longlong(query)))
		runtime.SetFinalizer(tmpValue, (*std_core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackItemTemplate40f1b6_IsTextureProvider
func callbackItemTemplate40f1b6_IsTextureProvider(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "isTextureProvider"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewItemTemplateFromPointer(ptr).IsTextureProviderDefault())))
}

func (ptr *ItemTemplate) IsTextureProviderDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.ItemTemplate40f1b6_IsTextureProviderDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackItemTemplate40f1b6_KeyPressEvent
func callbackItemTemplate40f1b6_KeyPressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyPressEvent"); signal != nil {
		(*(*func(*std_gui.QKeyEvent))(signal))(std_gui.NewQKeyEventFromPointer(event))
	} else {
		NewItemTemplateFromPointer(ptr).KeyPressEventDefault(std_gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *ItemTemplate) KeyPressEventDefault(event std_gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6_KeyPressEventDefault(ptr.Pointer(), std_gui.PointerFromQKeyEvent(event))
	}
}

//export callbackItemTemplate40f1b6_KeyReleaseEvent
func callbackItemTemplate40f1b6_KeyReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "keyReleaseEvent"); signal != nil {
		(*(*func(*std_gui.QKeyEvent))(signal))(std_gui.NewQKeyEventFromPointer(event))
	} else {
		NewItemTemplateFromPointer(ptr).KeyReleaseEventDefault(std_gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *ItemTemplate) KeyReleaseEventDefault(event std_gui.QKeyEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6_KeyReleaseEventDefault(ptr.Pointer(), std_gui.PointerFromQKeyEvent(event))
	}
}

//export callbackItemTemplate40f1b6_MouseDoubleClickEvent
func callbackItemTemplate40f1b6_MouseDoubleClickEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseDoubleClickEvent"); signal != nil {
		(*(*func(*std_gui.QMouseEvent))(signal))(std_gui.NewQMouseEventFromPointer(event))
	} else {
		NewItemTemplateFromPointer(ptr).MouseDoubleClickEventDefault(std_gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *ItemTemplate) MouseDoubleClickEventDefault(event std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6_MouseDoubleClickEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(event))
	}
}

//export callbackItemTemplate40f1b6_MouseMoveEvent
func callbackItemTemplate40f1b6_MouseMoveEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseMoveEvent"); signal != nil {
		(*(*func(*std_gui.QMouseEvent))(signal))(std_gui.NewQMouseEventFromPointer(event))
	} else {
		NewItemTemplateFromPointer(ptr).MouseMoveEventDefault(std_gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *ItemTemplate) MouseMoveEventDefault(event std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6_MouseMoveEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(event))
	}
}

//export callbackItemTemplate40f1b6_MousePressEvent
func callbackItemTemplate40f1b6_MousePressEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mousePressEvent"); signal != nil {
		(*(*func(*std_gui.QMouseEvent))(signal))(std_gui.NewQMouseEventFromPointer(event))
	} else {
		NewItemTemplateFromPointer(ptr).MousePressEventDefault(std_gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *ItemTemplate) MousePressEventDefault(event std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6_MousePressEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(event))
	}
}

//export callbackItemTemplate40f1b6_MouseReleaseEvent
func callbackItemTemplate40f1b6_MouseReleaseEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseReleaseEvent"); signal != nil {
		(*(*func(*std_gui.QMouseEvent))(signal))(std_gui.NewQMouseEventFromPointer(event))
	} else {
		NewItemTemplateFromPointer(ptr).MouseReleaseEventDefault(std_gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *ItemTemplate) MouseReleaseEventDefault(event std_gui.QMouseEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6_MouseReleaseEventDefault(ptr.Pointer(), std_gui.PointerFromQMouseEvent(event))
	}
}

//export callbackItemTemplate40f1b6_MouseUngrabEvent
func callbackItemTemplate40f1b6_MouseUngrabEvent(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "mouseUngrabEvent"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewItemTemplateFromPointer(ptr).MouseUngrabEventDefault()
	}
}

func (ptr *ItemTemplate) MouseUngrabEventDefault() {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6_MouseUngrabEventDefault(ptr.Pointer())
	}
}

//export callbackItemTemplate40f1b6_OpacityChanged
func callbackItemTemplate40f1b6_OpacityChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "opacityChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackItemTemplate40f1b6_ParentChanged
func callbackItemTemplate40f1b6_ParentChanged(ptr unsafe.Pointer, vqq unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "parentChanged"); signal != nil {
		(*(*func(*std_quick.QQuickItem))(signal))(std_quick.NewQQuickItemFromPointer(vqq))
	}

}

//export callbackItemTemplate40f1b6_ReleaseResources
func callbackItemTemplate40f1b6_ReleaseResources(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "releaseResources"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewItemTemplateFromPointer(ptr).ReleaseResourcesDefault()
	}
}

func (ptr *ItemTemplate) ReleaseResourcesDefault() {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6_ReleaseResourcesDefault(ptr.Pointer())
	}
}

//export callbackItemTemplate40f1b6_RotationChanged
func callbackItemTemplate40f1b6_RotationChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "rotationChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackItemTemplate40f1b6_ScaleChanged
func callbackItemTemplate40f1b6_ScaleChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "scaleChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackItemTemplate40f1b6_SmoothChanged
func callbackItemTemplate40f1b6_SmoothChanged(ptr unsafe.Pointer, vbo C.char) {
	if signal := qt.GetSignal(ptr, "smoothChanged"); signal != nil {
		(*(*func(bool))(signal))(int8(vbo) != 0)
	}

}

//export callbackItemTemplate40f1b6_StateChanged
func callbackItemTemplate40f1b6_StateChanged(ptr unsafe.Pointer, vqs C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "stateChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(vqs))
	}

}

//export callbackItemTemplate40f1b6_TextureProvider
func callbackItemTemplate40f1b6_TextureProvider(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "textureProvider"); signal != nil {
		return std_quick.PointerFromQSGTextureProvider((*(*func() *std_quick.QSGTextureProvider)(signal))())
	}

	return std_quick.PointerFromQSGTextureProvider(NewItemTemplateFromPointer(ptr).TextureProviderDefault())
}

func (ptr *ItemTemplate) TextureProviderDefault() *std_quick.QSGTextureProvider {
	if ptr.Pointer() != nil {
		tmpValue := std_quick.NewQSGTextureProviderFromPointer(C.ItemTemplate40f1b6_TextureProviderDefault(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackItemTemplate40f1b6_TouchEvent
func callbackItemTemplate40f1b6_TouchEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "touchEvent"); signal != nil {
		(*(*func(*std_gui.QTouchEvent))(signal))(std_gui.NewQTouchEventFromPointer(event))
	} else {
		NewItemTemplateFromPointer(ptr).TouchEventDefault(std_gui.NewQTouchEventFromPointer(event))
	}
}

func (ptr *ItemTemplate) TouchEventDefault(event std_gui.QTouchEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6_TouchEventDefault(ptr.Pointer(), std_gui.PointerFromQTouchEvent(event))
	}
}

//export callbackItemTemplate40f1b6_TouchUngrabEvent
func callbackItemTemplate40f1b6_TouchUngrabEvent(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "touchUngrabEvent"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewItemTemplateFromPointer(ptr).TouchUngrabEventDefault()
	}
}

func (ptr *ItemTemplate) TouchUngrabEventDefault() {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6_TouchUngrabEventDefault(ptr.Pointer())
	}
}

//export callbackItemTemplate40f1b6_TransformOriginChanged
func callbackItemTemplate40f1b6_TransformOriginChanged(ptr unsafe.Pointer, vqq C.longlong) {
	if signal := qt.GetSignal(ptr, "transformOriginChanged"); signal != nil {
		(*(*func(std_quick.QQuickItem__TransformOrigin))(signal))(std_quick.QQuickItem__TransformOrigin(vqq))
	}

}

//export callbackItemTemplate40f1b6_Update
func callbackItemTemplate40f1b6_Update(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "update"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewItemTemplateFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *ItemTemplate) UpdateDefault() {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6_UpdateDefault(ptr.Pointer())
	}
}

//export callbackItemTemplate40f1b6_UpdatePolish
func callbackItemTemplate40f1b6_UpdatePolish(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "updatePolish"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewItemTemplateFromPointer(ptr).UpdatePolishDefault()
	}
}

func (ptr *ItemTemplate) UpdatePolishDefault() {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6_UpdatePolishDefault(ptr.Pointer())
	}
}

//export callbackItemTemplate40f1b6_VisibleChanged
func callbackItemTemplate40f1b6_VisibleChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "visibleChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackItemTemplate40f1b6_WheelEvent
func callbackItemTemplate40f1b6_WheelEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "wheelEvent"); signal != nil {
		(*(*func(*std_gui.QWheelEvent))(signal))(std_gui.NewQWheelEventFromPointer(event))
	} else {
		NewItemTemplateFromPointer(ptr).WheelEventDefault(std_gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *ItemTemplate) WheelEventDefault(event std_gui.QWheelEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6_WheelEventDefault(ptr.Pointer(), std_gui.PointerFromQWheelEvent(event))
	}
}

//export callbackItemTemplate40f1b6_WidthChanged
func callbackItemTemplate40f1b6_WidthChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "widthChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackItemTemplate40f1b6_WindowChanged
func callbackItemTemplate40f1b6_WindowChanged(ptr unsafe.Pointer, window unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "windowChanged"); signal != nil {
		(*(*func(*std_quick.QQuickWindow))(signal))(std_quick.NewQQuickWindowFromPointer(window))
	}

}

//export callbackItemTemplate40f1b6_XChanged
func callbackItemTemplate40f1b6_XChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "xChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackItemTemplate40f1b6_YChanged
func callbackItemTemplate40f1b6_YChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "yChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackItemTemplate40f1b6_ZChanged
func callbackItemTemplate40f1b6_ZChanged(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "zChanged"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackItemTemplate40f1b6_ChildEvent
func callbackItemTemplate40f1b6_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*std_core.QChildEvent))(signal))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewItemTemplateFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *ItemTemplate) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackItemTemplate40f1b6_ConnectNotify
func callbackItemTemplate40f1b6_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewItemTemplateFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *ItemTemplate) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackItemTemplate40f1b6_CustomEvent
func callbackItemTemplate40f1b6_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewItemTemplateFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *ItemTemplate) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackItemTemplate40f1b6_DeleteLater
func callbackItemTemplate40f1b6_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewItemTemplateFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *ItemTemplate) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackItemTemplate40f1b6_Destroyed
func callbackItemTemplate40f1b6_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(obj))
	}
	qt.Unregister(ptr)

}

//export callbackItemTemplate40f1b6_DisconnectNotify
func callbackItemTemplate40f1b6_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewItemTemplateFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *ItemTemplate) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackItemTemplate40f1b6_EventFilter
func callbackItemTemplate40f1b6_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewItemTemplateFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *ItemTemplate) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.ItemTemplate40f1b6_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackItemTemplate40f1b6_ObjectNameChanged
func callbackItemTemplate40f1b6_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackItemTemplate40f1b6_TimerEvent
func callbackItemTemplate40f1b6_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*std_core.QTimerEvent))(signal))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewItemTemplateFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *ItemTemplate) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.ItemTemplate40f1b6_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}
