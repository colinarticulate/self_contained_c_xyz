#ifdef WIN32
#define EXPORT __declspec(dllexport)
#else
#define EXPORT extern "C" __attribute__((visibility("default"))) __attribute__((used))
#endif

//This is only for the sleep functions:
#include <chrono>
#include <thread>

#include <cstring>
#include <ctype.h>

#include <stdio.h>
#include <stdlib.h>

#include "data.hpp"
#include "c_demo.h"

using namespace std::chrono;

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
int add(int a, int b)
{
    return a + b;
}

EXPORT
char* capitalize(char *str) {
    strcpy(buffer, str);
    buffer[0] = toupper(buffer[0]);
    return buffer;
}



//Testing getting sending const char to Dart
EXPORT
const char *hello_world()
{
    return "Hello World";
}

//EXPORT
// Array* create_params_container(char** params, int n) {
//     Array* results = (Array*)malloc(sizeof(Array)*n);
//     for(int i=0; i<n; i++){
//         //char* location = ps_data.data[i].result.result;
//         //int result_size =  ps_data.data[i].result.size;
//         int param_size = strlen(params[i])+1;
//         results[i].array =(char*)malloc(sizeof(char)*param_size);
//         memcpy(results[i].array, params[i], param_size);

//         results[i].len = param_size;
//     }

//     return results;
// }

// EXPORT
// struct Array* get_some_parameters(){
//         std::string data_path("/not_yet_have_the_path/");
//         PS_DYNAMIC_DATA ps_data(params125, params125_size,
//                         params72, params72_size, 
//                         params80, params80_size,
//                         params91, params91_size,
//                         params105, params105_size,
//                         data_path);
//     Array* params = create_params_container(ps_data.data[0].params.p, ps_data.data[0].params.size);
// }

// EXPORT
// void delete_params_container(struct Array* results, int n) {
//     for(int i = 0; i < n; i++) {
//         free(results[i].array);
//     }

//     free(results);    
// }

// char** create_params_sender(char** params, int n){
//     char** results = (char**)malloc(sizeof(char*)*n);
//     for(int i=0; i<n; i++){
//         int param_size = strlen(params[i])+1;
//         results[i] =(char*)malloc(sizeof(char)*param_size);
//         memcpy(results[i], params[i], param_size);
//     }

//     return results;
// }





ArrayOfStrings* create_params_sender(char** _params, int n){
    ArrayOfStrings* params = (ArrayOfStrings*)malloc(sizeof(ArrayOfStrings));
    params->num_arrays = n;
    params->array = (char**)malloc(sizeof(char*)*n);
    for(int i=0; i<n; i++){
        int param_size = strlen(_params[i])+1;
        params->array[i] =(char*)malloc(sizeof(char)*param_size);
        memcpy(params->array[i], _params[i], param_size);
    }

    return params;
    printf("CREATED PARAMS\n");
}

EXPORT
ArrayOfStrings* get_some_parameters(char *data_path){
    //std::string data_path("/not_yet_have_the_path/");
    PS_DYNAMIC_DATA ps_data(params125, params125_size,
                    params72, params72_size, 
                    params80, params80_size,
                    params91, params91_size,
                    params105, params105_size,
                    data_path);
    ArrayOfStrings* params = create_params_sender(ps_data.data[0].params.p, ps_data.data[0].params.size);
    return params;
}

EXPORT
void delete_params_sender(ArrayOfStrings* params) {
    int n = params->num_arrays;

    for(int i = 0; i < n; i++) {
        free(params->array[i]);
    }

    free(params->array);
    free(params);    
}

EXPORT
ArrayOfStrings* demo_ps(char* data_path) {
    //std::string data_path(path);
    PS_DYNAMIC_DATA ps_data(params125, params125_size,
                        params72, params72_size, 
                        params80, params80_size,
                        params91, params91_size,
                        params105, params105_size,
                        data_path);


    //sequential_encapsulated(ps_data.data);
    //parallel_encapsulated(ps_data.data);
    //parallel_encapsualted_with_pthreads(data);
    //Array results[5];
    ArrayOfStrings* params = create_params_sender(ps_data.data[0].params.p, ps_data.data[0].params.size);
    return params;
    

}

EXPORT
void demo_ps_batch(std::string data_path) {
    BATCH_DYNAMIC_DATA ps_batch_data(batch_params72, batch_params72_size,
                                    batch_params80, batch_params80_size,
                                    batch_params91, batch_params91_size,
                                    batch_params105, batch_params105_size,
                                    batch_params125, batch_params125_size,
                                    data_path);
    
    //sequential_encapsulated_batch(ps_batch_data.data);
    parallel_encapsulated_batch(ps_batch_data.data);
    //parallel_encapsualted_batch_with_pthreads(batch_data);

}

//***************************************

ArrayOfStrings* create_results_sender(PS_DYNAMIC_DATA *ps_data){
    ArrayOfStrings* params = (ArrayOfStrings*)malloc(sizeof(ArrayOfStrings));
    int n = 5;
    params->num_arrays = n;
    params->array = (char**)malloc(sizeof(char*)*n);
    for(int i=0; i<n; i++){
        int param_size = strlen(ps_data->data[i].result.result)+1;
        params->array[i] =(char*)malloc(sizeof(char)*param_size);
        memcpy(params->array[i], ps_data->data[i].result.result, param_size);
    }

    return params;
}

ArrayOfStrings* create_results_sender_batch(BATCH_DYNAMIC_DATA *ps_data){
    ArrayOfStrings* params = (ArrayOfStrings*)malloc(sizeof(ArrayOfStrings));
    int n = 5;
    params->num_arrays = n;
    params->array = (char**)malloc(sizeof(char*)*n);
    for(int i=0; i<n; i++){
        int param_size = strlen(ps_data->data[i].result.result)+1;
        params->array[i] =(char*)malloc(sizeof(char)*param_size);
        memcpy(params->array[i], ps_data->data[i].result.result, param_size);
    }

    return params;
}

EXPORT
void delete_results_sender(ArrayOfStrings* results) {
    int n = results->num_arrays;

    for(int i = 0; i < n; i++) {
        free(results->array[i]);
    }

    free(results->array);
    free(results);    
}

EXPORT
void delete_results_sender_batch(ArrayOfStrings* results) {
    int n = results->num_arrays;

    for(int i = 0; i < n; i++) {
        free(results->array[i]);
    }

    free(results->array);
    free(results);    
}


EXPORT
ArrayOfStrings* ps_demo(const char* path) {
    printf("%s\n",path);
    std::string data_path(path);
    printf("%s\n", data_path.data());
    PS_DYNAMIC_DATA ps_data(params125, params125_size,
                        params72, params72_size, 
                        params80, params80_size,
                        params91, params91_size,
                        params105, params105_size,
                        data_path);

    //std::this_thread::sleep_for(std::chrono::milliseconds(1000));

    printf("DEBUGGING: replacing parallel_encapsulated with sequential_encapsulated in ps_demo call\n");
    //sequential_encapsulated(ps_data.data);
    
    parallel_encapsulated(ps_data.data);
    //parallel_encapsualted_with_pthreads(ps_data.data);
    //Array results[5];
    ArrayOfStrings* params = create_results_sender(&ps_data);
    return params;
    

}

EXPORT
ArrayOfStrings* ps_demo_sequential(const char* path) {
    std::string data_path(path);
    PS_DYNAMIC_DATA ps_data(params125, params125_size,
                        params72, params72_size, 
                        params80, params80_size,
                        params91, params91_size,
                        params105, params105_size,
                        data_path);

    //std::this_thread::sleep_for(std::chrono::milliseconds(1000));

    sequential_encapsulated(ps_data.data);
    //parallel_encapsulated(ps_data.data);
    //parallel_encapsualted_with_pthreads(ps_data.data);
    //Array results[5];
    ArrayOfStrings* params = create_results_sender(&ps_data);
    return params;
    

}

EXPORT
ArrayOfStrings* ps_batch_demo(const char* path) {
    std::string data_path(path);
    BATCH_DYNAMIC_DATA ps_batch_data(batch_params125, batch_params125_size,
                        batch_params72, batch_params72_size, 
                        batch_params80, batch_params80_size,
                        batch_params91, batch_params91_size,
                        batch_params105, batch_params105_size,
                        data_path);

    //std::this_thread::sleep_for(std::chrono::milliseconds(1000));

    //sequential_encapsulated(ps_data.data);
    parallel_encapsulated_batch(ps_batch_data.data);
    //parallel_encapsualted_with_pthreads(ps_data.data);
    //Array results[5];
    ArrayOfStrings* params = create_results_sender_batch(&ps_batch_data);
    return params;
    
}

EXPORT
ArrayOfStrings* ps_batch_demo_sequential(const char* path) {
    std::string data_path(path);
    BATCH_DYNAMIC_DATA ps_batch_data(batch_params125, batch_params125_size,
                        batch_params72, batch_params72_size, 
                        batch_params80, batch_params80_size,
                        batch_params91, batch_params91_size,
                        batch_params105, batch_params105_size,
                        data_path);

    //std::this_thread::sleep_for(std::chrono::milliseconds(1000));

    sequential_encapsulated_batch(ps_batch_data.data);
    //parallel_encapsulated_batch(ps_batch_data.data);
    //parallel_encapsualted_with_pthreads(ps_data.data);
    //Array results[5];
    ArrayOfStrings* params = create_results_sender_batch(&ps_batch_data);
    return params;
    

}