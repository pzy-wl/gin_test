package main

/*
#include <stdio.h>

void printint(int v) {
    printf("printint: %d\n", v);
}
int sum(int a,int b ){
	return a+b;
}
*/
import "C"

func main() {
	v := 163
	C.printint(C.int(v))
	C.printint(C.sum(C.int(5), C.int(6)))
}
