# study_Go_And_Rust
- for study go language  
- go lang 심화튜토리얼 :http://golang.site/
- golang tutorial study link : https://go-tour-kr.appspot.com/
- go lang ver.

![go](./go.JPG)

## go언어를 실행 할때는 터미널에서 go run 파일명.go

=======================================================================================
# Rust언어 실행을 하려면 🍙
## cargo new hello_world --bin 후에 LLDB

### boj를 rust로푸는 example 문제: a+b
<code>


    use std::io;

    fn main() {

    let mut s = String::new();

    io::stdin().read_line(&mut s).unwrap();
    let values:Vec<i32> = s
        .as_mut_str()
        .split_whitespace()
        .map(|s| s.parse().unwrap())
        .collect();
    println!("{}", values[0] + values[1]);  }

</code>
