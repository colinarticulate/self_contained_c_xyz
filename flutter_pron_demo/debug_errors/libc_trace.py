import os

library = "libc.so"
function_to_find = "fmemopen"

def copy_to_file(filename, data_list):
    data = "\n".join(data_list)
    with open(filename, "w") as f:
        f.write(data)

def shell(cmd):
    stream = os.popen(cmd)
    output = stream.read().strip("\n").split("\n")
    return output

def search_library(directory, library):
    cmd_find_libc = f"find {directory} -iname {library}"
    output = shell(cmd_find_libc)
    return output

def search_function_within_library(output, function):
     
    has_function =[]
    has_not_function = []
    for line in output:
        if line != "":
            cmd_search_function = f"nm -D {line} | grep '{function}'"
            search_output = shell(cmd_search_function)

            if function in search_output[0]:
                has_function.append(line)
                # print(f"{line}\t\t{search_output[0]}")
                
            else:
                has_not_function.append(line)

    return has_function, has_not_function


def main():
    # output = os.system("find /usr -iname libc.so")
    output = search_library("/home/dbarbera", "libc.so")
    has_fmem, has_not_fmem = search_function_within_library(output, "fmemopen")     
    print(f"has fmemopen: {len(has_fmem)}, has not: {len(has_not_fmem)}. Total: {len(output)}") 

    has_funopen, has_not_funopen = search_function_within_library(output, "funopen")     
    print(f"has funopen: {len(has_funopen)}, has not: {len(has_not_funopen)}. Total: {len(output)}") 

    copy_to_file("has_fmemopen.txt", has_fmem)
    copy_to_file("has_not_fmemopen.txt", has_not_fmem)
    
    print("working on it")

if __name__ == "__main__":
    main()
