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

func part1(input []string) string {
	modules := loadInput(input)
	var toProcess []Pulse
	// Could be refactored to find the loop instead of iterating over all 1,000 button presses.
	// But this is performant enough.
	for i := 0; i < 1000; i++ {
		toProcess = append(toProcess, Pulse{to: "broadcaster", pType: low, from: "button"})
		for len(toProcess) > 0 {
			pulse := toProcess[0]
			mod := modules[pulse.to]
			// fmt.Println("Sending pulse:", pulse.String())
			newPulses := mod.receivePulse(pulse)
			// fmt.Println(len(newPulses), "added on from", pulse.to)
			toProcess = append(toProcess, newPulses...)
			toProcess = toProcess[1:]
			// fmt.Println(toProcess)
		}
	}

	highCount, lowCount := getTotalPulses(modules)
	return fmt.Sprint(highCount * lowCount)
}

func part2(input []string) string {
	return fmt.Sprint(0)
}

type Destination interface {
	receivePulse(Pulse) []Pulse
	getModule() *Module
	registerSource(string)
}

type Broadcaster struct {
	mod *Module
}

func (b *Broadcaster) receivePulse(pulse Pulse) []Pulse {
	b.mod.pulsesReceived[pulse.pType] = b.mod.pulsesReceived[pulse.pType] + 1

	return newPulses(pulse.pType, b.mod.destinations, b.mod.id)
	// broadcastPulse(pulse, b.mod.destinations)
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

func (ff *FlipFlop) receivePulse(pulse Pulse) []Pulse {
	ff.mod.pulsesReceived[pulse.pType] = ff.mod.pulsesReceived[pulse.pType] + 1
	// fmt.Println("FlipFlop", ff.mod.id, "Received pulse from", pulse.from, "of type", pulse.pType, "isOn:", ff.isOn)
	var newPType PulseType
	if pulse.pType == low {
		ff.isOn = !ff.isOn
		if ff.isOn {
			newPType = high
		} else {
			newPType = low
		}

		// broadcastPulse(newPulse, ff.getModule().destinations)
		return newPulses(newPType, ff.mod.destinations, ff.mod.id)
	}

	return nil
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

func (c *Conjuction) receivePulse(pulse Pulse) []Pulse {
	c.mod.pulsesReceived[pulse.pType] = c.mod.pulsesReceived[pulse.pType] + 1
	c.recentPulse[pulse.from] = pulse.pType
	allHigh := true
	for _, pType := range c.recentPulse {
		if pType == low {
			allHigh = false
			break
		}
	}

	var newPType PulseType
	if allHigh {
		newPType = low
		// broadcastPulse(Pulse{from: c.mod.id, pType: low}, c.mod.destinations)
	} else {
		newPType = high
		// broadcastPulse(Pulse{from: c.mod.id, pType: high}, c.mod.destinations)
	}

	return newPulses(newPType, c.mod.destinations, c.mod.id)
}

func (c *Conjuction) getModule() *Module {
	return c.mod
}

func (c *Conjuction) registerSource(source string) {
	c.recentPulse[source] = low
}

type Outputter struct {
	mod *Module
}

func (o *Outputter) receivePulse(pulse Pulse) []Pulse {
	o.mod.pulsesReceived[pulse.pType] = o.mod.pulsesReceived[pulse.pType] + 1
	return nil
}

func (o *Outputter) getModule() *Module {
	return o.mod
}

func (o *Outputter) registerSource(source string) {
}

type Pulse struct {
	pType PulseType
	from  string
	to    string
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

func newPulses(pType PulseType, dests []Destination, from string) []Pulse {
	newPulses := make([]Pulse, len(dests))
	for i, dest := range dests {
		newPulses[i] = Pulse{pType: pType, from: from, to: dest.getModule().id}
	}
	return newPulses
}

func (p Pulse) String() string {
	var pt string
	if p.pType == high {
		pt = "high"
	} else {
		pt = "low"
	}
	return fmt.Sprintf("%s -%s-> %s", p.from, pt, p.to)
}

func getTotalPulses(mods map[string]Destination) (highCount, lowCount int) {
	for _, mod := range mods {
		// fmt.Println(mod.getModule().id, mod.getModule().pulsesReceived)
		highCount += mod.getModule().pulsesReceived[PulseType(high)]
		lowCount += mod.getModule().pulsesReceived[PulseType(low)]
		// fmt.Println("Pulses received by ", mod.getModule().id, high, low)
	}
	return highCount, lowCount
}
func loadInput(in []string) map[string]Destination {
	dests := make(map[string][]string)
	modules := make(map[string]Destination)
	for _, line := range in {
		id := strings.Split(line, " -> ")[0][1:]
		// fmt.Printf("'%s'", id)
		if line[0] == '%' {
			modules[id] = &FlipFlop{isOn: false, mod: &Module{id: id, pulsesReceived: make(map[PulseType]int)}}
		} else if line[0] == '&' {
			modules[id] = &Conjuction{recentPulse: make(map[string]PulseType), mod: &Module{id: id, pulsesReceived: make(map[PulseType]int)}}
		} else {
			id = "broadcaster"
			modules["broadcaster"] = &Broadcaster{mod: &Module{id: "broadcaster", pulsesReceived: make(map[PulseType]int)}}
		}

		targets := strings.Split(line, "-> ")[1]
		dests[id] = strings.Split(targets, ", ")
	}

	// fmt.Println(modules)

	for source, targets := range dests {
		for _, target := range targets {
			if modules[target] == nil {
				modules[target] = &Outputter{mod: &Module{id: target, pulsesReceived: make(map[PulseType]int)}}
			}
			// fmt.Println("Registering Source:", source, "Dest:", target)
			modules[source].getModule().registerDestination(modules[target])
			// fmt.Println("registering source", modules[source], modules[target])
			modules[target].registerSource(source)
		}
	}

	return modules
}
