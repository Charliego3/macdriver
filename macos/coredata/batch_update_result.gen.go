// AUTO-GENERATED CODE, DO NOT MODIFY

package coredata

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [BatchUpdateResult] class.
var BatchUpdateResultClass = _BatchUpdateResultClass{objc.GetClass("NSBatchUpdateResult")}

type _BatchUpdateResultClass struct {
	objc.Class
}

// An interface definition for the [BatchUpdateResult] class.
type IBatchUpdateResult interface {
	IPersistentStoreResult
	Result() objc.Object
	ResultType() BatchUpdateRequestResultType
}

// The result returned when executing a batch update request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsbatchupdateresult?language=objc
type BatchUpdateResult struct {
	PersistentStoreResult
}

func BatchUpdateResultFrom(ptr unsafe.Pointer) BatchUpdateResult {
	return BatchUpdateResult{
		PersistentStoreResult: PersistentStoreResultFrom(ptr),
	}
}

func (bc _BatchUpdateResultClass) Alloc() BatchUpdateResult {
	rv := objc.Call[BatchUpdateResult](bc, objc.Sel("alloc"))
	return rv
}

func BatchUpdateResult_Alloc() BatchUpdateResult {
	return BatchUpdateResultClass.Alloc()
}

func (bc _BatchUpdateResultClass) New() BatchUpdateResult {
	rv := objc.Call[BatchUpdateResult](bc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewBatchUpdateResult() BatchUpdateResult {
	return BatchUpdateResultClass.New()
}

func (b_ BatchUpdateResult) Init() BatchUpdateResult {
	rv := objc.Call[BatchUpdateResult](b_, objc.Sel("init"))
	return rv
}

// The result of a batch-update request, either the number of updated objects, the identifiers of the updated objects, or a status value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsbatchupdateresult/1404946-result?language=objc
func (b_ BatchUpdateResult) Result() objc.Object {
	rv := objc.Call[objc.Object](b_, objc.Sel("result"))
	return rv
}

// The type of result that Core Data returns from the request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsbatchupdateresult/1404900-resulttype?language=objc
func (b_ BatchUpdateResult) ResultType() BatchUpdateRequestResultType {
	rv := objc.Call[BatchUpdateRequestResultType](b_, objc.Sel("resultType"))
	return rv
}
