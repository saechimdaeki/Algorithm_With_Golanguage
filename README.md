# study_Go_And_Rust
- for study go language  
- go lang 심화튜토리얼 :http://golang.site/
- golang tutorial study link : https://go-tour-kr.appspot.com/
- go lang ver.

![go](./go.JPG)

## go언어를 실행 할때는 터미널에서 go run 파일명.go

### 왜 go언어를 해야하는가 ? 
- https://blog.jse.li/posts/torrent/
- https://blog.youngbin.xyz/2019-09-09-migrating-skhus-backend-from-nodejs-to-golang/
- https://github.com/golang/go/wiki/LearnServerProgramming
- https://www.buzzvil.com/ko/2018/02/12/tech-blog-go-%EC%84%9C%EB%B2%84-%EA%B0%9C%EB%B0%9C%ED%95%98%EA%B8%B0/

## go언어의 간단한 실용예
- http://golang.site/go/article/111-%EA%B0%84%EB%8B%A8%ED%95%9C-%EC%9B%B9-%EC%84%9C%EB%B2%84-HTTP-%EC%84%9C%EB%B2%84 (서버)
- http://pyrasis.com/book/GoForTheReallyImpatient/Unit58   (서버)

=================================================================================
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
