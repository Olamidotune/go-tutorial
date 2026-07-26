package main

import (
	"errors"
	"fmt"
)

const (
	freePlan = "free"
	proPlan  = "pro"
)

type cost struct {
	day   int
	value float64
}

func main() {
	messages := []string{
		"SpiderMan",
		"SuperMan",
		"Interstellar",
		"Sinners",
	}

	numbers := []float64{
		1.2,
		3.4,
		5.6,
		7.8,
		9.0,
		10.32,
		43,
		45,
		12.5,
	}

	fizzBuzz()
	send("Eri", 3)
	send("Armour", 1)
	send("Dotun", 0)
	test("John", 0, freePlan)
	test("Doe", 6, proPlan)

	messagesCost := getMessageCost(messages)
	defer fmt.Println("============================================")
	fmt.Printf("The costs of the messages are: %v\n", messagesCost)

	sumTest(numbers...)

	// getCostByDayTest([]cost{
	// 	{day: 0, value: 1.0},
	// 	{day: 1, value: 2.0},
	// 	{day: 0, value: 3.0},
	// 	{day: 2, value: 4.0},
	// })

	getCostByDayTest([]cost{
		{0, 1.0},
		{1, 2.0},
		{1, 3.1},
		{2, 2.5},
		{3, 3.6},
		{3, 2.7},
		{4, 3.34},
	})
}

func fizzBuzz() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Printf("%v fizzbuzz\n", i)
		} else if i%3 == 0 {
			fmt.Printf("fizz\n")
		} else if i%5 == 0 {
			fmt.Printf("buzz\n")
		}

	}
}

func getMessageWithRetires() [3]string {
	return [3]string{"Click here to sign up", "Please click me", "Guy click me na"}
}

func send(name string, doneAt int) {
	fmt.Printf("Sending to %s...", name)
	fmt.Println()

	messages := getMessageWithRetires()
	for i := 0; i < len(messages); i++ {
		msg := messages[i]
		fmt.Printf("Sending: %v\n", msg)
		if i == doneAt {
			fmt.Println("they responded!")
			break
		}
		if i == len(messages)-1 {
			println("Complete Failure")
		}
	}

}

func test(name string, doneAt int, plan string) {
	defer fmt.Println("============================================")
	fmt.Printf("Sending to %s...", name)
	fmt.Println()

	messages, err := getMessageWithRetriesForPlan(plan)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for i := 0; i < len(messages); i++ {
		msg := messages[i]
		fmt.Printf("Sending: '%v'", msg)
		fmt.Println()

		if i == doneAt {
			fmt.Println("they responded!")
			break
		}
		if i == len(messages)-1 {
			println("Complete Failure")
		}
	}

}

func getMessageWithRetriesForPlan(plan string) ([]string, error) {

	allMessages := getMessageWithRetires()

	if plan == proPlan {
		//[:] is to call all the data in the array
		return allMessages[:], nil
	}
	if plan == freePlan {
		//[0:2] means indexes 0 and 1. So if it is [3:9] it should be 3 and 8
		return allMessages[0:2], nil
	}
	return nil, errors.New("Unsupported Plan")
}

func getMessageCost(messages []string) []float64 {
	costs := make([]float64, len(messages))

	for i, x := range messages {
		costs[i] = float64(len(x)) * 0.01
	}

	return costs
}

func sum(numbers ...float64) float64 {
	total := 0.0
	for _, x := range numbers {
		total += x
	}

	fmt.Println(total)
	return total
}

func sumTest(numbers ...float64) {
	total := sum(numbers...)
	defer fmt.Println("============================================")
	fmt.Printf("Summing %v costs..", len(numbers))
	fmt.Printf("Bills for this month: %.2f\n", total)
	fmt.Printf("Total costs for %v items: %.2f\n", len(numbers), total)
}

func getCostByDay(costs []cost) []float64 {
	costValue := []float64{}
	for i := 0; i < len(costs); i++ {
		cost := costs[i]
		for cost.day >= len(costValue) {
			costValue = append(costValue, 0.0)
		}
		costValue[cost.day] += cost.value

	}

	return costValue
}

func getCostByDayTest(costs []cost) {
	costByDay := getCostByDay(costs)
	fmt.Printf("cost by day: %v\n", costByDay)
}
