package mapn

import "errors"

var ErrKeyAlreadyExist = errors.New("github.com/dywoq/dywoqlib/container/mapn: key already exists")
var ErrKeyNotFound = errors.New("github.com/dywoq/dywoqlib/container/mapn: key not found")