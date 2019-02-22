# Задача

Сортировка пузырьком

# Бенчмарк

```
BenchmarkBubbleSort64-4    	  500000	      2508 ns/op	       0 B/op	       0 allocs/op
BenchmarkBubbleSort128-4   	  200000	      8747 ns/op	       0 B/op	       0 allocs/op
BenchmarkBubbleSort256-4   	   50000	     32444 ns/op	       0 B/op	       0 allocs/op
```

# Результат

```
Unsorted: [95 218 71 84 54 237 93 141 146 107 177 73 85 244 77 173]
Sorted: [54 71 73 77 84 85 93 95 107 141 146 173 177 218 237 244]
```