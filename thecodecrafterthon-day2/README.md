# Base Converter (Go CLI Tool)

## Overview

The **Base Converter** is a Command-Line Interface (CLI) program written in Go.
It allows users to convert numbers between different number systems:

* Hexadecimal (HEX)
* Binary (BIN)
* Decimal (DEC)

---

## Features

* Convert HEX → Decimal
* Convert Binary → Decimal
* Convert Decimal → Binary & HEX
* Exit option
* Input validation for incorrect values

---

## How It Works

The program runs in a loop and follows this process:

```
Choose Base → Enter Value → Convert → Display Result → Repeat or Exit
```

---

##  How to Run

1. Open your terminal
2. Navigate to the project folder
3. Run:

```bash id="p8k2zn"
go run main.go
```

---

## Menu Options

When the program starts, you will see:

```
1: HEX
2: BIN
3: DEC
4: QUIT
```

---

## Conversions Explained

### 1. HEX → Decimal

* Input a hexadecimal number
* Output is its decimal equivalent

Example:

```
Input: 1A
Output: 26
```

---

### 2. Binary → Decimal

* Input a binary number
* Output is its decimal equivalent

Example:

```
Input: 1010
Output: 10
```

---

### 3. Decimal → Binary & HEX

* Input a decimal number
* Output:

  * Binary representation
  * Hexadecimal representation

Example:

```
Input: 15
Output:
1111
F
```

---

### 4. Quit

Exits the program:

```
=THANK YOU FOR YOUR TIME=-->
```

---

## Error Handling

* Invalid HEX input:

```
=INVALID HEX VALUE=
```

* Invalid Binary input:

```
=INVALID BIN VALUE=
```

* Invalid Decimal input:

```
=INVALID DEC VALUE=
```

* Invalid menu choice:

```
=*OUT OF BASES RANGE START AGAIN*=
```

---

## Loop Behavior

* The program continues running until **Quit (4)** is selected
* After each conversion, it returns to the menu

---

## What I Learn From This Project

* Number base conversion
* Using `strconv` package
* Handling user input in Go
* Error checking
* Building interactive CLI programs

---

## Author

Samuel Jireh
