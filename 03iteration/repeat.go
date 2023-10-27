package iteration

// import "fmt"

func Repeat(character string,j int) string {
	var repeated string
	var iterations int=5
	if(j!=0){
		iterations=j
	}
	for i := 0; i < iterations; i++ {
		repeated = repeated + character
	}
	return repeated
}