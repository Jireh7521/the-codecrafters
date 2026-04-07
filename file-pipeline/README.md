# Gopher Pipeline Processor

## Overview

This program is a **file processing pipeline written in Go**.
It reads a text file, applies a set of transformation rules to each line, and writes the processed result to a new file.

---

## How It Works

The program follows this flow:

```
Input File → Process Each Line → Generate Output → Write to File
```

---

## Input

The program takes **two command-line arguments**:

```
go run . input.txt output.txt
```

* `input.txt` → file to read from
* `output.txt` → file to write results to

---

## Transformation Rules

Each line in the input file is processed using the following rules **in order** as agreed by our team members:

1. **Trim whitespace**

   * Removes spaces at the beginning and end of a line

2. **Replace TODO**

   * Lines starting with `TODO:` become:

     ```
     ✦ ACTION:
     ```

3. **ALL CAPS → Title Case**

   * Example:

     ```
     HELLO WORLD → Hello World
     ```

4. **lowercase → UPPERCASE**

   * Example:

     ```
     hello world → HELLO WORLD
     ```

5. **Reverse words (incase "REVERSE" is present)**

   * Example:

     ```
     REVERSE this line → line this REVERSE
     ```

---

## Important Notes

* Empty lines are automatically removed
* Rules are applied **in sequence**
* Not all rules apply to every line
* Each line is processed independently

---

## Output Format

The output file contains:

### 1. Header

```
Gopher's Sentinel Field Report - Processed
```

### 2. Numbered Lines

```
1. FIRST LINE
2. SECOND LINE
```

### 3. Summary Section

```
Summary:
✦ Lines read    : X
✦ Lines written : X
✦ Lines removed : X
✦ Rules applied : 5
```

---

## Project Structure

```
gopher-pipeline/
│
├── main.go        → controls the program flow
├── processor.go   → applies transformation rules
├── utils.go       → helper functions
├── summary.go     → builds final output
```

---

## How to Run

1. Open terminal
2. Navigate to project folder
3. Run:

```
go run . input.txt output.txt
```

---

## Example

### Input:

```
   hello world
TODO: fix bug
THIS IS SHOUTING
REVERSE this line
```

### Output:

```
Gopher's Sentinel Field Report - Processed
1. HELLO WORLD
2. ✦ ACTION: fix bug
3. This Is Shouting
4. line this REVERSE
```

---

## What I Learn From This Project

* File reading and writing in Go
* String manipulation
* Conditional logic
* Working with slices
* Building structured programs

---

## Author

Samuel Jireh
