#include <stdio.h>

//Linux
FILE *crossplatformfmemopen (void *__s, size_t __len, const char *__modes){
    // FILE *crossplatformfmemopen (void *__s, unsigned long __len, const char *__modes) {
    return fmemopen (__s, __len, __modes);
}


// #include <stdio.h>
// #include <stdlib.h>
// #include <string.h>
// #include <errno.h>

// #define MIN(a, b) ((a) < (b) ? (a) : (b))

// typedef struct {
//     char* buf;
//     size_t size;
//     size_t pos;
//     int mode;
// } MemFile;

// static int mem_flush(FILE* fp) {
//     return 0;
// }

// static int mem_close(FILE* fp) {
//     MemFile* mf = (MemFile*) fp;

//     if (mf == NULL || mf->buf == NULL) {
//         return -1;
//     }

//     free(mf->buf);
//     free(mf);
//     return 0;
// }

// static int mem_read(void* ptr, size_t size, size_t count, FILE* fp) {
//     MemFile* mf = (MemFile*) fp;
//     size_t bytes = MIN(count * size, mf->size - mf->pos);

//     if (bytes == 0) {
//         return 0;
//     }

//     memcpy(ptr, mf->buf + mf->pos, bytes);
//     mf->pos += bytes;
//     return bytes / size;
// }

// static int mem_write(const void* ptr, size_t size, size_t count, FILE* fp) {
//     MemFile* mf = (MemFile*) fp;
//     size_t bytes = count * size;

//     if (mf->mode == 'r') {
//         errno = EBADF;
//         return -1;
//     }

//     if (mf->pos + bytes > mf->size) {
//         size_t new_size = mf->pos + bytes;
//         char* new_buf = (char*) realloc(mf->buf, new_size);

//         if (new_buf == NULL) {
//             errno = ENOMEM;
//             return -1;
//         }

//         mf->buf = new_buf;
//         mf->size = new_size;
//     }

//     memcpy(mf->buf + mf->pos, ptr, bytes);
//     mf->pos += bytes;
//     return bytes / size;
// }

// FILE* crossplatformfmemopen(void* buf, size_t size, const char* mode) {
//     printf("DEBUG: inside fmemopen");
//     MemFile* mf = (MemFile*) malloc(sizeof(MemFile));
//     printf("DEBUG: after MemFile");
//     if (mf == NULL) {
//         return NULL;
//     }

//     mf->buf = (char*) buf;
//     mf->size = size;
//     mf->pos = 0;
//     mf->mode = *mode;

//     switch (*mode) {
//         case 'r':
//             return funopen(mf, mem_read, mem_flush, NULL, mem_close);
//         case 'w':
//             return funopen(mf, NULL, mem_write, mem_flush, mem_close);
//         case 'a':
//             mf->pos = size;
//             return funopen(mf, NULL, mem_write, mem_flush, mem_close);
//         default:
//             free(mf);
//             errno = EINVAL;
//             return NULL;
//     }
// }

