// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MetadataQueryResultGroup] class.
var MetadataQueryResultGroupClass = _MetadataQueryResultGroupClass{objc.GetClass("NSMetadataQueryResultGroup")}

type _MetadataQueryResultGroupClass struct {
	objc.Class
}

// An interface definition for the [MetadataQueryResultGroup] class.
type IMetadataQueryResultGroup interface {
	objc.IObject
	ResultAtIndex(idx uint) objc.Object
	Subgroups() []MetadataQueryResultGroup
	Value() objc.Object
	Attribute() string
	Results() []objc.Object
	ResultCount() uint
}

// The NSMetadataQueryResultGroup class represents a collection of grouped attribute results returned by an NSMetadataQuery object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataqueryresultgroup?language=objc
type MetadataQueryResultGroup struct {
	objc.Object
}

func MetadataQueryResultGroupFrom(ptr unsafe.Pointer) MetadataQueryResultGroup {
	return MetadataQueryResultGroup{
		Object: objc.ObjectFrom(ptr),
	}
}

func (mc _MetadataQueryResultGroupClass) Alloc() MetadataQueryResultGroup {
	rv := objc.Call[MetadataQueryResultGroup](mc, objc.Sel("alloc"))
	return rv
}

func MetadataQueryResultGroup_Alloc() MetadataQueryResultGroup {
	return MetadataQueryResultGroupClass.Alloc()
}

func (mc _MetadataQueryResultGroupClass) New() MetadataQueryResultGroup {
	rv := objc.Call[MetadataQueryResultGroup](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMetadataQueryResultGroup() MetadataQueryResultGroup {
	return MetadataQueryResultGroupClass.New()
}

func (m_ MetadataQueryResultGroup) Init() MetadataQueryResultGroup {
	rv := objc.Call[MetadataQueryResultGroup](m_, objc.Sel("init"))
	return rv
}

// Returns the query result at a specific index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataqueryresultgroup/1410397-resultatindex?language=objc
func (m_ MetadataQueryResultGroup) ResultAtIndex(idx uint) objc.Object {
	rv := objc.Call[objc.Object](m_, objc.Sel("resultAtIndex:"), idx)
	return rv
}

// An array containing the result group’s subgroups. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataqueryresultgroup/1409929-subgroups?language=objc
func (m_ MetadataQueryResultGroup) Subgroups() []MetadataQueryResultGroup {
	rv := objc.Call[[]MetadataQueryResultGroup](m_, objc.Sel("subgroups"))
	return rv
}

// The result group’s value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataqueryresultgroup/1417674-value?language=objc
func (m_ MetadataQueryResultGroup) Value() objc.Object {
	rv := objc.Call[objc.Object](m_, objc.Sel("value"))
	return rv
}

// The result group’s attribute name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataqueryresultgroup/1411276-attribute?language=objc
func (m_ MetadataQueryResultGroup) Attribute() string {
	rv := objc.Call[string](m_, objc.Sel("attribute"))
	return rv
}

// An array containing the result group’s result objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataqueryresultgroup/1410191-results?language=objc
func (m_ MetadataQueryResultGroup) Results() []objc.Object {
	rv := objc.Call[[]objc.Object](m_, objc.Sel("results"))
	return rv
}

// The number of results returned by the result group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataqueryresultgroup/1414790-resultcount?language=objc
func (m_ MetadataQueryResultGroup) ResultCount() uint {
	rv := objc.Call[uint](m_, objc.Sel("resultCount"))
	return rv
}
