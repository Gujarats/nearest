package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Gujarats/API-Golang/database"
	mgo "gopkg.in/mgo.v2"
)

var mongo *mgo.Session

var logger *log.Logger

func init() {
	logger = log.New(os.Stderr,
		"Driver Model :: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	listConnection := database.SystemConnection()
	mongo = listConnection["mongodb"].(*mgo.Session)
	model := LoadTest{}
	loadTests := model.GetAllLoadTest(mongo, "loadTest1")
	uniqueDatas := findUniqueData(loadTests)

	fmt.Println("===========Data============")
	fmt.Println("unique Data count = ", len(uniqueDatas))
	fmt.Println("max duplicate = ", maxDuplicatData(uniqueDatas, loadTests))
	fmt.Println("min duplicate = ", minDuplicateData(uniqueDatas, loadTests))

	fmt.Println("===========Latency============")
	fmt.Println("min latency= ", minLatency(loadTests))
	fmt.Println("max latency= ", maxLatency(loadTests))
	fmt.Println("average latency= ", averageLatency(loadTests))

}

func averageLatency(loadTests []LoadTest) float64 {
	var total float64
	count := float64(len(loadTests))
	for _, loadTest := range loadTests {
		total += loadTest.Latency
	}

	return total / count

}

func minLatency(loadTests []LoadTest) float64 {
	min := 100000.0
	for _, loadTest := range loadTests {
		if min > loadTest.Latency {
			min = loadTest.Latency
		}
	}

	return min

}

func maxLatency(loadTests []LoadTest) float64 {
	max := 0.0
	for _, loadTest := range loadTests {
		if max < loadTest.Latency {
			max = loadTest.Latency
		}
	}

	return max

}

func findUniqueData(loadTests []LoadTest) []LoadTest {
	var results []LoadTest
	results = append(results, loadTests[0])

	if len(loadTests) > 0 {
		for i := 1; i < len(loadTests); i++ {
			loadTest := loadTests[i]
			if !isExist(loadTest, results) {
				// add new loadTest to results
				results = append(results, loadTest)
			}
		}
	}

	return results

}

func isExist(loadTest LoadTest, results []LoadTest) bool {
	exist := false
	for _, result := range results {
		if result.DriverId == loadTest.DriverId {
			return true
		}
	}

	return exist
}

func maxDuplicatData(uniqueDatas []LoadTest, loadTests []LoadTest) int {
	max := 0
	for _, uniqueData := range uniqueDatas {
		counter := 0
		for _, loadTest := range loadTests {
			if uniqueData.DriverId == loadTest.DriverId {
				counter++
			}
		}
		if max < counter {
			max = counter
		}
	}

	return max
}

func minDuplicateData(uniqueDatas []LoadTest, loadTests []LoadTest) int {
	min := 99999999
	for _, uniqueData := range uniqueDatas {
		counter := 0
		for _, loadTest := range loadTests {
			if uniqueData.DriverId == loadTest.DriverId {
				counter++
			}
		}
		if min > counter {
			min = counter
		}
	}

	return min

}