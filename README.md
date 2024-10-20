# keyboard

The goal of this program is to handle single keys and lines.  
There are a number of programs on github that handle this challenge.  

This program is light weight and it relies on well supported go libraries:  
 - golang/x/term
 - unix.termios

Several keys produce an escape sequence which can be up to characters in length.  
This program uses a state machine to sort out valid escape sequences and returns two integers for:  
 - type
 - key

key types:
 - Type 0: Runes  
 - Type 1: Alt + Rune  
 - Type 2: F1 - F4  
 - Type 3: cursors  
 - Type 7: F5 - F8  
 - Type 8: F9 - F12 (F11 changes screen)  
 - Type 9: insert, delete, page-up, page-down  

Note Cntl + key produces runes CNTL + c == 3 for example.  

## next steps

The key value is a rune (1 byte), the type should be another byte, so that a 16 bit integer number provides the exact key.  

Maybe there should be constants defining each key!  

Add a programmable squence?  

Create a line from alpha-numeric characters  

## tests

I tested the key strokes for:  
 - alphanumeric and symbol keys
 - cntl alpha
 - alt alpha
 - cursor keys
 - f1 through f12 keys
 - insert, delete, home, end, page-up and page-down keys
