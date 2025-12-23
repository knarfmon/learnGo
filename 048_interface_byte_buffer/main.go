package main

import (
	"bytes"
	"fmt"
)

//type person struct {
//	first string
//}
//
//func (p person) writeOut(w io.Writer) error {
//	_, err := w.Write([]byte(p.first))
//	return err
//}

func main() {
	//f, err := os.Create("output.txt")
	//if err != nil {
	//	log.Fatalf("error creating file: %s", err)
	//}
	//defer f.Close()
	//
	//s := []byte("Hello World")
	//
	//_, err = f.Write(s)
	//if err != nil {
	//	log.Fatalf("error writing to file: %s", err)
	//}

	b := bytes.NewBufferString("Hello ")
	fmt.Println(b.String())
	b.WriteString("Gophers")
	fmt.Println(b.String())
	b.Reset()
	b.WriteString("Hey, Hey, Hey")
	fmt.Println(b.String())

	b.Reset()
	b.Write([]byte("Happy, Happy"))
	fmt.Println(b.String())
}
