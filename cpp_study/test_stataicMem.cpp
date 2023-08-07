#include <iostream>
/*
*当我们声明类的成员为静态时，这意味着无论创建多少个类的对象，静态成员都只有一个副本。
*不能把静态成员的初始化放置在类的定义中，但是可以在类的外部通过使用范围解析运算符 :: 来重新声明静态变量从而对它进行初始化
*/
using namespace std;

class Box
{
public:
    Box(int l,int w=10,int h=20);
    void printInfo();
    static int  objCount;
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
   cout << "Total objects: " << Box::objCount << endl;

   return 0;
}
