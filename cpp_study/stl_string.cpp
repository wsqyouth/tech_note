一、string类字符串的介绍
    在程序设计中, 字符串的使用十分频繁, C语言类型字符串(简称C-串)在使用与字符串的处理上较为复杂, C++为了在程序设计中更加方便的使用字符串特
新增了一种string类型的字符串。
    string类字符串为STL(Standard Template Library, 标准模板库)中的一种自定义的数据类型, 相对于C-串来说string类型串具有一些明显的优势, 
    首先, 它在内存使用上是自动的, 需要多少, 开辟多少, 并且能够根据字符串大小的变化自动调整所开辟的内存; 
    此外, string串还提供了大量的方法以便更好的完成对字符串的各种操作。
    
二、声明一个string型字符串
    同普通变量一样, string类型的字符串在使用前也需进行声明, 并且也可以对其进行相关的初始化操作, 相关的声明以及初始化方法如下:
        string s;                                                       //声明一个string型字符串s
        string s(const string &str);                                    //声明string型字符串s并用另一个string型字符串str对其进行初始化
        string s(const string &str, size_type n);                       //将字符串str中起始位置n后的字符串作为字符串s的初始值
        string s(const string &str, size_type n, size_type m);          //将字符串str位置n起, 长为m的部分的字符作为字符串s的初始值
        string s(const char *cs) ;                                      //将C-串cs作为string串s的初始值
        string s(const char *cs, size_type n);                          //将C-串cs的前n个字符作为string串s的初始值
        string s(const char *cs, size_type n, size_type m);             //将C-串cs中位置n起, 长为m的部分的字符作为字符串s的初始值
        string s(size_type num, char c);                                //初始化字符串值为num个c字符
        string s(iterator begin, iterator end);                         //将区间[begin, end]内的字符作为字符串s的初始值
        
示例：
#include<iostream>

using namespace std ;

int main()
{
    char cs[] = "hello,world" ;         //声明并初始化一个C-串
    string str ;                        //声明一个string串
    str = "hello,world" ;               //对string串进行赋值
    cout<<"str="<<str<<endl ;

    //使用string类型初始化另一个string类型
    string s1(str) ;
    cout<<"str1="<<s1<<endl ;
    string s2(str, 2) ;
    cout<<"str2="<<s2<<endl ;
    string s3(str, 2, 5) ;
    cout<<"str3="<<s3<<endl ;

    //使用C-串初始化string类型串
    string s4(cs) ;
    cout<<"str4="<<s4<<endl ;
    string s5(cs, 2) ;
    cout<<"str5="<<s5<<endl ;
    string s6(cs, 2, 5) ;
    cout<<"str6="<<s6<<endl ;

    return 0 ;
}

三、字符串的输入输出
    除了已经学习的 ">>"、"cin.get()"和"cin.getline()"对字符串进进行输入外, string头文件中还定义了getline()函数用于输入string字符串。
    getline的函数原型如下:

        istream& getline ( istream &is , string &str , char delim );        //形式一
        istream& getline ( istream& , string& );                            //形式二
 
    getline的第一个函数参数为输入流对象; 第二个为待输入的字符串; 第三个是可选参数, 为自定义的终止符。
    当输入到该字符时表示输入完成, 程序只保存终止符前的输入内容, 当省略时默认以'\n'为终止符。需要说明的是, 终止符不会保存到输入的字符串中去。
示例：
#include<iostream>
using namespace std ;

int main()
{
    string s;
    getline(cin, s) ;           //使用默认的'\n'作为终止符
    cout<<s<<endl ;

    getline(cin, s, '!') ;      //以'!'作为终止符
    cout<<s<<endl;

    return 0 ;
}
四、string串的基本使用方法
    在string类型的字符串中, 字符串的处理得到极大的简化, 例如原本在C-串中的复制操作, 需要借助string.h中的strcpy()函数才能完成, 
    而在string串中只需一个'='进行赋值就能完成。更具体的如下:
    1>. 复制
            string s1 = "hello" ;
            string s2 = s1 ;        //复制
    2>. 连接
            string s1 = "hello" ;
            string s2 = "world" ;
            s1 += s2 ;                //连接
   3>. 比较
            string s1 = "hello" ;
            string s2 = "world" ;
            if(s1 < s2)
                cout<<"s1 < s2" ;    //比较
   4>. 倒置串
            string s = "hello" ;
            reverse(s.begin(), s.end()) ;        //需要包含algorithm头文件, #include<algorithm> 
   5>. 查找串
            string s = "hello" ;
            cout<<s.find("ll") ;        //返回子串第一次出现的位置
   6>. 替换
            string s = "hello" ;
            s.replace(0, 2, "aaa") ;    //将字符串s中下标0-2部分字符串替换为"aaa"
            
五、string的更多方法
    由于string类型的字符串自身提供的方法太多, 这里不能一一详述, 只选择一些常用的来进一步说明。
  1>. 获取字符串状态
            s.size()                //返回字符串大小
            s.length()              //返回字符串长度
            s.max_size()            //返回字符串最大长度
            s.clear()               //清空字符串
            s.empty()               //判断字符串是否为空
 2>. 修改字符串
            ①. append - 追加
                string s = "hello" ;
                s.append("world") ;        //将"world"追加到s中
    
            ②. push_back - 追加字符到字符串
                string s = "hello" ;
                s.push_back('!') ;        //将'!'追加字符到字符串s中
    
            ③. insert - 插入
                string s = "hello" ;
                s.insert(2, "www") ;    //将字符串"www"插入到字符串s中, 插入位置为2
            
            ④. erase - 从字符串中擦除一些字符
                string s = "hello" ;
                s.erase(1, 2) ;            //从下标为1处向后擦去2个字符
            
            ⑤. swap - 与另一字符串交换内容
                string s1 = "hello" ;
                string s2 = "world" ;
                s1.swap(s2) ;            //将s1与s2中的字符串进行交换
