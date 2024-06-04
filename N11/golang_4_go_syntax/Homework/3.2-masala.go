package main

func main() {
    num1 := 45
    num2 := 12
    num3 := 78

    if num1 < num2 {
        if num2 < num3 {
            println(num1)
            println(num2)
            println(num3)
        } else {
            if num1 < num3 {
                println(num1)
                println(num3)
                println(num2)
            } else {
                println(num3)
                println(num1)
                println(num2)
            }
        }
    } else {
        if num1 < num3 {
            println(num2)
            println(num1)
            println(num3)
        } else {
            if num2 < num3 {
                println(num2)
                println(num3)
                println(num1)
            } else {
                println(num3)
                println(num2)
                println(num1)
            }
        }
    }
}
