# Chapter 3 TODO List

- Infix involving `NULL` is a mess: does `NULL < 1` yield `false` or `NULL` or error?
- Allow `+` to be a legitimate prefix operator (should be trivial to implement)
- Add more functions to repl, like moving input cursor, history commands, etc
- All the rest left to be done in previous chapters

### What's Done

- I think trivial conversions among `boolean`, `integer`, `float` are fine for monkey lang, especially when it's not static typed; thus, I am allowing something like `true + true`, `1==true`, etc; it's just like how `python` behaves
- `float` type parse and eval stage
- ~~positive int value evaluated as `unsigned int` if not in `let` statement and if possible (value not exceeding `math.MaxUint64` in the host language `golang`)~~ introducing an unsigned int will quickly get messy when applying infix calculation involving overflow :(
