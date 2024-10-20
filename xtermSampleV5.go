package main

// 2. add a loop

import (
    "fmt"
    "os"
	"log"
    "golang.org/x/term"
)

func main() {
    // switch stdin into 'raw' mode
    oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
    if err != nil {
        fmt.Println(err)
        return
    }


	for k:=1; k< 5; k++ {
		key, typ, err := GetKey()
		if err != nil {
			term.Restore(int(os.Stdin.Fd()), oldState)
			log.Fatalf("GetKey: %v\n", err)
		}
		fmt.Printf("key[%d]: %d typ: %d \r\n", k, key, typ)
	}

	term.Restore(int(os.Stdin.Fd()), oldState)
}


func GetKey()(key int, typ int, err error) {

    b:=make([]byte,1)
//	esc := make([]byte,3)
	state :=0
	max := 5
	for i:=0; i< max; i++ {
	    _, err := os.Stdin.Read(b)
   		if err != nil { return -1, -1, fmt.Errorf("read: %v", err)}

	fmt.Printf("state %d key: %d\r\n", state, int(b[0]))

		switch state {
		case 0:
			key = int(b[0])
			if b[0] != 27 {
				return key, 0, nil
			}
			state = 1
		case 1:
			key = int(b[0])
			if b[0] == 79 {
				state = 2
				break
			}
			if b[0] == 91 {
				state = 3
				break
			}
		// alt
			typ = 1

			return key, typ, nil

		// cursor
		case 2:
			key = int(b[0])
			return key, 2, nil

		// f keys 91
		case 3:
			key = int(b[0])
			if b[0] == 49 {
				state = 5
				break
			}
			if b[0] == 50 {
				state = 6
				break
			}
			if b[0] == 51 {
				state = 9
				break
			}
			if b[0] == 54 {
				state = 10
				break
			}

			return key, 3, nil

		case 4:
			key = int(b[0])
			return key, 4, nil

		case 5:
			key = int(b[0])
			state = 7

		case 6:
			key = int(b[0])
			state = 8

		case 7:
			if int(b[0]) == 126 {
				return key, 7, nil
			}
			return -1, -1,fmt.Errorf("state 7 Key: %d end code: %d", key, int(b[0]))

		case 8:
			if int(b[0]) == 126 {
				return key, 8, nil
			}
			return -1, -1,fmt.Errorf("state 8 Key: %d end code: %d", key, int(b[0]))

		case 9:
			if int(b[0]) == 126 {
				return key, 9, nil
			}
			return -1, -1,fmt.Errorf("state 9 Key: %d end code: %d", key, int(b[0]))

		case 10:
			if int(b[0]) == 126 {
				return key, 10, nil
			}
			return -1, -1,fmt.Errorf("state 10 Key: %d end code: %d", key, int(b[0]))

		default:
			return -1, -1,fmt.Errorf("invalid state: %d", state)
		}

	}

	return -1, -1,fmt.Errorf("invalid case: %d", state)
}
