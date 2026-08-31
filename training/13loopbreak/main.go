package main
import "fmt"

func main() {
	fmt.Println("lloops is go")

	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	// fmt.Println(days)
	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }
	for i, day := range days {
		if day == "Saturday" {
			fmt.Println("Breaking the loop at", day)
			break
		}
		if day == "Tuesday" {
			fmt.Println("Skipping", day)
			continue
		}
		fmt.Println("day", i, "is", day)
	}


	x := 1

	if x == 1 {
		goto Skip
	}

	fmt.Println("This will be skipped")

	Skip:
	fmt.Println("Goto skipped the previous line!")
}


//its illegal to use goto 