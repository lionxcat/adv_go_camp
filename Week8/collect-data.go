package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
)

func main() {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		panic(err.Error())
	}
	filenameReg := regexp.MustCompile(`redis_(\d+)\.txt`)
	secondsReg := regexp.MustCompile(`requests completed in (.+) seconds`)
	throughputReg := regexp.MustCompile(`throughput summary: (.+) requests per second`)
	latencyReg := regexp.MustCompile(`\s+(\d+\.\d+)\s+(\d+\.\d+)\s+(\d+\.\d+)\s+(\d+\.\d+)\s+(\d+\.\d+)\s+(\d+\.\d+)`)

	fmt.Print("payload")
	fmt.Printf("\t%[1]s-seconds\t%[1]s-throughput\t%[1]s-lat-avg\t%[1]s-lat-min\t%[1]s-lat-p50\t%[1]s-lat-p95\t%[1]s-lat-p99\t%[1]s-lat-max", "set")
	fmt.Printf("\t%[1]s-seconds\t%[1]s-throughput\t%[1]s-lat-avg\t%[1]s-lat-min\t%[1]s-lat-p50\t%[1]s-lat-p95\t%[1]s-lat-p99\t%[1]s-lat-max", "get")
	fmt.Println()

	for _, file := range files {
		fileName := file.Name()
		if !filenameReg.MatchString(fileName) {
			continue
		}
		payload := filenameReg.FindStringSubmatch(fileName)[1]
		f, err := os.Open(fileName)
		if err != nil {
			panic(err.Error())
		}
		defer f.Close()
		bytes, err := io.ReadAll(f)
		if err != nil {
			panic(err.Error())
		}

		seconds := secondsReg.FindAllStringSubmatch(string(bytes), 2)
		setSec := seconds[0][1]
		getSec := seconds[1][1]

		throughputs := throughputReg.FindAllStringSubmatch(string(bytes), 2)
		setThroughput := throughputs[0][1]
		getThroughput := throughputs[1][1]

		latencies := latencyReg.FindAllStringSubmatch(string(bytes), 2)
		setLatency := make(map[string]string, 6)
		setLatency["avg"] = latencies[0][1]
		setLatency["min"] = latencies[0][2]
		setLatency["p50"] = latencies[0][3]
		setLatency["p95"] = latencies[0][4]
		setLatency["p99"] = latencies[0][5]
		setLatency["max"] = latencies[0][6]
		getLatency := make(map[string]string, 6)
		getLatency["avg"] = latencies[1][1]
		getLatency["min"] = latencies[1][2]
		getLatency["p50"] = latencies[1][3]
		getLatency["p95"] = latencies[1][4]
		getLatency["p99"] = latencies[1][5]
		getLatency["max"] = latencies[1][6]

		fmt.Print(payload)
		fmt.Printf("\t%s\t%s", setSec, setThroughput)
		for i := 1; i < 7; i++ {
			fmt.Printf("\t%s", latencies[0][i])
		}
		fmt.Printf("\t%s\t%s", getSec, getThroughput)
		for i := 1; i < 7; i++ {
			fmt.Printf("\t%s", latencies[1][i])
		}
		fmt.Println()
	}
}
