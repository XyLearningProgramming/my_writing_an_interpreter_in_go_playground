# Writing An Interpreter In Go - My improved version

## What is this repo for?

I posted my version of the `monkey` language here (codes are in folders `my_*`), developed as me navigating through the `waiig` book by Thorsten Ball mainly because I'd like to share my enthusiasm of implementing the first interpreter of a programming language myself.

The `waiig` book is an excellent material for anyone who wanna know main parts of the interpreter without digging too deep into the theories, which is just what I want. But as I copy and try the implementation of the `monkey` language, I often think of features that could have been there and could have made the language more fun, so I just created my own version as well as an list of improvements and TODOs while reading through the book. If you are reading this book too, perhaps you will find some points in my list below suitable for your own "homework assignments".

Personally,I found these tweaks quite interesting and challenging (also not over-complicated). I therefore listed them here to share what I did and will do with anyone reading this.

## Features (What's done)

> The list below is a copy of `README.md`'s from the folders of each chapter

- Chapter 02

    1. `Lexer`: parse float && int type

- Chapter 03

    1. Conversions among `boolean`, `integer`, `float` are now available for our monkey lang, especially when it's not static typed; thus, I am allowing something like `true + true`, `1==true`, etc; it's just like how `python` behaves
    2. `float` type parse and eval stage
    3. ~~positive int value evaluated as `unsigned int` if not in `let` statement and if possible (value not exceeding `math.MaxUint64` in the host language `golang`)~~ introducing an unsigned int will quickly get messy when applying infix calculation involving overflow :(

- Chapter 04

    1. A __repl__ enabling backspace, history tracing, `exit()` command
    2. String literals now can start with either ' or "
    3. String literals now support backslash escaping; for instance, literals like below will work

        ```bash
        "boo\"foo"
        'boo""foo'
        'boo\\"foo'
        '\"'
        "\n\t\r"
        ```

    4. Arrays allow python-like indexing including striding:

        ```bash
        [0,1][1:] yields [1]
        [0,1,2][::-1] yields [2,1,0]
        ```


TODOs:

- Chapter 01

    1. Parse `rune` instead of `char` each time to enable Unicode
    2. Parse `float` correctly as one token
    3. Add `filename` and `lineno` to `token` struct
    4. Let `lexer` to accept `io.Reader` and file name as input

- Chapter 02

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

- Chapter 03

    1. Infix involving `NULL` is a mess: does `NULL < 1` yield `false` or `NULL` or error?
    2. Allow `+` to be a legitimate prefix operator (should be trivial to implement)
    3. Add more functions to repl, like moving input cursor, history commands, etc
    4. All the rest left to be done in previous chapters

- Chapter 04

    1. For loop as statement? for(`initialization_statement`; `test_expression`; `update_statement`) { `block_statements` }
    2. Reassign values to identifier? `let a= 1; a="b"`
