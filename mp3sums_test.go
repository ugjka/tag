package tag

import (
	"io/ioutil"
	"os"
	"testing"
)

const MP3Sum = "16d492b55318d4cbf5615b1e014989bdfc7586b8"

func TestMP3Sums(t *testing.T) {
	files, err := ioutil.ReadDir("testsum")
	if err != nil {
		t.Error("failed to read the testfiles")
	}
	for _, v := range files {
		f, err := os.Open("testsum/" + v.Name())
		if err != nil {
			t.Error("Could not open: ", v.Name())
		}
		sum, err := Sum(f)
		if err != nil {
			t.Error("Could not sum: ", v.Name())
		}
		if sum != MP3Sum {
			t.Errorf("%s: got sum %s, expected %s", v.Name(), sum, MP3Sum)
		}
	}
}
