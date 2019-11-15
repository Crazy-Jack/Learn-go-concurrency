# Notes for *Concurency in Go* 
Document my notes for learning [*Concurrency in Go*]() by Katherine Cox-Buday


## Chapter 1: Intro to Concurrency

1. Why Concurrency is hard to model?
- Data race: when different line of code try to access the same parameter concurrently, which will introduce vast variability of the program.
  + Introduced when programmers are thinking the problem sequentially.
  + Very insidious bug for concurrency bugs.

- Atomic: Atomic operations is the one that doesn't operate on the concurrent variables in a certain context.
 + e.g. i++ is atomic in the context A if A doesn't have other concurrent process to modify value of i.
 + Atomic serves a way to optimize concurrent code.


