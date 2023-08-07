/**
类的友元函数是定义在类外部，但有权访问类的所有私有（private）成员和保护（protected）成员。
尽管友元函数的原型有在类的定义中出现过，但是友元函数并不是成员函数。
如果要声明函数为一个类的友元，需要在类定义中该函数原型前使用关键字 friend
**/
#include <iostream>
using namespace std;

class Box
{
private:
    double width;
public:
    Box();
    void setWidth(double wid);
    friend void printWidth( Box & box);

};

Box::Box()
{
    width = 0.0;
}

void Box::setWidth(double wid)
{
    width = wid;
}
// 请注意：printWidth() 不是任何类的成员函数
void printWidth( Box & box )
{
   /* 因为 printWidth() 是 Box 的友元，它可以直接访问该类的任何成员 */
   cout << "Width of box : " << box.width <<endl;
}
int main() {
    Box boxObj;
	// 使用友元函数输出宽度
    printWidth(boxObj);
	// 使用成员函数设置宽度
    boxObj.setWidth(3.14);
    printWidth(boxObj);
	cout  << "test\n";
	return 0;
}
