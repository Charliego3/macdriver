// AUTO-GENERATED CODE, DO NOT MODIFY

package cloudkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Database] class.
var DatabaseClass = _DatabaseClass{objc.GetClass("CKDatabase")}

type _DatabaseClass struct {
	objc.Class
}

// An interface definition for the [Database] class.
type IDatabase interface {
	objc.IObject
	FetchAllSubscriptionsWithCompletionHandler(completionHandler func(subscriptions []Subscription, error foundation.Error))
	SaveSubscriptionCompletionHandler(subscription ISubscription, completionHandler func(subscription Subscription, error foundation.Error))
	FetchAllRecordZonesWithCompletionHandler(completionHandler func(zones []RecordZone, error foundation.Error))
	AddOperation(operation IDatabaseOperation)
	SaveRecordZoneCompletionHandler(zone IRecordZone, completionHandler func(zone RecordZone, error foundation.Error))
	DeleteRecordWithIDCompletionHandler(recordID IRecordID, completionHandler func(recordID RecordID, error foundation.Error))
	SaveRecordCompletionHandler(record IRecord, completionHandler func(record Record, error foundation.Error))
	DeleteSubscriptionWithIDCompletionHandler(subscriptionID SubscriptionID, completionHandler func(subscriptionID SubscriptionID, error foundation.Error))
	FetchSubscriptionWithIDCompletionHandler(subscriptionID SubscriptionID, completionHandler func(subscription Subscription, error foundation.Error))
	FetchRecordWithIDCompletionHandler(recordID IRecordID, completionHandler func(record Record, error foundation.Error))
	FetchRecordZoneWithIDCompletionHandler(zoneID IRecordZoneID, completionHandler func(zone RecordZone, error foundation.Error))
	PerformQueryInZoneWithIDCompletionHandler(query IQuery, zoneID IRecordZoneID, completionHandler func(results []Record, error foundation.Error))
	DeleteRecordZoneWithIDCompletionHandler(zoneID IRecordZoneID, completionHandler func(zoneID RecordZoneID, error foundation.Error))
	DatabaseScope() DatabaseScope
}

// An object that represents a collection of record zones and subscriptions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckdatabase?language=objc
type Database struct {
	objc.Object
}

func DatabaseFrom(ptr unsafe.Pointer) Database {
	return Database{
		Object: objc.ObjectFrom(ptr),
	}
}

func (dc _DatabaseClass) Alloc() Database {
	rv := objc.Call[Database](dc, objc.Sel("alloc"))
	return rv
}

func Database_Alloc() Database {
	return DatabaseClass.Alloc()
}

func (dc _DatabaseClass) New() Database {
	rv := objc.Call[Database](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDatabase() Database {
	return DatabaseClass.New()
}

func (d_ Database) Init() Database {
	rv := objc.Call[Database](d_, objc.Sel("init"))
	return rv
}

// Fetches all subscriptions from the current database. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckdatabase/1449110-fetchallsubscriptionswithcomplet?language=objc
func (d_ Database) FetchAllSubscriptionsWithCompletionHandler(completionHandler func(subscriptions []Subscription, error foundation.Error)) {
	objc.Call[objc.Void](d_, objc.Sel("fetchAllSubscriptionsWithCompletionHandler:"), completionHandler)
}

// Saves a specific subscription. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckdatabase/1449102-savesubscription?language=objc
func (d_ Database) SaveSubscriptionCompletionHandler(subscription ISubscription, completionHandler func(subscription Subscription, error foundation.Error)) {
	objc.Call[objc.Void](d_, objc.Sel("saveSubscription:completionHandler:"), objc.Ptr(subscription), completionHandler)
}

// Fetches all record zones from the current database. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckdatabase/1449112-fetchallrecordzoneswithcompletio?language=objc
func (d_ Database) FetchAllRecordZonesWithCompletionHandler(completionHandler func(zones []RecordZone, error foundation.Error)) {
	objc.Call[objc.Void](d_, objc.Sel("fetchAllRecordZonesWithCompletionHandler:"), completionHandler)
}

// Executes the specified operation in the current database. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckdatabase/1449116-addoperation?language=objc
func (d_ Database) AddOperation(operation IDatabaseOperation) {
	objc.Call[objc.Void](d_, objc.Sel("addOperation:"), objc.Ptr(operation))
}

// Saves a specific record zone. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckdatabase/1449108-saverecordzone?language=objc
func (d_ Database) SaveRecordZoneCompletionHandler(zone IRecordZone, completionHandler func(zone RecordZone, error foundation.Error)) {
	objc.Call[objc.Void](d_, objc.Sel("saveRecordZone:completionHandler:"), objc.Ptr(zone), completionHandler)
}

// Deletes a specific record. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckdatabase/1449122-deleterecordwithid?language=objc
func (d_ Database) DeleteRecordWithIDCompletionHandler(recordID IRecordID, completionHandler func(recordID RecordID, error foundation.Error)) {
	objc.Call[objc.Void](d_, objc.Sel("deleteRecordWithID:completionHandler:"), objc.Ptr(recordID), completionHandler)
}

// Saves a specific record. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckdatabase/1449114-saverecord?language=objc
func (d_ Database) SaveRecordCompletionHandler(record IRecord, completionHandler func(record Record, error foundation.Error)) {
	objc.Call[objc.Void](d_, objc.Sel("saveRecord:completionHandler:"), objc.Ptr(record), completionHandler)
}

// Deletes a specific subscription and delivers the deleted subscription’s identifier to a completion handler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckdatabase/1449120-deletesubscriptionwithid?language=objc
func (d_ Database) DeleteSubscriptionWithIDCompletionHandler(subscriptionID SubscriptionID, completionHandler func(subscriptionID SubscriptionID, error foundation.Error)) {
	objc.Call[objc.Void](d_, objc.Sel("deleteSubscriptionWithID:completionHandler:"), subscriptionID, completionHandler)
}

// Fetches a specific subscription and delivers it to a completion handler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckdatabase/1449106-fetchsubscriptionwithid?language=objc
func (d_ Database) FetchSubscriptionWithIDCompletionHandler(subscriptionID SubscriptionID, completionHandler func(subscription Subscription, error foundation.Error)) {
	objc.Call[objc.Void](d_, objc.Sel("fetchSubscriptionWithID:completionHandler:"), subscriptionID, completionHandler)
}

// Fetches a specific record. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckdatabase/1449126-fetchrecordwithid?language=objc
func (d_ Database) FetchRecordWithIDCompletionHandler(recordID IRecordID, completionHandler func(record Record, error foundation.Error)) {
	objc.Call[objc.Void](d_, objc.Sel("fetchRecordWithID:completionHandler:"), objc.Ptr(recordID), completionHandler)
}

// Fetches a specific record zone. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckdatabase/1449104-fetchrecordzonewithid?language=objc
func (d_ Database) FetchRecordZoneWithIDCompletionHandler(zoneID IRecordZoneID, completionHandler func(zone RecordZone, error foundation.Error)) {
	objc.Call[objc.Void](d_, objc.Sel("fetchRecordZoneWithID:completionHandler:"), objc.Ptr(zoneID), completionHandler)
}

// Searches for records matching a predicate in the specified record zone. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckdatabase/1449127-performquery?language=objc
func (d_ Database) PerformQueryInZoneWithIDCompletionHandler(query IQuery, zoneID IRecordZoneID, completionHandler func(results []Record, error foundation.Error)) {
	objc.Call[objc.Void](d_, objc.Sel("performQuery:inZoneWithID:completionHandler:"), objc.Ptr(query), objc.Ptr(zoneID), completionHandler)
}

// Deletes a specific record zone. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckdatabase/1449118-deleterecordzonewithid?language=objc
func (d_ Database) DeleteRecordZoneWithIDCompletionHandler(zoneID IRecordZoneID, completionHandler func(zoneID RecordZoneID, error foundation.Error)) {
	objc.Call[objc.Void](d_, objc.Sel("deleteRecordZoneWithID:completionHandler:"), objc.Ptr(zoneID), completionHandler)
}

// The type of database. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckdatabase/1640398-databasescope?language=objc
func (d_ Database) DatabaseScope() DatabaseScope {
	rv := objc.Call[DatabaseScope](d_, objc.Sel("databaseScope"))
	return rv
}
