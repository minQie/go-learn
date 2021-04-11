package _proxy

// 代理分 静态代理 和 动态代理

/*
静态代理

定义一个接口作为被代理和代理都要实现的标准
定义一个结构体作为被代理
定义一个结构体作为代理：包含一个被代理的实例，再要求实现的接口方法中，会调用被代理对象的接口实现，然后可以在被代理接口实现逻辑的上下加上代理的逻辑

实际调用方法实现时，是创建一个代理对象，去调用
*/

/*
动态代理

不同于 Java，Go 中无法比较方便的利用反射实现动态代理，但是我们可以利用 go generate 实现类似的效果，并且这样实现有两个比较大的好处，
一个是有静态代码检查，我们在编译期间就可以及早发现问题，第二个是性能会更好
（在真实的项目中并不推荐这么做，因为有点得不偿失，本文只是在探讨一种可能性，并且可以复习一下 go 语法树先关的知识点）

动态代理相比静态代理主要就是为了解决生产力，将我们从繁杂的重复劳动中解放出来，正好，在 Go 中 Generate 也是干这个活的

目的：为 被 格式如【// @proxy XxxInterface】修饰的结构体创建生成代理类的源码

思路：
    读取文件, 获取文件的 ast 语法树
    通过 NewCommentMap 构建 node 和 comment 的关系
    通过 comment 是否包含 @proxy 接口名 的接口，判断该节点是否需要生成代理类
    通过 Lookup 方法找到接口
    循环获取接口的每个方法的，方法名、参数、返回值信息
    将方法信息，包名、需要代理类名传递给构建好的模板文件，生成代理类
    最后用 format 包的方法格式化源代码
*/