package mapn

import "errors"

var ErrKeyAlreadyExist = errors.New("github.com/dywoq/dywoqlib/container/mapn: key already exists")
var ErrKeyNotFound = errors.New("github.com/dywoq/dywoqlib/container/mapn: key not found")
var ErrNegativeFixedLength = errors.New("github.com/dywoq/dywoqlib/container/mapn: negative fixed length")
var ErrFixedLengthOutOfBounds = errors.New("github.com/dywoq/dywoqlib/container/mapn: fixed length out of bounds")
var ErrOutOfBounds = errors.New("github.com/dywoq/dywoqlib/container/mapn: out of bounds")
