package after

import "fmt"

type Logger struct{}

func (l *Logger) Log(pesan string) {
	fmt.Println("Pesan: " , pesan)
}