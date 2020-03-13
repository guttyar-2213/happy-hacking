package main
import "fmt"
type ObjectOrientation struct {
    Property string
}
func (e *ObjectOrientation) Method() string {
    return "Property -> \"" + e.Property + "\""
}
func multiple_return(in string) (out string, leng int) {
    return in, len(in)
}
func function_call() {
    str, leng := multiple_return("SUGOKU STRING")
    fmt.Println(str)
    fmt.Println(leng)
}
func object_create() {
    object := ObjectOrientation { Property: "HONTO NI STRING" }
    fmt.Println(object.Method())
}
func array() {
    array := [5]int{1, 2, 3, 4, 5}
    fmt.Println(array)
}
func main() {
    array()
}
