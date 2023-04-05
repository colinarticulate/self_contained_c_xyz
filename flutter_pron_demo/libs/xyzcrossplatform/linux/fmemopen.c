#include <stdio.h>

//Linux
FILE *crossplatformfmemopen (void *__s, size_t __len, const char *__modes){
    // FILE *crossplatformfmemopen (void *__s, unsigned long __len, const char *__modes) {
    return fmemopen (__s, __len, __modes);
}