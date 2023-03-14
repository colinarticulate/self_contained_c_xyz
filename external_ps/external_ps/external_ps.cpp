//
//  external_ps.cpp
//  external_ps
//
//  Created by David Barbera on 14/03/2023.
//

#include <iostream>
#include "external_ps.hpp"
#include "external_psPriv.hpp"

void external_ps::HelloWorld(const char * s)
{
    external_psPriv *theObj = new external_psPriv;
    theObj->HelloWorldPriv(s);
    delete theObj;
};

void external_psPriv::HelloWorldPriv(const char * s) 
{
    std::cout << s << std::endl;
};

