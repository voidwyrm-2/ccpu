use lexer::new_lexer;

mod lexer;
mod tokens;

fn main() {
    let mut lexer = new_lexer(",,, I am going to freaking kill the borrow: checker:".to_string());

    let lexer_res = lexer.lex();

    if lexer_res.err != "\0".to_string() {
        println!("{}", lexer_res.err);
        return;
    }

    for token in lexer_res.tokens {
        println!("{}", token);
    }
}
