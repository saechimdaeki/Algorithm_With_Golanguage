# study_Go_And_Rust
- go언어에 갑자기 끌려버려서 공부하기위한 저장소입니다 뭔가 프언어(이문근)때 비슷한문법을 좀배워서 
- go언어를 공부해봅시다 쿄쿗! 

- go lang ver.
![go](./go.JPG)

## go언어를 실행 할때는 터미널에서 go run 파일명.go
=============================================================================================
# Rust언어 실행을 하려면 🍙
## cargo new hello_world --bin 후에 LLDB

boj를 rust로푸는 example
use std::io;

fn main() {
    let mut s = String::new();

    io::stdin().read_line(&mut s).unwrap();

    let values:Vec<i32> = s
        .as_mut_str()
        .split_whitespace()
        .map(|s| s.parse().unwrap())
        .collect();

    println!("{}", values[0] + values[1]);
}
