/*
Have the function SimpleAdding(num) add up all the numbers from 1 to num.
For example: if the input is 4 then your program should return 10 because 1 + 2 + 3 + 4 = 10.
For the test cases, the parameter num will be any number from 1 to 1000.
*/

package main
import "fmt"

func SimpleAdding(num int) int {

  // code goes here
  // Note: feel free to modify the return type of this function
  	return num * (num + 1) / 2


}

func main() {

  // do not modify below here, readline is our function
  // that properly reads in the input for you
  fmt.Println(SimpleAdding(readline()))

}
