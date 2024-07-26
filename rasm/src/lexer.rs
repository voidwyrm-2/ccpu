use crate::tokens::{new_token, Token, TokenTypes};

fn is_letter(c: char) -> bool {
    return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z');
}

fn is_digit(c: char) -> bool {
    return c >= '0' && c <= '9';
}

/*
impl is_letter for char {
    fn is_letter(&self) -> bool {
        return *self >= 'a' || *self <= 'z' || *self >= 'A' || *self <= 'Z'
    }
}
*/

pub struct LexerResult {
    pub tokens: Vec<Token>,
    pub err: String,
}

pub struct Lexer {
    text: String,
    cchar: char,
    idx: i32,
    ln: i32,
}

impl Lexer {
    fn advance(&mut self) {
        self.idx += 1;
        if self.idx < self.text.len() as i32 {
            self.cchar = self.text.as_bytes()[self.idx as usize] as char;
        } else {
            self.cchar = '\0';
        }
        if self.cchar == '\n' {
            self.ln += 1
        }
    }

    pub fn lex(&mut self) -> LexerResult {
        let mut tokens: Vec<Token> = vec![];

        while self.cchar != '\0' {
            //println!("{}: '{}'", self.idx, self.cchar);
            match self.cchar {
                ',' => {
                    tokens.push(new_token(
                        TokenTypes::Comma,
                        ",".to_string(),
                        self.idx,
                        self.idx,
                        self.ln,
                    ));
                    self.advance();
                }
                '.' => {}
                _ => {
                    if is_letter(self.cchar) {
                        //println!("'{}' is a letter", self.cchar);
                        //self.advance();
                        tokens.push(self.collect_ident());
                    } else if is_digit(self.cchar) {
                        println!("'{}' is a digit", self.cchar);
                        self.advance();
                    } else if self.cchar == ' ' || self.cchar == '\n' || self.cchar == '\t' {
                        self.advance();
                    } else {
                        return LexerResult {
                            tokens: vec![],
                            err: format!(
                                "error on line {}: illegal character '{}'",
                                self.ln + 1,
                                self.cchar
                            ),
                        };
                    }
                }
            }
        }

        return LexerResult {
            tokens,
            err: "\0".to_string(),
        };
    }

    fn collect_ident(&mut self) -> Token {
        let start = self.idx;
        let mut ident_str = "".to_string();

        while self.cchar != '\0' && is_letter(self.cchar) {
            ident_str.push(self.cchar);
            self.advance();
        }

        let mut tt = TokenTypes::Ident;

        if self.cchar == ':' {
            self.advance();
            tt = TokenTypes::Label;
        }

        return new_token(tt, ident_str, start, self.idx - 1, self.ln);
    }
}

pub fn new_lexer(text: String) -> Lexer {
    let mut l = Lexer {
        text,
        cchar: '\0',
        idx: -1,
        ln: 0,
    };

    l.advance();

    return l;
}
