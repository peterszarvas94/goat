package request

import "net/http"

func HxReswap(w http.ResponseWriter, swap string) {
	if swap != "" {
		w.Header().Set("HX-Reswap", swap)
	}
}
