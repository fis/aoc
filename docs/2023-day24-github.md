# AoC 2023 day 24

See [2023-notes.md](2023-notes.md) for all the other days.

## Part 1

Let's start with some notation for part 1. We have a pair of hailstones $A$ and
$B$. For hailstone $i \in \{A, B\}$, denote its initial position by
$\mathbf{p}_0^{(i)}$ and its velocity by $\mathbf{v}^{(i)}$. Then
$\mathbf{p}^{(i)}(t) = \mathbf{p}_0^{(i)} + t \mathbf{v}^{(i)}$ will be the
position of the hailstone at time $t$.

We're interested in some region where the coordinates $(x, y)$ are bounded by

$$\begin{aligned}
x_{\mathrm{min}} &\leq x \leq x_{\mathrm{max}} \\
y_{\mathrm{min}} &\leq y \leq y_{\mathrm{max}}
\end{aligned}$$

When might hailstone $A$ be within that region? Let's look at the $x$ coordinate
first:

```math
\begin{aligned}
x_{\mathrm{min}} &\leq [\mathbf{p}^{(A)}(t)]_x \leq x_{\mathrm{max}} \\
x_{\mathrm{min}} &\leq [\mathbf{p}_0^{(i)}]_x + t [\mathbf{v}^{(i)}]_x \leq x_{\mathrm{max}} \\
x_{\mathrm{min}} - [\mathbf{p}_0^{(i)}]_x &\leq t [\mathbf{v}^{(i)}]_x \leq x_{\mathrm{max}} - [\mathbf{p}_0^{(i)}]_x \\
\frac{x_{\mathrm{min}} - [\mathbf{p}_0^{(i)}]_x}{[\mathbf{v}^{(i)}]_x} &\leq t \leq \frac{x_{\mathrm{max}} - [\mathbf{p}_0^{(i)}]_x}{[\mathbf{v}^{(i)}]_x}
\end{aligned}
```

In the above, we've for convenience assumed $[\mathbf{v}^{(i)}]\_x > 0$; if not,
it's just a matter of swapping the endpoints. By doing the same for the $y$
coordinate and taking the intersection, we can derive an interval
$[t_{\mathrm{min}}^{(A)}, t_{\mathrm{max}}^{(A)}]$ that specifies the only time
hailstone $A$ can be within the region of interest. If the interval is empty, or
$t_{\mathrm{max}}^{(A)} < 0$, this hailstone is useless to us: it can never meet
another hailstone within the region.

Otherwise, let's turn our attention to hailstone $B$. When does it encounter the
path of hailstone $A$? If the two are not going in the same direction, there
will be some pair of times $t^{(A)}, t^{(B)}$ for which
$\mathbf{p}^{(A)}(t^{(A)}) = \mathbf{p}^{(B)}(t^{(B)})$. Could we solve for the
times?

$$\begin{aligned}
\mathbf{p}^{(A)}(t^{(A)}) &= \mathbf{p}^{(B)}(t^{(B)}) \\
\mathbf{p}_0^{(A)} + \mathbf{v}^{(A)} t^{(A)} &= \mathbf{p}_0^{(B)} + \mathbf{v}^{(B)} t^{(B)} \\
\begin{pmatrix}
\mathbf{v}^{(A)} & -\mathbf{v}^{(B)}
\end{pmatrix} \begin{pmatrix}
t^{(A)} \\
t^{(B)}
\end{pmatrix} &= \mathbf{p}_0^{(B)} - \mathbf{p}_0^{(A)}
\end{aligned}$$

This is just a linear system of equations. If we solve it, we get:

```math
\begin{aligned}
t^{(A)} &= \frac{
  [\mathbf{p}_0^{(A)}]_y [\mathbf{v}^{(B)}]_x
  - [\mathbf{p}_0^{(A)}]_x [\mathbf{v}^{(B)}]_y
  + [\mathbf{p}_0^{(B)}]_x [\mathbf{v}^{(B)}]_y
  - [\mathbf{p}_0^{(B)}]_y [\mathbf{v}^{(B)}]_x
}{
  [\mathbf{v}^{(A)}]_x [\mathbf{v}^{(B)}]_y
  - [\mathbf{v}^{(A)}]_y [\mathbf{v}^{(B)}]_x
} \\
t^{(B)} &= \frac{
  [\mathbf{p}_0^{(A)}]_y
  - [\mathbf{p}_0^{(B)}]_y
  + t^{(A)} [\mathbf{v}^{(A)}]_y
}{[\mathbf{v}^{(B)}]_y}
\end{aligned}
```

We could also easily solve for the $(x, y)$ coordinates of the intersection, but
there's no need to: we can just check if
$t^{(A)} \in [t_{\mathrm{min}}^{(A)}, t_{\mathrm{max}}^{(A)}]$. We do need to
compute $t^{(B)}$ as well, though, to verify that $t^{(B)} \geq 0$.

Phew. And that was the *simple* part. Now for part 2...

## Part 2

Let's keep the same notation, but add a third hailstone $C$ from our collection.
We'll of course also have the thrown rock, $R$. What we need for the puzzle
solution then is to figure out the value of $\mathbf{p}_0^{(R)}$.

This time, we'll use $t^{(i)}, i \in \{A,B,C\}$ for the specific times when our
rock obliterates hailstone $i$. Let's play with that a little:

```math
\begin{aligned}
\mathbf{p}_0^{(R)} + t^{(i)} \mathbf{v}^{(R)} &= \mathbf{p}_0^{(i)} + t^{(i)} \mathbf{v}^{(i)} \\
\mathbf{p}_0^{(R)} - \mathbf{p}_0^{(i)} &= -t^{(i)} \left( \mathbf{v}^{(R)} - \mathbf{v}^{(i)} \right) \\
\left( \mathbf{p}_0^{(R)} - \mathbf{p}_0^{(i)} \right) \times \left( \mathbf{v}^{(R)} - \mathbf{v}^{(i)} \right)
&= -t^{(i)} \left( \mathbf{v}^{(R)} - \mathbf{v}^{(i)} \right) \times \left( \mathbf{v}^{(R)} - \mathbf{v}^{(i)} \right) \\
\left( \mathbf{p}_0^{(R)} - \mathbf{p}_0^{(i)} \right) \times \left( \mathbf{v}^{(R)} - \mathbf{v}^{(i)} \right) &= 0 \\
\mathbf{p}_0^{(R)} \times \mathbf{v}^{(R)}
- \mathbf{p}_0^{(R)} \times \mathbf{v}^{(i)}
- \mathbf{p}_0^{(i)} \times \mathbf{v}^{(R)}
+ \mathbf{p}_0^{(i)} \times \mathbf{v}^{(i)} &= 0
\end{aligned}
```

In the above, we used to our advantage the property that
$\mathbf{a} \times \mathbf{a} = \mathbf{0}$.

The result still does have a nasty $\mathbf{p}_0^{(R)} \times \mathbf{v}^{(R)}$
term that makes it nonlinear. However, that term is the same for all $i$. The
result of the equation is also constant ($0$). Let's see what happens if we
equate the above for hailstones $A$ and $B$ and rearrange:

```math
\begin{aligned}
\mathbf{p}_0^{(R)} \times \mathbf{v}^{(R)}
- \mathbf{p}_0^{(R)} \times \mathbf{v}^{(A)}
- \mathbf{p}_0^{(A)} \times \mathbf{v}^{(R)}
+ \mathbf{p}_0^{(A)} \times \mathbf{v}^{(A)}
&= \mathbf{p}_0^{(R)} \times \mathbf{v}^{(R)}
- \mathbf{p}_0^{(R)} \times \mathbf{v}^{(B)}
- \mathbf{p}_0^{(B)} \times \mathbf{v}^{(R)}
+ \mathbf{p}_0^{(B)} \times \mathbf{v}^{(B)} \\
\mathbf{p}_0^{(A)} \times \mathbf{v}^{(A)}
- \mathbf{p}_0^{(R)} \times \mathbf{v}^{(A)}
- \mathbf{p}_0^{(A)} \times \mathbf{v}^{(R)}
&= \mathbf{p}_0^{(B)} \times \mathbf{v}^{(B)}
- \mathbf{p}_0^{(R)} \times \mathbf{v}^{(B)}
- \mathbf{p}_0^{(B)} \times \mathbf{v}^{(R)}
\end{aligned}
```

Oh, the $\mathbf{p}_0^{(R)} \times \mathbf{v}^{(R)}$ term drops out. If we do
the same steps for the $A$, $C$ pair, and group the unknowns
($\mathbf{p}_0^{(R)}, \mathbf{v}^{(R)}$) on one side and known values on the
other, and expand the cross products, we're left with a linear system of six
equations and six unknowns:

```math
\begin{aligned}
\mathbf{p}_0^{(R)} \times \left( \mathbf{v}^{(B)} - \mathbf{v}^{(A)} \right)
+ \left( \mathbf{p}_0^{(B)} - \mathbf{p}_0^{(A)} \right) \times \mathbf{v}^{(R)}
&= \mathbf{p}_0^{(B)} \times \mathbf{v}^{(B)} - \mathbf{p}_0^{(A)} \times \mathbf{v}^{(A)} \\
\mathbf{p}_0^{(R)} \times \left( \mathbf{v}^{(C)} - \mathbf{v}^{(A)} \right)
+ \left( \mathbf{p}_0^{(C)} - \mathbf{p}_0^{(A)} \right) \times \mathbf{v}^{(R)}
&= \mathbf{p}_0^{(C)} \times \mathbf{v}^{(C)} - \mathbf{p}_0^{(A)} \times \mathbf{v}^{(A)}
\end{aligned}
```

If we expand out the cross products, and use a shorthand notation where
$p_d^i = [\mathbf{p}_0^{(i)}]_d$, $v_d^i = [\mathbf{v}^{(i)}]_d$, and for
$i \in \{B, C\}, q \in \{p, v\}$ also $\Delta q_d^i = q_d^i - q_d^A$, we get the
following explicit form of the equation system:

$$
\begin{pmatrix}
0 & \Delta v_z^B & -\Delta v_y^B & 0 & -\Delta p_z^B & \Delta p_y^B \\
-\Delta v_z^B & 0 & \Delta v_x^B & \Delta p_z^B & 0 & -\Delta p_x^B \\
\Delta v_y^B & -\Delta v_x^B & 0 & -\Delta p_y^B & \Delta p_x^B & 0 \\
0 & \Delta v_z^C & -\Delta v_y^C & 0 & -\Delta p_z^C & \Delta p_y^C \\
-\Delta v_z^C & 0 & \Delta v_x^C & \Delta p_z^C & 0 & -\Delta p_x^C \\
\Delta v_y^C & -\Delta v_x^C & 0 & -\Delta p_y^C & \Delta p_x^C & 0
\end{pmatrix} \begin{pmatrix}
\mathbf{p}_0^{(R)} \\
\mathbf{v}^{(R)}
\end{pmatrix} = \begin{pmatrix}
p_y^B v_z^B - p_z^B v_y^B - p_y^A v_z^A + p_z^A v_y^A \\
p_z^B v_x^B - p_x^B v_z^B - p_z^A v_x^A + p_x^A v_z^A \\
p_x^B v_y^B - p_y^B v_x^B - p_x^A v_y^A + p_y^A v_x^A \\
p_y^C v_z^C - p_z^C v_y^C - p_y^A v_z^A + p_z^A v_y^A \\
p_z^C v_x^C - p_x^C v_z^C - p_z^A v_x^A + p_x^A v_z^A \\
p_x^C v_y^C - p_y^C v_x^C - p_x^A v_y^A + p_y^A v_x^A
\end{pmatrix}
$$

Since the matrix is so small, there's no need to do anything fancy. The Go
solution here does standard Gaussian elimination with partial pivoting, and
back substitution to solve for $\mathbf{p}_0^{(R)}$. Since $\mathbf{v}^{(R)}$
isn't actually needed, it uses a slightly different arrangement of the matrix so
that the values of $\mathbf{p}_0^{(R)}$ fall out first from the row echelon
form.

Finally, there's the question on how to choose the hailstones $A$, $B$ and $C$.
Honestly, probably just picking the first three would work, but out of an
abundance of caution (okay, superstition), the code picks them using the
following process:

1. Select the first hailstone in the list as $A$.
2. Select the first remaining hailstone whose velocity vector is not parallel to
   $A$'s as $B$.
3. Select the first remaining hailstone whose velocity vector is not in the same
   plane as that defined by $A$'s and $B$'s as $C$.
