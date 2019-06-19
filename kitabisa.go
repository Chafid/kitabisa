package main

import "os"
import "fmt"
import "strconv"
import "math"

func PrintFibonacci (noOfSeq int) []int {
	var fibonacciList []int
	for i := 0; i < noOfSeq; i++ {
		fib := FibonacciRecursion(i)
		fibonacciList = append (fibonacciList, fib)
	}
	return fibonacciList
}

func PrintList (numberList []int) {
	for i:= 0; i < len (numberList); i++ {
		fmt.Print(numberList[i])
		fmt.Print(" ")
	}
}

func PrintPrime (noOfSeq int) []int {
	var primeList []int
	for i := 0; len(primeList) < noOfSeq; i++ {
		if IsPrimeSqrt(i) {
			primeList = append (primeList, i)
		} 		
	}			
	return primeList		
}

func PrintMulti (firstArg, secondArg int) int {
	var multiResult int
	multiResult = firstArg * secondArg
	return multiResult
	//fmt.Println (multiResult)	
} 

func PrintAddition( firstArg, secondArg int) int {
	var addResult int
	addResult = firstArg + secondArg
	return addResult
}

func main () {
	argCount := len (os.Args[1:])
	var calcResult int
	var resultList []int
	if argCount == 2  {
		if noOfSeq, err:= strconv.Atoi(os.Args[1]); err == nil {
			if os.Args[2] == "p" {
				fmt.Println ("Run prime function") 
				resultList = PrintPrime(noOfSeq)
				PrintList(resultList)
			} else if os.Args[2] == "f" {
				fmt.Println ("Run Fibonacci function")
				resultList = PrintFibonacci(noOfSeq)
				PrintList(resultList)
			} else {
				fmt.Println ("Operation must be f (Fibonnaci) or p (Prime) ")	
			}
		} else {
			fmt.Println ("First argument must be number")
		} 

	} else if argCount == 3 {
		firstArg, firstErr := strconv.Atoi(os.Args[1])
		secondArg, secondErr := strconv.Atoi(os.Args[2])  
		if  firstErr == nil && secondErr == nil  {
			if os.Args[3] == "*" {
				calcResult = PrintMulti(firstArg, secondArg)	
				fmt.Println ("Run multiplication function")
				fmt.Println (calcResult)
			} else if os.Args[3] == "+" {
				fmt.Println ("Run addition function")
				calcResult = firstArg + secondArg
				fmt.Println (calcResult)
			} else {
				fmt.Println ("Operator must be * or +")	
			}
		} else {
			fmt.Println ("First two arguments must be numbers")
		}
	} else {
		fmt.Println ("Not enough argument")
		fmt.Println ("Usage: ")
		fmt.Println ("kitabisa a b + for addition")
		fmt.Println ("kitabisa a b * for multiplication")
		fmt.Println ("kitabisa a p for prime numbers")
		fmt.Println ("kitabisa a f for Fibonacci sequence")
	}
}


func FibonacciRecursion(n int) int {
    if n <= 1 {
        return n
    }
    return FibonacciRecursion(n-1) + FibonacciRecursion(n-2)
}

func IsPrimeSqrt(value int) bool {
    for i := 2; i <= int(math.Floor(math.Sqrt(float64(value)))); i++ {
        if value%i == 0 {
            return false
        }
    }
    return value > 1
}