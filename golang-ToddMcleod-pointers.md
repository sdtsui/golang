Pointers Advice, by Todd Mcleod
https://www.udemy.com/learn-how-to-code/?couponCode=letsgo
https://www.youtube.com/watch?v=38OM9V-2clg


https://blog.golang.org/index
  maps
  strings
  delcaration
  syntax

*val - 
&val - address

constant: 
  const p string = 

  iota: "within a constant declaration, successive untype integer constants"

memory address:
  & Where do you live?
  - put into location

  fmt.Scan()
  fmt.Scan(&a) - can be used w/ memory address

  Sprint - prints to a string
  
reference:
  * something more

  c *int -> it's a pointer to an int

dereferencing:
  var b *int = &a
  b //0x20818a220 
  *b //43

  a := 43 //shorthand
  fmt.Println(a)
  fmt.Println(&a)

  var b *int = &a 
  //typed delcaration
  fmt.Println(b) //ref
  fmt.Println(*b)//dref

  *b = 42
  //now it's 42