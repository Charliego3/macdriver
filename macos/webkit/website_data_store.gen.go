// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [WebsiteDataStore] class.
var WebsiteDataStoreClass = _WebsiteDataStoreClass{objc.GetClass("WKWebsiteDataStore")}

type _WebsiteDataStoreClass struct {
	objc.Class
}

// An interface definition for the [WebsiteDataStore] class.
type IWebsiteDataStore interface {
	objc.IObject
	FetchDataRecordsOfTypesCompletionHandler(dataTypes foundation.ISet, completionHandler func(arg0 []WebsiteDataRecord))
	RemoveDataOfTypesModifiedSinceCompletionHandler(dataTypes foundation.ISet, date foundation.IDate, completionHandler func())
	IsPersistent() bool
	HttpCookieStore() HTTPCookieStore
}

// An object that manages cookies, disk and memory caches, and other types of data for a web view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebsitedatastore?language=objc
type WebsiteDataStore struct {
	objc.Object
}

func WebsiteDataStoreFrom(ptr unsafe.Pointer) WebsiteDataStore {
	return WebsiteDataStore{
		Object: objc.ObjectFrom(ptr),
	}
}

func (wc _WebsiteDataStoreClass) Alloc() WebsiteDataStore {
	rv := objc.Call[WebsiteDataStore](wc, objc.Sel("alloc"))
	return rv
}

func WebsiteDataStore_Alloc() WebsiteDataStore {
	return WebsiteDataStoreClass.Alloc()
}

func (wc _WebsiteDataStoreClass) New() WebsiteDataStore {
	rv := objc.Call[WebsiteDataStore](wc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewWebsiteDataStore() WebsiteDataStore {
	return WebsiteDataStoreClass.New()
}

func (w_ WebsiteDataStore) Init() WebsiteDataStore {
	rv := objc.Call[WebsiteDataStore](w_, objc.Sel("init"))
	return rv
}

// Returns the default data store, which stores data persistently to disk. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebsitedatastore/1532937-defaultdatastore?language=objc
func (wc _WebsiteDataStoreClass) DefaultDataStore() WebsiteDataStore {
	rv := objc.Call[WebsiteDataStore](wc, objc.Sel("defaultDataStore"))
	return rv
}

// Returns the default data store, which stores data persistently to disk. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebsitedatastore/1532937-defaultdatastore?language=objc
func WebsiteDataStore_DefaultDataStore() WebsiteDataStore {
	return WebsiteDataStoreClass.DefaultDataStore()
}

// Fetches the specified types of records from the data store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebsitedatastore/1532932-fetchdatarecordsoftypes?language=objc
func (w_ WebsiteDataStore) FetchDataRecordsOfTypesCompletionHandler(dataTypes foundation.ISet, completionHandler func(arg0 []WebsiteDataRecord)) {
	objc.Call[objc.Void](w_, objc.Sel("fetchDataRecordsOfTypes:completionHandler:"), objc.Ptr(dataTypes), completionHandler)
}

// Returns the set of all the available data types. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebsitedatastore/1532929-allwebsitedatatypes?language=objc
func (wc _WebsiteDataStoreClass) AllWebsiteDataTypes() foundation.Set {
	rv := objc.Call[foundation.Set](wc, objc.Sel("allWebsiteDataTypes"))
	return rv
}

// Returns the set of all the available data types. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebsitedatastore/1532929-allwebsitedatatypes?language=objc
func WebsiteDataStore_AllWebsiteDataTypes() foundation.Set {
	return WebsiteDataStoreClass.AllWebsiteDataTypes()
}

// Creates a new data store object that stores website data in memory, and doesn’t write that data to disk. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebsitedatastore/1532934-nonpersistentdatastore?language=objc
func (wc _WebsiteDataStoreClass) NonPersistentDataStore() WebsiteDataStore {
	rv := objc.Call[WebsiteDataStore](wc, objc.Sel("nonPersistentDataStore"))
	return rv
}

// Creates a new data store object that stores website data in memory, and doesn’t write that data to disk. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebsitedatastore/1532934-nonpersistentdatastore?language=objc
func WebsiteDataStore_NonPersistentDataStore() WebsiteDataStore {
	return WebsiteDataStoreClass.NonPersistentDataStore()
}

// Removes website data that changed after the specified date. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebsitedatastore/1532938-removedataoftypes?language=objc
func (w_ WebsiteDataStore) RemoveDataOfTypesModifiedSinceCompletionHandler(dataTypes foundation.ISet, date foundation.IDate, completionHandler func()) {
	objc.Call[objc.Void](w_, objc.Sel("removeDataOfTypes:modifiedSince:completionHandler:"), objc.Ptr(dataTypes), objc.Ptr(date), completionHandler)
}

// A Boolean value that indicates whether this object stores data to disk. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebsitedatastore/1532928-persistent?language=objc
func (w_ WebsiteDataStore) IsPersistent() bool {
	rv := objc.Call[bool](w_, objc.Sel("isPersistent"))
	return rv
}

// The object that manages the HTTP cookies for your website. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebsitedatastore/2881956-httpcookiestore?language=objc
func (w_ WebsiteDataStore) HttpCookieStore() HTTPCookieStore {
	rv := objc.Call[HTTPCookieStore](w_, objc.Sel("httpCookieStore"))
	return rv
}
