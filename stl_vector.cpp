
向量 vector 是一种对象实体, 能够容纳许多其他类型相同的元素, 因此又被称为容器。
在使用它时, 需要包含头文件 vector, #include<vector>
vector 容器与数组相比其优点在于它能够根据需要随时自动调整自身的大小以便容下所要放入的元素。此外, vector 也提供了许多的方法来对自身进行操作。

一、向量的声明及初始化
vector 型变量的声明以及初始化的形式也有许多, 常用的有以下几种形式:

        vector<int> a ;                                //声明一个int型向量a
        vector<int> a(10) ;                            //声明一个初始大小为10的向量
        vector<int> a(10, 1) ;                         //声明一个初始大小为10且初始值都为1的向量
        vector<int> b(a) ;                             //声明并用向量a初始化向量b
        vector<int> b(a.begin(), a.begin()+3) ;        //将a向量中从第0个到第2个(共3个)作为向量b的初始值
二、元素的输入及访问
元素的输入和访问可以像操作普通的数组那样, 用cin>>进行输入, cout<<a[n]这样进行输出
或者使用方法 assign  insert push_back()
    //全部输出
    vector<int>::iterator t ;
    for(t=a.begin(); t!=a.end(); t++)
        cout<<*t<<" " ;
        
    *t 为指针的间接访问形式, 意思是访问t所指向的元素值。

三、方法
1>. a.size()                 //获取向量中的元素个数
2>. a.empty()                //判断向量是否为空
3>. a.clear()                //清空向量中的元素
4>. 复制
        a = b ;            //将b向量复制到a向量中
5>. 比较
        保持 ==、!=、>、>=、<、<= 的惯有含义 ;
        如: a == b ;    //a向量与b向量比较, 相等则返回1
6>. 插入 - insert
        ①、 a.insert(a.begin(), 1000);            //将1000插入到向量a的起始位置前
        ②、 a.insert(a.begin(), 3, 1000) ;        //将1000分别插入到向量元素位置的0-2处(共3个元素)
        ③、 vector<int> a(5, 1) ;
            vector<int> b(10) ;
            b.insert(b.begin(), a.begin(), a.end()) ;        //将a.begin(), a.end()之间的全部元素插入到b.begin()前
 7>. 删除 - erase
        ①、 b.erase(b.begin()) ;                     //将起始位置的元素删除
        ②、 b.erase(b.begin(), b.begin()+3) ;        //将(b.begin(), b.begin()+3)之间的元素删除
8>. 交换 - swap
        b.swap(a) ;            //a向量与b向量进行交换
        
        
-----
#include <iostream>
#include <iterator>
#include <vector>
using namespace std;

int main()
{
	//大小为10初值为0的向量
	vector<int> v1(10,0);
	//保守容量
	v1.reserve(20);
	cout << "capacity:" <<v1.capacity() << "size:" <<v1.size() <<endl;
	//填充
	for(int i=0;i<=9;i++)
	{
		v1.push_back(i);
	}
	//v1.assign(1,99);

	//将(v1.begin(), v1.begin()+3)之间的元素删除
	v1.erase(v1.begin()+1,v1.begin()+3);
    //特定位置插入
	v1.insert(v1.begin(),99);
	//顺序输出(end指向末尾位置的下一个)
	//vector<int>::iterator pos;
	//for(pos=v1.begin();pos != v1.end(); pos++)
	//{
	//	cout << *pos << " ";
	//}
	copy(v1.begin(),v1.end(),ostream_iterator<int>(cout," "));
	cout << endl;
	//首尾元素直接访问
	cout << "head value:" << v1.front() <<endl;
	cout << "tail value " << v1.back() <<endl;
	//随机访问
	cout << "v1[2]=" <<v1[2] <<endl;
	//容量
	cout << "capacity:" <<v1.capacity() <<endl;
	//大小
	cout << v1.size() <<endl;
	//清空向量中的元素
	v1.clear();
	 //将b向量复制到a向量中
	if(v1.empty())
	{
		cout << "is empty" <<endl;
	}
	return 0;
 }
