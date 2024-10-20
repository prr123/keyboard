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
