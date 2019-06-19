package main
import "testing"
import "reflect"
import "fmt"

func TestPrintAddition (t *testing.T)  {
	total := PrintAddition (2, 4)
	if total != 6{
		t.Errorf("Addition was incorrect, got: %d, want: %d.", total, 6)
	} else {
		fmt.Printf ("Addition is correct, 2 + 4 got %d, want %d.\n", total, 6)
	}
}

func TestPrintMultiply (t *testing.T) {
	total := PrintMulti (3, 5)
	if total != 15 {
		t.Errorf("Multiply was incorrect, got: %d, want: %d.", total, 15)	
	} else {
		fmt.Printf ("Multiply is correct, 3 x 5 got %d, want: %d.\n", total, 15)
	}
}

func TestPrintFibonacci (t *testing.T) {
	total := PrintFibonacci (5)
	//fmt.Println(total)
	expected := []int{0, 1, 1, 2, 3}
	if !reflect.DeepEqual (total, expected)  {
		t.Errorf("Fibonacci sequence is incorrect, got: %v, want %v.\n", total, expected)	
	} else {
		fmt.Printf ("Fibonacci sequence is correct, got: %v, want %v.\n", total, expected)
	}
}

func TestPrintPrime (t *testing.T) {
	total := PrintPrime (7)
	expected := []int{2, 3, 5, 7, 11, 13, 17}
	if !reflect.DeepEqual(total, expected) {
		t.Errorf ("Prime list is incorrect, got: %v, want %v.\n", total, expected)
	} else {
		fmt.Printf ("Prime sequence is correct, got: %v, want %v.\n", total, expected)
	}
}