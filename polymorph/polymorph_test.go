package polymorph

import (
	"reflect"
	"testing"
)

type testStruct struct {
	A int
	B string
}

type testInterface interface {
	Foo() int
}

type testImpl struct{}

func (testImpl) Foo() int { return 42 }

func TestTypeOfGeneric(t *testing.T) {
	typ := TypeOfGeneric[int]()
	if typ.Kind() != reflect.Int {
		t.Errorf("expected kind Int, got %v", typ.Kind())
	}
}

func TestKindOf(t *testing.T) {
	if KindOf[int]() != reflect.Int {
		t.Errorf("expected kind Int")
	}
	if KindOf[testStruct]() != reflect.Struct {
		t.Errorf("expected kind Struct")
	}
}

func TestPackagePath(t *testing.T) {
	path := PackagePath[testStruct]()
	if path == "" {
		t.Errorf("expected non-empty package path")
	}
}

func TestNumMethods(t *testing.T) {
	if NumMethods[testStruct]() != 0 {
		t.Errorf("expected 0 methods for testStruct")
	}
	if NumMethods[testImpl]() == 0 {
		t.Errorf("expected at least 1 method for testImpl")
	}
}

func TestNumFields(t *testing.T) {
	if NumFields[testStruct]() != 2 {
		t.Errorf("expected 2 fields for testStruct")
	}
	if NumFields[int]() != 0 {
		t.Errorf("expected 0 fields for int")
	}
}

func TestNillable(t *testing.T) {
	if !Nillable[*int]() {
		t.Errorf("expected pointer to be nillable")
	}
	if Nillable[int]() {
		t.Errorf("expected int to not be nillable")
	}
	if !Nillable[map[string]int]() {
		t.Errorf("expected map to be nillable")
	}
}

func TestImplements(t *testing.T) {
	if !Implements[testInterface, testImpl]() {
		t.Errorf("testImpl should implement testInterface")
	}
	if Implements[testStruct, testImpl]() {
		t.Errorf("testImpl should not implement testStruct")
	}
	if Implements[int, testImpl]() {
		t.Errorf("int is not an interface, should return false")
	}
}

func TestHasMethod(t *testing.T) {
	if !HasMethod[testImpl]("Foo") {
		t.Errorf("testImpl should have method Foo")
	}
	if HasMethod[testStruct]("Bar") {
		t.Errorf("testStruct should not have method Bar")
	}
}

func TestHasField(t *testing.T) {
	if !HasField[testStruct]("A") {
		t.Errorf("testStruct should have field A")
	}
	if HasField[testStruct]("Z") {
		t.Errorf("testStruct should not have field Z")
	}
	if HasField[int]("A") {
		t.Errorf("int should not have field A")
	}
}

func TestComparable(t *testing.T) {
	if !Comparable[int]() {
		t.Errorf("int should be comparable")
	}
	if Comparable[map[string]int]() {
		t.Errorf("map should not be comparable")
	}
}
