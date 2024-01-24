---
menu:
    after:
        name: graph
        weight: 1
title: Построение графа
---

# Построение графа

Нужно написать воркер, который будет строить граф на текущей странице, каждые 5 секунд
От 5 до 30 элементов, случайным образом. Все ноды графа должны быть связаны.

```go
type Node struct {
    ID int
    Name string
	Form string // "circle", "rect", "square", "ellipse", "round-rect", "rhombus"
    Links []*Node
}
```


## Mermaid Chart

[MermaidJS](https://mermaid-js.github.io/) is library for generating svg charts and diagrams from text.

## Пример

{{< columns >}}
```tpl
{{</*/* mermaid [class="text-center"]*/*/>}}
graph LR
A[circle] --> U[square]
A[circle] --> B[ellipse]
A[circle] --> O[rhombus]
B[ellipse] --> G[ellipse]
B[ellipse] --> Q[square]
B[ellipse] --> H[circle]
C[rhombus] --> R[circle]
C[rhombus] --> Q[square]
C[rhombus] --> L[circle]
D[square] --> A[circle]
D[square] --> L[circle]
D[square] --> B[ellipse]
F[circle] --> H[circle]
F[circle] --> P[rhombus]
F[circle] --> R[circle]
G[ellipse] --> M[circle]
G[ellipse] --> C[rhombus]
G[ellipse] --> D[square]
H[circle] --> Z[rect]
H[circle] --> L[circle]
H[circle] --> E[round-rect]
I[rect] --> D[square]
I[rect] --> G[ellipse]
I[rect] --> X[rhombus]
J[ellipse] --> E[round-rect]
J[ellipse] --> F[circle]
J[ellipse] --> D[square]
K[square] --> Q[square]
K[square] --> G[ellipse]
L[circle] --> P[rhombus]
L[circle] --> C[rhombus]
L[circle] --> J[ellipse]
M[circle] --> T[round-rect]
N[rhombus] --> L[circle]
N[rhombus] --> B[ellipse]
N[rhombus] --> E[round-rect]
O[rhombus] --> K[square]
O[rhombus] --> Y[circle]
P[rhombus] --> D[square]
P[rhombus] --> K[square]
P[rhombus] --> S[rect]
Q[square] --> I[rect]
Q[square] --> G[ellipse]
S[rect] --> V[ellipse]
S[rect] --> U[square]
S[rect] --> K[square]
T[round-rect] --> E[round-rect]
T[round-rect] --> K[square]
T[round-rect] --> Z[rect]
U[square] --> F[circle]
U[square] --> Y[circle]
U[square] --> R[circle]
V[ellipse] --> [[ellipse]
V[ellipse] --> Y[circle]
V[ellipse] --> M[circle]
W[ellipse] --> D[square]
W[ellipse] --> I[rect]
X[rhombus] --> A[circle]
X[rhombus] --> G[ellipse]
Y[circle] --> A[circle]

{{</*/* /mermaid */*/>}}
```

<--->

{{< mermaid >}}
graph LR
A[circle] --> U[square]
A[circle] --> B[ellipse]
A[circle] --> O[rhombus]
B[ellipse] --> G[ellipse]
B[ellipse] --> Q[square]
B[ellipse] --> H[circle]
C[rhombus] --> R[circle]
C[rhombus] --> Q[square]
C[rhombus] --> L[circle]
D[square] --> A[circle]
D[square] --> L[circle]
D[square] --> B[ellipse]
F[circle] --> H[circle]
F[circle] --> P[rhombus]
F[circle] --> R[circle]
G[ellipse] --> M[circle]
G[ellipse] --> C[rhombus]
G[ellipse] --> D[square]
H[circle] --> Z[rect]
H[circle] --> L[circle]
H[circle] --> E[round-rect]
I[rect] --> D[square]
I[rect] --> G[ellipse]
I[rect] --> X[rhombus]
J[ellipse] --> E[round-rect]
J[ellipse] --> F[circle]
J[ellipse] --> D[square]
K[square] --> Q[square]
K[square] --> G[ellipse]
L[circle] --> P[rhombus]
L[circle] --> C[rhombus]
L[circle] --> J[ellipse]
M[circle] --> T[round-rect]
N[rhombus] --> L[circle]
N[rhombus] --> B[ellipse]
N[rhombus] --> E[round-rect]
O[rhombus] --> K[square]
O[rhombus] --> Y[circle]
P[rhombus] --> D[square]
P[rhombus] --> K[square]
P[rhombus] --> S[rect]
Q[square] --> I[rect]
Q[square] --> G[ellipse]
S[rect] --> V[ellipse]
S[rect] --> U[square]
S[rect] --> K[square]
T[round-rect] --> E[round-rect]
T[round-rect] --> K[square]
T[round-rect] --> Z[rect]
U[square] --> F[circle]
U[square] --> Y[circle]
U[square] --> R[circle]
V[ellipse] --> [[ellipse]
V[ellipse] --> Y[circle]
V[ellipse] --> M[circle]
W[ellipse] --> D[square]
W[ellipse] --> I[rect]
X[rhombus] --> A[circle]
X[rhombus] --> G[ellipse]
Y[circle] --> A[circle]

{{< /mermaid >}}

{{< /columns >}}
