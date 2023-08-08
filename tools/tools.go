package tools

import (
	"fmt"
	"time"
)

func FechaMySQL() string {
	t := time.Now()
	return fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second()) //Los meses y dias que tengan un digito les obliga a tener un 0 a la izquierda
}
