# Estimator - Task Time Estimation Calculator

A humorous tool for developers and project managers that helps estimate task completion time using various "scientifically-backed" methods.

## Features

- **π Method**: Multiply initial estimate by π (3.14159...)
- **"+2 units" Method**: Add 2 units to your initial estimate
- **PERT Method**: Calculate using (O + 4M + P) / 6 formula with probability ranges
- **The "Last 10%" Law**: Account for the well-known effect where "90% of the work takes 90% of the time, and the remaining 10% takes another 90% of the time"
- **Hofstadter's Law**: "It always takes longer than you expect, even when you take into account Hofstadter's Law"
- **Combined Method**: For extreme caution (π + 2)
- **Share Results**: Ability to share a link with a specific estimation
- **Multilingual**: Supports English, Russian, and Chinese

## Live Demo

Try it out at: [eta.lex.la](https://eta.lex.la)

## Running Locally

```bash
# Locally
$ go run main.go

# Or using Podman
$ podman build --file build/Containerfile --tag lexfrei/estimator .
$ podman run --publish 8080:8080 lexfrei/estimator
```

## Disclaimer

This tool is created for entertainment purposes and does not claim scientific accuracy. Any resemblance to actual task estimation is purely coincidental... or is it? 🤔
