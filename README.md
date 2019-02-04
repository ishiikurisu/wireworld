# Wireworld

Library for Wireworld operations.

## About Wireworld

From [Wikipedia](https://en.wikipedia.org/wiki/Wireworld):

A Wireworld cell can be in one of four different states, usually numbered 0–3 in software:

0. empty
1. electron head (a.k.a. positive)
2. electron tail (a.k.a. negative)
3. conductor

As in all cellular automata, time proceeds in discrete steps called generations (sometimes "gens" or "ticks"). Cells behave as follows:

- empty - empty,
- electron head - electron tail,
- electron tail - conductor,
- conductor - electron head if exactly one or two of the neighbouring cells are electron heads, otherwise remains conductor.