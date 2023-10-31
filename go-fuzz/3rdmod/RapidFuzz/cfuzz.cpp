#include <stdlib.h>
#include <stdio.h>
#include "cfuzz.h"
#include "rapidfuzz/fuzz.hpp"
#include "rapidfuzz/details/types.hpp"

SCScoreAlignment* NewSCScoreAlignment(rapidfuzz::ScoreAlignment<double>* sa){
    SCScoreAlignment * p = (SCScoreAlignment *)malloc(sizeof(SCScoreAlignment));
    p->score = sa->score;
    p->src_start = sa->src_start;
    p->src_end = sa->src_end;
    p->dest_start = sa->dest_start;
    p->dest_end = sa->dest_end;
    return p;
}

SCScoreAlignment* SCPartialRatioAlignment(char* src, char* dst){
    auto sa = rapidfuzz::fuzz::partial_ratio_alignment(src, dst);
    SCScoreAlignment * p = NewSCScoreAlignment(&sa);
    return p;
}

void DeleteSCScoreAlignment(SCScoreAlignment* p){
    delete p;
}