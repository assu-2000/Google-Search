# Google-Search

This repository is a learning journey and practical showcase on how to design **robust**, **scalable**, and **high-performance systems** using **Go**.  

Each Git tag in this repo represents a progressive milestone:
- From basic logic
- To some Go **concurrency patterns**
- Ending with highly **optimized concurrent architectures** inspired by real-world systems.

The code simulates querying a search engine (e.g., "Google Search") with time measurement to benchmark performance on your own machine.

##  Learning Objectives

- Master Go's concurrency primitives: goroutines, channels, select, fan-out/fan-in, timeout control, etc.
- Design concurrent systems in increasing levels of complexity.
- Measure and reason about performance tradeoffs.
- See in real time how different implementations scale with increased parallelism.

##  Project Struct

- `google.go`: Core file that launches a search request and measures execution time.
- Tags like `v1.0`, `v2.0`, ..., `vN.N` reflect different stages:
  - `v1.0`: Basic sequential implementation.
  - `v2.0`: Introduces goroutines for concurrent queries.
  - `v3+`: Leverages timeout handling, replication.

Each version refactors the logic to demonstrate increasingly concurrency strategies.


## Requirements

* [Go](https://golang.org/doc/install) installed (tested with **Go 1.24.4**)
* Any Unix-based or Windows environment with terminal access

## ▶️ How to Run

1. Clone the repository:

   ```bash
   git clone https://github.com/assu-2000/Google-Search.git
   cd Google-Search
   ```

2. Checkout any tag to explore a specific version:

   ```bash
   git checkout v3.0 
   ```

3. Run the program:

   ```bash
   go run google.go
   ```

4. You'll see the simulated result along with execution time:

   ```
   [Video3 results for "golang" query Web2 results for "golang" query Image3 results for "golang" query]
   Time taken: 35.4189ms
   ```

##  Explore by Yourself

Use the tags to:

* Compare execution times
* Understand bottlenecks
* See how concurrency boosts performance
* Experiment on your own CPU

## Credit

* [Rob Pike’s concurrency patterns in Go](https://talks.golang.org/2012/concurrency.slide)

## Inspiration

* Real-world practices for building distributed and low-latency systems.



