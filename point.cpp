//
// Created by Jiangshan on 2019-03-24.
//

#include "point.h"

// Point class implementation
Point::Point(int _x, int _y){
    x = _x;
    y = _y;
}

int Point::sum(){
    return x+y;
}