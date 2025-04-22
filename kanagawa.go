package main

func kanagawaPalette(theme themeName) palette {
	var palette palette
	switch theme {
	case "kanagawa-lotus":

		palette = append(palette, hexToColor("#c84053"))
		palette = append(palette, hexToColor("#6f894e"))
		palette = append(palette, hexToColor("#77713f"))
		palette = append(palette, hexToColor("#4d699b"))
		palette = append(palette, hexToColor("#b35b79"))
		palette = append(palette, hexToColor("#597b75"))
		palette = append(palette, hexToColor("#545464"))
		palette = append(palette, hexToColor("#8a8980"))
		palette = append(palette, hexToColor("#d7474b"))
		palette = append(palette, hexToColor("#6e915f"))
		palette = append(palette, hexToColor("#836f4a"))
		palette = append(palette, hexToColor("#6693bf"))
		palette = append(palette, hexToColor("#624c83"))
		palette = append(palette, hexToColor("#5e857a"))
		palette = append(palette, hexToColor("#43436c"))
	case "kanagawa-wave":
		palette = append(palette, hexToColor("#c34043"))
		palette = append(palette, hexToColor("#76946a"))
		palette = append(palette, hexToColor("#c0a36e"))
		palette = append(palette, hexToColor("#7e9cd8"))
		palette = append(palette, hexToColor("#957fb8"))
		palette = append(palette, hexToColor("#6a9589"))
		palette = append(palette, hexToColor("#c8c093"))
		palette = append(palette, hexToColor("#727169"))
		palette = append(palette, hexToColor("#e82424"))
		palette = append(palette, hexToColor("#98bb6c"))
		palette = append(palette, hexToColor("#e6c384"))
		palette = append(palette, hexToColor("#7fb4ca"))
		palette = append(palette, hexToColor("#938aa9"))
		palette = append(palette, hexToColor("#7aa89f"))
		palette = append(palette, hexToColor("#dcd7ba"))
	case "kanagawa-dragon":

		palette = append(palette, hexToColor("#c4746e"))
		palette = append(palette, hexToColor("#8a9a7b"))
		palette = append(palette, hexToColor("#c4b28a"))
		palette = append(palette, hexToColor("#8ba4b0"))
		palette = append(palette, hexToColor("#a292a3"))
		palette = append(palette, hexToColor("#8ea4a2"))
		palette = append(palette, hexToColor("#c8c093"))
		palette = append(palette, hexToColor("#a6a69c"))
		palette = append(palette, hexToColor("#e46876"))
		palette = append(palette, hexToColor("#87a987"))
		palette = append(palette, hexToColor("#e6c384"))
		palette = append(palette, hexToColor("#7fb4ca"))
		palette = append(palette, hexToColor("#938aa9"))
		palette = append(palette, hexToColor("#7aa89f"))
		palette = append(palette, hexToColor("#c5c9c5"))
	}
	return palette
}
