// exercise 1.1
package main
import "fmt"
import "os"
func main() {
  for i := 0; i < len(os.Args); i++ {
    fmt.Println(os.Args[i], i)
  }
}
