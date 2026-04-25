package main

import "fmt"

func main (){
	fmt.Println("==== CALCULADORA CIENTÍFICA v1.0 ====")
	var a float64
    fmt.Println("Ingrese el primer numero:")
	fmt.Scanln( &a)
	var b float64
    fmt.Println("Ingrese el segundo numero:")
	fmt.Scanln( &b)
	var ope string
    fmt.Println("Ingresa la operación (+, -, *, /,^,!):")
	fmt.Scanln( &ope)
	
	
	switch ope {
	case "+":
		fmt.Println("SUMA")
		var su float64 = a + b
		fmt.Println("Resultado",a, ope ,b,"=",su,) 
	case "-":
		fmt.Println("RESTA")
		var re float64 = a - b
		fmt.Println("Resultado",a,ope,b,"=",re,)
	case "*":
		fmt.Println("MULTIPLICACION")
		var mul float64 = a * b
		fmt.Println("Resultado",a,ope,b,"=",mul,)
	case "/":
		fmt.Println("DIVISION")
		var div float64 = a/b
		fmt.Printf(`"Resultado =" %.2f %s %.2f: %.2f`, a , ope, b ,div)
	case "^":
		fmt.Println("POTENCIA")
		var pot float64 = 1
		for i := 0; i < int(b); i++ {
			pot *= a
		}
		fmt.Println("Resultado", a, ope, b, "=", pot)
	case "!":
		fmt.Println("FACTORIAL")
		var fact float64 = 1	
		for i := 1; i <= int(a); i++ {
			fact *= float64(i)
		}	
		fmt.Println("Resultado", a, ope, "=", fact)

	default:
		fmt.Println("WEON")
	}

	

    
	




}