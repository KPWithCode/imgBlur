package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"github.com/KPWithCode/imgBlur/primitive"
	
)

func main() {
	cmd := exec.Command("primitive", strings.Fields("-i martin2.png -o 1on1.png -n 100 -m 6")...)
	b, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}
type PrimitiveMode 

const (
	0=combo, 1=triangle, 2=rect, 3=ellipse, 4=circle, 5=rotatedrect, 6=beziers, 7=rotatedellipse, 8=polygon

)

func primitive(inputFile,outputFile string, numShapes int, mode int) 
