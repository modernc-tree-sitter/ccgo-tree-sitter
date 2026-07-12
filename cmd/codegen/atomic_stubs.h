/* Stubs for compiler atomic builtins that ccgo doesn't handle natively. */
#ifndef CCGO_TREE_SITTER_ATOMIC_STUBS_H
#define CCGO_TREE_SITTER_ATOMIC_STUBS_H
typedef unsigned int uint32_t;
static inline uint32_t __atomic_add_fetch(volatile uint32_t *p, uint32_t v, int m) { (void)m; *p += v; return *p; }
static inline uint32_t __atomic_sub_fetch(volatile uint32_t *p, uint32_t v, int m) { (void)m; *p -= v; return *p; }
#endif
