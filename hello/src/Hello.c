#include "Hello.h"
#include "stdio.h"

int HelloYou(char* name, char** dest)
{
    int n = asprintf(dest, "Hello %s", name);
    return n;
}