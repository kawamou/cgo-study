package main

/*
#include <stdio.h>

void my_print(void)
{
    printf("Hello, World!\n");
}
*/
import "C"

func main() {
	C.my_print()
}
