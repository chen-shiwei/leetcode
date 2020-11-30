# 递归三步（栈:入入栈(递归函数)出出栈）
1. 递归函数的功能
2. 递归终止条件
3. 递归函数的表达

# 递归优化
- 存储递归计算过程
  - 用hash、array
- 自底向上(递推)
    
    - ```go
        func f(n int) int {
            if n <= 2 {
                return n
            }     
            var a,b = 1,2
            var sum int
            for i:=0; i<=n; i++{
                sum = a + b
                a = b
                b = sum
            }
            return sum  
        } 
      ```


