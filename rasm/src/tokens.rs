pub enum TokenTypes {
    Instruction,
    Comma,
    Regcall,
    Immediate,
    Comment,
    Address,
    Ident,
    Label,
}

/*
let str_toktypes = vec![
    "Instruction",
    "Comma",
    "Regcall",
    "Immediate",
    "Comment",
    "Address",
];
*/

pub struct Token {
    _type: TokenTypes,
    literal: String,
    start: i32,
    end: i32,
    ln: i32,
}

/*
impl Token {
    fn len(&self) -> i32 {
        return (self.end - self.start) + 1;
    }

    /*
    fn istype(&self, _type: TokenTypes) -> bool {
        return self._type as i32 == _type as i32;
    }
    */
}
*/

impl std::fmt::Display for Token {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        let type_str: &str;
        match self._type {
            TokenTypes::Instruction => type_str = "Instruction",
            TokenTypes::Comma => type_str = "Comma",
            TokenTypes::Regcall => type_str = "Regcall",
            TokenTypes::Immediate => type_str = "Immediate",
            TokenTypes::Comment => type_str = "Comment",
            TokenTypes::Address => type_str = "Address",
            TokenTypes::Ident => type_str = "Ident",
            TokenTypes::Label => type_str = "Label",
        }

        write!(
            f,
            "(type: {}, lit: '{}', start: {}, end: {}, ln: {})",
            type_str, self.literal, self.start, self.end, self.ln
        )
    }
}

pub fn new_token(_type: TokenTypes, literal: String, start: i32, end: i32, ln: i32) -> Token {
    return Token {
        _type,
        literal,
        start,
        end,
        ln,
    };
}
