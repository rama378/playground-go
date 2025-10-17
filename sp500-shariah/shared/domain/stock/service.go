package stock

type WVMAService struct{}

func NewWVMAService() *WVMAService {
	return &WVMAService{}
}

func (s *WVMAService) Calculate(data []*Stock, period int) float64 {
	if len(data) < period || period <= 0 {
		return 0
	}

	var (
		totalWeightedPrice float64
		totalVolume        int64
	)

	for _, d := range data[len(data)-period:] {
		totalWeightedPrice += d.Close * float64(d.Volume)
		totalVolume += d.Volume
	}

	if totalVolume == 0 {
		return 0
	}

	return totalWeightedPrice / float64(totalVolume)
}
