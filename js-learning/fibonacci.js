// Fibonacci sequence generator, returning an array of numbers of length n
function fibonacci(n) {
    if (n === 1) {
        return [0]
    }

    let fib = [0, 1];
    for (let i = 2; i < n; i++) {
        fib[i] = fib[i - 2] + fib[i - 1];
    }

    return fib;
}

console.log(fibonacci(10));

// Fibonacci calculator, returning the nth number in the sequence
function fibonacciCalc(n) {
    if (n === 1) {
        return 0;
    } else if (n === 2) {
        return 1;
    }

    return fibonacciCalc(n - 2) + fibonacciCalc(n - 1);
}

console.log(fibonacciCalc(10));

// Fibonacci calculator, returning the number under nth index in the sequence
function fibonacciCalc2(n) {
    if (n < 2 && n >= 0) {
        return n;
    }

    return fibonacciCalc2(n - 2) + fibonacciCalc2(n - 1);
}

console.log(fibonacciCalc2(10));