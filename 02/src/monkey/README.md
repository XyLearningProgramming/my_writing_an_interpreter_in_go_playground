# Chapter 2 TODO list

1. More operators (eg. bitwise ops)?
    - prefix ops: `~NUMBER`
    - infix ops: `NUMBER & NUMBER`, `|`, `^` (XOR) etc
    - postfix ops: `<IDENT|NUMBER>++` && `<IDENT|NUMBER>--`

2. "Plus Equals" sign as an attribute statement? `+= `&& `-=`
    - Structure: `<IDENT> += <EXPR>`

3. Inline if sentence as an **expression**? How to implement? 
    - Example: `1?true:-1`
    - Structure: `<EXPR> ? <EXPR> : <EXPR>`

4. Arrow function as  **function**  **expression**?
    - `(<EXPR>, <EXPR>, ...) => <EXPR>`
    - The main difficulty is that :
        1. brackets `()` now has two meanings, maybe we need to create a new type of NODE as interface for brackets;
        2. NOT ONLY `=>` as an operator should check its left value, BUT ALSO normal expression should check the legitimacy of `()` expression;
    - Maybe it's over-complicated under current Pratt parsing strategy

5. Function name after `fn`?

Done:
    - `Lexer`: parse float && int type
