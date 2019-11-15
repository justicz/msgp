package msgp

// let's just assume appengine
// uses 64-bit hardware...
const smallint = false

func SafeString(b []byte) string {
	return string(b)
}

func SafeBytes(s string) []byte {
	return []byte(s)
}
