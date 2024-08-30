# Why Write Go?

## Speed Demon?

### Go is executes faster than...

- JS
- Python
- Ruby
- PHP

### Go compiles faster than...

- Rust
- C
- C++
- Java
- C#

## Interpreted vs VM vs Compiled

### From Slowest to Fastest (By Category)

- Interpreted

  - Python
  - Ruby
  - PHP
  - JS

- VM

  - Java
  - C#

- Compiled

  - Go
  - Rust
  - C
  - C++

- Go is a compiled language, but its runtime speed is more comparable to languages like Java and C#
- However, Go tends to use much less memory than Java and C#
  - because there isn't a need for a VM
- The primary reason for this?
  - the Go runtime

### What is the Go runtime?

- The Go runtime manages memory allocation and garbage collection

### Long Compile Times are really terrible to work with

- large projects ca n take minutes or hours to compile in languages like java
- go compiles in seconds
- https://xkcd.com/303/

### Ok, We've mentioned "Compilation" a lot...

### But _what is compilation_?

- we write code in a way that we can actually read and understand
- that's nice and all, but computers don't speak that language

### Compilation: Simplified

- takes our human readable code...
- and makes it machine readable
- it also allows us to run our code on any machine that can execute it
- (languages like python need the interpreter installed in order to run)

## Distributing interpreted code can be difficult

- requires the recipient to have dependencies installed
- also, distributing the code effectively makes it open source

## Go doesn't have this problem

- compile your code and ship it out

## What's Your Type?

### Static vs Dynamic Typing

- Static
  - if its declared as a bool, its always a bool
- Dynamic
  - oops! I changed my bool to a string
    - wait, now its an int
      - now its a float?

## How do we quantify speed?

### Computation Speed

- how fast can we do computations?

### Memory

- how bloated is the application?
