package main

type planet struct {
	planetType, species, culture, feature, aspect, pickle string
}

func (p *planet) generate() planet {
	p.planetType = generatePlanetType()
	p.species = generateSpecies()
	p.culture = generateCulture()
	p.feature = generateFeature()
	p.aspect = generateFeatureAspect()
	p.pickle = generatePickle()

	return *p
}

func (p *planet) render(req string) {
	switch req {
	case "type":
		renderOutput("Planet Type: " + p.planetType) //render type
	case "species":
		renderOutput("Planet Species: " + p.species) //render species
	case "culture":
		renderOutput("Planet Culture: " + p.culture) //render culture
	case "feature":
		renderOutput("Planet Feature: " + p.feature) //render feature
	case "aspect":
		renderOutput("Planet Aspect: " + p.aspect) //render aspect
	case "pickle":
		renderOutput("Planet Pickle: " + p.pickle) //render pickle
	default:
		renderOutput("Planet Type: " + p.planetType) //render type
		renderOutput("Planet Species: " + p.species) //render species
		renderOutput("Planet Culture: " + p.culture) //render culture
		renderOutput("Planet Feature: " + p.feature) //render feature
		renderOutput("Planet Aspect: " + p.aspect)   //render aspect
		renderOutput("Planet Pickle: " + p.pickle)   //render pickle
	}
}
