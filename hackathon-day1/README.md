# Go CLI Calculator

## Overview

This is a simple **Command-Line Calculator** written in Go.
It allows users to input two numbers and perform basic arithmetic operations interactively.

---

## Features

* Addition
* Subtraction
* Multiplication
* Division (with remainder)
* Help menu
* Exit option
* Input validation (handles invalid numbers)

---

## How It Works

The program runs in a loop and follows this process:

```text
Input A → Input B → Choose Operation → Display Result → Repeat or Exit
```

---

## How to Run

1. Open your terminal
2. Navigate to the folder containing the file
3. Run:

```bash
go run main.go
```

---

## User Input

You will be prompted to enter:

```text
Input A:
Input B:
```

Then choose an operation:

```text
1: Addition
2: Subtraction
3: Multiplication
4: Division
5: Help
6: Quit
```

---

## Operations Explained

### 1. Addition

Adds two numbers:

```text
A + B
```

---

### 2. Subtraction

Subtracts second number from the first:

```text
A - B
```

---

### 3. Multiplication

Multiplies both numbers:

```text
A * B
```

---

### 4. Division

Divides A by B and shows:

* Result
* Remainder

Example:

```text
=ANSWER=: 5 =REMAINDER=: 2
```
* If B = 0:

```text
=ERROR=
```

---

### 5. Help

Displays information about the calculator:

```text
Welcome To Pro Calculator
What do you want to calculate today?
We help in addition, subtraction, multiplication and division.
```

---

### 6. Quit

Exits the program:

```text
==THANK YOU FOR CALCULATING WITH US==-->
```

---

## Error Handling

* If you enter a non-number:

```text
=ENTER A VALID NUMBER=
```

* If you choose an invalid operation:

```text
=OUT OF OPERATING RANGE= START AGAIN
```

---

## Loop Behavior

* The program keeps running until you choose **Quit (6)**
* After every operation, it returns to the beginning

---

## What I Learn From This Project

* Using loops (`for`)
* Handling user input (`fmt.Scanf`, `fmt.Scan`)
* Conditional logic (`switch`)
* Error handling in Go
* Building interactive CLI programs

---

## Author

Samuel Jireh
