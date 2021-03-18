package main

import (
	"log"
	"strconv"
)

type sector struct {
	object, coords, size string
}

func (s *sector) generate() sector {
	rnd := generateNumber(1, 6)
	switch rnd {
	case 1:
		rndPlanet := generateNumber(1, 6)
		if rndPlanet < 4 {
			s.object = "Planet"
		} else if rndPlanet < 6 {
			s.object = "Twin Planets"
		} else {
			s.object = "1D6 Planetoids"
		}
	case 2:
		rndOutpost := generateNumber(1, 3)
		switch rndOutpost {
		case 1:
			s.object = "Asteroid Colony"
		case 2:
			s.object = "Space Station"
		case 3:
			s.object = "Shipyard"
		}
	case 3:
		s.object = "Nebula (Half fuel costs.)"
	case 4:
		s.object = "Asteroid Field (Double fuel costs.)"
	case 5:
		s.object = "Badlands (Damaging to ship.)"
	case 6:
		s.object = generateStrangeAnomaly()
	}

	rnd2 := generateNumber(1, 3)
	switch rnd2 {
	case 1:
		s.size = "Small (1x1)"
	case 2:
		s.size = "Medium (2x2)"
	case 3:
		s.size = "Large (3x3)"
	}

	s.coords = "D20: " + strconv.Itoa(generateNumber(1, 20)) + " | D6: " + strconv.Itoa(generateNumber(1, 6))

	return *s
}

func (s *sector) render(req string) {
	if s.object == "Planet" {
		renderOutput("--- Planet ---")
		p := planet{}
		p.generate()
		p.render("all")
	} else if s.object == "Twin Planets" {
		renderOutput("--- Twin Planets ---")
		p1 := planet{}
		p1.generate()
		p1.render("all")

		renderOutput("")
		renderOutput("---")
		renderOutput("")

		p2 := planet{}
		p2.generate()
		p2.render("all")
	} else {
		renderOutput(s.object)
	}

	renderOutput("Coordinates: " + s.coords)
	renderOutput("Size: " + s.size)
}

func generateStrangeAnomaly() string {
	anomalies, err := readNameFile("./data/sectors/anomaly.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	rnd := generateNumber(1, len(anomalies)-1)
	rnd2 := generateNumber(1, len(anomalies)-1)
	rnd3 := generateNumber(1, len(anomalies)-1)

	return anomalies[rnd] + " | " + anomalies[rnd2] + " | " + anomalies[rnd3]
}
