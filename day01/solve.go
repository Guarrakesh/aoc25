package day01
import (
	"fmt"
	"bufio"
	"strings"
	"strconv"
	"math"
)
func SolvePuzzle1(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	var sum int = 50
	var count =0
	for scanner.Scan() {
	   var text =  scanner.Text()
	   var direction = text[0]
	   var amount = text[1:]
	   amountNumber, error := strconv.Atoi(string(amount))
	   if error != nil {
		   fmt.Printf("Error reading direction: %v\n", error)
			return 0
	   }
	   fmt.Printf("%v %q\n", amountNumber, direction)
		if direction == 'L' {
			sum = (sum - amountNumber)%100
		} else {
			sum = (sum + amountNumber)%100
		
		}


		fmt.Printf("Sum = %v\n", sum)
		if sum == 0 {
			count++
		}
		
	}


	return  count
}

func SolvePuzzle2(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	var y []int
	var x []int
	y = append(y, 50)
	x = append(x, 50)

	var zeroCount =0
	var previous = 50
	for scanner.Scan() {
		var text = scanner.Text()
		var direction = text[0]
		var amount = text[1:]
		amountNumber, _ := strconv.Atoi(string(amount))

		if direction == 'L' {
			amountNumber = -amountNumber
		}

		steps := 0
		newValue := previous + amountNumber
		if (amountNumber > 0) {
			// move right - count the number of times we pass through 0 
			steps = int(math.Floor(float64(newValue)/100)) - int(math.Floor(float64(previous)/100))
		} else if (amountNumber < 0) {
			steps = int(math.Ceil(float64(newValue)/100)) - int(math.Ceil(float64(previous)/100))
			steps = int(math.Abs(float64(steps)))

		
		}

		zeroCount += steps 
		

		previous = (newValue % 100 + 100) % 100
	}


	return zeroCount

}
