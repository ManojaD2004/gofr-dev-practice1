package filter

import (
	t "github.com/ManojaD2004/types"
)

func HighMarks (oldD, newD *[]t.StudentType){
	for i := 0; i < len(*oldD); i++ {
		if !(*oldD)[i].Validate() {
			continue
		}
		a := true
		d := (*oldD)[i]
		for i2 := 0; i2 < len(d.Marks); i2++ {
			a = a && gt(d.Marks[i2].Mark, 75) && lt(d.Marks[i2].Mark, 100)
		}
		a = a && gt(d.Rollno, 50) && lt(d.Rollno, 75)
		if a {
			*newD = append(*newD, d)
		}
		// Your extra conversion logic goes here
	}
}
