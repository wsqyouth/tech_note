说明：2017/11/16 纯手打熟悉语法

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
