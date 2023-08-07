#include <iostream>
using namespace std;

class Line
{
public:
	Line(double len); //构造函数
	Line(const Line &obj);//拷贝构造函数
	void setLine(double len);
	double getLine(void);
	~Line();//析构函数
private:
	double *ptr;
};
 Line::Line(double len)
{
	 // 为指针分配内存
	 ptr = new double;
	 *ptr = len;
	 cout << "object is being created!" <<endl;
}
Line::Line(const Line &obj)
{
	 // 为指针分配内存
	 ptr = new double;
	 *ptr = *(obj.ptr);
	 cout << "object is being copied!" <<endl;
}
 Line::~Line()
{
	cout << "object is being deleted!" <<endl;
}
void Line::setLine(double len)
{
	 *ptr = len;
}
double Line::getLine(void)
{
	return *ptr;
}
int main()
{
	{
		Line lineObj(0.15);
		cout << "the default member value:"<<lineObj.getLine()<< endl;
		lineObj.setLine(3.1415926);
		cout << "the changed member value:"<<lineObj.getLine()<< endl;
		
		
		Line lineObjcopy(lineObj);
		cout << "the default member value:"<<lineObjcopy.getLine()<< endl;
	}
	cout << "Hello, world!" << endl;
    return 0;
}
