// #include <chrono>
// #include <thread>
// #include <pthread.h> 

// #include "ps_plus.h"
// #include "batch_plus.h"
#include "data.hpp"
#include "c_demo.h"

using namespace std::chrono;


int
main()
{   
    std::string data_path("/home/dbarbera/Repositories/self_contained_c_xyz/Attempt2/data/");

    PS_DYNAMIC_DATA ps_data(params125, params125_size,
                            params72, params72_size, 
                            params80, params80_size,
                            params91, params91_size,
                            params105, params105_size,
                            data_path);


    //sequential_encapsulated(ps_data.data);
    parallel_encapsulated(ps_data.data);
    //parallel_encapsualted_with_pthreads(data);


    BATCH_DYNAMIC_DATA ps_batch_data(batch_params72, batch_params72_size,
                                    batch_params80, batch_params80_size,
                                    batch_params91, batch_params91_size,
                                    batch_params105, batch_params105_size,
                                    batch_params125, batch_params125_size,
                                    data_path);
    
    //sequential_encapsulated_batch(ps_batch_data.data);
    parallel_encapsulated_batch(ps_batch_data.data);
    //parallel_encapsualted_batch_with_pthreads(batch_data);
    
    printf("Working on it...\n");

    return 0;
    

    
}