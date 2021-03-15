package main

import (
	"log"
	"math/rand"
	"strconv"
	"time"
)

func generateNumber(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn((max+1)-min) + min
}

func generateLocation() string {
	locations, err := readNameFile("./data/missions/location.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return locations[generateNumber(0, len(locations)-1)]
}

func generateLocationAspect() string {
	aspects, err := readNameFile("./data/missions/aspect.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return aspects[generateNumber(0, len(aspects)-1)]
}

func generateSnag() string {
	snags, err := readNameFile("./data/missions/snag.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return snags[generateNumber(0, len(snags)-1)]
}

func generateEvent(t int) (string, string) {
	renderOutput("[--- Event ---](fg:green)")

	switch t {
	case 1:
		return generateScuffleEvent()
	case 2:
		return generateSocialEvent()
	case 3:
		return generateEncounterEvent()
	case 4:
		return generateDifficultyEvent()
	}

	return "ERROR", "Something is very wrong with the Event Generator."
}

func generateScuffleEvent() (string, string) {
	scuffles, err := readNameFile("./data/events/scuffle.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	tactics, err := readNameFile("./data/events/tactic.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	enemy := scuffles[generateNumber(0, len(scuffles)-1)]
	tactic := tactics[generateNumber(0, len(tactics)-1)]

	return enemy, tactic
}

func generateSocialEvent() (string, string) {
	socials, err := readNameFile("./data/events/social.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	bearings, err := readNameFile("./data/events/bearing.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	social := socials[generateNumber(0, len(socials)-1)]
	bearing := bearings[generateNumber(0, len(bearings)-1)]

	return social, bearing
}

func generateEncounterEvent() (string, string) {
	encounters, err := readNameFile("./data/events/encounter.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return encounters[generateNumber(0, len(encounters)-1)], generateFlavor()
}

func generateDifficultyEvent() (string, string) {
	difficulties, err := readNameFile("./data/events/difficulty.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return difficulties[generateNumber(0, len(difficulties)-1)], generateFlavor()
}

func generateFlavor() string {
	flavors01, err := readNameFile("./data/events/flavor01.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	flavors02, err := readNameFile("./data/events/flavor02.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	flavors03, err := readNameFile("./data/events/flavor03.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	flavorList := generateNumber(1, 3)
	flavor := ""

	switch flavorList {
	case 1:
		flavor = flavors01[generateNumber(0, len(flavors01)-1)]
	case 2:
		flavor = flavors02[generateNumber(0, len(flavors02)-1)]
	case 3:
		flavor = flavors03[generateNumber(0, len(flavors03)-1)]
	}

	return flavor
}

func generateTreasure() string {
	aspects, err := readNameFile("./data/monsters/treasure01.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	features, err := readNameFile("./data/monsters/treasure02.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	forms, err := readNameFile("./data/monsters/treasure03.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	aspectId := generateNumber(0, len(aspects)-1)
	featureId := generateNumber(0, len(features)-1)
	formId := generateNumber(0, len(forms)-1)

	return aspects[aspectId] + " " + features[featureId] + " " + forms[formId]
}

func generateHazard() string {
	hazardType := generateNumber(1, 3)
	hazards, err := readNameFile("./data/monsters/hazard0" + strconv.Itoa(hazardType) + ".names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	return hazards[generateNumber(0, len(hazards)-1)]
}

func generateShipName() string {
	names1, err := readNameFile("./data/shipnames/01.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	names2, err := readNameFile("./data/shipnames/02.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	names3, err := readNameFile("./data/shipnames/03.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	names4, err := readNameFile("./data/shipnames/04.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	names5, err := readNameFile("./data/shipnames/05.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	names6, err := readNameFile("./data/shipnames/06.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	names7, err := readNameFile("./data/shipnames/07.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	output := ""
	nameType := generateNumber(0, 2)

	switch nameType {
	case 0:

		rnd := generateNumber(0, len(names3))
		rnd2 := generateNumber(0, len(names4))

		output = names3[rnd] + names4[rnd2]
	case 1:

		rnd := generateNumber(0, len(names1))
		rnd2 := generateNumber(0, len(names2))

		output = "The " + names1[rnd] + " " + names2[rnd2]
	case 2:

		rnd := generateNumber(0, len(names5))
		rnd2 := generateNumber(0, len(names6))
		rnd3 := generateNumber(0, len(names7))

		output = "The " + names5[rnd] + " " + names6[rnd2] + " " + names7[rnd3]
	}

	return output
}

func generateShipPerk() string {
	rnd := generateNumber(1, 6)
	perks, err := readNameFile("./data/ships/perk0" + strconv.Itoa(rnd) + ".names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return perks[generateNumber(0, len(perks)-1)]
}

func generateShipQuirk() string {
	rnd := generateNumber(1, 3)
	quirks, err := readNameFile("./data/ships/quirk0" + strconv.Itoa(rnd) + ".names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return quirks[generateNumber(0, len(quirks)-1)]
}

func generateShipOrigin() string {
	origins, err := readNameFile("./data/ships/origin.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return origins[generateNumber(0, len(origins)-1)]
}

func generateSuddenEvent() (string, string) {
	rnd := generateNumber(1, 6)
	one, two := "", ""
	switch rnd {
	case 1:
		one, two = generateScuffleEvent()
	case 2:
		one, two = generateEncounterEvent()
	case 3:
		one = generateLocationAspect()
		two = generateFlavor()
	case 4:
		one, two = generateDifficultyEvent()
	case 5:
		one, two = generateSocialEvent()
	case 6:
		one = "Snag"
		two = generateSnag()
	}
	return one, two
}

func generatePlanet() planet {
	p := planet{}
	p.generate()
	return p
}

func generatePlanetType() string {
	rnd := generateNumber(1, 3)
	types, err := readNameFile("./data/planets/type0" + strconv.Itoa(rnd) + ".names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	return types[generateNumber(0, len(types)-1)]
}

func generateSpecies() string {
	rnd := generateNumber(1, 2)
	prefixes, err := readNameFile("./data/planets/prefix0" + strconv.Itoa(rnd) + ".names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	rnd2 := generateNumber(1, 2)
	suffixes, err := readNameFile("./data/planets/suffix0" + strconv.Itoa(rnd2) + ".names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	return prefixes[generateNumber(0, len(prefixes)-1)] + suffixes[generateNumber(0, len(suffixes)-1)]
}

func generateCulture() string {
	rnd := generateNumber(1, 3)
	cultures, err := readNameFile("./data/planets/culture0" + strconv.Itoa(rnd) + ".names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	rnd2 := generateNumber(1, 6)
	if rnd2 < 4 {
		return cultures[generateNumber(0, len(cultures)-1)]
	} else if rnd2 < 6 {
		return cultures[generateNumber(0, len(cultures)-1)] + " [-](fg:green) " + cultures[generateNumber(0, len(cultures)-1)]
	} else {
		return cultures[generateNumber(0, len(cultures)-1)] + " [-><-](fg:red) " + cultures[generateNumber(0, len(cultures)-1)]
	}
}

func generateFeature() string {
	rnd := generateNumber(1, 3)
	features, err := readNameFile("./data/planets/feature0" + strconv.Itoa(rnd) + ".names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	return features[generateNumber(0, len(features)-1)]
}

func generateFeatureAspect() string {
	rnd := generateNumber(1, 3)
	aspects, err := readNameFile("./data/planets/aspect0" + strconv.Itoa(rnd) + ".names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	return aspects[generateNumber(0, len(aspects)-1)]
}

func generatePickle() string {
	pickles, err := readNameFile("./data/planets/pickle.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	return pickles[generateNumber(0, len(pickles)-1)]
}

func generateWeather() string {
	weathers, err := readNameFile("./data/spaceencounters/weather.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	return weathers[generateNumber(0, len(weathers)-1)]
}

func generateDistressSignal() string {
	distresses, err := readNameFile("./data/spaceencounters/distress.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	return distresses[generateNumber(0, len(distresses)-1)]
}

func generateAnotherShip() string {
	rnd := generateNumber(1, 2)
	ships, err := readNameFile("./data/spaceencounters/ship0" + strconv.Itoa(rnd) + ".names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	return ships[generateNumber(0, len(ships)-1)]
}

func generateShipStatus() string {
	rnd := generateNumber(1, 2)
	statuses, err := readNameFile("./data/spaceencounters/status0" + strconv.Itoa(rnd) + ".names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	return statuses[generateNumber(0, len(statuses)-1)]
}

func generateCreature() (string, string) {
	creatures, err := readNameFile("./data/spaceencounters/creature.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	bearings, err := readNameFile("./data/spaceencounters/bearing.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	return creatures[generateNumber(0, len(creatures)-1)], bearings[generateNumber(0, len(bearings)-1)]
}

func generateIssue() (string, string) {
	issues, err := readNameFile("./data/spaceencounters/issue.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	severities, err := readNameFile("./data/spaceencounters/severity.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return severities[generateNumber(0, len(severities)-1)], issues[generateNumber(0, len(issues)-1)]
}

func generateStrangeEncounter() string {
	stranges, err := readNameFile("./data/spaceencounters/strange.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	return stranges[generateNumber(0, len(stranges)-1)]
}
