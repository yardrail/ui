package main

const (
	searchHeroIconSize    = 18
	searchCompactIconSize = 15
)

func searchExamples() []Example {
	return []Example{
		{Name: "plain", Component: searchExamplePlain()},
		{Name: "hero", Component: searchExampleHero()},
		{Name: "compact", Component: searchExampleCompact()},
	}
}
