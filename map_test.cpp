#include <map>
#include <string>
#include <iostream>
using namespace std;

int main()
{
    map<int,string> mapStudent;
    mapStudent[1] = "student_one";
    mapStudent[2] = "student_Two";
    mapStudent[3] = "student_Three";
    
    //使用反向迭代器 
    map<int,string>::reverse_iterator iter;
    for(iter = mapStudent.rbegin(); iter != mapStudent.rend(); iter++)
        std::cout << iter->first <<" "<<iter->second << std::endl;
        
    //使用数组 （可覆盖 ）
    int iSize = mapStudent.size();
    for(int i = 1; i <= iSize; i++)
        cout << i << mapStudent[i] <<endl;
        
    
    return 0;
}
