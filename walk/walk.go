package walk

import (
	"fmt"
	"os"
	"path/filepath"
)

type Walker struct {
	count uint64
	dist  [10]uint64
}

func Walk(root string) Walker {

	count := uint64(0)
	dist := [10]uint64{}

	first := func(num int64) byte {
		for num > 9 {
			num /= 10
		}
		return byte(num)
	}

	filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			count++
			dist[first(info.Size())]++
			return nil
		})

	return Walker{count, dist}
}

func (w *Walker) Add(add Walker) *Walker {

	w.count += add.count
	for i, n := range w.dist {
		w.dist[i] = n + add.dist[i]
	}
	return w
}

func (w *Walker) Ratio(n int) float64 {

	if n >= 0 && n <= 9 {
		return float64(w.dist[n]) / float64(w.count)
	}
	return 0
}

func (w *Walker) Strings() []string {

	digits := make([]string, 10)

	for i := range w.dist {
		digits[i] = fmt.Sprintf("%d (%.3f%%)", i, w.Ratio(i)*100)
	}

	return digits
}
