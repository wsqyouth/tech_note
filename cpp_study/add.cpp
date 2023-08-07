说明：2017/11/16 纯手打熟悉语法
数据封装：
公有成员addVal 和 getVal 是对外的接口，用户需要知道它们以便使用类。私有成员 total 是对外隐藏的，用户不需要了解它，但它又是类能正常工作所必需的。
-----
#include <iostream>
using namespace std;

class Adder
{
  public:
    Adder(int i=0)
    {
       total = i;
    }
    
    int getVal()
    {
        return total;
    }
    
    void addVal(int x);
  private:
    int total;
};

void Adder::addVal(int x)
{
    total += x;
}

int main()
{
    Adder obj;
    cout << "before:"<<obj.getVal()<<endl;
    obj.addVal(20);
    cout << "after:"<<obj.getVal()<<endl;
   
    return 0;
}
