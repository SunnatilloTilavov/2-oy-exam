package main

import "fmt"

func main(){
	var str string
	var strArray [3]string
	for i:=0;i<3;i++{
	fmt.Println("so'zni kiriting")
	fmt.Scan(&str)
	strArray[i]=str
	}

	str=(Prefix(strArray))
	fmt.Println("Prefix:",str)
}

func Prefix(strArray [3]string)string{
	runes1 := []rune(strArray[0])
	runes2 := []rune(strArray[1])
	runes3 := []rune(strArray[2])

	minLength := len(runes1)
    if len(runes2) < minLength {
        minLength = len(runes2)
    }
    if len(runes3) < minLength {
        minLength = len(runes3)
    }

	var b,c string
	d:=""
	a:=""

	for i:=0;i<minLength;i++{
		a+=string(runes1[i])
		b+=string(runes2[i])
		c+=string(runes3[i])

		if runes1[0]!=runes2[0]||runes2[0]!=runes3[0]{
			break			
		}

		if a==b&&b==c{
			d=a
		}else {
			break
		}
	} 
	return d
}