#ifndef CROSSPLATFORM_H_
#define CROSSPLATFORM_H_

#ifdef __cplusplus
extern "C"
{
#endif


#include <stdio.h>
#include <stddef.h>


FILE *crossplatformfmemopen (void *__s, size_t __len, const char *__modes);
// FILE *crossplatformfmemopen (void *__s, unsigned long __len, const char *__modes);
int crossplatformgetline (char **lineptr, size_t *n, FILE *stream);


#ifdef __cplusplus
}
#endif

#endif // #ifndef FMEMOPEN_H_