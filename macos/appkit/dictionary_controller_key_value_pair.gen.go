// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DictionaryControllerKeyValuePair] class.
var DictionaryControllerKeyValuePairClass = _DictionaryControllerKeyValuePairClass{objc.GetClass("NSDictionaryControllerKeyValuePair")}

type _DictionaryControllerKeyValuePairClass struct {
	objc.Class
}

// An interface definition for the [DictionaryControllerKeyValuePair] class.
type IDictionaryControllerKeyValuePair interface {
	objc.IObject
	LocalizedKey() string
	SetLocalizedKey(value string)
	Key() string
	SetKey(value string)
	Value() objc.Object
	SetValue(value objc.IObject)
	IsExplicitlyIncluded() bool
}

// A set of methods implemented by arranged objects to give access to information about those objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdictionarycontrollerkeyvaluepair?language=objc
type DictionaryControllerKeyValuePair struct {
	objc.Object
}

func DictionaryControllerKeyValuePairFrom(ptr unsafe.Pointer) DictionaryControllerKeyValuePair {
	return DictionaryControllerKeyValuePair{
		Object: objc.ObjectFrom(ptr),
	}
}

func (dc _DictionaryControllerKeyValuePairClass) Alloc() DictionaryControllerKeyValuePair {
	rv := objc.Call[DictionaryControllerKeyValuePair](dc, objc.Sel("alloc"))
	return rv
}

func DictionaryControllerKeyValuePair_Alloc() DictionaryControllerKeyValuePair {
	return DictionaryControllerKeyValuePairClass.Alloc()
}

func (dc _DictionaryControllerKeyValuePairClass) New() DictionaryControllerKeyValuePair {
	rv := objc.Call[DictionaryControllerKeyValuePair](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDictionaryControllerKeyValuePair() DictionaryControllerKeyValuePair {
	return DictionaryControllerKeyValuePairClass.New()
}

func (d_ DictionaryControllerKeyValuePair) Init() DictionaryControllerKeyValuePair {
	rv := objc.Call[DictionaryControllerKeyValuePair](d_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdictionarycontrollerkeyvaluepair/1531961-localizedkey?language=objc
func (d_ DictionaryControllerKeyValuePair) LocalizedKey() string {
	rv := objc.Call[string](d_, objc.Sel("localizedKey"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdictionarycontrollerkeyvaluepair/1531961-localizedkey?language=objc
func (d_ DictionaryControllerKeyValuePair) SetLocalizedKey(value string) {
	objc.Call[objc.Void](d_, objc.Sel("setLocalizedKey:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdictionarycontrollerkeyvaluepair/1527207-key?language=objc
func (d_ DictionaryControllerKeyValuePair) Key() string {
	rv := objc.Call[string](d_, objc.Sel("key"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdictionarycontrollerkeyvaluepair/1527207-key?language=objc
func (d_ DictionaryControllerKeyValuePair) SetKey(value string) {
	objc.Call[objc.Void](d_, objc.Sel("setKey:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdictionarycontrollerkeyvaluepair/1532277-value?language=objc
func (d_ DictionaryControllerKeyValuePair) Value() objc.Object {
	rv := objc.Call[objc.Object](d_, objc.Sel("value"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdictionarycontrollerkeyvaluepair/1532277-value?language=objc
func (d_ DictionaryControllerKeyValuePair) SetValue(value objc.IObject) {
	objc.Call[objc.Void](d_, objc.Sel("setValue:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdictionarycontrollerkeyvaluepair/1526717-explicitlyincluded?language=objc
func (d_ DictionaryControllerKeyValuePair) IsExplicitlyIncluded() bool {
	rv := objc.Call[bool](d_, objc.Sel("isExplicitlyIncluded"))
	return rv
}
