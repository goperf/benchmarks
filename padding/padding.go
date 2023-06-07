package padding

/*
+--+--+--+--+--+--+--+--+---+---+
| x | o | o | o | o | o | o | o |
+--+--+--+--+--+--+--+--+---+---+
| x | x | x | x | x | x | x | x |
+--+--+--+--+--+--+--+--+---+---+
| x | x | x | x | o | o | o | o |
+--+--+--+--+--+--+--+--+---+---+
24 bytes
*/
type plainStruct struct {
	myBool bool // 1 byte
	// [... 7 bytes]
	myFloat float64 // 8 bytes
	myInt   int32   // 4 bytes
	// [... 4 bytes]
} // 24 bytes

/*
+--+--+--+--+--+--+--+--+---+---+
| x | x | x | x | x | x | x | x |
+--+--+--+--+--+--+--+--+---+---+
| x | x | x | x | x | o | o | o |
+--+--+--+--+--+--+--+--+---+---+
16 bytes
*/
type optimizedStruct struct {
	myFloat float64 // 8 bytes
	myBool  bool    // 1 byte
	myInt   int32   // 4 bytes
	// [... 3 bytes]
} // 16 bytes
