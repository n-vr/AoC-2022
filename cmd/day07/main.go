package main

import (
	"log"
	"sort"
	"strconv"

	aoc "github.com/n-vr/AoC-2022"
	"github.com/n-vr/AoC-2022/cmd/day07/filesystem"
)

const (
	totalDiskSpace     int = 70000000
	necessaryDiskSpace int = 30000000
	maxDirSize         int = 100000
)

func main() {
	log.SetFlags(0)

	challengeOne := aoc.NewChallenge(1, 1, solvePartOne)

	err := challengeOne.Solve()
	if err != nil {
		log.Fatal(err)
	}

	log.Print(challengeOne)

	challengeTwo := aoc.NewChallenge(1, 2, solvePartTwo)

	err = challengeTwo.Solve()
	if err != nil {
		log.Fatal(err)
	}

	log.Print("-----\n", challengeTwo)
}

func solvePartOne(input []string) (string, error) {
	fs := filesystem.NewFilesystem()

	cmds := filesystem.ParseCommands(input)

	for _, cmd := range cmds {
		fs.Execute(cmd)
	}

	dirSizes := fs.GetDirectorySizes()

	filteredSizes := filterDirSizes(dirSizes)

	sum := sumDirSizes(filteredSizes)

	return strconv.Itoa(sum), nil
}

func solvePartTwo(input []string) (string, error) {
	fs := filesystem.NewFilesystem()

	cmds := filesystem.ParseCommands(input)

	for _, cmd := range cmds {
		fs.Execute(cmd)
	}

	dirSizes := fs.GetDirectorySizes()

	sort.Slice(dirSizes, func(i, j int) bool {
		return dirSizes[i] < dirSizes[j]
	})

	totalUsed := fs.Root.Size
	sizeToFree := necessaryDiskSpace - (totalDiskSpace - totalUsed)

	for _, size := range dirSizes {
		if size >= sizeToFree {
			return strconv.Itoa(size), nil
		}
	}

	return strconv.Itoa(0), nil
}

// Filters sizes to be below maxDirSize.
func filterDirSizes(sizes []int) []int {
	var filteredSizes []int
	for _, ds := range sizes {
		if ds < maxDirSize {
			filteredSizes = append(filteredSizes, ds)
		}
	}
	return filteredSizes
}

// Returns the sum of all directory sizes.
func sumDirSizes(sizes []int) (sum int) {
	for _, filteredSize := range sizes {
		sum += filteredSize
	}
	return
}
