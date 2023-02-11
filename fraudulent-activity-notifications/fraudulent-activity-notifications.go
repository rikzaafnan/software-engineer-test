package fraudulentactivitynotifications

import (
	"fmt"
	"log"
	"sort"
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
	fmt.Println("debit test 2 :", debits2)
	fmt.Println("days test 2 :", days2)
	fmt.Println("hasil test 2 :", result2)

	fmt.Println("================")
	// copntoh test 3
	debits3 := []int{10, 20, 30, 40, 50}
	days3 := 3

	// getMEdian := getMedian(debits, days)
	result3 := activityNotifications(debits3, days3)
	fmt.Println("debit test 3 :", debits3)
	fmt.Println("days test 3 :", days3)
	fmt.Println("hasil test 3 :", result3)

}

func activityNotifications(debits []int, days int) int {

	//Return if d is too much
	if days+1 > len(debits) {
		return 0
	}

	notices := 0
	var median float64

	// looping for all debit
	for i := 0; i < len(debits); i++ {

		// chekc condition if days > debit
		if days+1 > len(debits) {
			break
		}

		// initialize array for check median
		var countArrForMedian = make([]int, 0)
		if days+i < len(debits) {

			// get data array to check median
			for j := 0; j < days; j++ {
				// countArr[debits[i]]++

				// check if data == total len debit
				if i+j == len(debits) {
					break

				}

				countArrForMedian = append(countArrForMedian, debits[i+j])

			}
		}

		// check array median lebih besar dari 0
		if len(countArrForMedian) > 0 {

			// get median
			median = getMedian(countArrForMedian)

			// check jika median < debit
			if float64(debits[i+days]) >= 2*median {
				notices = notices + 1
			}

			// check median > debit
			if float64(debits[i+days]) < 2*median {
				notices = notices + 0
			}

		}

	}

	return notices

}

func getMedian(arrayForMedian []int) float64 {

	sort.Ints(arrayForMedian)

	var median float64
	l := len(arrayForMedian)
	if l == 0 {
		return 0
	} else if l%2 == 0 {

		duaNilaiTengah := arrayForMedian[l/2-1] + arrayForMedian[l/2]

		median = float64(float64(duaNilaiTengah) / 2)
	} else {
		median = float64(arrayForMedian[l/2])
	}

	return median

}
