//覆盖：
//在基类中定义了一个非虚拟函数，然后在派生类中又定义了一个同名同参数同返回类型的函数，这就是覆盖了。
//在派生类对象上直接调用这个函数名，只会调用派生类中的那个。
#include <iostream>

using namespace std;

class A
{
	public:
	  	void ShowMessage();
};


class B:public A
{
	public:
	 	void ShowMessage();

};

 
void A::ShowMessage()
{

  cout<<"Hello,This is A.\n";
  return;
}

void B::ShowMessage()
{
  cout<<"Hello,This is B.\n";
  return;
}

int main()
{
	A* p;
	B memb;
	
    *p = memb;
    p->ShowMessage();
	
    memb.ShowMessage();

  return 0;

}
-------------------------------
//重载：
//有两个或多个函数名相同的函数，但是函数的形参列表不同。在调用相同函数名的函数时，根据形参列表确定到底该调用哪一个函数。
#include <iostream>
using namespace std;

class A
{
	public:
	  void ShowMessage();
	  void ShowMessage(string str);
};

 
void A::ShowMessage()
{
	cout<<"Hi,This is A.\n";
 	return;
}

void A::ShowMessage(string str)
{
  cout<<str<<endl;
  return;
}

 
int main()
{
  A mem;
  mem.ShowMessage();
  mem.ShowMessage("Hello.How are you?\n");
	
  return 0;
}

-------------------------------
//多态：
//在基类中定义了一个虚拟函数，然后在派生类中又定义一个同名，同参数表的函数，这就是多态。多态是这3种情况中唯一采用动态绑定技术的一种情况。
//也就是说，通过一个基类指针来操作对象，如果对象是基类对象，就会调用基类中的那个函数，如果对象实际是派生类对象，就会调用派声雷中的那个函数
//调用哪个函数并不由函数的参数表决定，而是由函数的实际类型决定。(一处声明，多处实现)
#include <iostream>

using namespace std;

class A
{
	public:
  		 virtual void ShowMessage();
};

class B:public A
{
	public:
 	    void ShowMessage();
};

void A::ShowMessage()
{
  cout<<"This is A.\n";
  return;
}

void B::ShowMessage()
{
  cout<<"This is B.\n";
  return;
}

int main()
{
  A* p;
  p=new A();
  p->ShowMessage();
  p=new B();
  p->ShowMessage();

  return 0;
}

 
