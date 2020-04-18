package main

import (
	"errors"
	"fmt"
	"github.com/pluralsight-fundamentals/webservice/controllers"
	"github.com/pluralsight-fundamentals/webservice/models"
	"net/http"
	"reflect"
	"runtime"
)

type HTTPRequest struct {
	Method string
}

//const (
//	//pi = 3.1412
//	first= iota+6
//	second = 7<<iota //7*2
//)
//type User struct {
//	ID int
//	name string
//	address string
//}
func main() {
	/*
		fmt.Println("First Way to use pointer")
		fmt.Println("Hello World")
		// we need to initalize a pointer otherwise get an error
		var firstName *string = new(string)
		*firstName = "Shivani Sharma" // De-referencing
		fmt.Println(firstName)        // return the address
		fmt.Println(*firstName)

		fmt.Println("Another way")

		firstName1 := "Arthur"
		fmt.Println(firstName1)
		ptr := &firstName1
		fmt.Println(ptr)  // address of first_name1
		fmt.Println(*ptr) // value of ptr using de-referencing operator
		firstName1 = "Sharma"
		fmt.Println(ptr, *ptr)
		var i = complex(4, 4)
		re, ima := real(i), imag(i)
		fmt.Println(i, re, ima)

		fmt.Println(pi+3,pi+2)
		fmt.Println(pi+3)*/

	// IOTA
	//fmt.Println(first,second)

	// Collections

	/*	var arr [3]int
		arr[0] = 5
		arr[1] = 2
		arr[2] = 35

		// pre-initialized
		arr1:= [5]int{4,6,4,5}
		fmt.Println(arr,arr1)
		sli:= arr1[:]
		fmt.Println(sli)
		sli[2]=9
		sli[3]=56
		fmt.Println(sli,arr1)
		sli1 := []int{2,6,7,7,4,8}
		fmt.Println(sli1)
		sli1= append(sli1,42)
		fmt.Println(sli1)*/

	// Map
	/*	m:= map[string]int{"jun":54}
		m["Hello"] = 7
		m["Hello2"] = 754
		delete(m,"jun")// deleted by key
		fmt.Println(m)
		u := models.User{
			ID: 1,
			FirstName: "Shivani",
			LastName: "Sharma",
		}
		fmt.Println(u)
		port:= 3000
		port,err := webserver(port)
		fmt.Println(err,port)*/

	/*	controllers.RegisterControllers()
		http.ListenAndServe(":3000",nil)*/

	//var i int
	/*	for i < 5 {
			fmt.Println("Value : ",i)
			i++
			if i == 3{
				fmt.Println("Continue ................... ")
				continue
			}
			if i==4 {
				fmt.Println("Breaking ...................... ")
				break
			}
		}
	*/
	/*	for ;i<5;i++ {
		fmt.Println("Value : ",i)
	}*/
	//sli := []int{1, 3, 5, 7, 9, 11, 13, 23, 33, 43, 19}
	/*	for i:=0;i<len(sli);i++ {
		fmt.Println("Value Of Slice : ",sli[i])
	}*/
	/*	for i, v := range sli {
			fmt.Println(i, v)
		}
		port_Def := map[string]int {"Port_1":9090,"Port_2":8901,"Port_3":1234,"Port_4":3456}
		for k,v := range port_Def {
			fmt.Println(k,v)
		}*/
	//panic("Something Bad Happened !")

	u1 := models.User{
		ID:        3,
		FirstName: "Shivani",
		LastName:  "Sharma",
	}
	u2 := models.User{
		ID:        4,
		FirstName: "Shivani",
		LastName:  "Sharma",
	}
	fmt.Println("OPERATING SYSTEM :",runtime.GOOS)
	if u1.ID == u2.ID {
		fmt.Println("Same User Ids",reflect.TypeOf(u1))
	} else {
		fmt.Println("Different User Ids")
	}

	r := HTTPRequest{Method: "GET"}
	switch r.Method {
	case "GET":
		fmt.Println("Get Request")
		break
	case "POST":
		fmt.Println("Post Request")
		break
	case "PUT":
		fmt.Println("Put Request")
		break
	case "DELETE":
		fmt.Println("Delete Request")
		break
	default:
		fmt.Println("no request")
	}

	controllers.RegisterControllers()
	_ = http.ListenAndServe(":3000", nil)
}

func webserver(port int) (int, error) {
	fmt.Println(`Starting Server with port:`, port, "And currently processing")
	fmt.Println("Server Started _-_-_-_-_-_-_-_-")
	return port, errors.New("Something got wrong while starting server")
}
