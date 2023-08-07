#include <iostream>       
#include <vector>
using namespace std;
int main(void)
{
    vector<int>array;
    array.push_back(100);
    array.push_back(300);
    array.push_back(300);
    array.push_back(300);
    array.push_back(300);
    array.push_back(500);
    vector<int>::iterator itor;
    for(itor=array.begin();itor!=array.end();itor++)
    {
        if(*itor==300)
        {
            itor=array.erase(itor); //左值是必要的，否则会变为野指针
        }
    }
    for(itor=array.begin();itor!=array.end();itor++)
    {
            cout<<*itor<<"";
    }
  return 0;
}
