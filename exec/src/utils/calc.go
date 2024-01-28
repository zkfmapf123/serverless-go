package utils

func ByteToMB(b int64) float64 {
	return float64(b) / (1024 * 1024)
}
