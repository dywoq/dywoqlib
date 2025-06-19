package attribute

var event func()

// SetEvent sets custom event to attributes-functions.
// If you want to reset the event, you can use SetEvent(nil).
//
// Example without the custom event:
// 	func check() {
//		attribute.Deprecated() // or any attribute-function (attribute.Todo)
// 	}
//
// 	func main() {
//		check()
//	}
//
// Example with custom event:
// 	func check() {
//		attribute.SetEvent(func() { fmt.Println("custom event") })
//		attribute.Deprecated()
// 	}
//
// 	func main() {
//		check() 
//	}
//
func SetEvent(value func()) {
	event = value
}
