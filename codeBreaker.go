package main
var private_code string

func setCode(code string){
  private_code = code
}

func puntoFama(s string)string{
  aux := ""

  for i := 0; i < len(s); i++ {
    if (s[i]==private_code[i]) {
      aux+= "x"
    }
}
for i := 0; i < len(s); i++ {
  for j:=0; j < len(s); j++ {
    if j!=i {
      if (s[j]==private_code[i]) {
        aux+= "-"
        break;
      }
    }
  }
}
  return aux
}
