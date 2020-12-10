package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type vertex struct {
	color         string
	isContainedIn []edge
	Contains      []edge
}

type edge struct {
	count int
	v     *vertex
}

type vertices map[string]*vertex

func main() {
	f, err := os.Open("data")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	vv := make(vertices)
	for scanner.Scan() {
		t := scanner.Text()
		// behold the dumbest parsing
		if strings.Contains(t, "no other bags.") {
			continue
		}
		c := t[:strings.Index(t, " bags")]

		vv.addVertex(strings.TrimSpace(c))
		others := strings.Split(
			t[strings.LastIndex(t, "contain ")+8:], ",")
		for _, o := range others {
			o = strings.TrimLeft(o, " ")
			o = strings.TrimRight(o, ".")
			o = strings.TrimRight(o, "bag")
			o = strings.TrimRight(o, "bags")
			space := strings.Index(o, " ")
			count, _ := strconv.Atoi(o[:space])
			vv.addVertex(strings.TrimSpace(o[space:]))
			vv.addEdge(strings.TrimSpace(o[space:]), strings.TrimSpace(c), count)
		}
	}
	crossed := make(map[string]bool)
	colorsCrossed(vv["shiny gold"], crossed)
	var crossedTotal int
	for range crossed {
		crossedTotal++
	}

	fmt.Println(crossedTotal)
	fmt.Println(bagsContained(vv["shiny gold"]) - 1)
}

func colorsCrossed(v *vertex, crossed map[string]bool) {
	for _, e := range v.isContainedIn {
		crossed[e.v.color] = true
		colorsCrossed(e.v, crossed)
	}
}

func bagsContained(v *vertex) int {
	counter := 1
	for _, e := range v.Contains {
		counter += e.count * bagsContained(e.v)
	}
	return counter
}

func (vv vertices) addVertex(color string) {
	if _, ok := vv[color]; !ok {
		vv[color] = &vertex{
			color: color,
		}
	}
	return
}

func (vv vertices) addEdge(from, to string, count int) {
	// add both iscontainedin edge and contains edge
	isIn := edge{
		count: count,
		v:     vv[to],
	}
	fromNode := vv[from]
	fromNode.isContainedIn = append(fromNode.isContainedIn, isIn)
	contains := edge{
		count: count,
		v:     fromNode,
	}
	vv[to].Contains = append(vv[to].Contains, contains)
	return
}
