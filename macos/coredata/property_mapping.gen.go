// AUTO-GENERATED CODE, DO NOT MODIFY

package coredata

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PropertyMapping] class.
var PropertyMappingClass = _PropertyMappingClass{objc.GetClass("NSPropertyMapping")}

type _PropertyMappingClass struct {
	objc.Class
}

// An interface definition for the [PropertyMapping] class.
type IPropertyMapping interface {
	objc.IObject
	Name() string
	SetName(value string)
	ValueExpression() foundation.Expression
	SetValueExpression(value foundation.IExpression)
	UserInfo() foundation.Dictionary
	SetUserInfo(value foundation.Dictionary)
}

// A mapping instance that specifies in a model how to map from a property in a source entity to a property in a destination entity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspropertymapping?language=objc
type PropertyMapping struct {
	objc.Object
}

func PropertyMappingFrom(ptr unsafe.Pointer) PropertyMapping {
	return PropertyMapping{
		Object: objc.ObjectFrom(ptr),
	}
}

func (pc _PropertyMappingClass) Alloc() PropertyMapping {
	rv := objc.Call[PropertyMapping](pc, objc.Sel("alloc"))
	return rv
}

func PropertyMapping_Alloc() PropertyMapping {
	return PropertyMappingClass.Alloc()
}

func (pc _PropertyMappingClass) New() PropertyMapping {
	rv := objc.Call[PropertyMapping](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPropertyMapping() PropertyMapping {
	return PropertyMappingClass.New()
}

func (p_ PropertyMapping) Init() PropertyMapping {
	rv := objc.Call[PropertyMapping](p_, objc.Sel("init"))
	return rv
}

// The name of the property in the destination entity for the property mapping. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspropertymapping/1506748-name?language=objc
func (p_ PropertyMapping) Name() string {
	rv := objc.Call[string](p_, objc.Sel("name"))
	return rv
}

// The name of the property in the destination entity for the property mapping. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspropertymapping/1506748-name?language=objc
func (p_ PropertyMapping) SetName(value string) {
	objc.Call[objc.Void](p_, objc.Sel("setName:"), value)
}

// The value expression for the property mapping. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspropertymapping/1506819-valueexpression?language=objc
func (p_ PropertyMapping) ValueExpression() foundation.Expression {
	rv := objc.Call[foundation.Expression](p_, objc.Sel("valueExpression"))
	return rv
}

// The value expression for the property mapping. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspropertymapping/1506819-valueexpression?language=objc
func (p_ PropertyMapping) SetValueExpression(value foundation.IExpression) {
	objc.Call[objc.Void](p_, objc.Sel("setValueExpression:"), objc.Ptr(value))
}

// The user info for the property mapping. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspropertymapping/1506516-userinfo?language=objc
func (p_ PropertyMapping) UserInfo() foundation.Dictionary {
	rv := objc.Call[foundation.Dictionary](p_, objc.Sel("userInfo"))
	return rv
}

// The user info for the property mapping. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspropertymapping/1506516-userinfo?language=objc
func (p_ PropertyMapping) SetUserInfo(value foundation.Dictionary) {
	objc.Call[objc.Void](p_, objc.Sel("setUserInfo:"), value)
}
