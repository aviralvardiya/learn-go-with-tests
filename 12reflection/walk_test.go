package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Bobby"},
			[]string{"Bobby"},
		},
		{
			"struct with two string field",
			struct {
				Name string
				City string
			}{"Bobby", "Delhi"},
			[]string{"Bobby", "Delhi"},
		},
		{
			"struct with a non string field",
			struct {
				Name string
				Age  int
			}{"Bobby", 22},
			[]string{"Bobby"},
		},
		// {
		// 	"struct with nested fields",
		// 	struct{
		// 		Name string
		// 		Profile struct{
		// 			City string
		// 			Age int
		// 		}
		// 	}{"Bobby",struct{
		// 		City string
		// 		Age int
		// 	}{"Delhi",22}},
		// 	[]string{"Bobby","Delhi"},
		// },
		{
			"struct with nested fields",
			Person{
				"Bobby",
				Profile{"Delhi", 22}},
			[]string{"Bobby", "Delhi"},
		},
		{
			"pointer to things",
			&Person{
				"Bobby",
				Profile{"Delhi", 22}},
			[]string{"Bobby", "Delhi"},
		},
		{
			"slices",
			[]Profile{
				{"Delhi", 22},
				{"Mumbai", 26},
			},
			[]string{"Delhi", "Mumbai"},
		},
		{
			"arrays",
			[2]Profile{
				{"Delhi", 22},
				{"Mumbai", 26},
			},
			[]string{"Delhi", "Mumbai"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			expected := test.ExpectedCalls
			var got []string

			x := test.Input

			Walk(x, func(intput string) {
				got = append(got, intput)
			})

			if !reflect.DeepEqual(expected, got) {
				t.Errorf("got %v want %v", got, expected)
			}

		})
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Baa",
		}

		var got []string
		Walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Moo")
		assertContains(t, got, "Baa")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{"Berlin", 32}
			aChannel <- Profile{"Katowice", 44}
			close(aChannel)
		}()

		var got []string
		want := []string{"Berlin", "Katowice"}

		Walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{"Berlin", 32}, Profile{"Katowice", 44}
		}

		var got []string
		want := []string{"Berlin", "Katowice"}

		Walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
	}
}
