# Largest Rectangle in Histogram - Optimization Analysis

## Problem Overview
Find the area of the largest rectangle that can be formed in a histogram represented by an array of heights.

## Current Implementation Analysis

### Current Approach: Two-Pass with Left/Right Arrays
```go
func largestRectangleArea(heights []int) int {
    // Pass 1: Find left boundaries (closest smaller elements)
    left := make([]int, len(heights))
    // Initialize and populate left array...
    
    // Pass 2: Find right boundaries (closest smaller elements) 
    right := make([]int, len(heights))
    // Initialize and populate right array...
    
    // Pass 3: Calculate maximum area
    for i := 0; i < len(heights); i++ {
        area := (right[i] - left[i] - 1) * heights[i]
        ans = max(ans, area)
    }
}
```

**Time Complexity:** O(n)  
**Space Complexity:** O(n)  
**Passes:** 3 (left array, right array, final calculation)

## Optimization Strategies

### 1. 🏆 Single-Pass Stack Solution (Recommended)

The most significant optimization is to eliminate multiple passes by using a monotonic stack approach:

```go
func largestRectangleAreaOptimized(heights []int) int {
    stack := make([]int, 0, len(heights)) // Pre-allocate capacity
    maxArea := 0
    
    // Process all heights + one extra iteration for cleanup
    for i := 0; i <= len(heights); i++ {
        // Current height (0 for the extra iteration)
        currentHeight := 0
        if i < len(heights) {
            currentHeight = heights[i]
        }
        
        // Process all bars that are taller than current height
        for len(stack) > 0 && currentHeight < heights[stack[len(stack)-1]] {
            // Pop the tallest bar
            height := heights[stack[len(stack)-1]]
            stack = stack[:len(stack)-1]
            
            // Calculate width
            width := i
            if len(stack) > 0 {
                width = i - stack[len(stack)-1] - 1
            }
            
            // Update maximum area
            maxArea = max(maxArea, height*width)
        }
        
        // Push current index
        if i < len(heights) {
            stack = append(stack, i)
        }
    }
    
    return maxArea
}
```

**Benefits:**
- **Fewer iterations:** 1 pass instead of 3
- **Better cache locality:** Process data sequentially
- **Cleaner logic:** Calculate area immediately when bounds are known
- **Memory efficiency:** Single stack instead of two arrays

### 2. Memory Optimization for Current Approach

If keeping the existing approach, optimize memory allocation:

```go
// More efficient initialization
left := make([]int, len(heights))
for i := range left {
    left[i] = -1  // More idiomatic than separate loop
}

// Pre-allocate stack capacity
stack := make([]int, 0, len(heights))
```

### 3. Early Termination

Add edge case handling to avoid unnecessary computation:

```go
func largestRectangleArea(heights []int) int {
    if len(heights) == 0 {
        return 0
    }
    if len(heights) == 1 {
        return heights[0]
    }
    
    // Main algorithm...
}
```

## Performance Comparison

| Approach | Time | Space | Passes | Cache Efficiency |
|----------|------|--------|--------|------------------|
| Current (Two-pass) | O(n) | O(n) | 3 | Moderate |
| Single-pass Stack | O(n) | O(n) | 1 | High |

## Algorithm Explanation: Single-Pass Stack

### Key Insight
Use a **monotonic increasing stack** to maintain indices of bars in ascending order of heights. When we encounter a shorter bar, we know we've found the right boundary for all taller bars in the stack.

### Step-by-Step Process:
1. **Maintain stack** of indices where heights are in ascending order
2. **When current height < stack top height:**
   - Pop the stack (this is our target bar)
   - Calculate width using current position and remaining stack top
   - Update maximum area
3. **Push current index** to stack
4. **Extra iteration** with height 0 to process remaining bars

### Visual Example:
```
Heights: [2, 1, 5, 6, 2, 3]
         0  1  2  3  4  5

i=0: stack=[0], heights=[2]
i=1: height=1 < heights[0]=2, pop 0, area=2*1=2, stack=[1]
i=2: stack=[1,2], heights=[1,5]  
i=3: stack=[1,2,3], heights=[1,5,6]
i=4: height=2 < heights[3]=6, pop 3, area=6*1=6
     height=2 < heights[2]=5, pop 2, area=5*2=10
     stack=[1,4]
i=5: stack=[1,4,5]
i=6: (cleanup) pop all remaining, final max area = 10
```

## Implementation Considerations

### Stack Operations
- Use slice with pre-allocated capacity: `make([]int, 0, len(heights))`
- Avoid frequent reallocations

### Boundary Handling
- Add extra iteration with height 0 to ensure all bars are processed
- Handle empty stack case when calculating width

### Edge Cases
- Empty histogram: return 0
- Single bar: return height of that bar
- All same heights: return height × length

## Testing Strategy

The optimized solution should pass all existing test cases. Key test scenarios:
- Edge cases (empty, single bar)
- Monotonic sequences (ascending, descending)
- Complex patterns (mountain, valley, plateau)
- Performance with large uniform arrays

## Conclusion

The **single-pass stack solution** provides the best optimization by:
- Reducing algorithm complexity from 3 passes to 1 pass
- Improving cache locality and memory access patterns  
- Maintaining the same O(n) time complexity with better constants
- Using cleaner, more intuitive logic

This optimization is particularly beneficial for large datasets where reducing the number of passes significantly improves performance.