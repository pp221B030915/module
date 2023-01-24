package main 
import "fmt"
type Student struct{
	courses [10]string 
}
func main(){
	var st Student
	var com int
	for{
		fmt.Scan(&com)
		if com == 1{
			var course string 
			var num int
			fmt.Scan(&course)
			fmt.Scan(&num)
			st.courses[num] = course
			fmt.Println(st.courses) 
		}
		if com == 2{
			break
		}
}
}