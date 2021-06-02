package _322_重新安排行程

import (
	"sort"
)

type Airport struct {
	airport string
	visited bool
}

type Airports []Airport

func (a Airports) Len() int {
	return len(a)
}

func (a Airports) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
	return
}
func (a Airports) Less(i, j int) bool {
	return a[i].airport < a[j].airport
}

func findItinerary(tickets [][]string) []string {
	var (
		path    []string
		fn      func(curTicket string, ticketN int) bool
		targets = make(map[string]Airports)
	)

	for _, ticket := range tickets {
		if _, ok := targets[ticket[0]]; !ok {
			targets[ticket[0]] = make(Airports, 0)
		}
		targets[ticket[0]] = append(targets[ticket[0]], Airport{
			airport: ticket[1],
			visited: false,
		})
	}

	for _, target := range targets {
		sort.Sort(target)
	}

	path = append(path, "JFK")

	fn = func(curTicket string, ticketN int) bool {
		if ticketN == len(tickets) {
			return true
		}
		for i, target := range targets[curTicket] {
			if target.visited {
				continue
			}
			path = append(path, target.airport)
			targets[curTicket][i].visited = true
			if fn(target.airport, ticketN+1) {
				return true
			}
			targets[curTicket][i].visited = false
			path = path[:len(path)-1]
		}

		return false
	}

	fn("JFK", 0)

	return path
}
