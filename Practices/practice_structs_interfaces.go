// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message

package main
import "fmt"

type Subject struct{
    name string
    marks int
}
type User struct{
    name string
    roll string
    rank int
    cgpa float64
    subject Subject
}


func main() {
//   user1:=User{
//       name:"Ashish" ,
//       roll: "2022UG1025",
//       rank: 1,
//       cgpa: 8.87,
//   }
  userArray:=[]User{}
  
  for i:=0;i<10;i++{
      tempUser:=User{
          name:"Ashish "+ string(rune(i+65)),
          roll:fmt.Sprintf("2022UG102%d",i),
          rank:i,
          cgpa:8.8+float64(i)/100,
          subject:Subject{
              name: "Maths",
              marks: 56,
          },
      }
      userArray=append(userArray,tempUser)
  }
  fmt.Println(userArray)
  
}