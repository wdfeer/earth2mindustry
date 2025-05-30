package px

type Pixel struct {
	R, G, B uint8
}

func (self Pixel) Equal(other Pixel) bool {
	return self.R == other.R && self.G == other.G && self.B == other.B
}
