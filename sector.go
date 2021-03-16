package main

type sector struct {
	aspect, feature, form string
}

func (s *sector) generate() sector {

	return *s
}

func (s *sector) render(req string) {
}
