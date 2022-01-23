package channelbw

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type channelTest struct {
	in  [][]uint64
	out uint
}
type channelTestCase [][]uint64
type channelTestCases []channelTestCase

var showTimes = [][]uint64{
	{800, 1200, 4},
	{600, 800, 6},
	{900, 1400, 5},
	{1500, 1600, 11},
	{800, 900, 8},
}

func TestSolve1(t *testing.T) {
	assert.Equal(t, 12, solve1(showTimes), "they should be equal")
	//t.Error("expected", 12, "got", solve1(showTimes))
	//fmt.Println("Max BW needed with solve1 = ")
}

func TestSolve2(t *testing.T) {
	assert.Equal(t, 12, solve2(showTimes), "they should be equal")
	//t.Error("expected", 12, "got", solve2(showTimes))
}

func TestSolveFromCSV(t *testing.T) {
	file, err := os.Open("bandwidth.csv")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	// there could be out of memory errors here if the CSV file is large
	// At any point it's enough if a single test-case is available in memory
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
	file.Close()

	if len(txtlines) == 0 {
		log.Fatalf("empty test cases file")
	}

	// first line is number of tests
	tests, err := strconv.Atoi(txtlines[0])
	if err != nil {
		log.Fatalf("Bad input at line 1 !")
	}

	numLines := 1
	for i := 0; i < tests; i++ {
		numChannels, err := strconv.Atoi(strings.TrimSpace(txtlines[numLines]))
		if err != nil {
			log.Fatalf("Bad input for channels for test %d", i)
		}
		numLines++

		channels := txtlines[numLines : numChannels+numLines]
		channelTimeBWs := make([][]uint64, numChannels)

		for c, channel := range channels {
			vals := strings.Split(channel, ",")
			channelTimeBWs[c] = make([]uint64, 3)

			for v, val := range vals {
				e, err := strconv.Atoi(strings.TrimSpace(val))
				if err != nil {
					log.Fatalf("Bad input for test %d %s", i, err)
				}
				channelTimeBWs[c][v] = uint64(e)
			}
		}
		numLines += numChannels

		// next line is expected result from test
		want, err := strconv.Atoi(txtlines[numLines])
		if err != nil {
			log.Fatalf("Bad input at line %d", numLines)
		}
		numLines++

		testname := fmt.Sprintf("Solve1 %d,%d,%d", i+1, numChannels, want)
		t.Run(testname, func(t *testing.T) {
			ans := solve1(channelTimeBWs)
			if ans != want {
				t.Errorf("got %d, want %d", ans, want)
			}
		})

		testname = fmt.Sprintf("Solve2 %d,%d,%d", i+1, numChannels, want)
		t.Run(testname, func(t *testing.T) {
			ans := solve2(channelTimeBWs)
			if ans != want {
				t.Errorf("got %d, want %d", ans, want)
			}
		})
	}
}
