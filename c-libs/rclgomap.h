/** 
 * Copyright (c) 2014 rxi
 *
 * This library is free software; you can redistribute it and/or modify it
 * under the terms of the MIT license. See LICENSE for details.
 */

#ifndef MAP_H
#define MAP_H

#include <string.h>
#include <inttypes.h>

#define MAP_VERSION "0.1.0"

// Default values recommended by <http://isthe.com/chongo/tech/comp/fnv/>.
const uint32_t Prime; //   16777619
const uint32_t Seed; // 2166136261

struct map_node_t;
typedef struct map_node_t map_node_t;

typedef struct {
  map_node_t **buckets;
  unsigned nbuckets, nnodes;
} map_base_t;

typedef struct {
  unsigned bucketidx;
  map_node_t *node;
} map_iter_t;


#define map_t(T)\
  struct { map_base_t base; T *ref; T tmp; }


#define map_init(m)\
  memset(m, 0, sizeof(*(m)))


#define map_deinit(m)\
  map_deinit_(&(m)->base)


#define map_get(m, key)\
  ( (m)->ref = map_get_(&(m)->base, key) )


#define map_set(m, key, value)\
  ( (m)->tmp = (value),\
    map_set_(&(m)->base, key, &(m)->tmp, sizeof((m)->tmp)) )


#define map_remove(m, key)\
  map_remove_(&(m)->base, key)


#define map_iter(m)\
  map_iter_()


#define map_next(m, iter)\
  map_next_(&(m)->base, iter)


uint32_t fnv1a(unsigned char);
uint32_t map_hash(uint64_t);
void map_deinit_(map_base_t *m);
void *map_get_(map_base_t *m, uint64_t key);
int map_set_(map_base_t *m, uint64_t key, void *value, int vsize);
void map_remove_(map_base_t *m, uint64_t key);
map_iter_t map_iter_(void);
uint64_t map_next_(map_base_t *m, map_iter_t *iter);


typedef map_t(void*) map_void_t;
typedef map_t(char*) map_str_t;
typedef map_t(int) map_int_t;
typedef map_t(char) map_char_t;
typedef map_t(float) map_float_t;
typedef map_t(double) map_double_t;

#endif