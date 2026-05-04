package filter

const (
	// EightBits is used to get RGB into 8 bits representation.
	EightBits = 257

	// Alpha value for RGBA images.
	Alpha = 255

	// MaxCharacters for random name generation.
	MaxCharacters = 7

	// MaxIntensity for negative filter.
	MaxIntensity = 255

	// LetterRunes for random name generation.
	LetterRunes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// Red wave length for gray filter.
	RedWaveLength = 0.21

	// GreenWaveLength for gray filter.
	GreenWaveLength = 0.72

	// BlueWaveLength for gray filter.
	BlueWaveLength = 0.07

	// RGBA max value.
	RGBA = 255

	// Sepia constants.
	SepiaRedForRed   = 0.393
	SepiaGreenForRed = 0.769
	SepiaBlueForRed  = 0.189

	SepiaRedForGreen   = 0.349
	SepiaGreenForGreen = 0.686
	SepiaBlueForGreen  = 0.168

	SepiaRedForBlue   = 0.272
	SepiaGreenForBlue = 0.534
	SepiaBlueForBlue  = 0.131

	// Sketch constants.
	IntensityFactor uint8 = 120
	HighestValue    uint8 = 255
	MeanValue       uint8 = 150
	LowestValue     uint8 = 0

	// Red filter constants.
	MaxRed         = 255
	MinGreenForRed = 200
	MinBlueForRed  = 200

	// Green filter constants.
	MaxGreen        = 200
	MinRedForGreen  = 200
	MinBlueForGreen = 200

	// Blue filter constants.
	MaxBlue         = 255
	MinRedForBlue   = 180
	MinGreenForBlue = 180
)
