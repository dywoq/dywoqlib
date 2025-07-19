// when Windows workflow runs, it sends this error:
// Run go test -bench=. ./...
// .
// FAIL	. [setup failed]
// no Go files in D:\a\dywoqlib\dywoqlib
// FAIL
// Error: Process completed with exit code 1.
//
// by adding doc.go, it makes the root directory a valid go package,
// even it it doesn't contain any executable code and allowing to execute 
// without the "no Go files" error for the root.

package dywoqlib
