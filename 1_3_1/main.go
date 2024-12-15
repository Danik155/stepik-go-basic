// package main

// import "fmt"

// func main() {
// 	var number int
// 	_, _ = fmt.Scan(&number)
// 	fmt.Println(number)
// 	number--
// 	fmt.Println(number)
// 	number--
// 	fmt.Println(number)
// }

// package main

// import "fmt"

// func main() {
// 	var a, b, c, d int
// 	_, _ = fmt.Scan(&a)
// 	_, _ = fmt.Scan(&b)
// 	_, _ = fmt.Scan(&c)
// 	_, _ = fmt.Scan(&d)
// 	fmt.Println(a + b + c + d)
// }

// package main

// import "fmt"

// func main() {
// 	var a, b, n int
// 	_, _ = fmt.Scan(&a)
// 	_, _ = fmt.Scan(&b)
// 	_, _ = fmt.Scan(&n)
// 	fmt.Println(a*n+b*n/100, b*n%100)
// }

// package main

// import "fmt"

// func main() {
// 	var a int
// 	_, _ = fmt.Scan(&a)
// 	fmt.Println("Следующее за числом", a, "число:", a+1)
// 	fmt.Println("Для числа", a, "предыдущее число:", a-1)
// }

// package main

// import "fmt"

// func main() {
// 	var a int
// 	_, _ = fmt.Scan(&a)
// 	fmt.Println(a / 100 + a % 100 / 10 * 10 + a % 10 * 100)
// }

// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	var a, b, c, d float64
// 	_, _ = fmt.Scan(&a, &b, &c, &d)
// 	fmt.Println(math.Abs(a-c) + math.Abs(b-d))
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var a, b int
// 	_, _ = fmt.Scan(&a, &b)
// 	if a%b == 0 {
// 		fmt.Println("YES")
// 	} else {
// 		fmt.Println("NO")
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	var a int
// 	_, _ = fmt.Scan(&a)
// 	if a/100 < a%100/10 && a%100/10 < a%10 {
// 		fmt.Println("YES")
// 	} else {
// 		fmt.Println("NO")
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	var a int
// 	var b string
// 	_, _ = fmt.Scan(&a, &b)
// 	if a >= 12 && a <= 18 && b == "m" {
// 		fmt.Println("YES")
// 	} else {
// 		fmt.Println("NO")
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	var a int
// 	_, _ = fmt.Scan(&a)
// 	i := 1
// 	for i <= a {
// 		fmt.Println(i)
// 		i *= 2
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	var a int
// 	_, _ = fmt.Scan(&a)
// 	s := 0
// 	for a != 0 {
// 		if a%2 == 0 && a%3 != 0 {
// 			s += a
// 		}
// 	}
// 	fmt.Println(a)
// }

// package main

// import "fmt"

// func main() {
// 	var a float64
// 	fmt.Scan(&a)
// 	m := a
// 	b := 0
// 	for a != 0 {
// 		if a > m {
// 			m = a
// 			b++
// 		}
// 		fmt.Scan(&a)
// 	}
// 	fmt.Println(b)
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var a, b, c, d, e int
// 	_, _ = fmt.Scan(&a, &b, &c, &d, &e)
// 	sum := 0
// 	for i := 0; i <= 1000; i++ {
// 		if i-e == 0 {
// 			continue
// 		}
// 		if (a*i*i*i+b*i*i+c*i+d)/(i-e) == 0 {
// 			sum++
// 		}
// 	}
// 	fmt.Println(sum)
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var a, m1, m2, t int
// 	fmt.Scan(&m1, &t)
// 	if t > m1 {
// 		m2 = m1
// 		m1 = t
// 	} else {
// 		m2 = t
// 	}
// 	for {
// 		fmt.Scan(&a)
// 		if a == 0 {
// 			break
// 		}
// 		if a > m1 {
// 			m2 = m1
// 			m1 = a
// 		} else if a > m2 {
// 			m2 = a
// 		}
// 	}
// 	fmt.Println(m2)
// }

// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	var a, b, c, x1, x2 float64
// 	_, _ = fmt.Scan(&a, &b, &c)
// 	d := b*b - 4*a*c
// 	if d > 0 {
// 		if (-b+math.Sqrt(d))/(2*a) > (-b-math.Sqrt(d))/(2*a) {
// 			fmt.Println((-b - math.Sqrt(d)) / (2 * a))
// 			fmt.Println((-b + math.Sqrt(d)) / (2 * a))
// 		} else {
// 			fmt.Println((-b + math.Sqrt(d)) / (2 * a))
// 			fmt.Println((-b - math.Sqrt(d)) / (2 * a))
// 		}
// 	} else if d == 0 {
// 		fmt.Println((-b) / (2 * a))
// 	}
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var a, b int
// 	_, _ = fmt.Scan(&a)
// 	s := 0
// 	for i := 0; i < a; i++ {
// 		_, _ = fmt.Scan(&b)
// 		if b%10 == 0 {
// 			s += b
// 		}
// 	}
// 	fmt.Println(s)
// }

// package main

// import "fmt"

// func main() {
// 	var l, a int
// 	fmt.Scan(&l)
// 	s := make([]int, l)
// 	for i := 0; i < l; i++ {
// 		fmt.Scan(&a)
// 		s[i] = a
// 	}
// 	var t int
// 	for i := 0; i < l; i++ {
// 		if i == l {
// 			s[0] = t
// 		} else {
// 			t = s[i+1]
// 			s[i+1] = t
// 		}
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	var a, t int
// 	_, _ = fmt.Scan(&a)
// 	array := make([]int, a)
// 	for i := 0; i < a; i++ {
// 		fmt.Scan(&t)
// 		array[i] = t
// 	}
// 	var min int
// 	for i, v := range array {
// 		if i == 0 {
// 			min = v
// 		} else if v < min {
// 			min = v
// 		}
// 	}
// 	for i := range array {
// 		array[i] -= min
// 	}
// 	for _, v := range array {
// 		fmt.Print(v, " ")
// 	}
// }

// package main

// import "fmt"

// func main() {

//     var n int
//     _, _ = fmt.Scan(&n)

//     array := make([][]int, n)
//     for i, _ := range array {
//         array[i] = make([]int, n)
//     }

//     for i, _ := range array {
//         for ii, _ := range array[i] {
//             if i + ii == len(array) - 1 {
//                 array[i][ii] = 1
//             } else if i + ii < len(array) - 1 {
//                 array[i][ii] = 0
//             } else if i + ii > len(array) - 1 {
//                 array[i][ii] = 2
//             }
//         }
//     }

//     for i, _ := range array {
//         for _, vv := range array[i] {
//             fmt.Print(vv, " ")
//         }
//         fmt.Println()
//     }

// }

// package main

// import (
// 	"fmt"
// )

// func main() {

// 	var s string

// 	_, _ = fmt.Scan(&s)

// 	f := 0

// 	for i := 0; i < len(s); i++ {
// 		if s[i] >= 'A' && s[i] <= 'Z' || s[i] >= 'a' && s[i] <= 'z' || s[i] >= '0' && s[i] <= '9' || s[i] == '_' {
//             fmt.Println("CONTINUE YES YES YES")
// 			continue
// 		} else {
// 			f = 1
// 		}

// 		if f == 0 {
// 			fmt.Println("YES")
// 		} else {
// 			fmt.Println("NO")
// 		}
// 	}

// }

package main

import (
	"fmt"
)

// BEGIN (write your solution here)
func DomainForLocale(domain, locale string) string {
	if domain == "" {
		return fmt.Sprintf("en.%s", domain)
	} else {
		return fmt.Sprintf("%s.%s", locale, domain)
	}
}

func main() {
	fmt.Println(DomainForLocale("1", "2"))
}
