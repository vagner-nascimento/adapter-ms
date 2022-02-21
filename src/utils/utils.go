package utils

import "time"

func Sleep(seconds int64) {
	time.Sleep(time.Duration(seconds) * time.Second)
}

func FillMap(targ map[string]string, sourc ...map[string]string) {
	for _, s := range sourc {
		for k, v := range s {
			targ[k] = v
		}
	}
}
