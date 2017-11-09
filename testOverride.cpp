#include <iostream>
using namespace std;
class Box
{
	public:
		void setWidth(double w);
		void setHeight(double h);
		double getArea();
	    //类成员函数的运算符重载
		Box operator+(const Box&b);
	private:
		double width;
		double height;
		double area;
};
void  Box::setWidth(double w)
{
	width = w;
}
void Box::setHeight(double h)
{
	height = h;
}
double Box::getArea()
{
	return width*height;
}
// 重载+运算符，用于把两个Box对象相加
Box Box::operator+(const Box&b)
{
	Box box;
	box.width = this->width + b.width;
	box.height = this->height + b.height;
	return box;
}
int main()
{
  
	Box box,box1;
	Box boxAll;
	//box对象的宽高
	box.setWidth(2.5);
	box.setHeight(3.2);
	//box1对象的宽高
	box1.setWidth(10.1);
	box1.setHeight(10.1);
	//利用相同对象的重载运算符，把box和box1对象的宽高相加
	boxAll = box + box1;
	cout << "Area:"<<box.getArea()<<endl;
	cout << "Area1:"<<box1.getArea()<<endl;
	cout << "Both Area:"<<boxAll.getArea()<<endl;
    return 0;
}
