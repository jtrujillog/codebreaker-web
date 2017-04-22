package main

import "testing"

func Test4x(t *testing.T){
    setCode("1234")
  expected := "xxxx"
  actual := puntoFama("1234")
  if actual != expected{
    t.Errorf("Test failed, expected:  '%s' ,got: '%s'", expected, actual )
  }
}

func Test3x(t *testing.T){
    setCode("1234")
  expected := "xxx"
  actual := puntoFama("1238")
  if actual != expected{
    t.Errorf("Test failed, expected:  '%s' ,got: '%s'", expected, actual )
  }
}

func Test2x(t *testing.T){
    setCode("1234")
  expected := "xx"
  actual := puntoFama("1258")
  if actual != expected{
    t.Errorf("Test failed, expected:  '%s' ,got: '%s'", expected, actual )
  }
}

func Test1x(t *testing.T){
    setCode("1234")
  expected := "x"
  actual := puntoFama("1658")
  if actual != expected{
    t.Errorf("Test failed, expected:  '%s' ,got: '%s'", expected, actual )
  }
}

func TestVacio(t *testing.T){
    setCode("1234")
  expected := ""
  actual := puntoFama("0658")
  if actual != expected{
    t.Errorf("Test failed, expected:  '%s' ,got: '%s'", expected, actual )
  }
}

func Test1x1_(t *testing.T){
    setCode("1234")
  expected := "x-"
  actual := puntoFama("1458")
  if actual != expected{
    t.Errorf("Test failed, expected:  '%s' ,got: '%s'", expected, actual )
  }
}
func Test1x2_(t *testing.T){
    setCode("1234")
  expected := "x--"
  actual := puntoFama("1348")
  if actual != expected{
    t.Errorf("Test failed, expected:  '%s' ,got: '%s'", expected, actual )
  }
}
func Test1x3_(t *testing.T){
    setCode("1234")
  expected := "x---"
  actual := puntoFama("1342")
  if actual != expected{
    t.Errorf("Test failed, expected:  '%s' ,got: '%s'", expected, actual )
  }
}
func Test4_(t *testing.T){
    setCode("1234")
  expected := "----"
  actual := puntoFama("4321")
  if actual != expected{
    t.Errorf("Test failed, expected:  '%s' ,got: '%s'", expected, actual )
  }
}
func Test2x1_(t *testing.T){
  setCode("1234")
  expected := "xx-"
  actual := puntoFama("1249")
  if actual != expected{
    t.Errorf("Test failed, expected:  '%s' ,got: '%s'", expected, actual )
  }
}
func Test2x2_(t *testing.T){
  setCode("1234")
  expected := "xx--"
  actual := puntoFama("1324")
  if actual != expected{
    t.Errorf("Test failed, expected:  '%s' ,got: '%s'", expected, actual )
  }
}
