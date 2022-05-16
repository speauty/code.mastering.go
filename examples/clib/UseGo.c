#include <stdio.h>
#include "../cshare/usedByC.h"

// gcc -o UseGo UseGo.c ../cshare/usedByC.o
int main(int argc, char **argv) {
 GoInt x = 12;
 GoInt y = 23;
 printf("call a go func");
 PrintMsg();


 GoInt p = Multiply(x, y);
 printf("Product: %d\n", (int)p);
}