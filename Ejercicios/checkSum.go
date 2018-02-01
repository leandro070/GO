/*
Have the function CheckNums(num1,num2) take both parameters being passed and
return the string true if num2 is greater than num1, otherwise return the string false.
If the parameter values are equal to each other then return the string -1.
*/

package main
import "fmt"

func CheckNums(num1 int, num2 int) string {

    if num2 > num1 {
        return "true"
    } else if num2 < num1 {
        return "false"
    } else {
        return "-1"
    }

}

func main() {

  // do not modify below here, readline is our function
  // that properly reads in the input for you
  fmt.Println(CheckNums(readline()))

}
