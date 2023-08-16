// AUTO-GENERATED CODE, DO NOT MODIFY

package contacts

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Group] class.
var GroupClass = _GroupClass{objc.GetClass("CNGroup")}

type _GroupClass struct {
	objc.Class
}

// An interface definition for the [Group] class.
type IGroup interface {
	objc.IObject
	Name() string
	Identifier() string
}

// An immutable object that represents a group of contacts. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cngroup?language=objc
type Group struct {
	objc.Object
}

func GroupFrom(ptr unsafe.Pointer) Group {
	return Group{
		Object: objc.ObjectFrom(ptr),
	}
}

func (gc _GroupClass) Alloc() Group {
	rv := objc.Call[Group](gc, objc.Sel("alloc"))
	return rv
}

func Group_Alloc() Group {
	return GroupClass.Alloc()
}

func (gc _GroupClass) New() Group {
	rv := objc.Call[Group](gc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewGroup() Group {
	return GroupClass.New()
}

func (g_ Group) Init() Group {
	rv := objc.Call[Group](g_, objc.Sel("init"))
	return rv
}

// Returns a predicate to find subgroups in the specified parent group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cngroup/1403176-predicateforsubgroupsingroupwith?language=objc
func (gc _GroupClass) PredicateForSubgroupsInGroupWithIdentifier(parentGroupIdentifier string) foundation.Predicate {
	rv := objc.Call[foundation.Predicate](gc, objc.Sel("predicateForSubgroupsInGroupWithIdentifier:"), parentGroupIdentifier)
	return rv
}

// Returns a predicate to find subgroups in the specified parent group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cngroup/1403176-predicateforsubgroupsingroupwith?language=objc
func Group_PredicateForSubgroupsInGroupWithIdentifier(parentGroupIdentifier string) foundation.Predicate {
	return GroupClass.PredicateForSubgroupsInGroupWithIdentifier(parentGroupIdentifier)
}

// Returns a predicate to find groups with the specified identifiers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cngroup/1402953-predicateforgroupswithidentifier?language=objc
func (gc _GroupClass) PredicateForGroupsWithIdentifiers(identifiers []string) foundation.Predicate {
	rv := objc.Call[foundation.Predicate](gc, objc.Sel("predicateForGroupsWithIdentifiers:"), identifiers)
	return rv
}

// Returns a predicate to find groups with the specified identifiers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cngroup/1402953-predicateforgroupswithidentifier?language=objc
func Group_PredicateForGroupsWithIdentifiers(identifiers []string) foundation.Predicate {
	return GroupClass.PredicateForGroupsWithIdentifiers(identifiers)
}

// Returns a predicate to find groups in the specified container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cngroup/1402794-predicateforgroupsincontainerwit?language=objc
func (gc _GroupClass) PredicateForGroupsInContainerWithIdentifier(containerIdentifier string) foundation.Predicate {
	rv := objc.Call[foundation.Predicate](gc, objc.Sel("predicateForGroupsInContainerWithIdentifier:"), containerIdentifier)
	return rv
}

// Returns a predicate to find groups in the specified container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cngroup/1402794-predicateforgroupsincontainerwit?language=objc
func Group_PredicateForGroupsInContainerWithIdentifier(containerIdentifier string) foundation.Predicate {
	return GroupClass.PredicateForGroupsInContainerWithIdentifier(containerIdentifier)
}

// The name of the group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cngroup/1403371-name?language=objc
func (g_ Group) Name() string {
	rv := objc.Call[string](g_, objc.Sel("name"))
	return rv
}

// The unique identifier for a group on the device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cngroup/1403172-identifier?language=objc
func (g_ Group) Identifier() string {
	rv := objc.Call[string](g_, objc.Sel("identifier"))
	return rv
}
