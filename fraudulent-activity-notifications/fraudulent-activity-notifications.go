package fraudulentactivitynotifications

import (
	"fmt"
	"log"
)

// const days = 3

func Fraudulentactivitynotifications() {

	log.Println("hehehe")

	// copntoh test 1
	debits := []int{2, 3, 4, 2, 3, 6, 8, 4, 5}
	days := 5
	// getMEdian := getMedian(debits, days)
	result := activityNotifications(debits, days)
	fmt.Println("debit test 1 :", debits)
	fmt.Println("days test 1 :", days)
	fmt.Println("hasil test 1 :", result)

	fmt.Println("================")

	// copntoh test 2
	debits2 := []int{1, 2, 3, 4, 4}
	days2 := 4
	// getMEdian := getMedian(debits, days)
	result2 := activityNotifications(debits2, days2)
	fmt.Println("debit test 1 :", debits2)
	fmt.Println("days test 3 :", days2)
	fmt.Println("hasil test 1 :", result2)

	fmt.Println("================")
	// copntoh test 3
	debits3 := []int{10, 20, 30, 40, 50}
	days3 := 3

	// notif literasi-1 = 1,
	// notif literasi-2 = 1,

	// getMEdian := getMedian(debits, days)
	result3 := activityNotifications(debits3, days3)
	fmt.Println("debit test 3 :", debits3)
	fmt.Println("days test 3 :", days3)
	fmt.Println("hasil test 3 :", result3)

}

func getMedian(countArr []int, days int) int {
	var sum = 0
	total := 0
	// log.Println("get median")

	for i := 0; i < len(countArr); i++ {

		sum += countArr[i]

		if sum*2 == days {

			total = 1*2 + 1
			// fmt.Println("sum : ", sum)
			// fmt.Println("get median : ", total)
			return total
		}

		if sum*2 > days {
			total = 1 * 2
			// fmt.Println("sum : ", sum)
			// fmt.Println("get median : ", total)
			return total
		}

	}

	return total

}

func activityNotifications(debits []int, days int) int {

	//Return if d is too much
	if days+1 >= len(debits) {
		return 0
	}

	// fmt.Println("debit : ", debits)
	// fmt.Println("days : ", days)

	countArr := make([]int, 210)
	notices := 0

	for i := 0; i < days; i++ {
		countArr[debits[i]]++
	}

	for i := days; i < len(debits); i++ {
		getMedian := getMedian(countArr, days)

		if debits[i] >= getMedian {
			return notices + 1
		}

		if i == len(debits)-1 {
			break
		}

		countArr[debits[i-days]]--
		countArr[debits[i]]++

	}
	// fmt.Println(countArr)
	// fmt.Println(notices)
	return notices

}
