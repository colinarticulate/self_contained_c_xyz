#include <stdlib.h>
#include <stdio.h>

int crossplatformgetline (char **lineptr, size_t *n, FILE *stream) {

    return getline(lineptr, n, stream);
}