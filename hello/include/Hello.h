#ifndef HEADER_HELLO_H
#define HEADER_HELLO_H

// prevent mangling
#ifdef __cplusplus
extern "C" {
#endif

int HelloYou(char* name, char** dest);

#ifdef __cplusplus
}
#endif

#endif