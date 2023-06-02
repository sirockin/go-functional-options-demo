---
marp: true
class:
- invert

---

# The Go Functional Options Constructor Pattern

### Dave Sirockin
#### 02/06/2023
---

# Some Common Constructor Patterns

- Simple Parameter List 
  - more than a few is a mess
  - defaults to nil, 
  - not easy to read
  - not extensible or maintainable
- Options Struct passed in 
  - okay but not good for handling defaults
- Builder Pattern (`NewFoo().WithX(x).WithY(y)`) 
  - can't pass out errors

---

# Functional Options Pattern Advantages

- clear defaults
- configurable
- enables simple and complex validations
- maintainable
- nils not required
- errors passed out

# References

- https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis
- https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html
