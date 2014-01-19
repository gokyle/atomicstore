package atomicstore

import "bytes"
import "io/ioutil"
import "os"
import "testing"

func writeFile(t *testing.T, msg []byte) {
	err := AtomicWrite("tmp.out", msg, 0600)
	if err != nil {
		t.Fatalf("atomicstore: failed to write file (%v)", err)
	}
}

var msg = []byte("This is a test message.")

func TestAtomicWrite(t *testing.T) {
	for i := 0; i < 128; i++ {
		go writeFile(t, msg)
	}
	writeFile(t, msg)
}

func TestRead(t *testing.T) {
	for i := 0; i < 10; i++ {
		in, err := ioutil.ReadFile("tmp.out")
		if err != nil {
			t.Fatalf("atomicstore: failed to read file (%v)", err)
		} else if !bytes.Equal(in, msg) {
			t.Fatal("atomicstore: read file fails to match the message")
		}
	}
	os.Remove("tmp.out")
}
