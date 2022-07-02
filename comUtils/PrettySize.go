package comUtils

import "fmt"

func PrettySize(size int) string {
	GB := 1024 * 1024 * 1024 * 1.0
	MB := 1024 * 1024 * 1.0
	KB := 1024 * 1.0
	sizef := float64(size)
	if sizef >= GB {
		return fmt.Sprintf("%.2f", sizef/GB) + " MB"
	} else if sizef >= MB {
		return fmt.Sprintf("%.2f", sizef/MB) + " MB"
	} else {
		return fmt.Sprintf("%.2f", sizef/KB) + " KB"
	}
}
