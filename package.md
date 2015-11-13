## package的内容
首先有3个概念要理清楚<br>
1. package名：在.go文件第一句声明的pacakge的名
2. 路径名：.go文件所在的路径（相对于GOROOT或者GOPATH）
3. 文件名：.go文件的文件名（这个其实对整个过程不起任何的作用，完全可以忽略）
我们知道go install会安装第三方的项目，也可以生成自己的工程项目，但是很多新手一直搞不清楚中间的路径关系，我测试了一些，然后结合自己的理解，写了一点记录性的东西。
首先是一个项目的名字，我写了一个testPackage，目录如下
* testPackage
  * sql
    * mydriver
      * mydriver.go(pacakge mydriver)
    * sql.go(package sql)
> Note:这里mydriver.go是什么真的无所谓，关键是里面的pacakge，最好要是和他的上一层路径是一样的，也就是和mydriver一样
那么在sql.go中如果要引入mydriver这个package，就需要

`
  import "testPackage/sql/mydriver"
`

这里就可以看出，其实一切都是相对于(GOROOT和GOPATH来的)，如果你这样想，import的是最后编译出来的.a文件，就会认为与他们之间的相对路径不存在任何的问题了。（我之前一直在纠结这个问题）
## 最后看一下编译后文件
- testPacakge
  - sql
    - mydriver.a
  - sql.a
  
## 结论
1. import的包只和有关路径名有关
2. 最后生成的.a文只和上一文件夹名有关系
3. 代码里的使用还是和package有关系的
4. .go文件名在整个过程没有起到任何的作用
