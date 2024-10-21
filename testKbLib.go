package main

// 2. add a loop


// v6 add select for timeout
// v7 experiment with termios0
// v8 change into library routines

import (
    "fmt"
//    "os"
	"log"
//	"time"
//	tos "github.com/pkg/term/termios"
//    "golang.org/x/term"
	kb "goDemo/keys/xterm/kbLib"
)


func main() {

    // switch stdin into 'raw' mode
	oldstate, err := kb.InitKB()

	for k:=1; k< 21; k++ {

		keyRec, err := kb.GetKey()
		if err != nil {
			err1 := kb.RestoreKB(oldstate)
			if err1 != nil {log.Fatalf("error -- GetKey err: %v Restore: %v\n", err, err1)}
			log.Fatalf("error -- GetKey: %v\n", err)
		}
		fmt.Printf("key[%d]: %d typ: %d \r\n", k, keyRec.Key, keyRec.Typ)
		if keyRec.Key == 27 {break}
	}
	fmt.Printf("*** exiting ***\r\n")
	err = kb.RestoreKB(oldstate)
	if err !=nil {log.Fatalf("error -- Rstore: %v\n", err)}
}

