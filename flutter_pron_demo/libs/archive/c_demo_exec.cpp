// #include <chrono>
// #include <thread>
// #include <pthread.h> 

// #include "ps_plus.h"
// #include "batch_plus.h"
#include <chrono>
#include <thread>

#include "data.hpp"
#include "c_demo.h"

using namespace std::chrono;

struct Array
{
	char* array;
	int len;
};

struct ArrayOfArray
{
    Array* array;
    int len;
};

typedef struct 
{
    char** array;
    int num_arrays;
} ArrayOfStrings;

void delete_results_container(Array* results, int n){
    
    for(int i = 0; i < n; i++) {
        free(results[i].array);
    }

    free(results);
}

void delete_params_container(Array* results, int n) {
    for(int i = 0; i < n; i++) {
        free(results[i].array);
    }

    free(results);    
}

Array* create_results_container(PS_DYNAMIC_DATA* ps_data) {
    Array* results = (Array*)malloc(sizeof(Array)*5);
    for(int i=0; i<5; i++){
        //char* location = ps_data.data[i].result.result;
        //int result_size =  ps_data.data[i].result.size;
        results[i].array =(char*)malloc(sizeof(char)*ps_data->data[i].result.size+1);
        memcpy(results[i].array, ps_data->data[i].result.result,ps_data->data[i].result.size);

        //results[i].len = ps_data.data[i].result.size;
    }

    return results;
}

Array* create_params_container(char** params, int n) {
    Array* results = (Array*)malloc(sizeof(Array)*n);
    for(int i=0; i<n; i++){
        //char* location = ps_data.data[i].result.result;
        //int result_size =  ps_data.data[i].result.size;
        int param_size = strlen(params[i])+1;
        results[i].array =(char*)malloc(sizeof(char)*param_size);
        memcpy(results[i].array, params[i], param_size);

        results[i].len = param_size;
    }

    return results;
}



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
}

void delete_params_sender(ArrayOfStrings* params) {
    int n = params->num_arrays;

    for(int i = 0; i < n; i++) {
        free(params->array[i]);
    }

    free(params->array);
    free(params);    
}



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

void delete_results_sender(ArrayOfStrings* results) {
    int n = results->num_arrays;

    for(int i = 0; i < n; i++) {
        free(results->array[i]);
    }

    free(results->array);
    free(results);    
}



ArrayOfStrings* ps_demo(const char* data_path) {
    //std::string data_path(path);
    PS_DYNAMIC_DATA ps_data(params125, params125_size,
                        params72, params72_size, 
                        params80, params80_size,
                        params91, params91_size,
                        params105, params105_size,
                        data_path);
    PS_DYNAMIC_DATA *ps_data_ptr = &ps_data;
   
    //std::this_thread::sleep_for(std::chrono::milliseconds(1000));

    //sequential_encapsulated(ps_data.data);
    parallel_encapsulated(ps_data_ptr->data);
    //parallel_encapsualted_with_pthreads(data);
    //Array results[5];

    

    ArrayOfStrings* params = create_results_sender(ps_data_ptr);
    return params;
    

}


int
main()
{   
    //std::string data_path("/home/dbarbera/Repositories/self_contained_c_xyz/Attempt2/data/");
    //std::string data_path("/home/dbarbera/.local/share/com.example.flutter_ps_plus_demo/ps_plus/");
    std::string data_path("./data/");
    // PS_DYNAMIC_DATA ps_data(params125, params125_size,
    //                         params72, params72_size, 
    //                         params80, params80_size,
    //                         params91, params91_size,
    //                         params105, params105_size,
    //                         data_path);



    //sequential_encapsulated(ps_data.data);
    //parallel_encapsulated(ps_data.data);
    //parallel_encapsualted_with_pthreads(data);


    BATCH_DYNAMIC_DATA ps_batch_data(batch_params72, batch_params72_size,
                                    batch_params80, batch_params80_size,
                                    batch_params91, batch_params91_size,
                                    batch_params105, batch_params105_size,
                                    batch_params125, batch_params125_size,
                                    data_path);
    
    //sequential_encapsulated_batch(ps_batch_data.data);
    //parallel_encapsulated_batch(ps_batch_data.data);
    //parallel_encapsualted_batch_with_pthreads(batch_data);

    // Array* results = create_results_container(&ps_data);//don't pass by value or we incurr segmentation fault as it will invoke
    //                                                     //the destructor when running out of scope inside this function and also
    //                                                     //again when exiting main and, hence, those mallocs
    //                                                     // would have been freed already !!!-> seg fault!!!



    // //printf("Working on it...\n");

    // delete_results_container(results,5);

    // Array* arr_params = create_params_container(ps_data.data[0].params.p, ps_data.data[0].params.size);
    // printf("Working on it...\n");
    // delete_params_container(arr_params, ps_data.data[0].params.size);

    // ArrayOfStrings* params = create_params_sender(ps_data.data[0].params.p, ps_data.data[0].params.size);
    // printf("Working on it...\n");
    // delete_params_sender(params);

   
    ArrayOfStrings* results = ps_demo(data_path.data());
    printf("Working on it...\n");
    delete_results_sender(results);
    return 0;
    

    
}