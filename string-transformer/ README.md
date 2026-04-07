# Sentinel String Transformer

## Overview

The **Sentinel String Transformer** is a Command-Line Interface (CLI) tool written in Go.
It allows users to input text and apply different string transformations interactively.

---

## Features

* Convert text to uppercase
* Convert text to lowercase
* Capitalise text
* Smart Title Case formatting
* Convert text to snake_case
* Reverse characters in each word
* Exit program

---

## How It Works

The program runs in a loop and follows this flow:

```
Enter Text → Choose Operation → Transform Text → Display Result → Repeat or Exit
```

---

## How to Run

1. Open your terminal
2. Navigate to the project folder
3. Run:

```
go run main.go
```

---

## User Input

1. Enter any text:

```
hello world example
```

2. Choose an operation:

```text id="3r2nxl"
1: ToUpper
2: ToLower
3: Capitalise
4: Title
5: Snake
6: Reverse
7: Exit
```

---

## Operations Explained

### 1. ToUpper

Converts all characters to uppercase:

```
hello world → HELLO WORLD
```

---

### 2. ToLower

Converts all characters to lowercase:

```
HELLO WORLD → hello world
```

---

### 3. Capitalise

Capitalises each word:

```
hello world → Hello World
```

---

### 4. Title (Smart Title Case)

* Capitalises important words
* Keeps small words lowercase (like "and", "the", "of")

Example:

```
the lord of the rings → The Lord of the Rings
```

---

### 5. Snake Case

Converts text to `snake_case`:

```
Hello World! → hello_world
```

---

### 6. Reverse

Reverses characters in each word:

```
hello world → olleh dlrow
```

---

### 7. Exit

Stops the program:

```
Shutting down String Transformer. Goodbye.
```

---

## Error Handling

* If no text is provided:

```
✗ No text provided.
```

* If an invalid option is selected:

```
✗ Unknown command
```

---

## Special Features

### Smart Title Case

The program avoids capitalising small words like:

```
a, an, the, and, in, of, etc.
```

---

### Snake Case Cleaning

* Removes punctuation
* Converts spaces to underscores

---

### Unicode Support

Uses Go’s `unicode` package to safely handle characters.

---

## Loop Behavior

* Program keeps running until user selects **Exit (7)**
* After every operation, it returns to the main menu

---

## What I Learn From Doing This Project

* String manipulation in Go
* Working with slices and loops
* Using `bufio` for input
* Regular expressions (`regexp`)
* Unicode handling
* CLI program structure

---

## Author

Samuel Jireh
