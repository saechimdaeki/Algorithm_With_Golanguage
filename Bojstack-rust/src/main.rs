#![allow(non_snake_case)]

macro_rules! parse_line {
  ($($t: ty),+) => ({
    let mut line = String::new();
    std::io::stdin().read_line(&mut line).unwrap();
    let mut iter = line.split_whitespace();
    ($(iter.next().unwrap().parse::<$t>().unwrap()),+)
  })
}

fn main() {
  let mut v = Vec::new();
  let n = parse_line!(usize);
  for _ in 0 .. n {
    let mut line = String::new();
    std::io::stdin().read_line(&mut line).unwrap();
    let mut tokens = line.split_whitespace();
    match tokens.next().unwrap() {
      "push" => v.push(tokens.next().unwrap().parse::<i32>().unwrap()),
      "pop"  => println!("{}", v.pop().unwrap_or(-1)),
      "size" => println!("{}", v.len()),
      "empty"=> println!("{}", if v.is_empty() {1} else {0}),
      "top"  => println!("{}", v.last().unwrap_or(&-1)),
      _ => (),
    }
  }
}


/*
    study용으로 참고하겠습니다. bivoje49님 
*/