// AUTO-GENERATED CODE, DO NOT MODIFY

package cloudkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Query] class.
var QueryClass = _QueryClass{objc.GetClass("CKQuery")}

type _QueryClass struct {
	objc.Class
}

// An interface definition for the [Query] class.
type IQuery interface {
	objc.IObject
	SortDescriptors() []foundation.SortDescriptor
	SetSortDescriptors(value []foundation.ISortDescriptor)
	Predicate() foundation.Predicate
	RecordType() RecordType
}

// A query that describes the criteria to apply when searching for records in a database. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckquery?language=objc
type Query struct {
	objc.Object
}

func QueryFrom(ptr unsafe.Pointer) Query {
	return Query{
		Object: objc.ObjectFrom(ptr),
	}
}

func (q_ Query) InitWithRecordTypePredicate(recordType RecordType, predicate foundation.IPredicate) Query {
	rv := objc.Call[Query](q_, objc.Sel("initWithRecordType:predicate:"), recordType, objc.Ptr(predicate))
	return rv
}

// Creates a query with the specified record type and predicate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckquery/1413119-initwithrecordtype?language=objc
func Query_InitWithRecordTypePredicate(recordType RecordType, predicate foundation.IPredicate) Query {
	return QueryClass.Alloc().InitWithRecordTypePredicate(recordType, predicate)
}

func (qc _QueryClass) Alloc() Query {
	rv := objc.Call[Query](qc, objc.Sel("alloc"))
	return rv
}

func Query_Alloc() Query {
	return QueryClass.Alloc()
}

func (qc _QueryClass) New() Query {
	rv := objc.Call[Query](qc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewQuery() Query {
	return QueryClass.New()
}

func (q_ Query) Init() Query {
	rv := objc.Call[Query](q_, objc.Sel("init"))
	return rv
}

// The sort descriptors for organizing the query’s results. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckquery/1413121-sortdescriptors?language=objc
func (q_ Query) SortDescriptors() []foundation.SortDescriptor {
	rv := objc.Call[[]foundation.SortDescriptor](q_, objc.Sel("sortDescriptors"))
	return rv
}

// The sort descriptors for organizing the query’s results. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckquery/1413121-sortdescriptors?language=objc
func (q_ Query) SetSortDescriptors(value []foundation.ISortDescriptor) {
	objc.Call[objc.Void](q_, objc.Sel("setSortDescriptors:"), value)
}

// The predicate to use for matching records. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckquery/1413112-predicate?language=objc
func (q_ Query) Predicate() foundation.Predicate {
	rv := objc.Call[foundation.Predicate](q_, objc.Sel("predicate"))
	return rv
}

// The record type to search. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckquery/1413117-recordtype?language=objc
func (q_ Query) RecordType() RecordType {
	rv := objc.Call[RecordType](q_, objc.Sel("recordType"))
	return rv
}
