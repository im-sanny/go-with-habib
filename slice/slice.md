## 🎯 The 7 Golden Rules of Slices

### 1. **The Reallocation Rule**

> "Append causes reallocation ONLY when `len == cap`, creating a NEW array"

**When you see `append()`**: Always check `if len == cap` in your mind

### 2. **The Sharing Rule**

> "Slice assignment (`y := x`) shares the underlying array, but append with reallocation breaks the sharing"

**Remember**: They're connected until reallocation separates them!

### 3. **The Capacity Rule**

> "Slicing operations (`x[a:b]`) inherit capacity from the parent slice"

```go
arr := [5]int{1,2,3,4,5}
slice := arr[1:3]  // [2,3], but cap=4 (can access arr[3] and arr[4])
```

### 4. **The Modification Rule**

> "Modifications affect ALL slices sharing the same array until reallocation"

**Test trick**: If they share array, changes propagate; if not, they don't!

### 5. **The Length vs Capacity Rule**

> "Length is what you SEE, capacity is what you CAN use without reallocation"

**Mental model**:

- `len` = occupied seats
- `cap` = total seats in the theater

### 6. **The Make Rule**

> "`make([]T, len, cap)` creates slice with length `len` and optional capacity `cap`"

```go
s := make([]int, 2, 5)  // [0,0] but can append 3 more without reallocation
```

### 7. **The Zero Value Rule**

> "`var s []int` creates nil slice, `s := []int{}` creates empty non-nil slice"

```go
var a []int        // nil, len=0, cap=0
b := []int{}       // empty, len=0, cap=0 (but not nil)
c := make([]int,0) // empty, len=0, cap=0 (but not nil)
```

## 🚨 Common Pitfalls & Quick Checks

### Quick Check 1: "Will they stay connected?"

```go
x := []int{1,2,3}
y := x
// Ask: Is cap > len? If YES → they stay connected after append
```

### Quick Check 2: "Does append create new array?"

```go
// Before append, ask: len == cap?
// If YES → new array, if NO → same array
```

### Quick Check 3: "Will modification affect both?"

```go
// Ask: Did ANY append cause reallocation between them?
// If NO → changes affect both, if YES → changes affect only one
```

## 📝 Exam Cheat Sheet (Mental Checklist)

When solving slice problems, ask:

1. **Initial state**: What's the len/cap?
2. **Assignment**: Are they sharing array? (`y := x` → YES)
3. **Append**: `len == cap`? → REALLOCATION!
4. **Modification**: Do they still share array?
5. **Final output**: Trace each slice's current array

## 🎯 Pro Tips for Tests

**Watch for these trick questions:**

- Slicing beyond length but within capacity (`x[:4]` when len=3, cap=5)
- Multiple appends in sequence
- Chains of slice assignments (`a := b; c := a`)
- Append with multiple values (`append(x, 4, 5, 6)`)

**Memory Aid**: "**L**ength **C**apacity **R**eallocation **S**haring"

- Check **L**ength and **C**apacity
- Determine if **R**eallocation happens
- Track **S**haring relationships
