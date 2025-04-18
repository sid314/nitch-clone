package main

func kanagawaPalette(theme ThemeName) Palette {
	var palette Palette
	switch theme {
	case "kanagawa-lotus":

		palette = append(palette, HexToColor("#c84053"))
		palette = append(palette, HexToColor("#6f894e"))
		palette = append(palette, HexToColor("#77713f"))
		palette = append(palette, HexToColor("#4d699b"))
		palette = append(palette, HexToColor("#b35b79"))
		palette = append(palette, HexToColor("#597b75"))
		palette = append(palette, HexToColor("#545464"))
		palette = append(palette, HexToColor("#8a8980"))
		palette = append(palette, HexToColor("#d7474b"))
		palette = append(palette, HexToColor("#6e915f"))
		palette = append(palette, HexToColor("#836f4a"))
		palette = append(palette, HexToColor("#6693bf"))
		palette = append(palette, HexToColor("#624c83"))
		palette = append(palette, HexToColor("#5e857a"))
		palette = append(palette, HexToColor("#43436c"))
	case "kanagawa-wave":
		palette = append(palette, HexToColor("#c34043"))
		palette = append(palette, HexToColor("#76946a"))
		palette = append(palette, HexToColor("#c0a36e"))
		palette = append(palette, HexToColor("#7e9cd8"))
		palette = append(palette, HexToColor("#957fb8"))
		palette = append(palette, HexToColor("#6a9589"))
		palette = append(palette, HexToColor("#c8c093"))
		palette = append(palette, HexToColor("#727169"))
		palette = append(palette, HexToColor("#e82424"))
		palette = append(palette, HexToColor("#98bb6c"))
		palette = append(palette, HexToColor("#e6c384"))
		palette = append(palette, HexToColor("#7fb4ca"))
		palette = append(palette, HexToColor("#938aa9"))
		palette = append(palette, HexToColor("#7aa89f"))
		palette = append(palette, HexToColor("#dcd7ba"))
	case "kanagawa-dragon":

		palette = append(palette, HexToColor("#c4746e"))
		palette = append(palette, HexToColor("#8a9a7b"))
		palette = append(palette, HexToColor("#c4b28a"))
		palette = append(palette, HexToColor("#8ba4b0"))
		palette = append(palette, HexToColor("#a292a3"))
		palette = append(palette, HexToColor("#8ea4a2"))
		palette = append(palette, HexToColor("#c8c093"))
		palette = append(palette, HexToColor("#a6a69c"))
		palette = append(palette, HexToColor("#e46876"))
		palette = append(palette, HexToColor("#87a987"))
		palette = append(palette, HexToColor("#e6c384"))
		palette = append(palette, HexToColor("#7fb4ca"))
		palette = append(palette, HexToColor("#938aa9"))
		palette = append(palette, HexToColor("#7aa89f"))
		palette = append(palette, HexToColor("#c5c9c5"))
	}
	return palette
}
