package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	got := Add(2, 3)
	want := 5
	if got != want {
		t.Errorf("got: %d, want: %d", got, want)
	}
}

// godoc -http=:6060 and http://localhost:6060/pkg/learn-go/integers/ you can see the example for add function
func ExampleAdd() {
	sum := Add(3, 4)
	fmt.Println(sum)
	// Output: 7
}
