package main

// 2. add a loop


// v6 add select for timeout
// v7 experiment with termios0
// v8 change into library routines

import (
    "fmt"
    "os"
	"log"
//	"time"
	tos "github.com/pkg/term/termios"
    "golang.org/x/term"
//	kb "goDemo/keys/xterm/kbLib"
)

type KeyEv struct {
    key int
    typ int
}

func main() {

    // switch stdin into 'raw' mode
	oldstate, err := InitKB()

	for k:=1; k< 21; k++ {

		keyRec, err := GetKey()
		if err != nil {
			err1 := RestoreKB(oldstate)
			if err1 != nil {log.Fatalf("error -- GetKey err: %v Restore: %v\n", err, err1)}
			log.Fatalf("error -- GetKey: %v\n", err)
		}
		fmt.Printf("key[%d]: %d typ: %d \r\n", k, keyRec.key, keyRec.typ)
		if keyRec.key == 27 {break}
	}
	fmt.Printf("*** exiting ***\r\n")
	err = RestoreKB(oldstate)
	if err !=nil {log.Fatalf("error -- Rstore: %v\n", err)}
}

func InitKB()(oldState *term.State, err error) {
    oldState, err = term.MakeRaw(int(os.Stdin.Fd()))
    if err != nil {
        return nil, fmt.Errorf("makeraw: %v", err)
    }
	return oldState, nil
}

func RestoreKB(state *term.State) (err error) {

	err = term.Restore(int(os.Stdin.Fd()), state)
	return err
}


func GetKey()(keyRec KeyEv, err error) {

	res :=-1
    b:=make([]byte,1)
	state :=0
	max := 5
	key := -1
	fd := os.Stdin.Fd()

	for i:=0; i< max; i++ {

	    _, err := os.Stdin.Read(b)
   		if err != nil {
			keyRec.typ = -1
			keyRec.key = -1
			return keyRec, fmt.Errorf("read: %v", err)
		}
		nchars, err := tos.Tiocinq(fd)
		if err != nil {
			keyRec.typ = -1
			keyRec.key = -1
			return keyRec, fmt.Errorf("tiocinq: %v", err)
		}

		res = int(b[0])
//	fmt.Printf("state %d key: %d nchars: %d\r\n", state, res, nchars)

		switch state {
		case 0:
			key = res
			if res != 27 {
				keyRec.typ = 0
				keyRec.key = key
				return keyRec, nil
			}
			if nchars > 0 {
				state = 1
			} else {
				keyRec.typ = 0
				keyRec.key = key
				return keyRec, nil
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
			return keyRec, nil

		// cursor
		case 2:
			key = res
			keyRec.typ = 2
			keyRec.key = key
			return keyRec, nil

		// f keys 91
		case 3:
			key = res
			switch res {
			case 49:
				state = 5
			case 50:

				if nchars == 2 {state = 6; break;}
				if nchars == 1 {state = 9; break;}
				keyRec.typ = -1
				keyRec.key = -1
				return keyRec, fmt.Errorf("tiocinq: %v", err)
			case 51:
				state = 9
			case 53:
				state = 9
			case 54:
				state = 9
			default:
				keyRec.typ = 3
				keyRec.key = key
				return keyRec, nil
			}


		case 4:
			key = res
			keyRec.typ = 4
			keyRec.key = key
			return keyRec, nil

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
				return keyRec, nil
			}
			keyRec.typ = -1
			keyRec.key = -1
			return keyRec, fmt.Errorf("state 7 Key: %d end code: %d", key, res)

		case 8:
			if res == 126 {
				keyRec.typ = 8
				keyRec.key = key
				return keyRec, nil
			}
			keyRec.typ = -1
			keyRec.key = -1
			return keyRec, fmt.Errorf("state 8 Key: %d end code: %d", key, res)

		case 9:
			if res == 126 {
				keyRec.typ = 9
				keyRec.key = key
				return keyRec, nil
			}
			keyRec.typ = -1
			keyRec.key = -1
			return keyRec, fmt.Errorf("state 9 Key: %d end code: %d", key, res)

		case 10:
			if res == 126 {
				keyRec.typ = 10
				keyRec.key = key
				return keyRec, nil
			}
			keyRec.typ = -1
			keyRec.key = -1
			return keyRec, fmt.Errorf("state 10 Key: %d end code: %d", key, res)

		case 11:
			keyRec.typ = 0
			keyRec.key = key
			return keyRec, nil

		default:
			keyRec.typ = -1
			keyRec.key = -1
			return keyRec, fmt.Errorf("invalid state: %d", state)
		}

	}

		keyRec.typ = -1
		keyRec.key = -1
		return keyRec, fmt.Errorf("unable to parse state")
}
