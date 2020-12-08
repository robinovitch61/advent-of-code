package main

// This is the first code I've ever written in golang
// plz no judgement future Leo who is better at go

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "strconv"
)

type Bag struct {
    number int
    color string
    holdsGold bool
}

func readFile(path string) []string {
    var contents []string

    file, err := os.Open(path)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        contents = append(contents, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    return contents
}

func parseBagsString(bagsString string) []Bag {
    var bagsLess = strings.ReplaceAll(bagsString, " bags", "")
    var bagLess = strings.ReplaceAll(bagsLess, " bag", "")
    var periodLess = strings.ReplaceAll(bagLess, ".", "")
    var bags = strings.Split(periodLess, ", ")
    if strings.Join(bags, "") == "no others" {
        return []Bag{}
    } else {
        var allBags []Bag
        for _, bag := range bags {
            var split = strings.Split(bag, " ")
            var count, _ = strconv.Atoi(split[0])
            var color = strings.Join(split[1:], " ")
            allBags = append(allBags, Bag{count, color, false})
        }
        return allBags
    }
}

func parseBag(bag string) (string, []Bag) {
    var split = strings.Split(bag, " bags contain ")
    var colorOfBag, bagsString = split[0], split[1]
    return colorOfBag, parseBagsString(bagsString)
}

func makeBagOBags(bags []string) map[string][]Bag {
    var bagOBags = make(map[string][]Bag)

    for _, bag := range bags {
        var colorOfBag, bagsInBag = parseBag(bag)
        bagOBags[colorOfBag] = bagsInBag
    }

    return bagOBags
}

func holdsSweetSweetGold(color string, bags map[string][]Bag) bool {
    if bags[color] == nil {
        return false
    } else if color == "shiny gold" {
        return true
    } else {
        for _, bag := range bags[color] {
            if bag.holdsGold {
                return true
            } else if holdsSweetSweetGold(bag.color, bags) {
                bag.holdsGold = true
                return true
            }
        }
        return false
    }
}

func numBagsInside(color string, bags map[string][]Bag) int {
    if bags[color] == nil {
        return 0
    } else {
        var bagsInside = 0
        for _, bag := range bags[color] {
            bagsInside = bagsInside + bag.number + (bag.number * numBagsInside(bag.color, bags))
        }
        return bagsInside
    }
}

func main() {
    // --- Day 7: Handy Haversacks ---
    // You land at the regional airport in time for your next flight. In fact, it looks like you'll even have time to grab some food: all flights are currently delayed due to issues in luggage processing.
    // 
    // Due to recent aviation regulations, many rules (your puzzle input) are being enforced about bags and their contents; bags must be color-coded and must contain specific quantities of other color-coded bags. Apparently, nobody responsible for these regulations considered how long they would take to enforce!
    // 
    // For example, consider the following rules:
    // 
    // light red bags contain 1 bright white bag, 2 muted yellow bags.
    // dark orange bags contain 3 bright white bags, 4 muted yellow bags.
    // bright white bags contain 1 shiny gold bag.
    // muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
    // shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
    // dark olive bags contain 3 faded blue bags, 4 dotted black bags.
    // vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
    // faded blue bags contain no other bags.
    // dotted black bags contain no other bags.
    // These rules specify the required contents for 9 bag types. In this example, every faded blue bag is empty, every vibrant plum bag contains 11 bags (5 faded blue and 6 dotted black), and so on.
    // 
    // You have a shiny gold bag. If you wanted to carry it in at least one other bag, how many different bag colors would be valid for the outermost bag? (In other words: how many colors can, eventually, contain at least one shiny gold bag?)
    // 
    // In the above rules, the following options would be available to you:
    // 
    // A bright white bag, which can hold your shiny gold bag directly.
    // A muted yellow bag, which can hold your shiny gold bag directly, plus some other bags.
    // A dark orange bag, which can hold bright white and muted yellow bags, either of which could then hold your shiny gold bag.
    // A light red bag, which can hold bright white and muted yellow bags, either of which could then hold your shiny gold bag.
    // So, in this example, the number of bag colors that can eventually contain at least one shiny gold bag is 4.
    // 
    // How many bag colors can eventually contain at least one shiny gold bag? (The list of rules is quite long; make sure you get all of it.)

    //fileContents := readFile("./07_test_input.txt")
    fileContents := readFile("./07_input.txt")
    allBags := makeBagOBags(fileContents)

    var bagsHoldGold []bool
    for color, _ := range allBags {
        if color == "shiny gold" {
            bagsHoldGold = append(bagsHoldGold, false)
        } else {
            bagsHoldGold = append(bagsHoldGold, holdsSweetSweetGold(color, allBags))
        }
    }

    numHoldGold := 0
    for _, holdsGold := range bagsHoldGold {
        if holdsGold == true {
            numHoldGold += 1
        }
    }
    fmt.Println(numHoldGold)

    // --- Part Two ---
    // It's getting pretty expensive to fly these days - not because of ticket prices, but because of the ridiculous number of bags you need to buy!
    // 
    // Consider again your shiny gold bag and the rules from the above example:
    // 
    // faded blue bags contain 0 other bags.
    // dotted black bags contain 0 other bags.
    // vibrant plum bags contain 11 other bags: 5 faded blue bags and 6 dotted black bags.
    // dark olive bags contain 7 other bags: 3 faded blue bags and 4 dotted black bags.
    // So, a single shiny gold bag must contain 1 dark olive bag (and the 7 bags within it) plus 2 vibrant plum bags (and the 11 bags within each of those): 1 + 1*7 + 2 + 2*11 = 32 bags!
    // 
    // Of course, the actual rules have a small chance of going several levels deeper than this example; be sure to count all of the bags, even if the nesting becomes topologically impractical!
    // 
    // Here's another example:
    // 
    // shiny gold bags contain 2 dark red bags.
    // dark red bags contain 2 dark orange bags.
    // dark orange bags contain 2 dark yellow bags.
    // dark yellow bags contain 2 dark green bags.
    // dark green bags contain 2 dark blue bags.
    // dark blue bags contain 2 dark violet bags.
    // dark violet bags contain no other bags.
    // In this example, a single shiny gold bag must contain 126 other bags.
    // 
    // How many individual bags are required inside your single shiny gold bag?

    fmt.Println(numBagsInside("shiny gold", allBags))
}

