
//参考自徐晓鑫《后台开发》,记录之。

#include <set> 
#include <string>
#include <iterator>
#include <iostream>
#include <string.h> //strcmp
using namespace std;

struct strLess{
    bool operator() (const char *s1,const char *s2) const {
        return strcmp(s1,s2) < 0;
    }
};
void printSet(set<int> s){
    copy(s.begin(), s.end(), ostream_iterator<int>(cout,", "));
    cout<<endl;
}
int main()
{
    /*创建set对象，共5种方式,如果无自定义函数对象,则采用系统默认方式*/
    //方式1：创建空的set对象，元素类型为int
    set<int> s1;
    //方式2：创建空的set对象，元素类型为char*,比较函数对象（即排序准则）为自定义strLess
    set<const char*,strLess> s2(strLess);
    //方式3：利用set对象s1,拷贝生成set对象s2
    set<int> s3(s1);
    //方式4：用迭代区间[&first,&last]所指的元素，创建一个set对象
    int iArray[] = {13,32,19};
    set<int> s4(iArray,iArray+3);
    //方式5：用迭代区间[&first,&last]所指的元素，及比较函数对象strLess,创建一个set对象
    const char* szArray[] = {"hello","world","bird"};
    set<const char*,strLess> s5(szArray,szArray+3,strLess());
    
    /*元素插入：共3种方式*/
    //方式1：插入value，返回pair配对对象
    cout<<"s1.insert() :"<<endl;
    for (int i=4;i>=0;i--)
        s1.insert(i*10); //set会自动调整顺序
    printSet(s1);
    //根据.second判断是否插入成功（注：value不能与set容器内元素重复）
    cout<<"s1.insert(20).second = "<<endl;
    if (s1.insert(20).second){
        cout<<"insert ok"<<endl;
        printSet(s1);
    }
    else
        cout<<"insert failed"<<endl;
    
    cout<<"s1.insert(50).second = "<<endl;
    if (s1.insert(50).second){
        cout<<"insert ok"<<endl;
        printSet(s1);
    }
    else
        cout<<"insert failed"<<endl;
    //根据p对象.second判断是否插入成功    
    cout<<"pair<set<int>::iterator,bool> p;\n p=s1.insert(60);\n if(p.second):"<<endl;
    pair<set<int>::iterator,bool> p;
    p = s1.insert(60);
    if(p.second){
        cout<<"insert ok"<<endl;
        printSet(s1);
    }
    else
        cout<<"insert Failed"<<endl;
    
    /*元素删除：共4种方式*/
    //方式1：移除set容器内元素值为value的所有元素，并返回移除元素个数
    cout<<"\ns1.erase(70) = "<<endl;
    s1.erase(70);
    printSet(s1);
    cout<<"s1.erase(60) = "<<endl;
    s1.erase(60);
    printSet(s1);
    //方式2：移除pos位置的元素，无返回值
    cout<<"set<int>::iterator iter = s1.begin();\ns1.erase(iter) = "<<endl;
    set<int>::iterator iter = s1.begin();
    s1.erase(iter);
    printSet(s1);
    
    /*元素查找：共2种方式*/
    //方式1：count(value)返回set对象内元素值为value的元素个数
    cout <<"\ns1.count(10) = "<<s1.count(10)<<",s1.count(80) = " <<s1.count(80)<<endl;
    //方式2：find(value)返回value所在位置，找不到则返回end()
    cout <<"s1.find(10): ";
    if(s1.find(10) != s1.end())
        cout <<"find it"<<endl;
    else
        cout<<"not found!"<<endl;
    
    cout <<"s1.find(80): ";
    if(s1.find(80) != s1.end())
        cout <<"found it"<<endl;
    else
        cout<<"not found!"<<endl;
    /*其他常用函数*/
    cout<<"\ns1.empty()="<<s1.empty()<<"  s1.size()"<<s1.size()<<endl;
    set<int> s9;
    s9.insert(1000);
    cout<<"s1.swap(s9) :" <<endl;
    s1.swap(s9);
    cout<<"s1:"<<endl;
    printSet(s1);
    cout<<"s9:"<<endl;
    printSet(s9);
    return 0;
}

输出：
s1.insert() :
0, 10, 20, 30, 40, 
s1.insert(20).second = 
insert failed
s1.insert(50).second = 
insert ok
0, 10, 20, 30, 40, 50, 
pair<set<int>::iterator,bool> p;
 p=s1.insert(60);
 if(p.second):
insert ok
0, 10, 20, 30, 40, 50, 60, 

s1.erase(70) = 
0, 10, 20, 30, 40, 50, 60, 
s1.erase(60) = 
0, 10, 20, 30, 40, 50, 
set<int>::iterator iter = s1.begin();
s1.erase(iter) = 
10, 20, 30, 40, 50, 

s1.count(10) = 1,s1.count(80) = 0
s1.find(10): find it
s1.find(80): not found!

s1.empty()=0  s1.size()5
s1.swap(s9) :
s1:
1000, 
s9:
10, 20, 30, 40, 50, 


