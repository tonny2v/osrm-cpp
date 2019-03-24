//
// Created by Jiangshan on 2019-03-24.
//
#include "test.h"
#include "point.h"


// non-Point Class functions
int sum(int x, int y){
    return x+y;
}

// wrapper of class function
int p_sum(int x, int y) {
    Point p = Point(x, y);
    return p.sum();
}