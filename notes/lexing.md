# Lexing

## 1.1 Lexical Analysis

- Lexical analysis (lexing) -> turn source code into tokens

- Done by a lexer, tokenizer, or scanner.

- Tokens: small, categorizable data structures.

- Tokens are fed into the "parser", which creates the AST (Abstract Syntax Tree).

Example:

"let x = 5 + 5;"

Gets transformed by the lexer into

```
[
    LET,
    IDENTIFIER("x"),
    EQUAL_SIGN,
    INTEGER(5),
    PLUS_SIGN,
    INTEGER(5),
    SEMICOLON
]
```

- Whitespaces in Monkey are not significant, so they don't have a token.

- Tokens vary between lexer implementations.

- A lexer can attach line number, column and filename to a token. This helps for better error output. We won't do it here.

## 1.2 Defining our tokens

- Tokens usually represent things like: 
    - numbers
    - variable names (identifiers)
    - keywords (let, fn, if)
    - special characters (parenthesis, semicolons, equal sign)

See `src/token/token.go` to see the list of tokens.