/*
Have the function FirstFactorial(num) take the num parameter being passed and return the factorial
of it (e.g. if num = 4, return (4 * 3 * 2 * 1)). For the test cases, the range will be between 1
and 18 and the input will always be an integer. 
*/

package main
import "fmt"

func FirstFactorial(num int) int {
    if num<=1{
     return 1
    } else{
    num = num * FirstFactorial(num-1)
  // code goes here
  // Note: feel free to modify the return type of this function
  return num

    }

}

func main() {

  // do not modify below here, readline is our function
  // that properly reads in the input for you
  fmt.Println(FirstFactorial(readline()))

}
