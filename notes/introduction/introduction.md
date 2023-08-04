# Algorithms


## What is an Algorithm?

An Algorithm is a step-by-step unambiguous instruction to solve a given problem and obtain a legitimate output in a finite time.

Just like a recipe in cooking, it tells your computer exactly what steps it needs to take to arrive at the desired result.

In programming and algorithms, non-ambiguity means the instructions or steps are clear, precise, and unambiguous. Each instruction has exactly one meaning and leaves no room for misinterpretation.

## Design & Analysis

<img width="924" alt="image" src="https://github.com/RaghavTheGreat1/algorithms/assets/28825619/81c0c3d9-0e5c-4355-90d1-fce300e16576">

## Analysis of Algorithms

Analysis of algorithms is the process of finding the computational complexity of algorithms. It involves determining the amount of resources (time, space) necessary to execute them. The main goal of analyzing algorithms is to understand the theoretical performance of an algorithm.

Typically, we focus on the following:

1. Time Complexity: This represents the amount of time taken by an algorithm to run. Usually represented with Big-O notation, it helps to analyze the worst-case, average-case, and best-case scenarios in terms of time complexity.

2. Space Complexity: This represents the amount of memory an algorithm needs to run to completion. Some algorithms might be fast, but they require a lot of memory. Others may require lesser memory but might take more time to complete.

3. Trade-off between Time and Space: Sometimes you could get improvement in time complexity by using more memory, or save memory by accepting slightly slower execution times (e.g., using caching to remember previously computed results). This trade-off is studied in the analysis of algorithms.

4. Scaling to larger inputs: Analysis also predicts how the algorithm performs as it scales with larger input size.

Understanding and being able to analyze the complexity of an algorithm can help in choosing the most efficient algorithm for a specific task, accounting for the constraints of the actual hardware being used and the specificities of the actual data you're handling.

### Running Time Analysis

Also known as time complexity analysis, is a theoretical assessment that measures the time taken by an algorithm to run as a function of the **size of its input**. In other words, it provides an estimate of how long an algorithm will take to execute given a certain amount of data to process and  often discussed in **Big O Notation**.

### Types of Analysis of Algorithms

1. **Worst-Case Analysis**: This measures the maximum time taken by an algorithm to solve a problem. This is denoted as the Big-O Notation (O(n)) and gives an upper bound of the time complexity.

2. **Average-Case Analysis**: This measures the average time taken by an algorithm to solve a problem. This type of analysis is complex as it requires us to consider all possible inputs and their probabilities. Due to the complexity and extensive computation, average case analysis is not frequently used. This is denoted by Theta Notation (θ(n))

3. **Best-Case Analysis**: This measures the minimum time taken by an algorithm to solve a problem. This is denoted as the Omega Notation (Ω(n)) and gives a lower bound of the time complexity.

### Orders of Growth

Order of growth is a measure used in computational complexity theory to describe the performance or complexity of an algorithm. It's commonly expressed with Big O Notation, which describes how the time execution or space usage grows with the input size.

Here are some common orders of growth from smallest to largest, along with their Big O notation and layman explanation:

1. O(1) - Constant Time: Regardless of the size of the input data, the time taken remains constant.

2. O(log n) - Logarithmic Time: Execution time increases logarithmically with the input size. Common in algorithms that cut down their input size with each step like binary search.

3. O(n) - Linear Time: The running time increases at most linearly with the size of the input. An example would be a simple loop running over an array.

4. O(n log n) - Linearithmic Time: Better than quadratic but worse than linear. A well-known algorithm with this time complexity is the Fast Fourier Transform.

5. O(n^2) - Quadratic Time: Every increase in the input size doubles the time taken. Simple sorting algorithms like bubble sort fall into this category.

6. O(n^3) - Cubic Time: The execution time grows cubically for an increase in input size. An example would be the naive matrix multiplication algorithm.

7. O(2^n) - Exponential Time: Execution time doubles for each addition to the input data set. The classic recursive calculation of Fibonacci numbers is an example.

The goal in most cases is to aim for the least order of growth, as it scales best with larger inputs.
