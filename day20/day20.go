package day20

import (
	"fmt"
	"strings"

	"github.com/mikeconroy/advent-of-code-23/utils"
)

func Run() (string, string) {
	input := utils.ReadFileIntoSlice("day20/input")
	return part1(input), part2(input)
}

// Structured so each Module can broadcast pulses to the next destination.
// Needs to be refactored so the main loop orchestrates the pulses
// As we need to process pulses in order.
// Broadcast -> a,b,c needs to complete sending pulses before a then sends out pulses.
// Change the broadcastPulse func -> return a list of pulses to send instead.
// Then have the main loop store these in an array to process in order.
func part1(input []string) string {
	modules := loadInput(input)
	for key, val := range modules {
		fmt.Println(key, val.getModule())
	}

	fmt.Println("Inverter", modules["inv"])
	return fmt.Sprint(0)
}

func loadInput(in []string) map[string]Destination {
	dests := make(map[string][]string)
	modules := make(map[string]Destination)
	for _, line := range in {
		id := strings.Split(line, " -> ")[0][1:]
		fmt.Printf("'%s'", id)
		if strings.Contains(line, "broadcaster") {
			id = "broadcaster"
			modules["broadcaster"] = &Broadcaster{mod: &Module{id: "broadcaster", pulsesReceived: make(map[PulseType]int)}}
		} else if line[0] == '%' {
			modules[id] = &FlipFlop{isOn: false, mod: &Module{id: id, pulsesReceived: make(map[PulseType]int)}}
		} else if line[0] == '&' {
			modules[id] = &Conjuction{recentPulse: make(map[string]PulseType), mod: &Module{id: id, pulsesReceived: make(map[PulseType]int)}}
		}

		targets := strings.Split(line, "-> ")[1]
		dests[id] = strings.Split(targets, ", ")
	}

	fmt.Println(modules)

	for source, targets := range dests {
		for _, target := range targets {
			modules[source].getModule().registerDestination(modules[target])
			modules[target].registerSource(source)
		}
	}

	return modules
}

func part2(input []string) string {
	return fmt.Sprint(0)
}

type Destination interface {
	receivePulse(Pulse)
	getModule() *Module
	registerSource(string)
}

type Broadcaster struct {
	mod *Module
}

func (b *Broadcaster) receivePulse(pulse Pulse) {
	mod := b.getModule()
	mod.pulsesReceived[pulse.pType] = mod.pulsesReceived[pulse.pType] + 1
	broadcastPulse(pulse, mod.destinations)
}

func (b *Broadcaster) getModule() *Module {
	return b.mod
}
func (b *Broadcaster) registerSource(source string) {
}

type FlipFlop struct {
	isOn bool
	mod  *Module
}

func (ff *FlipFlop) receivePulse(pulse Pulse) {
	if pulse.pType == low {
		newPulse := Pulse{from: ff.mod.id}
		if ff.isOn {
			newPulse.pType = high
		} else {
			newPulse.pType = low
		}

		broadcastPulse(newPulse, ff.getModule().destinations)
		ff.isOn = !ff.isOn
	}
}

func (ff *FlipFlop) getModule() *Module {
	return ff.mod
}
func (ff *FlipFlop) registerSource(source string) {
}

type Conjuction struct {
	recentPulse map[string]PulseType // A map of Module IDs -> Type of Pulse last received.
	mod         *Module
}

func (c *Conjuction) receivePulse(pulse Pulse) {
	c.recentPulse[pulse.from] = pulse.pType
	allHigh := true
	for _, pType := range c.recentPulse {
		if pType == low {
			allHigh = false
			break
		}
	}
	if allHigh {
		broadcastPulse(Pulse{from: c.mod.id, pType: low}, c.mod.destinations)
	} else {
		broadcastPulse(Pulse{from: c.mod.id, pType: high}, c.mod.destinations)
	}
}

func (c *Conjuction) getModule() *Module {
	return c.mod
}

func (c *Conjuction) registerSource(source string) {
	c.recentPulse[source] = low
}

type Pulse struct {
	pType PulseType
	from  string
}

type PulseType int

const (
	low = iota
	high
)

type Module struct {
	id             string
	destinations   []Destination
	pulsesReceived map[PulseType]int
}

func (m *Module) registerDestination(dest Destination) {
	m.destinations = append(m.destinations, dest)
}

func broadcastPulse(pulse Pulse, destinations []Destination) {
	for _, dest := range destinations {
		dest.receivePulse(pulse)
	}
}
