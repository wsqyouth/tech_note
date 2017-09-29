
#include <iostream>
/*
*静态成员函数即使在类对象不存在的情况下也能被调用，静态函数只要使用类名加范围解析运算符 :: 就可以访问。
*静态成员函数只能访问静态成员数据、其他静态成员函数和类外部的其他函数。
*静态成员函数没有 this 指针，只能访问静态成员（包括静态成员变量和静态成员函数）
*/
using namespace std;

class Box
{
public:
    Box(int l,int w=10,int h=20);
    void printInfo();
    static int  objCount;
    static int printCount();
private:
    int length;
    int width;
    int height;
};

Box::Box(int l,int w,int h)
{
    length = l;
    width = w;
    height = h;
    objCount++;  //静态成员在类的所有对象中是共享的。
    cout << "constructor is called" << endl;
}
int Box::objCount = 0;
//静态成员函数只能访问静态成员（包括静态成员变量和静态成员函数）
int Box::printCount()
{
    return objCount;
}

void Box::printInfo()
{
    cout << "length =" <<length<<endl
         << "width =" <<width <<endl
         << "height ="<<height<<endl;
}
int main(void)
{
   Box Box1(3.3);    // 声明 box1(利用参数初始化法测试)
   Box Box2(8.5, 6.0, 2.0);    // 声明 box2
   
   //输出对象信息
   Box1.printInfo();
   Box2.printInfo();
   
   // 输出对象的总数
   cout << "Total objects: " << Box::printCount() << endl;

   return 0;
}
