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
