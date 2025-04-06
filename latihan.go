package main
import "fmt"

func pangkatModularTerbalik(a, b int)int{
	var hasil, i, j, a1 int
	j = b
	a1= 1
	hasil = 0
	for b > 0{
		a1 = a1*a
		b--
	}
	for a1 > 0{
		i = a1%10
		hasil = (hasil*10)+i
		a1 = a1/10
	}
	return hasil%j
}
func main(){
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Print(pangkatModularTerbalik(a,b))
}