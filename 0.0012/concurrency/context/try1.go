package context

import "time"

func doingSomeWork() (string, error) {
	time.Sleep(time.Microsecond * 600)
	return "SOme HEavy WoRk DoNe", nil
}
