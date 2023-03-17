// #ifdef WIN32
// #define EXPORT __declspec(dllexport)
// #else
// #define EXPORT extern "C" __attribute__((visibility("default"))) __attribute__((used))
//#endif

// #include <chrono>
// #include <thread>

#include <stdint>
//#include <stdio.h>
//#include <stdlib.h>

// #if _WIN32
// #include <windows.h>
// #else
// #include <pthread.h>
// #include <unistd.h>
// #endif

#define EXPORT 

#if defined(__cplusplus)
extern "C" {
#endif

static char buffer[1024];

struct Array
{
	char* array;
	int len;
};

typedef struct 
{
    char** array;
    int num_arrays;
} ArrayOfStrings;

EXPORT 
int add(int a, int b);

EXPORT
char* capitalize(char *str);

EXPORT
intptr_t sum_long_running(intptr_t a, intptr_t b);

EXPORT
const char *hello_world();

EXPORT
ArrayOfStrings* get_some_parameters(char *data_path);

EXPORT
void delete_params_sender(ArrayOfStrings* params);

EXPORT
void delete_results_sender(ArrayOfStrings* results);

EXPORT
void delete_results_sender_batch(ArrayOfStrings* results);

EXPORT
ArrayOfStrings* ps_demo(const char* path);

EXPORT
ArrayOfStrings* ps_demo_sequential(const char* path);

EXPORT
ArrayOfStrings* ps_batch_demo(const char* path);

EXPORT
ArrayOfStrings* ps_batch_demo_sequential(const char* path);

#if defined(__cplusplus)
/* end 'extern "C"' wrapper */
}
#endif