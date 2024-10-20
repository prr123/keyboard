package main

// 2. add a loop


// v6 add select for timeout
// v7 experiment with termios0

import (
    "fmt"
    "os"
	"log"
//	"time"
	tos "github.com/pkg/term/termios"
    "golang.org/x/term"
)

type KeyEv struct {
	key int
	typ int
	Err error
}

func main() {

//	var keyRec KeyEv
    // switch stdin into 'raw' mode
    oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
    if err != nil {
        fmt.Println(err)
        return
    }

	for k:=1; k< 11; k++ {
		keyRec := GetKey()

		if keyRec.Err != nil {
			term.Restore(int(os.Stdin.Fd()), oldState)
			log.Fatalf("GetKey: %v\n", err)
		}
		fmt.Printf("key[%d]: %d typ: %d \r\n", k, keyRec.key, keyRec.typ)
		if keyRec.key == 27 {break}
	}
	fmt.Printf("*** exiting ***\r\n")
	term.Restore(int(os.Stdin.Fd()), oldState)
}

func getInp(inp chan int) {

    b:=make([]byte,1)
	_, err := os.Stdin.Read(b)
	if err != nil {
		inp <- -1
	}
	inp <- int(b[0])
}


//func GetKey()(key int, typ int, err error) {
func GetKey()(keyRec KeyEv) {

	res :=-1
    b:=make([]byte,1)
	state :=0
	max := 5
	key := -1
	fd := os.Stdin.Fd()


//	inp :=make(chan int, 1)
	for i:=0; i< max; i++ {

	    _, err := os.Stdin.Read(b)
   		if err != nil {
			keyRec.typ = -1
			keyRec.key = -1
			keyRec.Err = fmt.Errorf("read: %v", err)
			return keyRec
		}
		nchars, err := tos.Tiocinq(fd)
		if err != nil {
			keyRec.typ = -1
			keyRec.key = -1
			keyRec.Err = fmt.Errorf("tiocinq: %v", err)
			return keyRec
		}

		res = int(b[0])
//	fmt.Printf("state %d key: %d nchars: %d\r\n", state, res, nchars)

		switch state {
		case 0:
			key = res
			if res != 27 {
				keyRec.typ = 0
				keyRec.key = key
				keyRec.Err = nil
				return keyRec
			}
			if nchars > 0 {
				state = 1
			} else {
				fmt.Printf("exiting\r\n")
				keyRec.typ = 0
				keyRec.key = key
				keyRec.Err = nil
				return keyRec
			}

		case 1:

			key = res
			if res == 79 {
				state = 2
				break
			}
			if res == 91 {
				state = 3
				break
			}
		// alt
			keyRec.typ = 1
			keyRec.key = key
			keyRec.Err = nil
			return keyRec

		// cursor
		case 2:
			key = res
			keyRec.typ = 2
			keyRec.key = key
			keyRec.Err = nil
			return keyRec

		// f keys 91
		case 3:
			key = res
			switch res {
			case 49:
				state = 5
			case 50:
				state = 6
			case 51:
				state = 9
			case 54:
				state = 10
			default:
				keyRec.typ = 3
				keyRec.key = key
				keyRec.Err = nil
				return keyRec
			}


		case 4:
			key = res
			keyRec.typ = 4
			keyRec.key = key
			keyRec.Err = nil
			return keyRec

		case 5:
			key = res
			state = 7

		case 6:
			key = res
			state = 8

		case 7:
			if res == 126 {
				keyRec.typ = 7
				keyRec.key = key
				keyRec.Err = nil
				return keyRec
			}
			keyRec.typ = -1
			keyRec.key = -1
			keyRec.Err = fmt.Errorf("state 7 Key: %d end code: %d", key, int(b[0]))
			return keyRec

		case 8:
			if res == 126 {
				keyRec.typ = 8
				keyRec.key = key
				keyRec.Err = nil
				return keyRec
			}
			keyRec.typ = -1
			keyRec.key = -1
			keyRec.Err = fmt.Errorf("state 8 Key: %d end code: %d", key, int(b[0]))
			return keyRec

		case 9:
			if res == 126 {
				keyRec.typ = 9
				keyRec.key = key
				keyRec.Err = nil
				return keyRec
			}
			keyRec.typ = -1
			keyRec.key = -1
			keyRec.Err = fmt.Errorf("state 9 Key: %d end code: %d", key, int(b[0]))
			return keyRec

		case 10:
			if res == 126 {
				keyRec.typ = 10
				keyRec.key = key
				keyRec.Err = nil
				return keyRec
			}
			keyRec.typ = -1
			keyRec.key = -1
			keyRec.Err = fmt.Errorf("state 10 Key: %d end code: %d", key, int(b[0]))
			return keyRec

		case 11:
			keyRec.typ = 0
			keyRec.key = key
			keyRec.Err = nil
			return keyRec

		default:
			keyRec.typ = -1
			keyRec.key = -1
			keyRec.Err = fmt.Errorf("invalid state: %d", state)
			return keyRec
		}

	}

		keyRec.typ = -1
		keyRec.key = -1
		keyRec.Err = fmt.Errorf("unable to parse state")
		return keyRec
}
