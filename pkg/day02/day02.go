package day02

type ColorCount struct {
	Count int    `parser:"@Int"`
	Color string `parser:"@('red' | 'green' | 'blue')"`
}

type Draw struct {
	ColorCounts []*ColorCount `parser:"@@ ( ',' @@ )*"`
}

type Game struct {
	Id    int     `parser:"'Game' @Int ':'"`
	Draws []*Draw `parser:"@@ (';' @@)*"`
}

func (game *Game) IsLegal(maxColors map[string]int) bool {
	for _, draw := range game.Draws {
		for _, colorCount := range draw.ColorCounts {
			if colorCount.Count > maxColors[colorCount.Color] {
				return false
			}
		}
	}
	return true
}

func (game *Game) Power() int {
	maxColors := make(map[string]int)

	for _, draw := range game.Draws {
		for _, colorCount := range draw.ColorCounts {
			if colorCount.Count > maxColors[colorCount.Color] {
				maxColors[colorCount.Color] = colorCount.Count
			}
		}
	}

	result := 1
	for _, maxCount := range maxColors {
		result *= maxCount
	}
	return result
}
