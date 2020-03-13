package main
import (
  "fmt"
  "unsafe"
  // "reflect"
)
func main() {
  i := 5
  p := unsafe.Pointer(uintptr(unsafe.Pointer(&i)) + 1)
  fmt.Println(*((*int)(p)), i)
}
