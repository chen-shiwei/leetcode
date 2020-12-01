package _9_单词搜索

import "fmt"

func exist(board [][]byte, word string) bool {
    x:=len(board)
    y:=len(board[0])
    // 记录走过的位置
    boardShadow:=make([][]bool,x)
    for i:=0;i<x;i++{
        boardShadow[i] = make([]bool,y)
    }

    var check  func(i,j int,wordIdx int) bool
    check = func(i,j int,wordIdx int) bool{
        if board[i][j] != word[wordIdx]{
            return false
        }
        if boardShadow[i][j] {
            return false
        }
        // 匹配到最后的位置
        if wordIdx == len(word) - 1{
            return true
        }
        boardShadow[i][j] = true
        // 清除本次走过对的位置
        defer func() {
            boardShadow[i][j] = false
        }()
        // top
        if i - 1 >= 0 && check(i-1,j,wordIdx+1){
            return true
        }
        // bottom
        if i+1 < x && check(i+1,j,wordIdx+1){
            return true
        }
        // left
        if j-1 >= 0 && check(i,j-1,wordIdx+1){
            return true
        }
         // right
        if j+1 < y && check(i,j+1,wordIdx+1){
            return true
        }

        return false
    }

    for i:=0;i<x;i++{
        for j:=0;j<y;j++{
            if check(i,j,0){
                return true
            }
        }
    }
    return false
}
