#include <chrono>
#include <thread>
#include <pthread.h> 

#include "ps_plus.h"
#include "batch_plus.h"
#include "data.hpp"

using namespace std::chrono;


int
main()
{
    struct PS_Data data[5];
    load_data(params125, params125_size,
              params72, params72_size, 
              params80, params80_size,
              params91, params91_size,
              params105, params105_size, 
              data);
    //load_data(params125, params125, params125, params125, params125, data);


    // DATA data(5);
    // data.load(params125);

    sequential_encapsulated(data);
    parallel_encapsulated(data);
    //parallel_encapsualted_with_pthreads(data);

    // int i=4;
    // XYZ_PocketSphinx ps1;
    // ps1.init((void*)(data[i].jsgf.buffer), (size_t)data[i].jsgf.size, (void*)data[i].wav.buffer, (size_t)data[i].wav.size, (int)data[i].params.size, data[i].params.p);
    // ps1.init_recognition();
    // ps1.recognize_from_buffered_file();
    // printf("%s\n",ps1._result);
    // ps1.terminate();

    struct PS_Batch_Data batch_data[5];
    load_data_batch(batch_params72, batch_params72_size,
                    batch_params80, batch_params80_size,
                    batch_params91, batch_params91_size,
                    batch_params105, batch_params105_size,
                    batch_params125, batch_params125_size,
                    batch_data);
    
    sequential_encapsulated_batch(batch_data);
    parallel_encapsulated_batch(batch_data);
    //parallel_encapsualted_batch_with_pthreads(batch_data);
    
    printf("Working on it...\n");

    return 0;
    

    
}