/*
这里模块的目的是说，假如参数是树上的任意节点id，父传了，子都会传，而希望得到的是除去子的，详见下例

假如有一棵树型结构
  1
 / \
2   3

当参数为 [1, 2, 3] 时，最终入库的希望只有 1
*/
package tree
