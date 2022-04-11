# Chapter 4 TODO list

- For loop as statement? for(`initialization_statement`; `test_expression`; `update_statement`) { `block_statements` }
- Reassign values to identifier? `let a= 1; a="b"`

### What's Done

- String literals now can start with either ' or "
- String literals now support backslash escaping; for instance, literals like below will work

```bash
"boo\"foo"
'boo""foo'
'boo\\"foo'
'\"'
"\n\t\r"
```

- Arrays allow python-like indexing including striding:

```bash
[0,1][1:] yields [1]
[0,1,2][::-1] yields [2,1,0]
```
