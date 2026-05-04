package filter

// GetFilterList returns a list of available filters with their metadata.
func GetFilterList() []FilterInfo {
	return []FilterInfo{
		{Name: "gray", Description: "Gray scale filter", Apply: ApplyGray},
		{Name: "sepia", Description: "Sepia filter", Apply: ApplySepia},
		{Name: "negative", Description: "Negative filter", Apply: ApplyNegative},
		{Name: "sketch", Description: "Sketch filter", Apply: ApplySketch},
		{Name: "red", Description: "Red scale filter", Apply: ApplyRed},
		{Name: "green", Description: "Green scale filter", Apply: ApplyGreen},
		{Name: "blue", Description: "Blue scale filter", Apply: ApplyBlue},
		{Name: "mirror", Description: "Mirror/flip filter", Apply: ApplyMirror},
		{Name: "sharp", Description: "Sharp filter", Apply: ApplySharp},
		{Name: "blur", Description: "Blur filter", Apply: ApplyBlur},
	}
}

// GetFilterByName returns a filter function by name.
func GetFilterByName(name string) (FilterFunc, bool) {
	filters := GetFilterList()
	for _, f := range filters {
		if f.Name == name {
			return f.Apply, true
		}
	}

	return nil, false
}
