#ifndef FMEMOPEN_H_
#define FMEMOPEN_H_

#ifdef __cplusplus
extern "C"
{
#endif


#include <stdio.h>
#include <stddef.h>


FILE *crossplatformfmemopen (void *__s, size_t __len, const char *__modes);
// FILE *crossplatformfmemopen (void *__s, unsigned long __len, const char *__modes);


#ifdef __cplusplus
}
#endif

#endif // #ifndef FMEMOPEN_H_